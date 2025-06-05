import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, Button, Alert, StyleSheet, ScrollView, Modal, TouchableOpacity, ActivityIndicator, Switch } from 'react-native';
import { Picker } from '@react-native-picker/picker';
import DateTimePicker from '@react-native-community/datetimepicker';
import { DeviceEventEmitter } from 'react-native';
const AdminDashboard = ({ navigation }) => {
  const [activeYear, setActiveYear] = useState(1);
  const [activeDept, setActiveDept] = useState('CSE');
  const [config, setConfig] = useState({});
  const [saving, setSaving] = useState(false);
  const [message, setMessage] = useState('');
  const [departments, setDepartments] = useState([]);
  const [showDeptModal, setShowDeptModal] = useState(false);
  const [loadingDepts, setLoadingDepts] = useState(true);
  const [drawerOpen, setDrawerOpen] = useState(false);

  const fetchDepartments = async () => {
    try {
      setLoadingDepts(false);
      let response = await fetch('http://10.150.255.205:8080/admin/api/departments');
      
      if (!response.ok) {
        response = await fetch('http://10.150.255.205:8080/admin/api/config?year=1');
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const configData = await response.json();
        if (configData && configData.departments) {
          setDepartments(configData.departments);
          return;
        }
        throw new Error("No departments found in config");
      }
      
      const data = await response.json();
      
      if (!data.success && !data.departments) {
        throw new Error(data.error || "Failed to load departments");
      }
      
      const depts = data.departments || data;
      
      const sortedDepts = Array.isArray(depts) ? 
        depts.sort((a, b) => {
          if (a.code === 'CSE') return -1;
          if (b.code === 'CSE') return 1;
          return a.code.localeCompare(b.code);
        }) : [
          
          { code: 'CSE', name: 'Computer Science' },
          { code: 'IT', name: 'Information Technology' },
          { code: 'ECE', name: 'Electronics and Communication' },
          { code: 'EEE', name: 'Electrical and Electronics' },
          { code: 'MECH', name: 'Mechanical' },
        ];
      
      setDepartments(sortedDepts);
    } catch (error) {
      console.error('Error fetching departments:', error);
      setDepartments([
        { code: 'CSE', name: 'Computer Science' },
        { code: 'IT', name: 'Information Technology' },
        { code: 'ECE', name: 'Electronics and Communication' },
        { code: 'EEE', name: 'Electrical and Electronics' },
        { code: 'MECH', name: 'Mechanical' },
      ]);
    } finally {
      setLoadingDepts(false);
    }
  };

  useEffect(() => {
  const fetchData = async () => {
    await fetchDepartments();
    await fetchConfig(activeYear, activeDept);
  };
  
  fetchData();
  
  // Add listener for department updates
  const subscription = DeviceEventEmitter.addListener('departmentsUpdated', () => {
    fetchDepartments();
  });

  return () => {
    subscription.remove();
  };
}, [activeYear, activeDept]);

  const fetchConfig = async (year, department) => {
    try {
      setConfig({});
      const response = await fetch(
        `http://10.150.255.205:8080/admin/api/config?year=${year}&department=${department}`
      );
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      const processedConfig = calculateDependentValues(data);
      setConfig(processedConfig);
    } catch (error) {
      console.error('Error fetching config:', error);
      setMessage(`Failed to load configuration: ${error.message}`);
      
      const defaultConfig = calculateDependentValues(getDefaultConfig(year, department));
      setConfig(defaultConfig);
    }
  };

  const getDefaultConfig = (year, department) => {
    let redemptionRatio = 10;
    if (year === 2) redemptionRatio = 50;
    else if (year === 3) redemptionRatio = 90;
    else if (year === 4) redemptionRatio = 120;

    if (department === 'CSE') redemptionRatio += 5;
    else if (department === 'IT') redemptionRatio += 3;

    return {
      redemptionRatio,
      subjectsCount: 6,
      redemptionVariance: [1, 2, 4, 6],
      sectionMaxMarks: [5, 3, 2, 1],
      year,
      department
    };
  };

  const calculateDependentValues = (config) => {
    if (!config) return getDefaultConfig(activeYear, activeDept);
    
    const sectionPoints = config.sectionMaxMarks?.map((maxMarks, index) => {
      return maxMarks * config.redemptionRatio * config.redemptionVariance[index];
    }) || [];
    
    const pointsPerMark = sectionPoints.map((points, idx) => {
      return points / config.sectionMaxMarks[idx];
    });
    
    const minPointsForHalf = pointsPerMark.map(point => point / 2);
    const maxPointsPerSubject = sectionPoints.reduce((sum, val) => sum + val, 0);
    
    return {
      ...config,
      sectionPoints,
      pointsPerMark,
      minPointsForHalf,
      maxPointsPerSubject
    };
  };

  const handleSave = async () => {
    try {
      setSaving(true);
      const configToSave = {
        redemptionRatio: config.redemptionRatio,
        subjectsCount: config.subjectsCount,
        redemptionVariance: config.redemptionVariance,
        sectionMaxMarks: config.sectionMaxMarks,
        year: activeYear,
        department: activeDept
      };
      
      const response = await fetch('http://10.150.255.205:8080/admin/api/config', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          year: activeYear,
          department: activeDept,
          config: configToSave
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText);
      }

      const result = await response.json();
      
      if (!result.success) {
        throw new Error(result.error || "Failed to save config");
      }

      setMessage(`${activeDept} Year ${activeYear} configuration saved successfully`);
      await fetchConfig(activeYear, activeDept);
    } catch (error) {
      console.error('Error saving config:', error);
      setMessage(`Error saving configuration: ${error.message}`);
    } finally {
      setSaving(false);
    }
  };

  const handleIndependentValueChange = (field, value) => {
    const newConfig = { 
      ...config,
      [field]: Number(value)
    };
    const updatedConfig = calculateDependentValues(newConfig);
    setConfig(updatedConfig);
  };

  const handleIndependentArrayChange = (field, index, value) => {
    const newArray = [...config[field]];
    newArray[index] = Number(value);
    const newConfig = {
      ...config,
      [field]: newArray
    };
    const updatedConfig = calculateDependentValues(newConfig);
    setConfig(updatedConfig);
  };

  const renderArrayInputs = (field, labels, editable = true) => {
    return (
      <View style={styles.arrayContainer}>
        <Text style={styles.label}>{field}:</Text>
        {config[field]?.map((value, index) => (
          <View key={`${field}-${index}`} style={styles.arrayInput}>
            <Text style={styles.arrayLabel}>{labels ? labels[index] : `Item ${index + 1}`}</Text>
            <TextInput
              style={[styles.input, !editable && styles.readOnlyInput]}
              value={String(value)}
              onChangeText={editable ? 
                (text) => handleIndependentArrayChange(field, index, text) : 
                undefined}
              keyboardType="numeric"
              editable={editable}
            />
          </View>
        ))}
      </View>
    );
  };

  const renderValueInput = (field, editable = true) => {
    return (
      <View>
        <Text style={styles.label}>{field}:</Text>
        <TextInput
          style={[styles.input, !editable && styles.readOnlyInput]}
          value={String(config[field] || '')}
          onChangeText={editable ? 
            (text) => handleIndependentValueChange(field, text) : 
            undefined}
          keyboardType="numeric"
          editable={editable}
        />
      </View>
    );
  };

  return (
    <View style={{ flex: 1 }}>
      {/* New Header */}
      <View style={styles.header}>
        <TouchableOpacity 
          style={styles.menuButton} 
          onPress={() => setDrawerOpen(true)}
        >
          <Text style={{ color: 'white', fontSize: 24 }}>☰</Text>
        </TouchableOpacity>
        <Text style={styles.headerTitle}>Admin Dashboard</Text>
        <View style={{ width: 48 }} />
      </View>

      {/* Drawer */}
      {drawerOpen && (
        <>
          <View style={styles.overlay} onTouchEnd={() => setDrawerOpen(false)} />
          <View style={styles.drawer}>
            <TouchableOpacity 
              style={styles.drawerItem}
              onPress={() => {
                setDrawerOpen(false);
              }}
            >
              <Text style={styles.drawerItemText}>Dashboard</Text>
            </TouchableOpacity>
            <TouchableOpacity 
              style={styles.drawerItem}
              onPress={() => {
                setDrawerOpen(false);
                navigation.navigate('RedemptionDates', { 
                  department: activeDept, 
                  year: activeYear 
                });
              }}
            >
              <Text style={styles.drawerItemText}>Redemption Dates</Text>
            </TouchableOpacity>

            <TouchableOpacity 
              style={styles.drawerItem}
              onPress={() => {
                setDrawerOpen(false);
                navigation.navigate('DepartmentManagement');
              }}
            >
              <Text style={styles.drawerItemText}>Department Management</Text>
            </TouchableOpacity>

            <TouchableOpacity 
                    style={styles.drawerItem}
                    onPress={() => {
                      setDrawerOpen(false);
                      navigation.navigate('MainApp');
                    }}
                  >
                    <Text style={styles.drawerItemText}>Student Page</Text>
                  </TouchableOpacity>

          </View>
        </>
      )}

      <ScrollView contentContainerStyle={[styles.container, { paddingTop: 10 }]}>
        <Text style={styles.title}>Reward System Configuration</Text>
        
        <View style={styles.selectorContainer}>
  <View style={styles.selector}>
    <Text style={styles.selectorLabel}>Department</Text>
    <View style={styles.pickerContainer}>
      {loadingDepts ? (
        <ActivityIndicator style={styles.picker} />
      ) : (
        <Picker
          selectedValue={activeDept}
          style={styles.picker}
          onValueChange={(itemValue) => setActiveDept(itemValue)}
          dropdownIconColor="#6c5ce7"
        >
          {departments.map(dept => (
            <Picker.Item 
              key={dept.code} 
              label={`${dept.code} - ${dept.name}`} 
              value={dept.code} 
            />
          ))}
        </Picker>
      )}
    </View>
  </View>
  
  <View style={styles.selector}>
    <Text style={styles.selectorLabel}>Year</Text>
    <View style={styles.yearButtons}>
      {[1, 2, 3, 4].map(year => (
        <TouchableOpacity
          key={year}
          style={[
            styles.yearButton,
            activeYear === year && styles.yearButtonActive
          ]}
          onPress={() => setActiveYear(year)}
        >
          <Text style={[
            styles.yearButtonText,
            activeYear === year && styles.yearButtonTextActive
          ]}>
            Year {year}
          </Text>
        </TouchableOpacity>
      ))}
    </View>
  </View>
</View>

        
        <Text style={styles.configTitle}>
          Configuration for {activeDept} - Year {activeYear}
        </Text>

        <View style={styles.singleColumnContainer}>
  {renderValueInput('redemptionRatio')}
  {renderValueInput('subjectsCount')}
  {renderValueInput('maxPointsPerSubject', false)}
</View>
      <View style={styles.twoColumnContainer}>
  <View style={styles.columnLeft}>
    {renderArrayInputs('redemptionVariance', ['Variance 1', 'Variance 2', 'Variance 3', 'Variance 4'])}
  </View>
  <View style={styles.columnRight}>
    {renderArrayInputs('sectionMaxMarks', ['5 Marks', '3 Marks', '2 Marks', '1 Mark'])}
  </View>
</View>

<View style={styles.twoColumnContainer}>
  <View style={styles.columnLeft}>
    {renderArrayInputs('sectionPoints', ['5 Marks', '3 Marks', '2 Marks', '1 Mark'], false)}
  </View>
  <View style={styles.columnRight}>
    {renderArrayInputs('pointsPerMark', ['5 Marks', '3 Marks', '2 Marks', '1 Mark'], false)}
  </View>
</View>

<View style={styles.singleColumnContainer}>
  {renderArrayInputs('minPointsForHalf', ['5 Marks', '3 Marks', '2 Marks', '1 Mark'], false)}
</View>
        <Button 
          title={`Save ${activeDept} Year ${activeYear} Config`} 
          onPress={handleSave} 
          disabled={saving} 
        />

        {message ? <Text style={styles.message}>{message}</Text> : null}

        <Modal
          visible={showDeptModal}
          animationType="slide"
          transparent={false}
          onRequestClose={() => setShowDeptModal(false)}
        >
          <View style={styles.modalContainer}>
            <DepartmentManagement 
              departments={departments}
              onDepartmentAdded={() => {
                fetchDepartments();
                setShowDeptModal(false);
              }}
              onClose={() => setShowDeptModal(false)}
            />
          </View>
        </Modal>
      </ScrollView>
    </View>
  );
};

