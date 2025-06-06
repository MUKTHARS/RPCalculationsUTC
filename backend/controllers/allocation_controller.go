package controllers

import (
	"encoding/json"
	"fmt"

	//  "log"
	"math"
	"net/http"
	"strconv"

	"Rp/backend/config"
	"Rp/backend/models"
)

func AllocatePoints(totalPoints int, department string, year int) models.StudentAllocation {

	configData := config.GetConfigForYearAndDepartment(year, department)
	var allocation models.StudentAllocation
	allocation.TotalPoints = totalPoints
	allocation.Department = department
	allocation.Year = year

	subjectsCount := configData.SubjectsCount
	allocation.Subjects = make([]models.SubjectAllocation, subjectsCount)
	for i := range allocation.Subjects {
		allocation.Subjects[i] = models.SubjectAllocation{
			SectionPoints:     []int{0, 0, 0, 0},
			SectionMarks:      []float64{0, 0, 0, 0},
			TotalSubjectMarks: 0,
		}
	}

	redemptionRatio := config.RedemptionRatio
	redemptionVariance := config.RedemptionVariance
	sectionMaxMarks := config.SectionMaxMarks
	sectionsCount := len(sectionMaxMarks)

	maxPossible := 0
	for i := 0; i < sectionsCount; i++ {
		maxPossible += int(sectionMaxMarks[i]*float64(redemptionRatio)*float64(redemptionVariance[i])) * subjectsCount
	}

	pointsLeft := totalPoints
	if totalPoints > maxPossible {
		pointsLeft = maxPossible
		allocation.PointsRemaining = totalPoints - maxPossible
		fmt.Printf("Total points exceed max possible. Capped to %d. Extra: %d\n", maxPossible, allocation.PointsRemaining)
	} else {
		fmt.Printf("Total points within limit. Proceeding with %d points.\n", totalPoints)
	}

	for sectionIdx := 0; sectionIdx < sectionsCount; sectionIdx++ {
		sectionMaxPoints := int(sectionMaxMarks[sectionIdx] * float64(redemptionRatio) * float64(redemptionVariance[sectionIdx]))
		minPointsForHalf := int(float64(redemptionRatio) * float64(redemptionVariance[sectionIdx]) / 2)

		fmt.Printf("Section %d max points: %d\n", sectionIdx, sectionMaxPoints)

		for subjIdx := 0; subjIdx < subjectsCount; subjIdx++ {
			if pointsLeft <= 0 {
				fmt.Println("No points left to allocate.")
				break
			}

			alloc := 0
			if pointsLeft >= sectionMaxPoints {
				alloc = sectionMaxPoints
			} else {
				pointsPerMark := float64(redemptionRatio) * float64(redemptionVariance[sectionIdx])
				fullMarks := math.Floor(float64(pointsLeft) / pointsPerMark)
				fullPoints := int(fullMarks * pointsPerMark)
				remainingAfterFull := pointsLeft - fullPoints

				if remainingAfterFull >= minPointsForHalf {
					alloc = fullPoints + minPointsForHalf
				} else {
					alloc = fullPoints
				}
			}

			if alloc > 0 {
				allocation.Subjects[subjIdx].SectionPoints[sectionIdx] = alloc
				pointsLeft -= alloc
				fmt.Printf("Allocated %d points to subject %d section %d. Points left: %d\n", alloc, subjIdx, sectionIdx, pointsLeft)
			}
		}
	}

	for subjIdx := 0; subjIdx < subjectsCount; subjIdx++ {
		subjectTotal := 0.0
		for sectionIdx := 0; sectionIdx < sectionsCount; sectionIdx++ {
			points := allocation.Subjects[subjIdx].SectionPoints[sectionIdx]
			if points == 0 {
				continue
			}

			pointsPerMark := float64(redemptionRatio) * float64(redemptionVariance[sectionIdx])
			minForHalf := pointsPerMark / 2

			fullMarks := math.Floor(float64(points) / pointsPerMark)
			remainingPoints := float64(points) - (fullMarks * pointsPerMark)

			var halfMark float64
			if remainingPoints >= minForHalf {
				halfMark = 0.5
			}

			marks := math.Min(fullMarks+halfMark, sectionMaxMarks[sectionIdx])
			allocation.Subjects[subjIdx].SectionMarks[sectionIdx] = marks
			subjectTotal += marks
			allocation.TotalMarks += marks

			fmt.Printf("Subject %d section %d: %d points => %.1f marks (Full: %.1f, Half: %.1f)\n",
				subjIdx, sectionIdx, points, marks, fullMarks, halfMark)
		}
		allocation.Subjects[subjIdx].TotalSubjectMarks = subjectTotal
		fmt.Printf("Total marks for subject %d: %.1f\n", subjIdx, subjectTotal)
	}

	if totalPoints <= maxPossible {
		allocation.PointsRemaining = pointsLeft
	}

	fmt.Printf("Final points remaining: %d\n", allocation.PointsRemaining)
	fmt.Printf("Total marks allocated: %.1f\n", allocation.TotalMarks)

	return allocation
}

