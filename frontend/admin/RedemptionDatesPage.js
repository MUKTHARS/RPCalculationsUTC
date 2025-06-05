import React, { useState, useEffect } from 'react';
import { View, Text, Button, Alert, TouchableOpacity, StyleSheet, ActivityIndicator, Switch, ScrollView } from 'react-native';
import DateTimePicker from '@react-native-community/datetimepicker';
import { Picker } from '@react-native-picker/picker';
import { DeviceEventEmitter } from 'react-native';

const RedemptionDatesPage = ({ navigation }) => {
  const [department, setDepartments] = useState([]);
  const [selectedDepartment, setSelectedDepartment] = useState('CSE');
  const [selectedYear, setSelectedYear] = useState(1);
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [config, setConfig] = useState({
    startDate: new Date(),
    endDate: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000),
    isExtended: false,
    extendedDate: null,
    showTimePicker: false,
    currentEditing: null
  });
  const [showStartDatePicker, setShowStartDatePicker] = useState(false);
  const [showEndDatePicker, setShowEndDatePicker] = useState(false);
  const [showExtendedDatePicker, setShowExtendedDatePicker] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [loadingDepts, setLoadingDepts] = useState(true);

  // Fetch departments when component mounts
 useEffect(() => {
    const fetchDepartments = async () => {
      try {
        setLoadingDepts(true);
        let response = await fetch('http://10.150.255.205:8080/admin/api/departments');
        
        if (!response.ok) {
          response = await fetch('http://10.150.255.205:8080/admin/api/config?year=1');
          if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
          }
          
          const configData = await response.json();
          if (configData && configData.departments) {
            const sortedDepts = configData.departments.sort((a, b) => {
              if (a.code === 'CSE') return -1;
              if (b.code === 'CSE') return 1;
              return a.code.localeCompare(b.code);
            });
            setDepartments(sortedDepts);
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
        if (sortedDepts.length > 0 && !selectedDepartment) {
          setSelectedDepartment(sortedDepts[0].code);
        }
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
    
    fetchDepartments();
  }, []);

  // Load redemption dates when department or year changes
  useEffect(() => {
    if (selectedDepartment && selectedYear) {
      loadRedemptionDates();
    }
  }, [selectedDepartment, selectedYear]);

  const formatISTDateTime = (date) => {
    if (!date) return 'Select date';
    const displayDate = new Date(date);
    return displayDate.toLocaleDateString('en-IN', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
    }) + ' ' + displayDate.toLocaleTimeString('en-IN', {
      hour: '2-digit',
      minute: '2-digit',
      hour12: true
    });
  };

  const handleTimeChange = (event, selectedTime) => {
  if (selectedTime) {
    let newDate;
    if (config.currentEditing === 'start') {
      newDate = new Date(config.startDate);
    } else if (config.currentEditing === 'end') {
      newDate = new Date(config.endDate);
    } else {
      newDate = new Date(config.extendedDate);
    }
    
    newDate.setHours(selectedTime.getHours());
    newDate.setMinutes(selectedTime.getMinutes());
    
    if (config.currentEditing === 'start') {
      setConfig(prev => ({ ...prev, startDate: newDate }));
    } else if (config.currentEditing === 'end') {
      setConfig(prev => ({ ...prev, endDate: newDate }));
    } else {
      setConfig(prev => ({ ...prev, extendedDate: newDate }));
    }
  }
  setConfig(prev => ({ ...prev, showTimePicker: false, currentEditing: null }));
};
const loadRedemptionDates = async () => {
    try {
      setIsLoading(true);
      const response = await fetch(
        `http://10.150.255.205:8080/admin/api/redemption-dates?year=${selectedYear}&department=${selectedDepartment}`
      );
      
      if (!response.ok) {
        throw new Error('Failed to fetch redemption dates');
      }
      
      const data = await response.json();
      const configData = data.config || data;
      
      const parseDate = (dateString) => {
        if (!dateString) return new Date();
        const date = new Date(dateString);
        return new Date(date.getTime() + date.getTimezoneOffset() * 60 * 1000);
      };

      setConfig({
        startDate: parseDate(configData.startDate),
        endDate: parseDate(configData.endDate),
        isExtended: configData.isExtended || false,
        extendedDate: configData.extendedDate ? parseDate(configData.extendedDate) : null,
        showTimePicker: false,
        currentEditing: null
      });
    } catch (error) {
      console.error('Error loading redemption dates:', error);
      const now = new Date();
      setConfig({
        startDate: now,
        endDate: new Date(now.getTime() + 30 * 24 * 60 * 60 * 1000),
        isExtended: false,
        extendedDate: null,
        showTimePicker: false,
        currentEditing: null
      });
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    loadRedemptionDates();
  }, [selectedYear, department]);

const handleSaveDates = async () => {
  try {
    setIsLoading(true);
    
    // Validate dates
    if (config.startDate >= config.endDate) {
      throw new Error('Start date must be before end date');
    }

    if (config.isExtended && (!config.extendedDate || config.extendedDate <= config.endDate)) {
      throw new Error('Extended date must be after the original end date');
    }

     const convertToUTCDate = (date) => {
      return new Date(Date.UTC(
        date.getFullYear(),
        date.getMonth(),
        date.getDate(),
        date.getHours(),
        date.getMinutes(),
        date.getSeconds()
      ));
    };

    const response = await fetch(`http://10.150.255.205:8080/admin/api/redemption-dates`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        year: selectedYear,
        department: selectedDepartment,
        config: {
          startDate: convertToUTCDate(config.startDate).toISOString(),
          endDate: convertToUTCDate(config.endDate).toISOString(),
          isExtended: config.isExtended,
          extendedDate: config.isExtended ? convertToUTCDate(config.extendedDate).toISOString() : null
        }
      })
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || 'Failed to save redemption dates');
    }

    Alert.alert('Success', 'Redemption dates saved successfully!');
    loadRedemptionDates(); 

     DeviceEventEmitter.emit('redemptionDatesUpdated', {
      department,
      year: selectedYear
    });
  } catch (error) {
    console.error('Error saving redemption dates:', error);
    Alert.alert('Error', error.message || 'Failed to save redemption dates');
  } finally {
    setIsLoading(false);
  }
};
const convertToIST = (date) => {
  if (!date) return null;
  const offset = 5.5 * 60 * 60 * 1000; // IST is UTC+5:30
  return new Date(date.getTime() + offset);
};
const handleExtendDeadline = async () => {
  try {
    setIsLoading(true);
    
    if (!config.extendedDate || config.extendedDate <= config.endDate) {
      throw new Error('Extended date must be after the original end date');
    }

    const response = await fetch(`http://10.150.255.205:8080/admin/api/redemption-dates/extend`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        year: selectedYear,
        department,
        newDate: config.extendedDate.toISOString()
      })
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || 'Failed to extend deadline');
    }

    Alert.alert('Success', 'Deadline extended successfully!');
    loadRedemptionDates(); // Refresh the dates
  } catch (error) {
    console.error('Error extending deadline:', error);
    Alert.alert('Error', error.message || 'Failed to extend deadline');
  } finally {
    setIsLoading(false);
  }
};

  return (
    <View style={{ flex: 1 }}>
      {/* Header and year selector */}
      <View style={styles.header}>
  <TouchableOpacity 
    style={styles.menuButton} 
    onPress={() => setDrawerOpen(true)}
  >
    <Text style={{ color: 'white', fontSize: 24 }}>☰</Text>
  </TouchableOpacity>
  <Text style={styles.headerTitle}>Redemption Dates</Text>
  <View style={{ width: 48 }} />
</View>

{/* Add drawer code similar to AdminDashboard */}
{drawerOpen && (
  <>
    <View style={styles.overlay} onTouchEnd={() => setDrawerOpen(false)} />
    <View style={styles.drawer}>
      <TouchableOpacity 
        style={styles.drawerItem}
        onPress={() => {
          setDrawerOpen(false);
          navigation.navigate('AdminDashboard');
        }}
      >
        <Text style={styles.drawerItemText}>Dashboard</Text>
      </TouchableOpacity>
      <TouchableOpacity 
        style={styles.drawerItem}
        onPress={() => {
          setDrawerOpen(false);
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

      {/* Department and Year Selector - Updated to match AdminDashboard */}
      <View style={styles.selectorContainer}>
        <View style={styles.selector}>
          <Text style={styles.selectorLabel}>Department</Text>
          {loadingDepts ? (
            <ActivityIndicator style={styles.picker} />
          ) : (
            <Picker
              selectedValue={selectedDepartment}
              style={styles.picker}
              onValueChange={(itemValue) => setSelectedDepartment(itemValue)}
              dropdownIconColor="#6c5ce7"
            >
              {department.map(dept => (
                <Picker.Item 
                  key={dept.code} 
                  label={`${dept.code} - ${dept.name}`} 
                  value={dept.code} 
                />
              ))}
            </Picker>
          )}
        </View>
        
        <View style={styles.selector}>
          <Text style={styles.selectorLabel}>Year</Text>
          <View style={styles.yearButtons}>
            {[1, 2, 3, 4].map(year => (
              <TouchableOpacity
                key={year}
                style={[
                  styles.yearButton,
                  selectedYear === year && styles.yearButtonActive
                ]}
                onPress={() => setSelectedYear(year)}
              >
                <Text style={[
                  styles.yearButtonText,
                  selectedYear === year && styles.yearButtonTextActive
                ]}>
                  Year {year}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
        </View>
      </View>


      <ScrollView 
  contentContainerStyle={[styles.container, { flexGrow: 1 }]}
  style={{ flex: 1 }}
>
        <Text style={styles.subtitle}>{selectedDepartment} - Year {selectedYear}</Text>
        
        {/* Date controls */}
        <View style={styles.dateControl}>
          <Text style={styles.label}>Start Date:</Text>
          <TouchableOpacity 
  style={styles.dateButton}
  onPress={() => setShowStartDatePicker(true)}
>
  <Text>{formatISTDateTime(config.startDate)}</Text>
</TouchableOpacity>
          {showStartDatePicker && (
  <DateTimePicker
    value={config.startDate}
    mode="date"
    display="default"
    onChange={(event, date) => {
      setShowStartDatePicker(false);
      if (date) {
        setConfig(prev => ({ 
          ...prev, 
          startDate: date,
          // Ensure end date is after start date
          endDate: date >= prev.endDate ? new Date(date.getTime() + 24 * 60 * 60 * 1000) : prev.endDate,
         showTimePicker: true,
          currentEditing: 'start'
        }));
      }
    }}
  />
)}
{config.showTimePicker && config.currentEditing === 'start' && (
  <DateTimePicker
    value={config.startDate}
    mode="time"
    display="default"
    onChange={handleTimeChange}
  />
)}
        </View>

        <View style={styles.dateControl}>
          <Text style={styles.label}>End Date:</Text>
          <TouchableOpacity 
  style={styles.dateButton}
  onPress={() => setShowEndDatePicker(true)}
>
  <Text>{formatISTDateTime(config.endDate)}</Text>
</TouchableOpacity>
          {showEndDatePicker && (
  <DateTimePicker
    value={config.endDate}
    mode="date"
    display="default"
    minimumDate={new Date(config.startDate.getTime() + 24 * 60 * 60 * 1000)} // Next day after start date
    onChange={(event, date) => {
      setShowEndDatePicker(false);
      if (date) {
        setConfig(prev => ({ 
          ...prev, 
          endDate: date,
          // Reset extended date if it's now before the new end date
          extendedDate: prev.extendedDate && prev.extendedDate <= date ? null : prev.extendedDate,
          showTimePicker: true,
          currentEditing: 'end'
        }));
      }
    }}
  />
)}
{config.showTimePicker && config.currentEditing === 'end' && (
  <DateTimePicker
    value={config.endDate}
    mode="time"
    display="default"
    onChange={handleTimeChange}
  />
)}
        </View>

        <View style={styles.switchControl}>
          <Text style={styles.label}>Extended Deadline:</Text>
          <Switch
            value={config.isExtended}
            onValueChange={(value) => setConfig(prev => ({ 
              ...prev, 
              isExtended: value,
              extendedDate: value ? prev.endDate : null
            }))}
          />
        </View>

        {config.isExtended && (
          <View style={styles.dateControl}>
            <Text style={styles.label}>Extended Date:</Text>
           <TouchableOpacity 
  style={styles.dateButton}
  onPress={() => setShowExtendedDatePicker(true)}
>
  <Text>{formatISTDateTime(config.extendedDate)}</Text>
</TouchableOpacity>
            {showExtendedDatePicker && (
  <DateTimePicker
    value={config.extendedDate || new Date(config.endDate.getTime() + 24 * 60 * 60 * 1000)}
    mode="date"
    display="default"
    minimumDate={new Date(config.endDate.getTime() + 24 * 60 * 60 * 1000)} // Next day after end date
    onChange={(event, date) => {
      setShowExtendedDatePicker(false);
      if (date) {
        setConfig(prev => ({ ...prev, extendedDate: date,
          showTimePicker: true,
          currentEditing: 'extended'
         }));
      }
    }}
  />
)}
{config.showTimePicker && config.currentEditing === 'extended' && (
  <DateTimePicker
    value={config.extendedDate || new Date(config.endDate.getTime() + 24 * 60 * 60 * 1000)}
    mode="time"
    display="default"
    onChange={handleTimeChange}
  />
)}
          </View>
        )}

        <View style={styles.buttonContainer}>
          <Button
            title="Save Dates"
            onPress={handleSaveDates}
            disabled={isLoading}
          />
        </View>

        {config.isExtended && config.extendedDate && (
          <View style={styles.buttonContainer}>
            <Button
              title="Extend Deadline"
              onPress={handleExtendDeadline}
              disabled={isLoading}
              color="orange"
            />
          </View>
        )}
      </ScrollView>
    </View>
  );
};


const styles = StyleSheet.create({
  container: {
    
    padding: 20,
    backgroundColor: '#fff',
  },
  title: {
    fontSize: 22,
    fontWeight: 'bold',
    marginBottom: 10,
    textAlign: 'center',
  },
  subtitle: {
    fontSize: 18,
    marginBottom: 20,
    textAlign: 'center',
    color: '#555',
  },
  dateControl: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 15,
    padding: 15,
    backgroundColor: '#f8f9fa',
    borderRadius: 8,
  },
    label: {
    fontSize: 16,
    fontWeight: '600',
    marginTop: 15,
    marginBottom: 8,
    color: '#2c3e50',
  },
  dateButton: {
    padding: 12,
    backgroundColor: '#fff',
    borderWidth: 1,
    borderColor: '#ddd',
    borderRadius: 5,
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
  buttonContainer: {
    marginTop: 10,
    marginBottom: 15,
  },
  // New styles
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
  yearSelector: {
  flexDirection: 'row',
  alignItems: 'center',
  marginBottom: 15,
  padding: 15,
  backgroundColor: '#f8f9fa',
  borderRadius: 8,
},
yearPicker: {
  flex: 1,
  height: 50,
  color:'black',
},

  selector: {
    marginBottom: 20,
  },
 

  departmentBadge: {
  backgroundColor: '#e0f2fe',
  padding: 8,
  borderRadius: 20,
  alignSelf: 'center',
  marginBottom: 10,
  borderWidth: 1,
  borderColor: '#bae6fd',
},
departmentText: {
  color: '#0369a1',
  fontWeight: '600',
  fontSize: 14,
},
  
yearButtons: {
  flexDirection: 'row',
  justifyContent: 'center',
},
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
  picker: {
    height: 50,
    backgroundColor: '#fff',
    borderRadius: 8,
    borderWidth: 1,
    borderColor: '#ddd',
    color: 'black',  // Ensure text color is black for visibility
  },
  pickerContainer: {
    backgroundColor: '#f8f9fa',
    borderRadius: 10,
    borderWidth: 1,
    borderColor: '#e9ecef',
    overflow: 'hidden',
    marginBottom: 10,  // Add some spacing
  },
  selectorLabel: {
    fontSize: 16,
    fontWeight: '600',
    marginBottom: 12,
    textAlign: 'center',
    color: 'black',
    textTransform: 'uppercase',
    letterSpacing: 1,
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
export default RedemptionDatesPage;