const DepartmentManagement = ({ departments, onDepartmentAdded, onClose }) => {
  const [newDeptCode, setNewDeptCode] = useState('');
  const [newDeptName, setNewDeptName] = useState('');
  const [message, setMessage] = useState('');
  const [saving, setSaving] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [deptToDelete, setDeptToDelete] = useState('');

  const handleAddDepartment = async () => {
    if (!newDeptCode || !newDeptName) {
      setMessage('Both code and name are required');
      return;
    }

    setSaving(true);
    try {
      const response = await fetch('http://10.150.255.205:8080/admin/api/departments/add', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          code: newDeptCode.toUpperCase(),
          name: newDeptName,
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText);
      }

      const result = await response.json();
      
      if (!result.success) {
        throw new Error(result.error || "Failed to add department");
      }

      setMessage('Department added successfully');
      setNewDeptCode('');
      setNewDeptName('');
      onDepartmentAdded();

      DeviceEventEmitter.emit('departmentsUpdated');
    } catch (error) {
      setMessage(`Error: ${error.message}`);
    } finally {
      setSaving(false);
    }
  };

  const handleDeleteDepartment = async () => {
    if (!deptToDelete) {
      setMessage('Please select a department to delete');
      return;
    }

    if (deptToDelete === 'CSE') {
      setMessage('Cannot delete CSE department');
      return;
    }

    setDeleting(true);
    try {
      const response = await fetch('http://10.150.255.205:8080/admin/api/departments/delete', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          code: deptToDelete,
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText);
      }

      const result = await response.json();
      
      if (!result.success) {
        throw new Error(result.error || "Failed to delete department");
      }

      setMessage('Department deleted successfully');
      setDeptToDelete('');
      onDepartmentAdded();
    } catch (error) {
      setMessage(`Error: ${error.message}`);
    } finally {
      setDeleting(false);
    }
  };

  return (
    <View style={styles.departmentManagement}>
      <View style={styles.modalHeader}>
        <Text style={styles.modalTitle}>Department Management</Text>
        <TouchableOpacity onPress={onClose}>
          <Text style={styles.closeButton}>✕</Text>
        </TouchableOpacity>
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Add New Department</Text>
        <TextInput
          style={styles.input}
          placeholder="Department Code (e.g., CIVIL)"
          value={newDeptCode}
          onChangeText={setNewDeptCode}
        />
        <TextInput
          style={styles.input}
          placeholder="Department Name (e.g., Civil Engineering)"
          value={newDeptName}
          onChangeText={setNewDeptName}
        />
        <Button
          title="Add Department"
          onPress={handleAddDepartment}
          disabled={saving}
        />
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Delete Department</Text>
        <Picker
          selectedValue={deptToDelete}
          style={styles.picker}
          onValueChange={(itemValue) => setDeptToDelete(itemValue)}
        >
          <Picker.Item label="Select department to delete" value="" />
          {departments
            .filter(dept => dept.code !== 'CSE')
            .map(dept => (
              <Picker.Item 
                key={dept.code} 
                label={`${dept.code} - ${dept.name}`} 
                value={dept.code} 
              />
          ))}
        </Picker>
        <Button
          title="Delete Department"
          onPress={handleDeleteDepartment}
          disabled={deleting || !deptToDelete}
          color="#f44336"
        />
      </View>

      {message ? <Text style={styles.message}>{message}</Text> : null}
    </View>
  );
};

