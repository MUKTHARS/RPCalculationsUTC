import React, { useState, useEffect } from 'react';
import {
  SafeAreaView,
  Text,
  View,
  TextInput,
  TouchableOpacity,
  ScrollView,
  ActivityIndicator,
  Switch,
  Alert,
  
 StatusBar,
} from 'react-native';
import LoginScreen from './screens/LoginScreen';
import RedemptionDatesPage from './admin/RedemptionDatesPage';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { NavigationContainer } from '@react-navigation/native';
import PointsTable from './screens/PointsTable';
import AdminDashboard from './admin/admin_dashboard';
import DepartmentManagementPage from './admin/DepartmentManagementPage';
import styles from './styles/AppStyles';
import {Picker} from '@react-native-picker/picker';
const API_URL = 'http://10.0.2.2:8080'; 
const Stack = createNativeStackNavigator();

function MainAppScreen({ navigation }) {
  const [points, setPoints] = useState('');
  const [studentId, setStudentId] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState(null);
  const [error, setError] = useState('');
  const [useDatabase, setUseDatabase] = useState(true);
  const [subjectNames, setSubjectNames] = useState([]);
  const [editableAllocation, setEditableAllocation] = useState(null);
  const [isEditing, setIsEditing] = useState(false);
  const [studentSubjects, setStudentSubjects] = useState([]);
  const [departments, setDepartments] = useState([]);
  const [selectedDepartment, setSelectedDepartment] = useState(null);
  const [selectedYear, setSelectedYear] = useState(null);
  const [loadingDepts, setLoadingDepts] = useState(true);

  // Fetch departments from API
  const fetchDepartments = async () => {
    try {
      setLoadingDepts(true);
      const response = await fetch(`${API_URL}/admin/api/departments`);
      const data = await response.json();
      
      if (!response.ok || !data.success) {
        throw new Error(data.error || 'Failed to fetch departments');
      }
      
      const depts = data.departments || data;
      const sortedDepts = Array.isArray(depts) ? 
        depts.sort((a, b) => a.code.localeCompare(b.code)) : [
        { code: 'CSE', name: 'Computer Science' },
        { code: 'IT', name: 'Information Technology' },
        { code: 'ECE', name: 'Electronics and Communication' },
        { code: 'EEE', name: 'Electrical and Electronics' },
        { code: 'MECH', name: 'Mechanical' },
      ];
      
      setDepartments(sortedDepts);
      // if (sortedDepts.length > 0 && !selectedDepartment) {
      //   setSelectedDepartment(sortedDepts[0].code);
      // }
    } catch (err) {
      console.error('Error fetching departments:', err);
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

const fetchStudentSubjects = async (studentId) => {
  try {
    if (!studentId || studentId.length < 3) {
      // Clear values when no student ID or too short
      setStudentSubjects([]);
      setSelectedDepartment(null);
      setSelectedYear(null);
      return;
    }

    const response = await fetch(`${API_URL}/api/student-subjects/${studentId}`);
    const data = await response.json();
    
    if (!response.ok) {
      if (response.status === 404) {
        setStudentSubjects([]);
        setSelectedDepartment(null);
        setSelectedYear(null);
        return;
      }
      throw new Error(data.error || 'Failed to fetch student subjects');
    }
    
    if (!data.success) {
      throw new Error(data.error || 'Failed to fetch student subjects');
    }
    
    setStudentSubjects(data.subjects || []);

    // Get department and year from student_rewards endpoint
    const rewardsResponse = await fetch(`${API_URL}/api/student-rewards/${studentId}`);
    if (rewardsResponse.ok) {
      const rewardsData = await rewardsResponse.json();
     if (rewardsData.success) {
      setSelectedDepartment(rewardsData.department || null);
      setSelectedYear(rewardsData.year || null);
    }
    }
  } catch (err) {
    console.error('Error fetching student subjects:', err);
    setStudentSubjects([]);
    setSelectedDepartment(null);
    setSelectedYear(null);
  }
};

  const fetchSubjectNames = async () => {
    try {
      const response = await fetch(`${API_URL}/api/subjects`);
      const data = await response.json();
      
      if (!response.ok || !data.success) {
        throw new Error(data.error || 'Failed to fetch subject names');
      }
      
      setSubjectNames(data.subjects);
    } catch (err) {
      console.error('Error fetching subject names:', err);
      setSubjectNames([
        'Mathematics',
        'Physics',
        'Chemistry',
        'Biology',
        'History',
        'Geography',
      ]);
    }
  };

  const calculateAllocation = async () => {
    if (!studentId || !selectedDepartment || !selectedYear) {
    setError('Please enter a valid student ID');
    return;
  }
    setLoading(true);
    setError('');
    
    try {
      if (useDatabase) {
        if (!studentId.trim()) {
          setError('Please enter a student ID');
          setLoading(false);
          return;
        }
        
        await fetchStudentSubjects(studentId);
        
        const checkResponse = await fetch(`${API_URL}/api/student-allocation/${studentId}`);
        
        if (checkResponse.ok) {
          const existingData = await checkResponse.json();
          
          if (existingData.success && existingData.data) {
            setResult(existingData.data);
            setEditableAllocation(JSON.parse(JSON.stringify(existingData.data)));
            setLoading(false);
            return;
          }
        }
        
        const response = await fetch(`${API_URL}/api/calculate-db`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ 
            studentId: studentId,
            department: selectedDepartment,
            year: selectedYear
          }),
        });
        
        const data = await response.json();
        
        if (!response.ok || !data.success) {
          throw new Error(data.error || 'Failed to retrieve points from database');
        }
        
        setResult(data.data);
        setEditableAllocation(JSON.parse(JSON.stringify(data.data)));
      } else {
        // Manual points calculation (unchanged)
      }
    } catch (err) {
      setError(err.message || 'An error occurred during calculation');
    } finally {
      setLoading(false);
    }
  };

