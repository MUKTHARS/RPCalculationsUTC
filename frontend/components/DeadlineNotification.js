// components/DeadlineNotification.js
import React, { useEffect, useState } from 'react';
import { Alert, Modal, View, Text, TouchableOpacity, StyleSheet } from 'react-native';

const DeadlineNotification = ({ studentYear, department }) => {
  const [notification, setNotification] = useState(null);
  const [visible, setVisible] = useState(false);
const API_URL = 'http://10.150.255.205:8080'; 
  useEffect(() => {
    const checkDeadline = async () => {
      try {
        if (!studentYear || !department) return;
        
        const response = await fetch(`${API_URL}/admin/api/redemption-dates?year=${studentYear}&department=${department}`);
        const data = await response.json();
        
        if (data.startDate && data.endDate) {
          let deadline = new Date(data.endDate);
          if (data.isExtended && data.extendedDate) {
            deadline = new Date(data.extendedDate);
          }
          
          const now = new Date();
          const timeDiff = deadline.getTime() - now.getTime();
          const daysDiff = Math.ceil(timeDiff / (1000 * 3600 * 24));
          
          let message = '';
          
          if (daysDiff === 60) {
            message = `Attention Year ${studentYear} students! Redemption deadline is in 2 months (${deadline.toDateString()}).`;
          } else if (daysDiff === 30) {
            message = `Year ${studentYear} students: 1 month left until redemption deadline (${deadline.toDateString()}).`;
          } else if (daysDiff === 14) {
            message = `Two weeks remaining for Year ${studentYear} redemption deadline (${deadline.toDateString()}).`;
          } else if (daysDiff === 3) {
            message = `Year ${studentYear} students: Only 3 days left to redeem your points!`;
          } else if (daysDiff === 1) {
            message = `Final reminder for Year ${studentYear} students: Redemption deadline is tomorrow!`;
          } else if (daysDiff === 0) {
            message = `Today is the last day for Year ${studentYear} students to redeem points!`;
          } else if (daysDiff < 0) {
            message = `The redemption period for Year ${studentYear} has ended. Contact admin for extension.`;
          }
          
          if (message) {
            setNotification(message);
            setVisible(true);
          }
        }
      } catch (error) {
        console.error('Error checking deadline:', error);
      }
    };

    // Check immediately when component mounts
    checkDeadline();
    
    // Then check every hour
    const interval = setInterval(checkDeadline, 60 * 60 * 1000);
    
    return () => clearInterval(interval);
  }, [studentYear, department]);

  const handleClose = () => {
    setVisible(false);
    setNotification(null);
  };

  if (!notification || !visible) return null;

  return (
    <Modal
      transparent={true}
      visible={visible}
      animationType="slide"
      onRequestClose={handleClose}
    >
      <View style={styles.modalContainer}>
        <View style={styles.modalContent}>
          <Text style={styles.modalTitle}>Deadline Reminder</Text>
          <Text style={styles.modalText}>{notification}</Text>
          <TouchableOpacity style={styles.modalButton} onPress={handleClose}>
            <Text style={styles.modalButtonText}>OK</Text>
          </TouchableOpacity>
        </View>
      </View>
    </Modal>
  );
};

const styles = StyleSheet.create({
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0,0,0,0.5)',
  },
  modalContent: {
    backgroundColor: 'white',
    padding: 20,
    borderRadius: 10,
    width: '80%',
  },
  modalTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    marginBottom: 10,
    color: '#1E3A8A',
  },
  modalText: {
    fontSize: 16,
    marginBottom: 20,
  },
  modalButton: {
    backgroundColor: '#1E3A8A',
    padding: 10,
    borderRadius: 5,
    alignSelf: 'flex-end',
  },
  modalButtonText: {
    color: 'white',
    fontWeight: 'bold',
  },
});

export default DeadlineNotification;





// testing notification
// // components/DeadlineNotification.js
// import React, { useEffect, useState } from 'react';
// import { Alert, Modal, View, Text, TouchableOpacity, StyleSheet } from 'react-native';

// const DeadlineNotification = ({ studentYear, department }) => {
//   const [notification, setNotification] = useState(null);
//   const [visible, setVisible] = useState(false);
//   const [isTesting, setIsTesting] = useState(true); // Set to true for testing

//   useEffect(() => {
//     const checkDeadline = async () => {
//       try {
//         if (!studentYear || !department) return;
        
//         if (isTesting) {
//           // TESTING MODE - Simulate different time differences
//           const testDaysDiff = 3; // Change this value to test different scenarios
//           // Possible test values: 60, 30, 14, 3, 1, 0, -1
          
//           let message = '';
          