// CalculateSubjectAllocation calculates marks for a single subject given points
// This is used by the UI component for real-time calculation
func CalculateSubjectAllocation(points int, department string, year int) models.SubjectResponse {
	configData := config.GetConfigForYearAndDepartment(year, department)
	response := models.SubjectResponse{
		Success: true,
	}

	sectionPoints := make([]int, len(configData.SectionPoints))

	pointsRemaining := points

	for sectionIdx := 0; sectionIdx < len(configData.SectionPoints) && pointsRemaining > 0; sectionIdx++ {
		sectionMax := configData.SectionPoints[sectionIdx]

		if pointsRemaining >= sectionMax {
			sectionPoints[sectionIdx] = sectionMax
			pointsRemaining -= sectionMax
		} else {
			sectionPoints[sectionIdx] = pointsRemaining
			pointsRemaining = 0
		}
	}

	totalMarks := 0.0

	for sectionIdx := 0; sectionIdx < len(configData.SectionPoints); sectionIdx++ {
		points := sectionPoints[sectionIdx]
		if points == 0 {
			continue
		}

		pointsPerMark := configData.PointsPerMark[sectionIdx]
		minForHalf := configData.MinPointsForHalf[sectionIdx]

		fullMarks := math.Floor(float64(points) / pointsPerMark)
		remainingPoints := float64(points) - (fullMarks * pointsPerMark)

		var halfMark float64
		if remainingPoints >= minForHalf {
			halfMark = 0.5
		}

		marks := math.Min(fullMarks+halfMark, configData.SectionMaxMarks[sectionIdx])
		totalMarks += marks
	}

	response.Marks = totalMarks
	response.SectionPoints = sectionPoints

	return response
}

func PrintAllocation(allocation models.StudentAllocation) {

	configData := config.GetConfigForYearAndDepartment(allocation.Year, allocation.Department)
	fmt.Printf("\nStudent with %d reward points (Year %d, %s):\n",
		allocation.TotalPoints, allocation.Year, allocation.Department)
	fmt.Println("Subject Breakdown:")

	for i := 0; i < configData.SubjectsCount; i++ {
		subj := allocation.Subjects[i]
		subjectName := "Subject " + strconv.Itoa(i+1)
		if i < len(configData.SubjectNames) {
			subjectName = configData.SubjectNames[i]
		}

		fmt.Printf("%s: Points [%3d, %3d, %3d, %3d] | Marks [%3.1f, %3.1f, %3.1f, %3.1f] | Total Marks: %3.1f\n",
			subjectName,
			subj.SectionPoints[0], subj.SectionPoints[1], subj.SectionPoints[2], subj.SectionPoints[3],
			subj.SectionMarks[0], subj.SectionMarks[1], subj.SectionMarks[2], subj.SectionMarks[3],
			subj.TotalSubjectMarks)
	}

	fmt.Printf("\nTotal Marks Across All Subjects: %.1f\n", allocation.TotalMarks)
	fmt.Printf("Points Remaining: %d\n", allocation.PointsRemaining)
	fmt.Println("----------------------------------------")
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
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
		Points     int    `json:"points"`
		Department string `json:"department"`
		Year       int    `json:"year"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	points := requestData.Points
	dept := requestData.Department
	if dept == "" {
		dept = "CSE"
	}
	year := requestData.Year
	if year == 0 {
		year = 1
	}

	if points < 0 {
		SendErrorResponse(w, "Points cannot be negative", http.StatusBadRequest)
		return
	}

	allocation := AllocatePoints(points, dept, year)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ApiResponse{
		Success: true,
		Data:    allocation,
	})
}

func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.ApiResponse{
		Success: false,
		Error:   message,
	}

	json.NewEncoder(w).Encode(response)
}

func GetDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get departments from config or database
	depts, err := config.GetAllDepartments()
	if err != nil {
		SendErrorResponse(w, fmt.Sprintf("Failed to get departments: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert to the expected format
	departments := make([]map[string]string, 0)
	for _, dept := range depts {
		departments = append(departments, map[string]string{
			"code": dept["code"],
			"name": dept["name"],
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"departments": departments,
	})
}

func AddDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var requestData struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if requestData.Code == "" || requestData.Name == "" {
		SendErrorResponse(w, "Department code and name are required", http.StatusBadRequest)
		return
	}

	if err := config.AddDepartment(requestData.Code, requestData.Name); err != nil {
		SendErrorResponse(w, "Failed to add department", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func DeleteDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var requestData struct {
		Code string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if requestData.Code == "" {
		SendErrorResponse(w, "Department code is required", http.StatusBadRequest)
		return
	}

	if err := config.DeleteDepartment(requestData.Code); err != nil {
		SendErrorResponse(w, "Failed to delete department", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