//  useEffect(() => {
//     if (route.params?.studentId) {
//       setStudentId(route.params.studentId);
//       fetchStudentSubjects(route.params.studentId);
//     }
//     fetchDepartments();
//     fetchSubjectNames();
//   }, [route.params?.studentId]);

  useEffect(() => {

    fetchDepartments();
    fetchSubjectNames();
  }, []);

 return (
  <SafeAreaView style={styles.container}>
    <StatusBar barStyle="light-content" backgroundColor="#1E3A8A" />
    
    {/* Professional Header */}
    <View style={styles.header}>
      <View style={styles.headerContent}>
        <Text style={styles.headerTitle}>Reward Points System</Text>
        <Text style={styles.headerSubtitle}>Academic Performance Tracker</Text>
      </View>
      <TouchableOpacity 
        style={styles.adminButton}
        onPress={() => navigation.navigate('AdminDashboard')}
        activeOpacity={0.8}
      >
        <Text style={styles.adminButtonText}>Admin</Text>
      </TouchableOpacity>
    </View>

    {/* Make the entire screen scrollable */}
    <ScrollView 
      style={styles.mainScrollView}
      showsVerticalScrollIndicator={false}
      contentContainerStyle={styles.scrollContent}
    >
      {/* Input Section */}
      <View style={styles.inputSection}>
        <Text style={styles.sectionTitle}>Student Information</Text>
        
        {useDatabase && (
          <View style={styles.formContainer}>
            <View style={styles.inputGroup}>
              <Text style={styles.inputLabel}>Student ID</Text>
              <TextInput
                style={styles.textInput}
                placeholder="Enter student ID"
                value={studentId}
                onChangeText={(text) => {
                  setStudentId(text);
                  if (text.length > 3) {
                    fetchStudentSubjects(text);
                  }
                }}
                placeholderTextColor="#9CA3AF"
              />
            </View>
            
            <View style={styles.infoDisplay}>
              <View style={styles.infoItem}>
                <Text style={styles.infoLabel}>Department:</Text>
                <Text style={styles.infoValue}>
                  {selectedDepartment === null ? '' : selectedDepartment}
                </Text>
              </View>
              <View style={styles.infoItem}>
                <Text style={styles.infoLabel}>Academic Year:</Text>
                <Text style={styles.infoValue}>
                  {selectedYear === null ? '' : `Year ${selectedYear}`}
                </Text>
              </View>
            </View>
          </View>
        )}
        
        <TouchableOpacity
          style={[styles.calculateButton, loading && styles.calculateButtonDisabled]}
          onPress={calculateAllocation}
          disabled={loading}
          activeOpacity={0.8}
        >
          {loading ? (
            <View style={styles.loadingButtonContent}>
              <ActivityIndicator size="small" color="#FFFFFF" />
              <Text style={styles.calculateButtonText}>Processing...</Text>
            </View>
          ) : (
            <Text style={styles.calculateButtonText}>Calculate Allocation</Text>
          )}
        </TouchableOpacity>
      </View>
      
      {/* Error Display */}
      {error ? (
        <View style={styles.errorContainer}>
          <View style={styles.errorIcon}>
            <Text style={styles.errorIconText}>!</Text>
          </View>
          <Text style={styles.errorText}>{error}</Text>
        </View>
      ) : null}
      
      {/* Results Section */}
      {result ? (
        <View style={styles.resultsSection}>
          <Text style={styles.resultsSectionTitle}>Allocation Results</Text>
          <View style={styles.resultsContainer}>
            <PointsTable 
              result={result}
              subjectNames={subjectNames} 
              studentSubjects={studentSubjects} 
              studentId={studentId}
              department={selectedDepartment}
              year={selectedYear}
              onPointsUpdate={(updatedResult) => {
                setEditableAllocation(updatedResult);
              }}
            />
          </View>
        </View>
      ) : (
        <View style={styles.emptyStateContainer}>
          <View style={styles.emptyStateIcon}>
            <Text style={styles.emptyStateIconText}>📊</Text>
          </View>
          <Text style={styles.emptyStateTitle}>Ready to Calculate</Text>
          <Text style={styles.emptyStateDescription}>
            Enter student details above and tap "Calculate Allocation" to view the reward points distribution.
          </Text>
        </View>
      )}
    </ScrollView>
  </SafeAreaView>
);
}


