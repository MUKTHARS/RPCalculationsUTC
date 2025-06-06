package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	// "log"
	// "math"
)

type YearlyConfig struct {
	RedemptionRatio     int       `json:"redemptionRatio"`
	SubjectsCount       int       `json:"subjectsCount"`
	MaxPointsPerSubject int       `json:"maxPointsPerSubject"`
	RedemptionVariance  []int     `json:"redemptionVariance"`
	SectionPoints       []int     `json:"sectionPoints"`
	SectionMaxMarks     []float64 `json:"sectionMaxMarks"`
	PointsPerMark       []float64 `json:"pointsPerMark"`
	MinPointsForHalf    []float64 `json:"minPointsForHalf"`
	SubjectNames        []string  `json:"subjectNames"`
	Department          string    `json:"department"`
}

type RedemptionConfig struct {
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	IsExtended   bool      `json:"isExtended"`
	ExtendedDate time.Time `json:"extendedDate"`
}

var (
	RedemptionRatio     = 10
	SubjectsCount       = 6
	MaxPointsPerSubject = 250
)

var (
	DB                 *sql.DB
	YearlyConfigs      = make(map[string]map[int]*YearlyConfig)
	DefaultDepartment  = "CSE"
	RedemptionVariance = []int{1, 2, 4, 6}
	SectionPoints      = []int{50, 60, 80, 60}
	SectionMaxMarks    = []float64{5, 3, 2, 1}
	PointsPerMark      = []float64{10, 20, 40, 60}
	MinPointsForHalf   = []float64{5, 10, 20, 30}

	// SubjectNames = []string{
	// 	"Mathematics",
	// 	"Physics",
	// 	"Chemistry",
	// 	"Biology",
	// 	"History",
	// 	"Geography",
	// }
)

const (
	dbUser     = "root"
	dbPassword = "1234"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "rewards_db"
)

