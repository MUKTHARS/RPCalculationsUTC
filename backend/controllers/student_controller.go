package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"Rp/backend/config"
	"Rp/backend/models"
)

func GetStudentSubjectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	studentID := r.URL.Path[len("/api/student-subjects/"):]
	if studentID == "" {
		SendErrorResponse(w, "Student ID is required", http.StatusBadRequest)
		return
	}

	// Initialize with default values
	totalPoints := 0
	pointsRemaining := 0

	// Try to get rewards data (but don't fail if not found)
	err := config.DB.QueryRow(`
        SELECT points, points_remaining 
        FROM student_rewards 
        WHERE LOWER(student_id) = LOWER(?)
    `, studentID).Scan(&totalPoints, &pointsRemaining)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Database error: %v", err)
		SendErrorResponse(w, "Error retrieving student data", http.StatusInternalServerError)
		return
	}

	// Get all subjects (this is the primary data we need)
	rows, err := config.DB.Query(`
        SELECT 
            ss.subject_code, 
            COALESCE(s.name, ss.subject_name) as subject_name,
            COALESCE(sp.points, 0) as points,
            COALESCE(sp.marks, 0) as marks
        FROM student_subjects ss
        LEFT JOIN subjects s ON ss.subject_code = s.code
        LEFT JOIN subject_points sp ON 
            LOWER(sp.student_id) = LOWER(ss.student_id) AND 
            sp.subject_code = ss.subject_code
        WHERE LOWER(ss.student_id) = LOWER(?)
    `, studentID)

	if err != nil {
		log.Printf("Database error: %v", err)
		SendErrorResponse(w, "Error retrieving student subjects", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type subjectData struct {
		Code   string  `json:"subjectCode"`
		Name   string  `json:"subjectName"`
		Points int     `json:"points"`
		Marks  float64 `json:"marks"`
	}

	var subjects []subjectData
	for rows.Next() {
		var subj subjectData
		if err := rows.Scan(&subj.Code, &subj.Name, &subj.Points, &subj.Marks); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		subjects = append(subjects, subj)
	}

	if len(subjects) == 0 {
		w.WriteHeader(http.StatusOK) // Change from NotFound to OK
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  true,
			"subjects": []interface{}{}, // Return empty array instead of error
			"message":  "No subjects found for this student",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":         true,
		"subjects":        subjects,
		"totalPoints":     totalPoints,
		"pointsRemaining": pointsRemaining,
	})
}
func GetStudentSubjectsOnlyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	studentID := r.URL.Path[len("/api/student-subjects-only/{student_id}"):]
	if studentID == "" {
		SendErrorResponse(w, "Student ID is required", http.StatusBadRequest)
		return
	}

	rows, err := config.DB.Query(`
        SELECT 
            subject_code, 
            COALESCE(subject_name, '') as subject_name
        FROM student_subjects
        WHERE LOWER(student_id) = LOWER(?)
    `, studentID)

	if err != nil {
		log.Printf("Database error: %v", err)
		SendErrorResponse(w, "Error retrieving student subjects", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type subjectData struct {
		Code string `json:"subjectCode"`
		Name string `json:"subjectName"`
	}

	var subjects []subjectData
	for rows.Next() {
		var subj subjectData
		if err := rows.Scan(&subj.Code, &subj.Name); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		subjects = append(subjects, subj)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"subjects": subjects,
	})
}

