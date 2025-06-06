import React, { useState } from 'react';
import { 
  View, 
  Text, 
  TextInput, 
  TouchableOpacity, 
  Alert, 
  StyleSheet,
  ActivityIndicator,
  ScrollView,
  StatusBar
} from 'react-native';

const UnfreezeStudentPage = ({ navigation }) => {
  const [studentId, setStudentId] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [drawerOpen, setDrawerOpen] = useState(false);

  const handleUnfreeze = async (unfreeze) => {
    if (!studentId) {
      Alert.alert('Error', 'Please enter a student ID');
      return;
    }

    setIsLoading(true);
    try {
      const response = await fetch('http://10.0.2.2:8080/admin/api/unfreeze-student', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          studentId: studentId,
          unfreeze: unfreeze
        }),
      });

      const data = await response.json();
      
      if (!response.ok || !data.success) {
        throw new Error(data.error || 'Failed to update student status');
      }

      Alert.alert('Success', `Student ${unfreeze ? 'unfrozen' : 'frozen'} successfully`);
      setStudentId('');
    } catch (error) {
      Alert.alert('Error', error.message || 'An error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <View style={{ flex: 1 }}>
      {/* Header with hamburger icon */}
      <View style={styles.header}>
        <TouchableOpacity 
          style={styles.menuButton} 
          onPress={() => setDrawerOpen(true)}
        >
          <Text style={{ color: 'white', fontSize: 24 }}>☰</Text>
        </TouchableOpacity>
        <Text style={styles.headerTitle}>Unfreeze Student</Text>
        <View style={{ width: 48 }} /> {/* Spacer for alignment */}
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
                // navigation.navigate('AdminDashboard');
              }}
            >
              <Text style={styles.drawerItemText}>Dashboard</Text>
            </TouchableOpacity>
            <TouchableOpacity 
              style={styles.drawerItem}
              onPress={() => {
                setDrawerOpen(false);
                navigation.navigate('RedemptionDates');
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
                            navigation.navigate('UnfreezeStudent');
                          }}
            >
              <Text style={styles.drawerItemText}>Unfreeze Student</Text>
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

      <ScrollView contentContainerStyle={styles.container}>
        <Text style={styles.subtitle}>Allow/Disallow a student to edit marks after deadline</Text>
        
        <View style={styles.inputContainer}>
          <Text style={styles.label}>Student ID</Text>
          <TextInput
            style={styles.input}
            placeholder="Enter Student ID"
            value={studentId}
            onChangeText={setStudentId}
          />
        </View>
        
        <View style={styles.buttonRow}>
          <TouchableOpacity
            style={[styles.button, styles.unfreezeButton]}
            onPress={() => handleUnfreeze(true)}
            disabled={isLoading}
          >
            {isLoading ? (
              <ActivityIndicator color="#fff" />
            ) : (
              <Text style={styles.buttonText}>Unfreeze Student</Text>
            )}
          </TouchableOpacity>
          
          <TouchableOpacity
            style={[styles.button, styles.freezeButton]}
            onPress={() => handleUnfreeze(false)}
            disabled={isLoading}
          >
            {isLoading ? (
              <ActivityIndicator color="#fff" />
            ) : (
              <Text style={styles.buttonText}>Freeze Student</Text>
            )}
          </TouchableOpacity>
        </View>
      </ScrollView>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flexGrow: 1,
    padding: 20,
    backgroundColor: '#f8f9fa',
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
  activeDrawerItem: {
    backgroundColor: '#f0f0f0',
  },
  drawerItemText: {
    fontSize: 16,
    color: '#2c3e50',
    fontWeight: '500',
  },
  activeDrawerItemText: {
    color: '#6c5ce7',
    fontWeight: '600',
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
  subtitle: {
    fontSize: 16,
    marginBottom: 30,
    color: '#6c757d',
    textAlign: 'center',
  },
  inputContainer: {
    marginBottom: 25,
    backgroundColor: '#ffffff',
    padding: 20,
    borderRadius: 12,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.08,
    shadowRadius: 4,
    elevation: 2,
  },
  label: {
    fontSize: 16,
    fontWeight: '600',
    marginBottom: 10,
    color: '#2c3e50',
  },
  input: {
    height: 50,
    borderColor: '#e9ecef',
    borderWidth: 1,
    borderRadius: 8,
    paddingHorizontal: 15,
    backgroundColor: '#ffffff',
    fontSize: 16,
  },
  buttonRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginTop: 20,
  },
  button: {
    flex: 0.48,
    height: 50,
    borderRadius: 8,
    justifyContent: 'center',
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 2,
  },
  unfreezeButton: {
    backgroundColor: '#27ae60',
  },
  freezeButton: {
    backgroundColor: '#e74c3c',
  },
  buttonText: {
    color: '#ffffff',
    fontWeight: '600',
    fontSize: 16,
  },
});

export default UnfreezeStudentPage;