func InitDB() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	// Create all tables first
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS config_year (
			year INT NOT NULL,
			department VARCHAR(50) NOT NULL,
			config_data TEXT NOT NULL,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (year, department)
	)`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS departments (
        code VARCHAR(10) PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
`)
	if err != nil {
		return err
	}

	var deptCount int
	err = DB.QueryRow("SELECT COUNT(*) FROM departments").Scan(&deptCount)
	if err != nil {
		return err
	}

	if deptCount == 0 {
		defaultDepts := map[string]string{
			"CSE":  "Computer Science",
			"IT":   "Information Technology",
			"ECE":  "Electronics and Communication",
			"EEE":  "Electrical and Electronics",
			"MECH": "Mechanical",
		}
		for code, name := range defaultDepts {
			_, err = DB.Exec("INSERT INTO departments (code, name) VALUES (?, ?)", code, name)
			if err != nil {
				return err
			}
		}
	}

	// Create subjects table with subject_type and credits columns
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS subjects (
			id INT AUTO_INCREMENT PRIMARY KEY,
			code VARCHAR(10) UNIQUE NOT NULL,
			name VARCHAR(100) NOT NULL,
			subject_type ENUM('core', 'addon', 'elective') NOT NULL DEFAULT 'core',
			is_active TINYINT(1) NOT NULL DEFAULT 1,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Create subject_points table with semester and year columns
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS subject_points (
			student_id VARCHAR(50) NOT NULL,
			subject_code VARCHAR(10) NOT NULL,
			semester INT NOT NULL DEFAULT 1,
			year INT NOT NULL DEFAULT 1,
			department VARCHAR(50) NOT NULL DEFAULT 'CSE',
			subject_name VARCHAR(100) NOT NULL,
			points INT NOT NULL DEFAULT 0,
			marks FLOAT NOT NULL DEFAULT 0,
			PRIMARY KEY (student_id, subject_code, semester, year, department)
		)
	`)
	if err != nil {
		return err
	}

	// Create student_totals table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS student_totals (
			student_id VARCHAR(50) PRIMARY KEY,
			total_points INT NOT NULL DEFAULT 0,
			total_marks FLOAT NOT NULL DEFAULT 0
		)
	`)
	if err != nil {
		return err
	}

	// Create student_subjects table
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS student_subjects (
            id INT AUTO_INCREMENT PRIMARY KEY,
            student_id VARCHAR(50) NOT NULL,
            subject_code VARCHAR(10) NOT NULL,
            semester INT NOT NULL,
            year INT NOT NULL,
            department VARCHAR(50) NOT NULL DEFAULT 'CSE',
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (student_id) REFERENCES student_rewards(student_id),
            FOREIGN KEY (subject_code) REFERENCES subjects(code),
            UNIQUE KEY (student_id, subject_code, semester, year,department)
        )
    `)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS redemption_dates (
        year INT NOT NULL,
        department VARCHAR(50) NOT NULL,
        start_date DATETIME NOT NULL,
        end_date DATETIME NOT NULL,
        is_extended BOOLEAN DEFAULT FALSE,
        extended_date DATETIME,
        PRIMARY KEY (year, department))
`)
	if err != nil {
		return fmt.Errorf("error creating redemption_dates table: %v", err)
	}

	// Check and add columns if needed
	var studentNameExists int
	err = DB.QueryRow(`
        SELECT COUNT(*) FROM information_schema.columns 
        WHERE table_name = 'student_rewards' AND column_name = 'student_name'
    `).Scan(&studentNameExists)
	if err == nil && studentNameExists == 0 {
		_, err = DB.Exec(`ALTER TABLE student_rewards ADD COLUMN student_name VARCHAR(100) NOT NULL DEFAULT '' AFTER student_id`)
		if err != nil {
			return err
		}
	}

	var studentEmailExists int
	err = DB.QueryRow(`
        SELECT COUNT(*) FROM information_schema.columns 
        WHERE table_name = 'student_rewards' AND column_name = 'student_email'
    `).Scan(&studentEmailExists)
	if err == nil && studentEmailExists == 0 {
		_, err = DB.Exec(`ALTER TABLE student_rewards ADD COLUMN student_email VARCHAR(100) NOT NULL DEFAULT '' AFTER student_name`)
		if err != nil {
			return err
		}
	}

	var canEditColumnExists int
    err = DB.QueryRow(`
        SELECT COUNT(*) FROM information_schema.columns 
        WHERE table_name = 'student_rewards' AND column_name = 'can_edit_after_deadline'
    `).Scan(&canEditColumnExists)
    if err == nil && canEditColumnExists == 0 {
        _, err = DB.Exec(`
            ALTER TABLE student_rewards 
            ADD COLUMN can_edit_after_deadline BOOLEAN NOT NULL DEFAULT FALSE
        `)
        if err != nil {
            return fmt.Errorf("failed to add can_edit_after_deadline column: %v", err)
        }
    }

	var subjectTypeExists int
	err = DB.QueryRow(`
        SELECT COUNT(*) FROM information_schema.columns 
        WHERE table_name = 'subjects' AND column_name = 'subject_type'
    `).Scan(&subjectTypeExists)
	if err == nil && subjectTypeExists == 0 {
		_, err = DB.Exec(`ALTER TABLE subjects ADD COLUMN subject_type ENUM('core', 'addon', 'elective') NOT NULL DEFAULT 'core'`)
		if err != nil {
			return err
		}
	}


	

	// Initialize default data
	var configCount int
	err = DB.QueryRow("SELECT COUNT(*) FROM config WHERE config_key NOT LIKE '%Year'").Scan(&configCount)
	if err != nil {
		return err
	}

	if configCount == 0 {
		defaultConfig := map[string]interface{}{
			"RedemptionRatio":     RedemptionRatio,
			"SubjectsCount":       SubjectsCount,
			"MaxPointsPerSubject": MaxPointsPerSubject,
			"RedemptionVariance":  RedemptionVariance,
			"SectionPoints":       SectionPoints,
			"SectionMaxMarks":     SectionMaxMarks,
			"PointsPerMark":       PointsPerMark,
			"MinPointsForHalf":    MinPointsForHalf,
		}

		for key, value := range defaultConfig {
			jsonValue, err := json.Marshal(value)
			if err != nil {
				return err
			}

			_, err = DB.Exec("INSERT INTO config (config_key, config_value) VALUES (?, ?)", key, string(jsonValue))
			if err != nil {
				return err
			}
		}
	}

	// Now load the config after all tables and data are initialized
	if err := LoadConfigFromDB(); err != nil {
		return fmt.Errorf("failed to load config from DB: %v", err)
	}
	if err := initDefaultData(); err != nil {
		return err
	}

	if err := LoadAllYearConfigs(); err != nil {
		return fmt.Errorf("failed to load configs from DB: %v", err)
	}

	LoadSubjectNames()

	return nil
}
func SaveRedemptionDates(year int, department string, config *RedemptionConfig) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	// Validate dates
	if config.StartDate.IsZero() || config.EndDate.IsZero() {
		return fmt.Errorf("both start and end dates are required")
	}

	if config.StartDate.After(config.EndDate) {
		return fmt.Errorf("start date must be before end date")
	}

	// If extended, validate extended date
	if config.IsExtended && config.ExtendedDate.IsZero() {
		return fmt.Errorf("extended date is required when isExtended is true")
	}

	if config.IsExtended && !config.ExtendedDate.IsZero() && config.ExtendedDate.Before(config.EndDate) {
		return fmt.Errorf("extended date must be after the original end date")
	}

	// Format dates as strings for MySQL
	startDateStr := config.StartDate.Format("2006-01-02 15:04:05")
	endDateStr := config.EndDate.Format("2006-01-02 15:04:05")
	var extendedDateStr interface{}
	if config.IsExtended && !config.ExtendedDate.IsZero() {
		extendedDateStr = config.ExtendedDate.Format("2006-01-02 15:04:05")
	} else {
		extendedDateStr = nil
	}

	// Use transaction to ensure data consistency
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	_, err = tx.Exec(`
        INSERT INTO redemption_dates (year, department, start_date, end_date, is_extended, extended_date)
        VALUES (?, ?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE 
            start_date = VALUES(start_date),
            end_date = VALUES(end_date),
            is_extended = VALUES(is_extended),
            extended_date = VALUES(extended_date)
    `, year, department, startDateStr, endDateStr, config.IsExtended, extendedDateStr)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("database error: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func GetRedemptionDates(year int, department string) (*RedemptionConfig, error) {
	var config RedemptionConfig
	var startDateStr, endDateStr string
	var extendedDateStr sql.NullString
	var isExtended bool

	err := DB.QueryRow(`
        SELECT start_date, end_date, is_extended, extended_date
        FROM redemption_dates
        WHERE year = ? AND department = ?
    `, year, department).Scan(
		&startDateStr,
		&endDateStr,
		&isExtended,
		&extendedDateStr,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Return default config if not found
			return &RedemptionConfig{
				StartDate:  time.Now(),
				EndDate:    time.Now().AddDate(0, 1, 0), // Default 1 month
				IsExtended: false,
			}, nil
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	// Parse the date strings into time.Time
	startDate, err := time.Parse("2006-01-02 15:04:05", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02 15:04:05", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	config = RedemptionConfig{
		StartDate:  startDate,
		EndDate:    endDate,
		IsExtended: isExtended,
	}

	if extendedDateStr.Valid {
		extendedDate, err := time.Parse("2006-01-02 15:04:05", extendedDateStr.String)
		if err != nil {
			return nil, fmt.Errorf("failed to parse extended date: %v", err)
		}
		config.ExtendedDate = extendedDate
	}

	return &config, nil
}

func ExtendRedemptionDeadline(year int, department string, newDate time.Time) error {
	if newDate.IsZero() {
		return fmt.Errorf("invalid date")
	}

	// First check if there's an existing config
	existing, err := GetRedemptionDates(year, department)
	if err != nil {
		return fmt.Errorf("failed to get existing redemption dates: %v", err)
	}

	// Validate the new date is after the original end date
	if newDate.Before(existing.EndDate) {
		return fmt.Errorf("extended date must be after the original end date")
	}

	_, err = DB.Exec(`
        UPDATE redemption_dates
        SET is_extended = TRUE,
            extended_date = ?
        WHERE year = ? AND department = ?
    `, newDate, year, department)

	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}

	return nil
}
func GetAllDepartments() ([]map[string]string, error) {
	rows, err := DB.Query("SELECT code, name FROM departments ORDER BY code")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []map[string]string
	for rows.Next() {
		var code, name string
		if err := rows.Scan(&code, &name); err != nil {
			return nil, err
		}
		departments = append(departments, map[string]string{
			"code": code,
			"name": name,
		})
	}
	return departments, nil
}

