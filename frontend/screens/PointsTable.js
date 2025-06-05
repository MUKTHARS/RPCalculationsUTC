import React, { useState, useEffect } from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  Alert,
  ActivityIndicator
} from 'react-native';
import styles from '../styles/PointsTableStyles';
import DeadlineNotification from '../components/DeadlineNotification';

const PointsTable = ({ result, subjectNames, studentId, onPointsUpdate, studentSubjects, department, year }) => {
  const [editingSubject, setEditingSubject] = useState(null);
  const [editedPoints, setEditedPoints] = useState('');
  const [isSaving, setIsSaving] = useState(false);
  const [localResult, setLocalResult] = useState(result);
  const [filteredSubjects, setFilteredSubjects] = useState([]);
  const [studentDepartment, setStudentDepartment] = useState(department);
  const [academicYear, setAcademicYear] = useState(year);
const [canEdit, setCanEdit] = useState(true);
  // Update local result when prop changes
const API_URL = 'http://10.0.2.2:8080'; 
useEffect(() => {
  const checkRedemptionDeadline = async () => {
  try {
    if (!department || !year) return;
    
    const response = await fetch(`${API_URL}/admin/api/redemption-dates?year=${year}&department=${department}`);
    const data = await response.json();
    
    if (data.endDate) {
      let deadline = new Date(data.endDate);
      if (data.isExtended && data.extendedDate) {
        deadline = new Date(data.extendedDate);
      }
      
      // Convert UTC time from database to local time
      const localDeadline = new Date(
        deadline.getUTCFullYear(),
        deadline.getUTCMonth(),
        deadline.getUTCDate(),
        deadline.getUTCHours(),
        deadline.getUTCMinutes(),
        deadline.getUTCSeconds()
      );
      
      // Get current local time
      const now = new Date();
      
      // Compare local times
      setCanEdit(now <= localDeadline);
    }
  } catch (error) {
    console.error('Error checking redemption deadline:', error);
    setCanEdit(true); // Default to true if there's an error
  }
};
  
  checkRedemptionDeadline();
}, [department, year]);

  useEffect(() => {
    setLocalResult(result);
    setStudentDepartment(department);
    setAcademicYear(year);
  }, [result,department, year]);
// In PointsTable.js, update the useEffect that maps student subjects:
useEffect(() => {
    if (studentId && studentSubjects && studentSubjects.length > 0) {
      const mappedSubjects = studentSubjects.map((subject, index) => {
        // Try to find by code first, then fall back to index
        let resultSubject = result?.subjects?.find(s => 
          s.subjectCode === subject.subjectCode
        );
        
        if (!resultSubject && result?.subjects?.[index]) {
          resultSubject = result.subjects[index];
        }

        const totalPoints = resultSubject?.sectionPoints?.reduce((sum, p) => sum + p, 0) || 0;
        const totalMarks = resultSubject?.totalSubjectMarks || 0;
        
        return {
          subjectCode: subject.subjectCode,
          subjectName: subject.subjectName,
          points: totalPoints,
          marks: totalMarks,
          semester: subject.semester,
          academicYear: academicYear, // Use the passed academic year
          allocation: {
            sectionPoints: resultSubject?.sectionPoints || [0, 0, 0, 0],
            totalSubjectMarks: totalMarks
          }
        };
      });

      setFilteredSubjects(mappedSubjects);
    } else {
      setFilteredSubjects([]);
    }
  }, [studentSubjects, result, academicYear]);
  // Map student subjects to include points and marks
  useEffect(() => {
    if (studentSubjects && studentSubjects.length > 0) {
      const mappedSubjects = studentSubjects.map((subject, index) => {
        // Try to find by code first, then fall back to index
        let resultSubject = result?.subjects?.find(s => 
          s.subjectCode === subject.subjectCode
        );
        
        if (!resultSubject && result?.subjects?.[index]) {
          // If not found by code, match by index
          resultSubject = result.subjects[index];
        }

        const totalPoints = resultSubject?.sectionPoints?.reduce((sum, p) => sum + p, 0) || 0;
        const totalMarks = resultSubject?.totalSubjectMarks || 0;
        
        console.log(`Mapping ${subject.subjectCode}:`, {
          points: totalPoints,
          marks: totalMarks
        });

        return {
          subjectCode: subject.subjectCode,
          subjectName: subject.subjectName,
          points: totalPoints,
          marks: totalMarks,
          allocation: {
            sectionPoints: resultSubject?.sectionPoints || [0, 0, 0, 0],
            totalSubjectMarks: totalMarks
          }
        };
      });

      console.log("Mapped subjects:", mappedSubjects);
      setFilteredSubjects(mappedSubjects);
    } else {
      setFilteredSubjects([]);
    }
  }, [studentSubjects, result]);

 const handleEditPoints = (subjectCode, currentPoints) => {
    if (!canEdit) {
      Alert.alert("Redemption Closed", "The redemption period has ended. Contact admin for extension.");
      return;
    }
    setEditingSubject(subjectCode);
    setEditedPoints(currentPoints.toString());
  };

  const getSubjectTotalPoints = (subject) => {
    if (!subject || !subject.allocation || !subject.allocation.sectionPoints) return 0;
    return subject.allocation.sectionPoints.reduce((sum, points) => sum + (points || 0), 0);
  };

  const handleSavePoints = async (subjectCode, subjectName, subjectIndex) => {
    if (!studentId) {
      Alert.alert("Error", "Student ID is required");
      return;
    }

      const currentSubject = filteredSubjects[subjectIndex];
    const semester = currentSubject?.semester || 1;
    const academicYear = currentSubject?.academicYear || 1; // Use academic year (1-4) instead of calendar year

    const points = parseInt(editedPoints, 10);

    // Client-side validation
    if (isNaN(points)) {
      Alert.alert("Error", "Please enter a valid number");
      return;
    }

    if (points < 0) {
      Alert.alert("Error", "Points cannot be negative");
      return;
    }

    if (points > 250) {
      Alert.alert("Error", "Points cannot exceed 250 per subject");
      return;
    }

    // Calculate the difference between new and old points
    const oldPoints = getSubjectTotalPoints(filteredSubjects[subjectIndex]);
    const pointsDifference = points - oldPoints;

    // Check if we have enough remaining points
    const currentRemaining = localResult.pointsRemaining || 0;
    if (pointsDifference > currentRemaining) {
      Alert.alert("Error", "Not enough remaining points available");
      return;
    }

    setIsSaving(true);
    try {
      // First, get the marks calculation from the backend
      const calcResponse = await fetch('http://10.0.2.2:8080/api/calculate-subject', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          points: points
        }),
      });
      
      if (!calcResponse.ok) {
        const errorText = await calcResponse.text();
        throw new Error(`Server returned ${calcResponse.status}: ${errorText}`);
      }
      
      const calcData = await calcResponse.json();
      
      if (!calcData.success) {
        throw new Error(calcData.error || 'Failed to calculate marks');
      }
      
      // Get the marks and section points from the backend calculation
      const totalSubjectMarks = calcData.marks;
      const sectionPoints = calcData.sectionPoints || [points, 0, 0, 0];
      
      // Update local state
      const updatedResult = JSON.parse(JSON.stringify(localResult));
      
      // Find or create the subject in the result
      let resultSubjectIndex = updatedResult.subjects.findIndex(
        s => s.subjectCode === subjectCode
      );
      
      if (resultSubjectIndex >= 0) {
        // Update existing subject
        updatedResult.subjects[resultSubjectIndex].sectionPoints = sectionPoints;
        updatedResult.subjects[resultSubjectIndex].totalSubjectMarks = totalSubjectMarks;
      } else {
        // Add new subject
        resultSubjectIndex = updatedResult.subjects.length;
        updatedResult.subjects.push({
          subjectCode,
          subjectName,
          sectionPoints,
          totalSubjectMarks
        });
      }
      
      // Recalculate totals
      updatedResult.totalMarks = updatedResult.subjects.reduce(
        (sum, subject) => sum + subject.totalSubjectMarks, 0);
      
      updatedResult.pointsRemaining = currentRemaining - pointsDifference;
      
      setLocalResult(updatedResult);
      if (onPointsUpdate) onPointsUpdate(updatedResult);

      // Update the filtered subjects display
      const updatedFilteredSubjects = [...filteredSubjects];
      updatedFilteredSubjects[subjectIndex] = {
        ...updatedFilteredSubjects[subjectIndex],
        points: points,
        marks: totalSubjectMarks,
        allocation: {
          sectionPoints,
          totalSubjectMarks
        }
      };
      setFilteredSubjects(updatedFilteredSubjects);

      // Save to backend with all required data
      const saveResponse = await fetch('http://10.0.2.2:8080/api/save-subject-points', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          studentId: studentId,
          subjectCode: subjectCode,
          subjectName: subjectName,
          points: points,
          marks: totalSubjectMarks,
          sectionPoints: sectionPoints,
          semester: semester,
          year: academicYear
        }),
      });

      if (!saveResponse.ok) {
        const errorText = await saveResponse.text();
        throw new Error(`Failed to save subject points: ${errorText}`);
      }

      const saveData = await saveResponse.json();
      if (!saveData.success) {
        throw new Error(saveData.error || 'Failed to save subject points');
      }

      setEditingSubject(null);
    } catch (error) {
      console.error("Error updating points:", error);
      Alert.alert("Error", error.message || "Failed to update points. Please try again.");
      setLocalResult(result);
    } finally {
      setIsSaving(false);
    }
  };

  const renderSectionPoints = (subject) => {
    if (!subject || !subject.allocation || !subject.allocation.sectionPoints) return null;
    
    return (
      <View style={styles.sectionPointsContainer}>
        {subject.allocation.sectionPoints.map((points, idx) => (
          points > 0 ? (
            <Text key={`section-${idx}`} style={styles.sectionPointsText}>
              Sec {idx+1}: {points}
            </Text>
          ) : null
        ))}
      </View>
    );
  };

  if (!localResult) {
    return <Text>No data available</Text>;
  }

  return (
    <View style={styles.tableContainer}>
      {studentDepartment && academicYear && (
        <View style={styles.departmentBadge}>
          <Text style={styles.departmentText}>
            {studentDepartment} - Year {academicYear}
          </Text>
        </View>
      )}
      <View style={styles.tableHeader}>
        <Text style={[styles.tableHeaderCell, { flex: 1 }]}>Code</Text>
        <Text style={[styles.tableHeaderCell, { flex: 3 }]}>Subject Name</Text>
        <Text style={[styles.tableHeaderCell, { flex: 1 }]}>Points</Text>
        <Text style={[styles.tableHeaderCell, { flex: 1 }]}>Marks</Text>
        <Text style={[styles.tableHeaderCell, { flex: 1 }]}>Action</Text>
      </View>

      {filteredSubjects.map((subject, index) => (
        <View key={`${subject.subjectCode}-${index}`} style={styles.tableRow}>
    <Text style={[styles.tableCell, { flex: 1 }]}>{subject.subjectCode}</Text>
    <Text style={[styles.tableCell, { flex: 2 }]}>{subject.subjectName}</Text>
    <Text style={[styles.tableCell, { flex: 1 }]}>Year {subject.academicYear}</Text>
    <Text style={[styles.tableCell, { flex: 1 }]}>Sem {subject.semester}</Text>
          {editingSubject === subject.subjectCode ? (
            <TextInput
              style={[styles.tableCell, styles.input]}
              value={editedPoints}
              onChangeText={setEditedPoints}
              keyboardType="numeric"
              autoFocus
            />
          ) : (
            <Text style={[styles.tableCell, { flex: 1 }]}>{subject.points}</Text>
          )}
          
          <Text style={[styles.tableCell, { flex: 1 }]}>
            {subject.marks.toFixed(1)}
          </Text>
          
          <View style={[styles.tableCell, { flex: 1 }]}>
            {editingSubject === subject.subjectCode ? (
              isSaving ? (
                <ActivityIndicator size="small" />
              ) : (
            <TouchableOpacity
                  onPress={() => handleSavePoints(subject.subjectCode, subject.subjectName, index)}
                  style={styles.saveButton}
                >
                  <Text style={styles.saveButtonText}>Save</Text>
                </TouchableOpacity>
              )
            ) : (
              <TouchableOpacity
  onPress={() => handleEditPoints(subject.subjectCode, subject.points)}
  style={[styles.editButton, !canEdit && styles.disabledButton]}
  disabled={!canEdit}
>
  <Text style={[styles.editButtonText, !canEdit && styles.disabledButtonText]}>
    Edit
  </Text>
</TouchableOpacity>
            )}
          </View>
        </View>
      ))}

      {/* Totals row */}
      <View style={styles.tableTotalRow}>
        <Text style={[styles.tableTotalCell, { flex: 4 }]}>Total</Text>
        <Text style={[styles.tableTotalCell, { flex: 1 }]}>
          {localResult.totalPoints - (localResult.pointsRemaining || 0)}
        </Text>
        <Text style={[styles.tableTotalCell, { flex: 1 }]}>
          {localResult.totalMarks.toFixed(1)}
        </Text>
        <Text style={[styles.tableTotalCell, { flex: 1 }]}></Text>
      </View>
      
      {/* Points remaining row */}
      <View style={styles.tableTotalRow}>
        <Text style={[styles.tableTotalCell, { flex: 4 }]}>Points Remaining</Text>
        <Text style={[styles.tableTotalCell, { flex: 2 }]}>
          {localResult.pointsRemaining || 0}
        </Text>
      </View>

      <DeadlineNotification 
      studentYear={academicYear} 
      department={studentDepartment} 
    />
    </View>
  );
};

export default PointsTable;