// export default function App() {
//   return (
//     <NavigationContainer>
//       <Stack.Navigator initialRouteName="Login">
//         <Stack.Screen 
//           name="Login" 
//           component={LoginScreen} 
//           options={{ headerShown: false }}
//         />
//         <Stack.Screen 
//           name="MainApp" 
//           component={MainAppScreen} 
//           options={{ headerShown: false }}
//         />
//         <Stack.Screen 
//           name="AdminDashboard" 
//           component={AdminDashboard} 
//           options={{ title: 'Admin Dashboard' }}
//         />
//         <Stack.Screen 
//           name="RedemptionDates" 
//           component={RedemptionDatesPage} 
//           options={{ title: 'Redemption Dates' }}
//         />
//         <Stack.Screen 
//           name="DepartmentManagement" 
//           component={DepartmentManagementPage} 
//         />
//       </Stack.Navigator>
//     </NavigationContainer>
//   );
// }



export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen 
          name="MainApp" 
          component={MainAppScreen} 
          options={{ headerShown: false }}
        />
        <Stack.Screen 
          name="AdminDashboard" 
          component={AdminDashboard} 
          options={{ title: 'Admin Dashboard' }}
        />
        <Stack.Screen 
  name="RedemptionDates" 
  component={RedemptionDatesPage} 
  options={{ title: 'Redemption Dates' }}
/>
<Stack.Screen 
  name="DepartmentManagement" 
  component={DepartmentManagementPage} 
/>


      </Stack.Navigator>
    </NavigationContainer>
  );
}