func AddDepartment(code, name string) error {
	_, err := DB.Exec("INSERT INTO departments (code, name) VALUES (?, ?)", code, name)
	return err
}

func DeleteDepartment(code string) error {
	_, err := DB.Exec("DELETE FROM departments WHERE code = ?", code)
	return err
}

func initDefaultData() error {
	// Check and initialize default subjects
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM subjects").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		defaultSubjects := []string{
			"Mathematics",
			"Physics",
			"Chemistry",
			"Biology",
			"History",
			"Geography",
		}
		for i, name := range defaultSubjects {
			code := fmt.Sprintf("SUB%d", i+1)
			_, err = DB.Exec("INSERT INTO subjects (code, name) VALUES (?, ?)", code, name)
			if err != nil {
				return err
			}
		}
	}

	// Initialize default config for all years and departments
	departments := []string{"CSE", "IT", "ECE", "EEE", "MECH"}
	years := []int{1, 2, 3, 4}

	for _, dept := range departments {
		for _, year := range years {
			var count int
			err := DB.QueryRow("SELECT COUNT(*) FROM config_year WHERE year = ? AND department = ?", year, dept).Scan(&count)
			if err != nil {
				return err
			}

			if count == 0 {
				defaultConfig := createDefaultConfig(year, dept)
				jsonValue, err := json.Marshal(defaultConfig)
				if err != nil {
					return err
				}

				_, err = DB.Exec("INSERT INTO config_year (year, department, config_data) VALUES (?, ?, ?)",
					year, dept, string(jsonValue))
				if err != nil {
					return err
				}

				// Also save to config table with department suffix
				if err := saveDepartmentConfigToMainTable(year, dept, defaultConfig); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func getYearSuffix(year int) string {
	switch year {
	case 1:
		return "FirstYear"
	case 2:
		return "SecondYear"
	case 3:
		return "ThirdYear"
	case 4:
		return "FinalYear"
	default:
		return ""
	}
}
func saveDepartmentConfigToMainTable(year int, department string, config *YearlyConfig) error {
	yearSuffix := getYearSuffix(year)
	deptSuffix := strings.ToUpper(department)

	// Save each config value with year and department suffix
	if err := saveConfigValue(fmt.Sprintf("RedemptionRatio%s%s", yearSuffix, deptSuffix), config.RedemptionRatio); err != nil {
		return err
	}
	if err := saveConfigValue(fmt.Sprintf("SubjectsCount%s%s", yearSuffix, deptSuffix), config.SubjectsCount); err != nil {
		return err
	}
	if err := saveConfigValue(fmt.Sprintf("MaxPointsPerSubject%s%s", yearSuffix, deptSuffix), config.MaxPointsPerSubject); err != nil {
		return err
	}
	if err := saveConfigArray(fmt.Sprintf("RedemptionVariance%s%s", yearSuffix, deptSuffix), config.RedemptionVariance); err != nil {
		return err
	}
	if err := saveConfigArray(fmt.Sprintf("SectionPoints%s%s", yearSuffix, deptSuffix), config.SectionPoints); err != nil {
		return err
	}
	if err := saveConfigArray(fmt.Sprintf("SectionMaxMarks%s%s", yearSuffix, deptSuffix), config.SectionMaxMarks); err != nil {
		return err
	}
	if err := saveConfigArray(fmt.Sprintf("PointsPerMark%s%s", yearSuffix, deptSuffix), config.PointsPerMark); err != nil {
		return err
	}
	if err := saveConfigArray(fmt.Sprintf("MinPointsForHalf%s%s", yearSuffix, deptSuffix), config.MinPointsForHalf); err != nil {
		return err
	}

	return nil
}
func createDefaultConfig(year int, department string) *YearlyConfig {
	// Base values that can vary by year/department
	redemptionRatio := 10
	if year == 2 {
		redemptionRatio = 10
	} else if year == 3 {
		redemptionRatio = 10
	} else if year == 4 {
		redemptionRatio = 10
	}

	// Department-specific adjustments
	if department == "CSE" {
		redemptionRatio += 5
	} else if department == "IT" {
		redemptionRatio += 3
	}

	config := &YearlyConfig{
		RedemptionRatio:    redemptionRatio,
		SubjectsCount:      6,
		RedemptionVariance: []int{1, 2, 4, 6},
		SectionMaxMarks:    []float64{5, 3, 2, 1},
		Department:         department,
	}

	// Calculate dependent values
	calculateDependentValues(config)
	return config
}
func calculateDependentValues(config *YearlyConfig) {
	config.SectionPoints = make([]int, len(config.SectionMaxMarks))
	for i := range config.SectionPoints {
		config.SectionPoints[i] = int(config.SectionMaxMarks[i] * float64(config.RedemptionRatio) * float64(config.RedemptionVariance[i]))
	}

	config.PointsPerMark = make([]float64, len(config.SectionPoints))
	for i := range config.PointsPerMark {
		config.PointsPerMark[i] = float64(config.SectionPoints[i]) / config.SectionMaxMarks[i]
	}

	config.MinPointsForHalf = make([]float64, len(config.PointsPerMark))
	for i := range config.MinPointsForHalf {
		config.MinPointsForHalf[i] = config.PointsPerMark[i] / 2
	}

	config.MaxPointsPerSubject = 0
	for _, points := range config.SectionPoints {
		config.MaxPointsPerSubject += points
	}
}
func LoadAllYearConfigs() error {
	rows, err := DB.Query("SELECT year, config_data FROM config_year ORDER BY year,department")
	if err != nil {
		return fmt.Errorf("error querying config_year: %v", err)
	}
	defer rows.Close()
	YearlyConfigs = make(map[string]map[int]*YearlyConfig)
	// Clear existing configs
	for k := range YearlyConfigs {
		delete(YearlyConfigs, k)
	}

	for rows.Next() {
		var year int
		var department, configData string
		if err := rows.Scan(&year, &configData); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

		var config YearlyConfig
		if err := json.Unmarshal([]byte(configData), &config); err != nil {
			return fmt.Errorf("error unmarshaling config for year %d: %v", year, err)
		}

		// Ensure all fields are properly initialized
		if config.RedemptionVariance == nil {
			config.RedemptionVariance = []int{1, 2, 4, 6}
		}
		if config.SectionMaxMarks == nil {
			config.SectionMaxMarks = []float64{5, 3, 2, 1}
		}

		calculateDependentValues(&config)

		// Recalculate dependent values
		config.SectionPoints = make([]int, len(config.SectionMaxMarks))
		for i := range config.SectionPoints {
			config.SectionPoints[i] = int(config.SectionMaxMarks[i] * float64(config.RedemptionRatio) * float64(config.RedemptionVariance[i]))
		}

		config.PointsPerMark = make([]float64, len(config.SectionPoints))
		for i := range config.PointsPerMark {
			config.PointsPerMark[i] = float64(config.SectionPoints[i]) / config.SectionMaxMarks[i]
		}

		config.MinPointsForHalf = make([]float64, len(config.PointsPerMark))
		for i := range config.MinPointsForHalf {
			config.MinPointsForHalf[i] = config.PointsPerMark[i] / 2
		}

		config.MaxPointsPerSubject = 0
		for _, points := range config.SectionPoints {
			config.MaxPointsPerSubject += points
		}

		if YearlyConfigs[department] == nil {
			YearlyConfigs[department] = make(map[int]*YearlyConfig)
		}

		YearlyConfigs[department][year] = &config
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error after iterating rows: %v", err)
	}

	return nil
}

func GetConfigForYearAndDepartment(year int, department string) *YearlyConfig {
	if deptConfigs, exists := YearlyConfigs[department]; exists {
		if config, exists := deptConfigs[year]; exists {
			// Return a copy to prevent modification of the original
			return &YearlyConfig{
				RedemptionRatio:     config.RedemptionRatio,
				SubjectsCount:       config.SubjectsCount,
				MaxPointsPerSubject: config.MaxPointsPerSubject,
				RedemptionVariance:  append([]int{}, config.RedemptionVariance...),
				SectionPoints:       append([]int{}, config.SectionPoints...),
				SectionMaxMarks:     append([]float64{}, config.SectionMaxMarks...),
				PointsPerMark:       append([]float64{}, config.PointsPerMark...),
				MinPointsForHalf:    append([]float64{}, config.MinPointsForHalf...),
				SubjectNames:        append([]string{}, config.SubjectNames...),
				Department:          config.Department,
			}
		}
	}

	// Return default config with calculated values if not found
	return createDefaultConfig(year, department)
}

func SaveConfigForYearAndDepartment(year int, department string, config *YearlyConfig) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	// Ensure department is set
	config.Department = department

	// Calculate dependent values to ensure consistency
	calculateDependentValues(config)

	jsonValue, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	// Save to config_year table
	_, err = tx.Exec(`
		INSERT INTO config_year (year, department, config_data) 
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE config_data = VALUES(config_data)
	`, year, department, string(jsonValue))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error saving config: %v", err)
	}

	// Save to main config table with department suffix
	if err := saveDepartmentConfigToMainTable(year, department, config); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	// Update in-memory config
	if YearlyConfigs[department] == nil {
		YearlyConfigs[department] = make(map[int]*YearlyConfig)
	}
	YearlyConfigs[department][year] = config

	return nil
}

func InitAdminDB() error {
	// Initialize admin configuration tables if needed
	// Add this after creating the config_year table
	years := []int{1, 2, 3, 4}
	for _, year := range years {
		var count int
		err := DB.QueryRow("SELECT COUNT(*) FROM config_year WHERE year = ?", year).Scan(&count)
		if err != nil {
			return fmt.Errorf("error checking for existing config for year %d: %v", year, err)
		}

		if count == 0 {
			defaultConfig := &YearlyConfig{
				RedemptionRatio:    10,
				SubjectsCount:      6,
				RedemptionVariance: []int{1, 2, 4, 6},
				SectionMaxMarks:    []float64{5, 3, 2, 1},
			}

			// Calculate dependent values
			defaultConfig.SectionPoints = make([]int, len(defaultConfig.SectionMaxMarks))
			for i := range defaultConfig.SectionPoints {
				defaultConfig.SectionPoints[i] = int(defaultConfig.SectionMaxMarks[i] * float64(defaultConfig.RedemptionRatio) * float64(defaultConfig.RedemptionVariance[i]))
			}

			defaultConfig.PointsPerMark = make([]float64, len(defaultConfig.SectionPoints))
			for i := range defaultConfig.PointsPerMark {
				defaultConfig.PointsPerMark[i] = float64(defaultConfig.SectionPoints[i]) / defaultConfig.SectionMaxMarks[i]
			}

			defaultConfig.MinPointsForHalf = make([]float64, len(defaultConfig.PointsPerMark))
			for i := range defaultConfig.MinPointsForHalf {
				defaultConfig.MinPointsForHalf[i] = defaultConfig.PointsPerMark[i] / 2
			}

			defaultConfig.MaxPointsPerSubject = 0
			for _, points := range defaultConfig.SectionPoints {
				defaultConfig.MaxPointsPerSubject += points
			}

			jsonValue, err := json.Marshal(defaultConfig)
			if err != nil {
				return fmt.Errorf("error marshaling default config for year %d: %v", year, err)
			}

			_, err = DB.Exec("INSERT INTO config_year (year, config_data) VALUES (?, ?)", year, string(jsonValue))
			if err != nil {
				return fmt.Errorf("error inserting default config for year %d: %v", year, err)
			}
		}
	}

	// Load all configs
	if err := LoadAllYearConfigs(); err != nil {
		return fmt.Errorf("failed to load configs from DB: %v", err)
	}
	return nil
}

func LoadConfigFromDB() error {
	// Load each configuration value from the database
	if err := loadConfigValue("RedemptionRatio", &RedemptionRatio); err != nil {
		return err
	}
	if err := loadConfigValue("SubjectsCount", &SubjectsCount); err != nil {
		return err
	}
	if err := loadConfigValue("MaxPointsPerSubject", &MaxPointsPerSubject); err != nil {
		return err
	}
	if err := loadConfigArray("RedemptionVariance", &RedemptionVariance); err != nil {
		return err
	}
	if err := loadConfigArray("SectionPoints", &SectionPoints); err != nil {
		return err
	}
	if err := loadConfigArray("SectionMaxMarks", &SectionMaxMarks); err != nil {
		return err
	}
	if err := loadConfigArray("PointsPerMark", &PointsPerMark); err != nil {
		return err
	}
	if err := loadConfigArray("MinPointsForHalf", &MinPointsForHalf); err != nil {
		return err
	}
	return nil
}

func loadConfigValue(key string, value interface{}) error {
	var configValue string
	err := DB.QueryRow("SELECT config_value FROM config WHERE config_key = ?", key).Scan(&configValue)
	if err != nil {
		return fmt.Errorf("error loading config value for %s: %v", key, err)
	}

	switch v := value.(type) {
	case *int:
		_, err := fmt.Sscanf(configValue, "%d", v)
		return err
	default:
		return json.Unmarshal([]byte(configValue), value)
	}
}

func loadConfigArray(key string, array interface{}) error {
	var configValue string
	err := DB.QueryRow("SELECT config_value FROM config WHERE config_key = ?", key).Scan(&configValue)
	if err != nil {
		return fmt.Errorf("error loading config array for %s: %v", key, err)
	}
	return json.Unmarshal([]byte(configValue), array)
}

func SaveConfigToDB() error {
	// Save each configuration value to the database
	if err := saveConfigValue("RedemptionRatio", RedemptionRatio); err != nil {
		return err
	}
	if err := saveConfigValue("SubjectsCount", SubjectsCount); err != nil {
		return err
	}
	if err := saveConfigValue("MaxPointsPerSubject", MaxPointsPerSubject); err != nil {
		return err
	}
	if err := saveConfigArray("RedemptionVariance", RedemptionVariance); err != nil {
		return err
	}
	if err := saveConfigArray("SectionPoints", SectionPoints); err != nil {
		return err
	}
	if err := saveConfigArray("SectionMaxMarks", SectionMaxMarks); err != nil {
		return err
	}
	if err := saveConfigArray("PointsPerMark", PointsPerMark); err != nil {
		return err
	}
	if err := saveConfigArray("MinPointsForHalf", MinPointsForHalf); err != nil {
		return err
	}
	return nil
}

func saveConfigValue(key string, value interface{}) error {
	var configValue string

	switch v := value.(type) {
	case int:
		configValue = fmt.Sprintf("%d", v)
	default:
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return err
		}
		configValue = string(jsonValue)
	}

	_, err := DB.Exec(`
        INSERT INTO config (config_key, config_value) 
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE config_value = VALUES(config_value)
    `, key, configValue)

	return err
}