const styles = StyleSheet.create({
 container: {
    padding: 20,
    backgroundColor: '#f8f9fa',
  },
  redemptionButton: {
    marginBottom: 25,
  },
  title: {
    fontSize: 20,
    fontWeight: '700',
    marginBottom: 30,
    textAlign: 'center',
    color: '#2c3e50',
    letterSpacing: 0.5,
  },
  configTitle: {
    fontSize: 20,
    fontWeight: '600',
    marginBottom: 25,
    textAlign: 'center',
    color: '#34495e',
    backgroundColor: '#ffffff',
    padding: 15,
    borderRadius: 12,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
 selectorContainer: {
    backgroundColor: '#ffffff',
    padding: 20,
    borderRadius: 15,
    marginBottom: 25,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 3 },
    shadowOpacity: 0.12,
    shadowRadius: 6,
    elevation: 4,
    borderWidth: 1,  // Add this for border
    borderColor: '#ddd',  // Change to a lighter color for better visibility
  },
  selector: {
    marginBottom: 20,
  },
  selectorLabel: {
    fontSize: 16,
    fontWeight: '600',
    marginBottom: 12,
    textAlign: 'center',
    color: 'black',
    textTransform: 'uppercase',
    
  },
  pickerContainer: {
    backgroundColor: '#f8f9fa',
    borderRadius: 10,
    borderWidth: 1,
    borderColor: '#e9ecef',
    overflow: 'hidden',
     marginBottom: 10, 
  },
  picker: {
    height: 50,
    backgroundColor: '#fff',
    borderRadius: 8,
    borderWidth: 1,
    borderColor: '#ddd',
    color: 'black',
  },
  manageButton: {
    padding: 12,
    backgroundColor: '#3498db',
    borderRadius: 8,
    marginLeft: 8,
    elevation: 2,
  },
  manageButtonText: {
    color: 'white',
    fontWeight: '600',
    fontSize: 14,
  },
  yearButtons: {
    flexDirection: 'row',
    justifyContent: 'center',
  },
  
  // Year button styles
  yearButton: {
    paddingHorizontal: 16,
    paddingVertical: 10,
    borderRadius: 20,
    backgroundColor: '#e9ecef',
    marginHorizontal: 4,
    minWidth: 70,
    alignItems: 'center',
  },
  yearButtonActive: {
    backgroundColor: '#6c5ce7',
  },
  yearButtonText: {
    fontSize: 14,
    fontWeight: '600',
    color: '#6c757d',
  },
  yearButtonTextActive: {
    color: '#ffffff',
  },
  label: {
    fontSize: 16,
    fontWeight: '600',
    marginTop: 15,
    marginBottom: 8,
    color: '#2c3e50',
  },
  input: {
    height: 45,
    borderColor: '#ddd',
    borderWidth: 1,
    marginBottom: 15,
    paddingHorizontal: 12,
    borderRadius: 8,
    backgroundColor: '#ffffff',
    fontSize: 16,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.05,
    shadowRadius: 2,
    elevation: 1,
  },
  readOnlyInput: {
    backgroundColor: '#f8f9fa',
    color: '#6c757d',
  },
  
  // New two-column layout styles
  twoColumnContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginBottom: 20,
  },
  columnLeft: {
    flex: 0.48,
  },
  columnRight: {
    flex: 0.48,
  },
  
  arrayContainer: {
    marginBottom: 20,
    backgroundColor: '#ffffff',
    padding: 15,
    borderRadius: 12,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.08,
    shadowRadius: 4,
    elevation: 2,
  },
  arrayInput: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 10,
  },
  arrayLabel: {
    width: 80,
    fontSize: 14,
    fontWeight: '500',
    color: '#495057',
  },
  
  // Single column for full-width inputs
  singleColumnContainer: {
    backgroundColor: '#ffffff',
    padding: 15,
    borderRadius: 12,
    marginBottom: 20,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.08,
    shadowRadius: 4,
    elevation: 2,
  },
  
  // Save button styling
  saveButtonContainer: {
    marginTop: 30,
    marginBottom: 20,
  },
  
  // Success/Error popup overlay
  popupOverlay: {
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    justifyContent: 'center',
    alignItems: 'center',
    zIndex: 1000,
  },
  popup: {
    backgroundColor: '#ffffff',
    padding: 25,
    borderRadius: 15,
    minWidth: 280,
    maxWidth: 320,
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 10 },
    shadowOpacity: 0.25,
    shadowRadius: 20,
    elevation: 10,
  },
  popupText: {
    fontSize: 16,
    textAlign: 'center',
    marginBottom: 20,
    color: '#2c3e50',
    lineHeight: 24,
  },
  popupButton: {
    backgroundColor: '#27ae60',
    paddingHorizontal: 25,
    paddingVertical: 12,
    borderRadius: 8,
    minWidth: 100,
  },
  popupButtonText: {
    color: '#ffffff',
    fontWeight: '600',
    textAlign: 'center',
    fontSize: 16,
  },
  
  message: {
    marginTop: 15,
    color: '#27ae60',
    textAlign: 'center',
    fontSize: 16,
    fontWeight: '500',
  },
  modalContainer: {
    flex: 1,
    padding: 20,
    backgroundColor: '#f8f9fa',
  },
  departmentManagement: {
    flex: 1,
  },
  modalHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 25,
    backgroundColor: '#ffffff',
    padding: 20,
    borderRadius: 12,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  modalTitle: {
    fontSize: 22,
    fontWeight: '700',
    color: '#2c3e50',
  },
  closeButton: {
    fontSize: 24,
    padding: 10,
    color: '#e74c3c',
    fontWeight: 'bold',
  },
  section: {
    marginBottom: 20,
    padding: 20,
    backgroundColor: '#ffffff',
    borderRadius: 12,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.08,
    shadowRadius: 4,
    elevation: 2,
  },
  sectionTitle: {
    fontSize: 18,
    fontWeight: '600',
    marginBottom: 15,
    color: '#2c3e50',
  },
  dateControl: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 15,
    padding: 15,
    backgroundColor: '#f8f9fa',
    borderRadius: 8,
    borderWidth: 1,
    borderColor: '#e9ecef',
  },
  switchControl: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 15,
    padding: 15,
    backgroundColor: '#f8f9fa',
    borderRadius: 8,
  },
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: 18,
    backgroundColor: '#6c5ce7',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 3 },
    shadowOpacity: 0.15,
    shadowRadius: 6,
    elevation: 6,
  },
  headerTitle: {
    color: 'white',
    fontSize: 20,
    fontWeight: '700',
    letterSpacing: 0.5,
  },
  menuButton: {
    padding: 12,
    borderRadius: 8,
    backgroundColor: 'rgba(255, 255, 255, 0.1)',
  },
  drawer: {
    position: 'absolute',
    top: 0,
    left: 0,
    width: 280,
    height: '100%',
    backgroundColor: '#ffffff',
    zIndex: 100,
    elevation: 8,
    shadowColor: '#000',
    shadowOffset: { width: 4, height: 0 },
    shadowOpacity: 0.25,
    shadowRadius: 8,
  },
  drawerItem: {
    padding: 20,
    borderBottomWidth: 1,
    borderBottomColor: '#f1f2f6',
    backgroundColor: '#ffffff',
  },
  drawerItemText: {
    fontSize: 16,
    color: '#2c3e50',
    fontWeight: '500',
  },
  overlay: {
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: 'rgba(0, 0, 0, 0.6)',
    zIndex: 99,
  },
});

export default AdminDashboard;