//           if (testDaysDiff === 60) {
//             message = `[TEST] Year ${studentYear} students! Redemption deadline is in 2 months.`;
//           } else if (testDaysDiff === 30) {
//             message = `[TEST] Year ${studentYear} students: 1 month left until redemption deadline.`;
//           } else if (testDaysDiff === 14) {
//             message = `[TEST] Year ${studentYear} students: Two weeks remaining.`;
//           } else if (testDaysDiff === 3) {
//             message = `[TEST] Year ${studentYear} students: Only 3 days left!`;
//           } else if (testDaysDiff === 1) {
//             message = `[TEST] Year ${studentYear} students: Deadline is tomorrow!`;
//           } else if (testDaysDiff === 0) {
//             message = `[TEST] Year ${studentYear} students: Last day today!`;
//           } else if (testDaysDiff < 0) {
//             message = `[TEST] Year ${studentYear} students: Deadline has passed!`;
//           }
          
//           if (message) {
//             setNotification(message);
//             setVisible(true);
//           }
//           return;
//         }

//         // REAL IMPLEMENTATION (unchanged)
//         const response = await fetch(`${API_URL}/admin/api/redemption-dates?year=${studentYear}&department=${department}`);
//         const data = await response.json();
        
//         if (data.startDate && data.endDate) {
//           const deadline = new Date(data.endDate);
//           if (data.isExtended && data.extendedDate) {
//             deadline = new Date(data.extendedDate);
//           }
          
//           const now = new Date();
//           const timeDiff = deadline.getTime() - now.getTime();
//           const daysDiff = Math.ceil(timeDiff / (1000 * 3600 * 24));
          
//           let message = '';
          
//           if (daysDiff === 60) {
//             message = `Year ${studentYear} students! Redemption deadline is in 2 months (${deadline.toDateString()}).`;
//           } else if (daysDiff === 30) {
//             message = `Year ${studentYear} students: 1 month left until redemption deadline (${deadline.toDateString()}).`;
//           } else if (daysDiff === 14) {
//             message = `Year ${studentYear} students: Two weeks remaining (${deadline.toDateString()}).`;
//           } else if (daysDiff === 3) {
//             message = `Year ${studentYear} students: Only 3 days left!`;
//           } else if (daysDiff === 1) {
//             message = `Year ${studentYear} students: Deadline is tomorrow!`;
//           } else if (daysDiff === 0) {
//             message = `Year ${studentYear} students: Last day today!`;
//           } else if (daysDiff < 0) {
//             message = `Year ${studentYear} students: Deadline has passed!`;
//           }
          
//           if (message) {
//             setNotification(message);
//             setVisible(true);
//           }
//         }
//       } catch (error) {
//         console.error('Error checking deadline:', error);
//       }
//     };

//     // Check immediately when component mounts
//     checkDeadline();
    
//     if (!isTesting) {
//       // In real mode, check every hour
//       const interval = setInterval(checkDeadline, 60 * 60 * 1000);
//       return () => clearInterval(interval);
//     }
//   }, [studentYear, department, isTesting]);

//   const handleClose = () => {
//     setVisible(false);
//     setNotification(null);
//   };

//   if (!notification || !visible) return null;

//   return (
//     <Modal
//       transparent={true}
//       visible={visible}
//       animationType="slide"
//       onRequestClose={handleClose}
//     >
//       <View style={styles.modalContainer}>
//         <View style={styles.modalContent}>
//           <Text style={styles.modalTitle}>Deadline Reminder</Text>
//           <Text style={styles.modalText}>{notification}</Text>
//           <TouchableOpacity 
//             style={styles.modalButton} 
//             onPress={handleClose}
//           >
//             <Text style={styles.modalButtonText}>OK</Text>
//           </TouchableOpacity>
//           {isTesting && (
//             <TouchableOpacity 
//               style={[styles.modalButton, { backgroundColor: '#666', marginTop: 10 }]}
//               onPress={() => {
//                 handleClose();
//                 setTimeout(() => {
//                   setNotification(notification.replace(/\[TEST\] /, '[TEST] '));
//                   setVisible(true);
//                 }, 1000);
//               }}
//             >
//               <Text style={styles.modalButtonText}>Test Again</Text>
//             </TouchableOpacity>
//           )}
//         </View>
//       </View>
//     </Modal>
//   );
// };

// const styles = StyleSheet.create({
//   modalContainer: {
//     flex: 1,
//     justifyContent: 'center',
//     alignItems: 'center',
//     backgroundColor: 'rgba(0,0,0,0.5)',
//   },
//   modalContent: {
//     backgroundColor: 'white',
//     padding: 20,
//     borderRadius: 10,
//     width: '80%',
//   },
//   modalTitle: {
//     fontSize: 18,
//     fontWeight: 'bold',
//     marginBottom: 10,
//     color: '#1E3A8A',
//   },
//   modalText: {
//     fontSize: 16,
//     marginBottom: 20,
//   },
//   modalButton: {
//     backgroundColor: '#1E3A8A',
//     padding: 10,
//     borderRadius: 5,
//     alignSelf: 'flex-end',
//   },
//   modalButtonText: {
//     color: 'white',
//     fontWeight: 'bold',
//   },
// });

// export default DeadlineNotification;