func saveConfigArray(key string, array interface{}) error {
	jsonValue, err := json.Marshal(array)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
        INSERT INTO config (config_key, config_value) 
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE config_value = VALUES(config_value)
    `, key, string(jsonValue))

	return err
}

func LoadSubjectNames() error {
	rows, err := DB.Query("SELECT name FROM subjects")
	if err != nil {
		return err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		names = append(names, name)
	}

	if len(names) > 0 {
		// Update all configs with subject names
		for dept, yearConfigs := range YearlyConfigs {
			for year, config := range yearConfigs {
				config.SubjectNames = make([]string, len(names))
				copy(config.SubjectNames, names)
				YearlyConfigs[dept][year] = config
			}
		}
	}

	return nil
}

func GetConfigForYear(year int) *YearlyConfig {
	return GetConfigForYearAndDepartment(year, "CSE")
}
func SaveConfigForYear(year int, config *YearlyConfig) error {
	// Ensure we have all required fields
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	// Default to CSE department if not specified
	department := config.Department
	if department == "" {
		department = "CSE"
		config.Department = department
	}

	// Calculate dependent values to ensure consistency
	config.SectionPoints = make([]int, len(config.SectionMaxMarks))
	for i := range config.SectionPoints {
		config.SectionPoints[i] = int(config.SectionMaxMarks[i] * float64(config.RedemptionRatio) * float64(config.RedemptionVariance[i]))
	}

	config.PointsPerMark = make([]float64, len(config.SectionPoints))
	for i := range config.PointsPerMark {
		config.PointsPerMark[i] = float64(config.SectionPoints[i]) / config.SectionMaxMarks[i]
	}

	config.MinPointsForHalf = make([]float64, len(config.PointsPerMark))
	for i := range config.MinPointsForHalf {
		config.MinPointsForHalf[i] = config.PointsPerMark[i] / 2
	}

	config.MaxPointsPerSubject = 0
	for _, points := range config.SectionPoints {
		config.MaxPointsPerSubject += points
	}

	jsonValue, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	// Use a transaction to ensure data consistency
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	// Save to config_year table (existing logic)
	_, err = tx.Exec("DELETE FROM config_year WHERE year = ? AND department = ?", year, department)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting old config: %v", err)
	}

	_, err = tx.Exec(`
        INSERT INTO config_year (year, department, config_data) 
        VALUES (?, ?, ?)
    `, year, department, string(jsonValue))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting new config: %v", err)
	}

	// Save individual values to config table with year and department suffixes
	yearSuffix := getYearSuffix(year)
	deptSuffix := strings.ToUpper(department)

	// Save each config value with year and department suffix
	if err := saveConfigValueWithSuffix(tx, "RedemptionRatio", yearSuffix+deptSuffix, config.RedemptionRatio); err != nil {
		tx.Rollback()
		return err
	}
	if err := saveConfigValueWithSuffix(tx, "SubjectsCount", yearSuffix+deptSuffix, config.SubjectsCount); err != nil {
		tx.Rollback()
		return err
	}
	if err := saveConfigValueWithSuffix(tx, "MaxPointsPerSubject", yearSuffix+deptSuffix, config.MaxPointsPerSubject); err != nil {
		tx.Rollback()
		return err
	}
	if err := saveConfigArrayWithSuffix(tx, "RedemptionVariance", yearSuffix+deptSuffix, config.RedemptionVariance); err != nil {
		tx.Rollback()
		return err
	}
	if err := saveConfigArrayWithSuffix(tx, "SectionPoints", yearSuffix+deptSuffix, config.SectionPoints); err != nil {
		tx.Rollback()
		return err
	}
	if err := saveConfigArrayWithSuffix(tx, "SectionMaxMarks", yearSuffix+deptSuffix, config.SectionMaxMarks); err != nil {
		tx.Rollback()
		return err
	}
	if err := saveConfigArrayWithSuffix(tx, "PointsPerMark", yearSuffix+deptSuffix, config.PointsPerMark); err != nil {
		tx.Rollback()
		return err
	}
	if err := saveConfigArrayWithSuffix(tx, "MinPointsForHalf", yearSuffix+deptSuffix, config.MinPointsForHalf); err != nil {
		tx.Rollback()
		return err
	}

	// Additionally save the Year 1 values to the main config table without suffix
	// This is to maintain backward compatibility with the existing code
	if year == 1 && department == "CSE" {
		if err := saveConfigValueToMainTable(tx, "RedemptionRatio", config.RedemptionRatio); err != nil {
			tx.Rollback()
			return err
		}
		if err := saveConfigValueToMainTable(tx, "SubjectsCount", config.SubjectsCount); err != nil {
			tx.Rollback()
			return err
		}
		if err := saveConfigValueToMainTable(tx, "MaxPointsPerSubject", config.MaxPointsPerSubject); err != nil {
			tx.Rollback()
			return err
		}
		if err := saveConfigArrayToMainTable(tx, "RedemptionVariance", config.RedemptionVariance); err != nil {
			tx.Rollback()
			return err
		}
		if err := saveConfigArrayToMainTable(tx, "SectionPoints", config.SectionPoints); err != nil {
			tx.Rollback()
			return err
		}
		if err := saveConfigArrayToMainTable(tx, "SectionMaxMarks", config.SectionMaxMarks); err != nil {
			tx.Rollback()
			return err
		}
		if err := saveConfigArrayToMainTable(tx, "PointsPerMark", config.PointsPerMark); err != nil {
			tx.Rollback()
			return err
		}
		if err := saveConfigArrayToMainTable(tx, "MinPointsForHalf", config.MinPointsForHalf); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	// Update in-memory config
	if YearlyConfigs[department] == nil {
		YearlyConfigs[department] = make(map[int]*YearlyConfig)
	}
	YearlyConfigs[department][year] = config

	return nil
}

// Helper function to save config values to the main config table without suffix
func saveConfigValueToMainTable(tx *sql.Tx, key string, value interface{}) error {
	var configValue string

	switch v := value.(type) {
	case int:
		configValue = fmt.Sprintf("%d", v)
	default:
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return err
		}
		configValue = string(jsonValue)
	}

	_, err := tx.Exec(`
        INSERT INTO config (config_key, config_value) 
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE config_value = VALUES(config_value)
    `, key, configValue)

	return err
}

// Helper function to save config arrays to the main config table without suffix
func saveConfigArrayToMainTable(tx *sql.Tx, key string, array interface{}) error {
	jsonValue, err := json.Marshal(array)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
        INSERT INTO config (config_key, config_value) 
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE config_value = VALUES(config_value)
    `, key, string(jsonValue))

	return err
}

func saveConfigValueWithSuffix(tx *sql.Tx, key, suffix string, value interface{}) error {
	fullKey := fmt.Sprintf("%s%s", key, suffix)
	var configValue string

	switch v := value.(type) {
	case int:
		configValue = fmt.Sprintf("%d", v)
	default:
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return err
		}
		configValue = string(jsonValue)
	}

	_, err := tx.Exec(`
        INSERT INTO config (config_key, config_value) 
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE config_value = VALUES(config_value)
    `, fullKey, configValue)

	return err
}

func saveConfigArrayWithSuffix(tx *sql.Tx, key, suffix string, array interface{}) error {
	fullKey := fmt.Sprintf("%s%s", key, suffix)
	jsonValue, err := json.Marshal(array)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
        INSERT INTO config (config_key, config_value) 
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE config_value = VALUES(config_value)
    `, fullKey, string(jsonValue))

	return err
}