func CalculateFromDBHandler(w http.ResponseWriter, r *http.Request) {
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
		StudentID string `json:"studentId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	studentID := requestData.StudentID
	if studentID == "" {
		SendErrorResponse(w, "Student ID is required", http.StatusBadRequest)
		return
	}

	points, err := getStudentPoints(studentID)
	if err != nil {
		log.Printf("Database error: %v", err)
		SendErrorResponse(w, "Error retrieving points from database", http.StatusInternalServerError)
		return
	}

	if points < 0 {
		SendErrorResponse(w, "Points cannot be negative", http.StatusBadRequest)
		return
	}

	
	var year int
var department string
err = config.DB.QueryRow(`
    SELECT year, department FROM student_rewards 
    WHERE student_id = ?
`, requestData.StudentID).Scan(&year, &department)

if err != nil {
    // Default values if not found
    year = 1
    department = "CSE"
}
	allocation := AllocatePoints(points, department, year)

	_, err = config.DB.Exec(`
		UPDATE student_rewards 
		SET total_marks = ?, points_remaining = ?
		WHERE student_id = ?
	`, allocation.TotalMarks, allocation.PointsRemaining, studentID)

	if err != nil {
		log.Printf("Database error updating student rewards: %v", err)
		SendErrorResponse(w, "Error updating student totals", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ApiResponse{
		Success: true,
		Data:    allocation,
	})
}

func getStudentPoints(studentID string) (int, error) {
	var points int
	err := config.DB.QueryRow("SELECT points FROM student_rewards WHERE student_id = ?", studentID).Scan(&points)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	return points, err
}
func SaveSubjectPointsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var requestData struct {
		StudentID     string  `json:"studentId"`
		SubjectCode   string  `json:"subjectCode"`
		SubjectName   string  `json:"subjectName"`
		Points        int     `json:"points"`
		Marks         float64 `json:"marks"`
		SectionPoints []int   `json:"sectionPoints"`
		Semester      int     `json:"semester"`
		Year          int     `json:"year"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Convert section points to JSON for storage
	sectionPointsJSON, err := json.Marshal(requestData.SectionPoints)
	if err != nil {
		SendErrorResponse(w, "Error processing section points", http.StatusInternalServerError)
		return
	}

	// Save to database
	_, err = config.DB.Exec(`
        INSERT INTO student_subject_points 
        (student_id, subject_code, subject_name, points, marks, section_points, semester, year)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
            subject_name = VALUES(subject_name),
            points = VALUES(points),
            marks = VALUES(marks),
            section_points = VALUES(section_points),
            semester = VALUES(semester),
            year = VALUES(year),
            updated_at = CURRENT_TIMESTAMP
    `,
		requestData.StudentID,
		requestData.SubjectCode,
		requestData.SubjectName,
		requestData.Points,
		requestData.Marks,
		string(sectionPointsJSON),
		requestData.Semester,
		requestData.Year)

	if err != nil {
		log.Printf("Database error saving subject points: %v", err)
		SendErrorResponse(w, "Error saving subject points", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ApiResponse{
		Success: true,
	})
}
func SaveAllSubjectsHandler(w http.ResponseWriter, r *http.Request) {
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
		models.SaveAllSubjectsRequest
		Semester int `json:"semester"`
		Year     int `json:"year"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Set default values if not provided
	if requestData.Semester == 0 {
		requestData.Semester = 1
	}
	if requestData.Year == 0 {
		requestData.Year = time.Now().Year()
	}

	tx, err := config.DB.Begin()
	if err != nil {
		log.Printf("Database error starting transaction: %v", err)
		SendErrorResponse(w, "Error starting transaction", http.StatusInternalServerError)
		return
	}

	// First update student_rewards with year information
	_, err = tx.Exec(`
        INSERT INTO student_rewards (student_id, year, points, total_marks, points_remaining)
        VALUES (?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE 
            year = VALUES(year),
            points = VALUES(points),
            total_marks = VALUES(total_marks),
            points_remaining = VALUES(points_remaining)
    `, requestData.StudentID, requestData.Year, requestData.TotalPoints,
        requestData.TotalMarks, requestData.PointsRemaining)
    if err != nil {
        tx.Rollback()
        log.Printf("Database error updating student year: %v", err)
        SendErrorResponse(w, "Error updating student year", http.StatusInternalServerError)
        return
    }

	// First delete all existing subject points for this student/semester/year
	// to ensure we have a clean slate
	_, err = tx.Exec(`
        DELETE FROM subject_points 
        WHERE student_id = ? AND semester = ? AND year = ?
    `, requestData.StudentID, requestData.Semester, requestData.Year)
	if err != nil {
		tx.Rollback()
		log.Printf("Database error clearing existing subject points: %v", err)
		SendErrorResponse(w, "Error clearing existing subject points", http.StatusInternalServerError)
		return
	}

	// Now insert all subjects (whether modified or not)
	for _, subject := range requestData.Subjects {
		// Create student-subject mapping
		_, err = tx.Exec(`
            INSERT INTO student_subjects (student_id, subject_code, semester, year)
            VALUES (?, ?, ?, ?)
            ON DUPLICATE KEY UPDATE
                semester = VALUES(semester),
                year = VALUES(year)
        `, requestData.StudentID, subject.SubjectCode, requestData.Semester, requestData.Year)
		if err != nil {
			tx.Rollback()
			log.Printf("Database error saving student subject mapping: %v", err)
			SendErrorResponse(w, "Error saving student subject mapping", http.StatusInternalServerError)
			return
		}

		// Save subject points with semester/year
		// This will insert ALL subjects, not just modified ones
		_, err = tx.Exec(`
            INSERT INTO subject_points 
                (student_id, subject_code, semester, year, subject_name, points, marks) 
            VALUES (?, ?, ?, ?, ?, ?, ?)
        `, requestData.StudentID, subject.SubjectCode, requestData.Semester,
			requestData.Year, subject.SubjectName, subject.Points, subject.Marks)
		if err != nil {
			tx.Rollback()
			log.Printf("Database error saving subject: %v", err)
			SendErrorResponse(w, "Error saving subject data", http.StatusInternalServerError)
			return
		}
	}

	// Save totals
	_, err = tx.Exec(`
        INSERT INTO student_totals (student_id, total_points, total_marks) 
        VALUES (?, ?, ?)
        ON DUPLICATE KEY UPDATE 
            total_points = VALUES(total_points),
            total_marks = VALUES(total_marks)
    `, requestData.StudentID, requestData.TotalPoints, requestData.TotalMarks)
	if err != nil {
		tx.Rollback()
		log.Printf("Database error saving totals: %v", err)
		SendErrorResponse(w, "Error saving total data", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Database error committing transaction: %v", err)
		SendErrorResponse(w, "Error saving data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ApiResponse{
		Success: true,
	})
}

func UpdateSubjectPointsHandler(w http.ResponseWriter, r *http.Request) {
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
        StudentID       string  `json:"studentId"`
        SubjectCode     string  `json:"subjectCode"`
        SubjectName     string  `json:"subjectName"`
        Points          int     `json:"points"`
        Marks           float64 `json:"marks"`
        TotalMarks      float64 `json:"totalMarks"`
        PointsRemaining int     `json:"pointsRemaining"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
        return
    }

    if requestData.StudentID == "" || requestData.SubjectCode == "" {
        SendErrorResponse(w, "Student ID and Subject Code are required", http.StatusBadRequest)
        return
    }

    if requestData.Points < 0 || requestData.Points > config.MaxPointsPerSubject {
        SendErrorResponse(w, fmt.Sprintf("Points must be between 0 and %d", config.MaxPointsPerSubject), http.StatusBadRequest)
        return
    }

    // Get department and year from the database
    var department string
    var year int
    err := config.DB.QueryRow(`
        SELECT department, year 
        FROM student_rewards 
        WHERE student_id = ?
    `, requestData.StudentID).Scan(&department, &year)
    if err != nil {
        // Default values if not found
        department = "CSE"
        year = 1
    }

    // Check redemption deadline before proceeding
    redemptionConfig, err := config.GetRedemptionDates(year, department)
    if err != nil {
        SendErrorResponse(w, "Failed to check redemption deadline", http.StatusInternalServerError)
        return
    }

    deadline := redemptionConfig.EndDate
    if redemptionConfig.IsExtended && !redemptionConfig.ExtendedDate.IsZero() {
        deadline = redemptionConfig.ExtendedDate
    }

    if time.Now().After(deadline) {
    var canEditAfterDeadline bool
    err = config.DB.QueryRow(`
        SELECT can_edit_after_deadline 
        FROM student_rewards 
        WHERE student_id = ?
    `, requestData.StudentID).Scan(&canEditAfterDeadline)
    
    if err != nil || !canEditAfterDeadline {
        SendErrorResponse(w, "Redemption period has ended. Contact admin for extension.", http.StatusForbidden)
        return
    }
}

    subjectName := requestData.SubjectName
    if subjectName == "" {
        codeIndex, err := strconv.Atoi(requestData.SubjectCode[3:])
        configData := config.GetConfigForYearAndDepartment(year, department)
        if err == nil && codeIndex > 0 && codeIndex <= len(configData.SubjectNames) {
            subjectName = configData.SubjectNames[codeIndex-1]
        } else {
            subjectName = "Subject " + requestData.SubjectCode[3:]
        }
    }

    tx, err := config.DB.Begin()
    if err != nil {
        log.Printf("Database error starting transaction: %v", err)
        SendErrorResponse(w, "Error starting transaction", http.StatusInternalServerError)
        return
    }

    _, err = tx.Exec(`
        INSERT INTO subject_points (student_id, subject_code, subject_name, points, marks) 
        VALUES (?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE 
            subject_name = VALUES(subject_name), 
            points = VALUES(points),
            marks = VALUES(marks)
    `, requestData.StudentID, requestData.SubjectCode, subjectName, requestData.Points, requestData.Marks)

    if err != nil {
        tx.Rollback()
        log.Printf("Database error saving subject points: %v", err)
        SendErrorResponse(w, "Error saving subject points", http.StatusInternalServerError)
        return
    }

    _, err = tx.Exec(`
        UPDATE student_rewards 
        SET total_marks = ?, points_remaining = ?
        WHERE student_id = ?
    `, requestData.TotalMarks, requestData.PointsRemaining, requestData.StudentID)

    if err != nil {
        tx.Rollback()
        log.Printf("Database error updating student rewards: %v", err)
        SendErrorResponse(w, "Error updating student totals", http.StatusInternalServerError)
        return
    }

    if err := tx.Commit(); err != nil {
        log.Printf("Database error committing transaction: %v", err)
        SendErrorResponse(w, "Error saving data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.ApiResponse{
        Success: true,
    })
}
func GetStudentRewardsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    studentID := r.URL.Path[len("/api/student-rewards/"):]
    if studentID == "" {
        SendErrorResponse(w, "Student ID is required", http.StatusBadRequest)
        return
    }

    var department *string // Use pointer to detect NULL
    var academicYear *int
    
    // Get the most common department and academic year from student_subjects
    err := config.DB.QueryRow(`
        SELECT department, year 
        FROM student_subjects 
        WHERE student_id = ?
        GROUP BY department, year
        ORDER BY COUNT(*) DESC
        LIMIT 1
    `, studentID).Scan(&department, &academicYear)

    if err != nil {
        if err == sql.ErrNoRows {
            // Return empty values instead of defaults
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success":    true,
                "department": nil,
                "year":       nil,
            })
            return
        } else {
            log.Printf("Database error: %v", err)
            SendErrorResponse(w, "Error retrieving student data", http.StatusInternalServerError)
            return
        }
    }

    // Only return values if they exist
    result := map[string]interface{}{
        "success": true,
    }
    
    if department != nil {
        result["department"] = *department
    } else {
        result["department"] = nil
    }
    
    if academicYear != nil {
        result["year"] = *academicYear
    } else {
        result["year"] = nil
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func GetSubjectPointsHandler(w http.ResponseWriter, r *http.Request) {
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

	studentID := r.URL.Path[len("/api/subject-points/"):]
	if studentID == "" {
		SendErrorResponse(w, "Student ID is required", http.StatusBadRequest)
		return
	}

	// Get query parameters for semester/year filtering
	queryParams := r.URL.Query()
	semester := queryParams.Get("semester")
	year := queryParams.Get("year")

	var rows *sql.Rows
	var err error

	// Build query based on provided filters
	if semester != "" && year != "" {
		// Get points for specific semester/year
		rows, err = config.DB.Query(`
            SELECT sp.subject_code, sp.subject_name, sp.points, sp.marks, 
                   sp.semester, sp.year
            FROM subject_points sp
            JOIN student_subjects ss ON sp.student_id = ss.student_id 
                AND sp.subject_code = ss.subject_code
                AND sp.semester = ss.semester
                AND sp.year = ss.year
            WHERE sp.student_id = ? AND sp.semester = ? AND sp.year = ?
            ORDER BY sp.subject_code
        `, studentID, semester, year)
	} else if year != "" {
		// Get all points for specific year
		rows, err = config.DB.Query(`
            SELECT sp.subject_code, sp.subject_name, sp.points, sp.marks, 
                   sp.semester, sp.year
            FROM subject_points sp
            JOIN student_subjects ss ON sp.student_id = ss.student_id 
                AND sp.subject_code = ss.subject_code
                AND sp.semester = ss.semester
                AND sp.year = ss.year
            WHERE sp.student_id = ? AND sp.year = ?
            ORDER BY sp.semester, sp.subject_code
        `, studentID, year)
	} else {
		// Get all points (backward compatibility)
		rows, err = config.DB.Query(`
            SELECT sp.subject_code, sp.subject_name, sp.points, sp.marks, 
                   sp.semester, sp.year
            FROM subject_points sp
            JOIN student_subjects ss ON sp.student_id = ss.student_id 
                AND sp.subject_code = ss.subject_code
                AND sp.semester = ss.semester
                AND sp.year = ss.year
            WHERE sp.student_id = ?
            ORDER BY sp.year, sp.semester, sp.subject_code
        `, studentID)
	}

	if err != nil {
		log.Printf("Database error: %v", err)
		SendErrorResponse(w, "Error retrieving subject points", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	pointsMap := make(map[string]interface{})
	for rows.Next() {
		var code, name string
		var points int
		var marks float64
		var semester, year int

		if err := rows.Scan(&code, &name, &points, &marks, &semester, &year); err != nil {
			log.Printf("Error scanning row: %v", err)
			SendErrorResponse(w, "Error processing data", http.StatusInternalServerError)
			return
		}

		key := fmt.Sprintf("%s-%d-%d", code, semester, year)
		pointsMap[key] = map[string]interface{}{
			"subjectCode": code,
			"subjectName": name,
			"points":      points,
			"marks":       marks,
			"semester":    semester,
			"year":        year,
		}
	}

	// Get student year if available
	var studentYear int
	err = config.DB.QueryRow(`
        SELECT year FROM student_rewards WHERE student_id = ?
    `, studentID).Scan(&studentYear)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error getting student year: %v", err)
	}

	response := map[string]interface{}{
		"success":     true,
		"data":        pointsMap,
		"studentId":   studentID,
		"studentYear": studentYear,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
