package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"Rp/backend/config"
	"Rp/backend/models"
)

func GetSubjectsHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := config.LoadSubjectNames(); err != nil {
		log.Printf("Error loading subject names: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	var subjectNames []string
rows, err := config.DB.Query("SELECT name FROM subjects")
if err == nil {
    defer rows.Close()
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err == nil {
            subjectNames = append(subjectNames, name)
        }
    }
}

response := models.SubjectsResponse{
    Success:  true,
    Subjects: subjectNames,
}

	json.NewEncoder(w).Encode(response)
}

func CalculateSubjectMarksHandler(w http.ResponseWriter, r *http.Request) {
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

	var requestData models.SubjectCalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if requestData.Points < 0 {
		SendErrorResponse(w, "Points cannot be negative", http.StatusBadRequest)
		return
	}

	if requestData.Points > config.MaxPointsPerSubject {
		SendErrorResponse(w, fmt.Sprintf("Points cannot exceed %d", config.MaxPointsPerSubject), http.StatusBadRequest)
		return
	}

	sectionPoints, totalMarks := calculateSubjectMarks(requestData.Points)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SubjectCalculationResponse{
		Success:       true,
		Points:        requestData.Points,
		SectionPoints: sectionPoints,
		Marks:         totalMarks,
	})
}

func calculateSubjectMarks(points int) ([4]int, float64) {
	var sectionPoints [4]int
	var totalMarks float64
	remainingPoints := points

	for sectionIdx := 0; sectionIdx < len(config.SectionPoints); sectionIdx++ {
		sectionMax := config.SectionPoints[sectionIdx]
		minForHalf := config.MinPointsForHalf[sectionIdx]
		pointsPerMark := config.PointsPerMark[sectionIdx]

		if remainingPoints <= 0 {
			break
		}

		alloc := 0
		if remainingPoints >= sectionMax {
			alloc = sectionMax
		} else {
			fullMarks := math.Floor(float64(remainingPoints) / pointsPerMark)
			fullPoints := int(fullMarks * pointsPerMark)
			remainingAfterFull := remainingPoints - fullPoints

			if float64(remainingAfterFull) >= minForHalf {
				alloc = fullPoints + int(minForHalf)
			} else {
				alloc = fullPoints
			}
		}

		if alloc > 0 {
			sectionPoints[sectionIdx] = alloc
			remainingPoints -= alloc

			fullMarks := math.Floor(float64(alloc) / pointsPerMark)
			remainingPointsForHalf := float64(alloc) - (fullMarks * pointsPerMark)

			var halfMark float64
			if remainingPointsForHalf >= minForHalf {
				halfMark = 0.5
			}

			marks := fullMarks + halfMark
			if marks > config.SectionMaxMarks[sectionIdx] {
				marks = config.SectionMaxMarks[sectionIdx]
			}
			totalMarks += marks
		}
	}

	return sectionPoints, totalMarks
}