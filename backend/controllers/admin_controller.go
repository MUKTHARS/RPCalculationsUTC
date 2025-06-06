package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Rp/backend/config"
)

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		SendErrorResponse(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get year and department from query params
	yearStr := r.URL.Query().Get("year")
	department := r.URL.Query().Get("department")
	if department == "" {
		department = "CSE" // Default department
	}

	if yearStr == "" {
		// Return all configs if no year specified
		allConfigs := make(map[string]interface{})
		for dept, yearConfigs := range config.YearlyConfigs {
			allConfigs[dept] = yearConfigs
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allConfigs)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		SendErrorResponse(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	configData := config.GetConfigForYearAndDepartment(year, department)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(configData)
}

func UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		SendErrorResponse(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Year   int                  `json:"year"`
		Config *config.YearlyConfig `json:"config"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if requestData.Year < 1 || requestData.Year > 4 {
		SendErrorResponse(w, "Year must be between 1 and 4", http.StatusBadRequest)
		return
	}

	if requestData.Config == nil {
		SendErrorResponse(w, "Config cannot be empty", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if requestData.Config.RedemptionRatio <= 0 {
		SendErrorResponse(w, "Redemption ratio must be positive", http.StatusBadRequest)
		return
	}

	if len(requestData.Config.RedemptionVariance) != 4 {
		SendErrorResponse(w, "Redemption variance must have exactly 4 values", http.StatusBadRequest)
		return
	}

	if len(requestData.Config.SectionMaxMarks) != 4 {
		SendErrorResponse(w, "Section max marks must have exactly 4 values", http.StatusBadRequest)
		return
	}

	if err := config.SaveConfigForYear(requestData.Year, requestData.Config); err != nil {
		SendErrorResponse(w, fmt.Sprintf("Failed to save config: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func GetRedemptionDatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		SendErrorResponse(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	yearStr := r.URL.Query().Get("year")
	department := r.URL.Query().Get("department")
	if department == "" {
		department = "CSE"
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		SendErrorResponse(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	config, err := config.GetRedemptionDates(year, department)
	if err != nil {
		SendErrorResponse(w, fmt.Sprintf("Failed to get redemption dates: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

func SaveRedemptionDatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		SendErrorResponse(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Year       int                      `json:"year"`
		Department string                   `json:"department"`
		Config     *config.RedemptionConfig `json:"config"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		SendErrorResponse(w, "Invalid request format: "+err.Error(), http.StatusBadRequest)
		return
	}

	if request.Config == nil {
		SendErrorResponse(w, "Config cannot be empty", http.StatusBadRequest)
		return
	}

	// Validate dates
	if request.Config.StartDate.IsZero() || request.Config.EndDate.IsZero() {
		SendErrorResponse(w, "Both start and end dates are required", http.StatusBadRequest)
		return
	}

	if request.Config.StartDate.After(request.Config.EndDate) {
		SendErrorResponse(w, "Start date must be before end date", http.StatusBadRequest)
		return
	}

	// If extended date is set, validate it
	if !request.Config.ExtendedDate.IsZero() && request.Config.ExtendedDate.Before(request.Config.EndDate) {
		SendErrorResponse(w, "Extended date must be after the original end date", http.StatusBadRequest)
		return
	}

	err := config.SaveRedemptionDates(request.Year, request.Department, request.Config)
	if err != nil {
		SendErrorResponse(w, "Failed to save redemption dates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func ExtendRedemptionDeadlineHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		SendErrorResponse(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Year       int       `json:"year"`
		Department string    `json:"department"`
		NewDate    time.Time `json:"newDate"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if request.NewDate.IsZero() {
		SendErrorResponse(w, "Invalid date", http.StatusBadRequest)
		return
	}

	if err := config.ExtendRedemptionDeadline(request.Year, request.Department, request.NewDate); err != nil {
		SendErrorResponse(w, fmt.Sprintf("Failed to extend deadline: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func UnfreezeStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Only POST method is allowed"})
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		StudentID string `json:"studentId"`
		Unfreeze  bool   `json:"unfreeze"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if request.StudentID == "" {
		http.Error(w, "Student ID is required", http.StatusBadRequest)
		return
	}

	// First check if student exists in student_rewards
	var existingStatus *bool
	err := config.DB.QueryRow("SELECT can_edit_after_deadline FROM student_rewards WHERE student_id = ?", request.StudentID).Scan(&existingStatus)

	// If no record exists or there was an error, we'll proceed with the update
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("Failed to check student status: %v", err), http.StatusInternalServerError)
		return
	}

	// Update or insert the student's unfreeze status
	_, err = config.DB.Exec(`
        INSERT INTO student_rewards (student_id, can_edit_after_deadline) 
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE 
            can_edit_after_deadline = ?
    `, request.StudentID, request.Unfreeze, request.Unfreeze)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update student status: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"unfreeze": request.Unfreeze,
	})
}
