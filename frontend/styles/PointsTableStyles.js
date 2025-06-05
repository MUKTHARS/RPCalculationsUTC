import { StyleSheet, Dimensions } from 'react-native';

const { width } = Dimensions.get('window');

export default StyleSheet.create({
  tableContainer: {
  backgroundColor: '#ffffff',
  borderRadius: 12,
  shadowColor: '#000',
  shadowOffset: {
    width: 0,
    height: 2,
  },
  shadowOpacity: 0.08,
  shadowRadius: 8,
  elevation: 4,
  margin: 0, 
  padding: 0, 
  
},

  departmentBadge: {
    backgroundColor: '#f8f9fa',
    paddingHorizontal: 16,
    paddingVertical: 12,
    borderTopLeftRadius: 12,
    borderTopRightRadius: 12,
    borderBottomWidth: 1,
    borderBottomColor: '#e9ecef',
  },

  departmentText: {
    fontSize: 14,
    fontWeight: '600',
    color: '#495057',
    textAlign: 'center',
  },

  tableHeader: {
    flexDirection: 'row',
    backgroundColor: '#6c757d',
    paddingVertical: 16,
    paddingHorizontal: 12,
    borderBottomWidth: 2,
    borderBottomColor: '#495057',
  },

  tableHeaderCell: {
    fontSize: 12,
    fontWeight: '700',
    color: '#ffffff',
    textAlign: 'center',
    textTransform: 'uppercase',
    letterSpacing: 0.5,
  },

  tableRow: {
    flexDirection: 'row',
    paddingVertical: 16,
    paddingHorizontal: 12,
    borderBottomWidth: 1,
    borderBottomColor: '#f1f3f4',
    alignItems: 'center',
    backgroundColor: '#ffffff',
  },

  tableCell: {
    fontSize: 12,
    color: '#212529',
    textAlign: 'center',
    lineHeight: 20,
  },

  tableTotalRow: {
    flexDirection: 'row',
    paddingVertical: 16,
    paddingHorizontal: 12,
    backgroundColor: '#f8f9fa',
    borderTopWidth: 2,
    borderTopColor: '#dee2e6',
    borderBottomLeftRadius: 12,
    borderBottomRightRadius: 12,
  },

  tableTotalCell: {
    fontSize: 15,
    fontWeight: '700',
    color: '#495057',
    textAlign: 'center',
  },

  input: {
    backgroundColor: '#ffffff',
    borderWidth: 2,
    borderColor: '#007bff',
    borderRadius: 8,
    paddingHorizontal: 12,
    paddingVertical: 8,
    fontSize: 14,
    color: '#212529',
    minHeight: 40,
  },

  editButton: {
    backgroundColor: '#007bff',
    paddingHorizontal: 11,
    
    paddingVertical: 8,
    paddingRight: 8,
    borderRadius: 24,
    minWidth: 50,
    alignItems: 'center',
    shadowColor: '#007bff',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.2,
    shadowRadius: 4,
    elevation: 3,
  },

  editButtonText: {
    color: '#ffffff',
    fontSize: 12,
    fontWeight: '600',
    textTransform: 'uppercase',
    letterSpacing: 0.5,
  },

  saveButton:{
    backgroundColor: '#28a745',
    paddingHorizontal: 11,
    
    paddingVertical: 8,
    paddingRight: 8,
    borderRadius: 24,
    minWidth: 50,
    alignItems: 'center',
    shadowColor: '#007bff',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.2,
    shadowRadius: 4,
    elevation: 3,
  },

  saveButtonText: {
    color: '#ffffff',
    fontSize: 12,
    fontWeight: '600',
    textTransform: 'uppercase',
    letterSpacing: 0.5,
  },

  saveAllButton: {
    backgroundColor: '#28a745',
    paddingVertical: 8,
    paddingHorizontal: 11,
    borderRadius: 8,
    marginHorizontal: 16,
    marginVertical: 12,
    alignItems: 'center',
    shadowColor: '#28a745',
    shadowOffset: {
      width: 0,
      height: 3,
    },
    shadowOpacity: 0.3,
    shadowRadius: 6,
    elevation: 5,
  },

  saveAllButtonText: {
    color: '#ffffff',
    fontSize: 16,
    fontWeight: '700',
    textTransform: 'uppercase',
    letterSpacing: 1,
  },

  sectionPointsContainer: {
    flexDirection: 'row',
    justifyContent: 'center',
    flexWrap: 'wrap',
    marginTop: 4,
    paddingHorizontal: 4,
  },

  sectionPointsText: {
    fontSize: 10,
    color: '#6c757d',
    backgroundColor: '#f8f9fa',
    paddingHorizontal: 6,
    paddingVertical: 2,
    borderRadius: 10,
    marginHorizontal: 2,
    overflow: 'hidden',
  },

  // Modal styles
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.6)',
  },

  modalContent: {
    backgroundColor: '#ffffff',
    borderRadius: 16,
    padding: 24,
    width: width * 0.9,
    maxWidth: 400,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 10,
    },
    shadowOpacity: 0.25,
    shadowRadius: 16,
    elevation: 10,
  },

  modalTitle: {
    fontSize: 20,
    fontWeight: '700',
    color: '#212529',
    textAlign: 'center',
    marginBottom: 20,
  },

  modalLabel: {
    fontSize: 14,
    fontWeight: '600',
    color: '#495057',
    marginBottom: 8,
  },

  modalInput: {
    backgroundColor: '#f8f9fa',
    borderWidth: 1,
    borderColor: '#dee2e6',
    borderRadius: 8,
    paddingHorizontal: 16,
    paddingVertical: 12,
    fontSize: 16,
    color: '#212529',
    marginBottom: 16,
  },

  modalButtonContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginTop: 8,
  },

  // Selector styles
  yearSelector: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#f8f9fa',
    borderRadius: 8,
    paddingHorizontal: 16,
    paddingVertical: 8,
    marginBottom: 16,
    marginHorizontal: 16,
  },

  deptSelector: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#f8f9fa',
    borderRadius: 8,
    paddingHorizontal: 16,
    paddingVertical: 8,
    marginBottom: 16,
    marginHorizontal: 16,
  },

  yearLabel: {
    fontSize: 14,
    fontWeight: '600',
    color: '#495057',
    marginRight: 12,
  },

  picker: {
    flex: 1,
    height: 40,
    color: '#212529',
  },

  yearPicker: {
    flex: 1,
    height: 40,
    color: '#212529',
  },

  // Loading and empty states
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 32,
  },

  loadingText: {
    fontSize: 16,
    color: '#6c757d',
    marginTop: 12,
  },

  emptyContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 32,
  },

  emptyText: {
    fontSize: 16,
    color: '#6c757d',
    textAlign: 'center',
    lineHeight: 24,
  },

  // Error states
  errorContainer: {
    backgroundColor: '#f8d7da',
    borderWidth: 1,
    borderColor: '#f5c6cb',
    borderRadius: 8,
    padding: 12,
    marginHorizontal: 16,
    marginVertical: 8,
  },

  errorText: {
    color: '#721c24',
    fontSize: 14,
    textAlign: 'center',
  },

  // Success states
  successContainer: {
    backgroundColor: '#d4edda',
    borderWidth: 1,
    borderColor: '#c3e6cb',
    borderRadius: 8,
    padding: 12,
    marginHorizontal: 16,
    marginVertical: 8,
  },

  successText: {
    color: '#155724',
    fontSize: 14,
    textAlign: 'center',
  },

  disabledButton: {
  backgroundColor: '#ccc',
  opacity: 0.7,
}
});