import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, Button, StyleSheet, TouchableOpacity, ActivityIndicator, ScrollView } from 'react-native';
import { Picker } from '@react-native-picker/picker';

const DepartmentManagementPage = ({ navigation }) => {
  const [departments, setDepartments] = useState([]);
  const [newDeptCode, setNewDeptCode] = useState('');
  const [newDeptName, setNewDeptName] = useState('');
  const [message, setMessage] = useState('');
  const [saving, setSaving] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [deptToDelete, setDeptToDelete] = useState('');
  const [loading, setLoading] = useState(true);

  const fetchDepartments = async () => {
    try {
      setLoading(true);
      let response = await fetch('http://10.0.2.2:8080/admin/api/departments');
      
      if (!response.ok) {
        response = await fetch('http://10.0.2.2:8080/admin/api/config?year=1');
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
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDepartments();
  }, []);

  const handleAddDepartment = async () => {
    if (!newDeptCode || !newDeptName) {
      setMessage('Both code and name are required');
      return;
    }

    setSaving(true);
    try {
      const response = await fetch('http://10.0.2.2:8080/admin/api/departments/add', {
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
      fetchDepartments();
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
      const response = await fetch('http://10.0.2.2:8080/admin/api/departments/delete', {
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
      fetchDepartments();
    } catch (error) {
      setMessage(`Error: ${error.message}`);
    } finally {
      setDeleting(false);
    }
  };

  return (
    <View style={{ flex: 1 }}>
      {/* Header */}
      <View style={styles.header}>
        <TouchableOpacity 
          style={styles.menuButton} 
          onPress={() => navigation.goBack()}
        >
          <Text style={{ color: 'white', fontSize: 24 }}>←</Text>
        </TouchableOpacity>
        <Text style={styles.headerTitle}>Department Management</Text>
        <View style={{ width: 48 }} />
      </View>

      <ScrollView contentContainerStyle={styles.container}>
        {loading ? (
          <ActivityIndicator size="large" style={styles.loader} />
        ) : (
          <>
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
          </>
        )}
      </ScrollView>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flexGrow: 1,
    padding: 20,
  },
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: 15,
    backgroundColor: '#6200ee',
  },
  headerTitle: {
    color: 'white',
    fontSize: 20,
    fontWeight: 'bold',
  },
  menuButton: {
    padding: 10,
  },
  section: {
    marginBottom: 20,
    padding: 15,
    backgroundColor: '#f8f9fa',
    borderRadius: 8,
  },
  sectionTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    marginBottom: 10,
  },
  input: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    marginBottom: 10,
    paddingHorizontal: 10,
    borderRadius: 5,
    backgroundColor: 'white',
  },
  picker: {
    height: 50,
    width: '100%',
    marginBottom: 15,
    backgroundColor: 'white',
  },
  message: {
    marginTop: 10,
    color: 'green',
    textAlign: 'center',
  },
  loader: {
    marginTop: 50,
  },
});

export default DepartmentManagementPage;