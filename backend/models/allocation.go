package models

type SubjectAllocation struct {
	SubjectCode       string    `json:"subjectCode"`
	SubjectName       string    `json:"subjectName"`
	SectionPoints     []int     `json:"sectionPoints"`
	SectionMarks      []float64 `json:"sectionMarks"`
	TotalSubjectMarks float64   `json:"totalSubjectMarks"`
}

type StudentAllocation struct {
	TotalPoints     int                  `json:"totalPoints"`
	Year            int                  `json:"year"`
	Subjects        []SubjectAllocation `json:"subjects"`
	TotalMarks      float64              `json:"totalMarks"`
	PointsRemaining int                  `json:"pointsRemaining"`
	Department string `json:"department"`
	StudentID string `json:"studentID"`
}
type AllocationResponse struct {
	Success bool              `json:"success"`
	Data    StudentAllocation `json:"data"`
	Error   string            `json:"error,omitempty"`
}
type ApiResponse struct {
	Success bool              `json:"success"`
	Data    StudentAllocation `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
}

type SubjectsResponse struct {
	Success  bool     `json:"success"`
	Subjects []string `json:"subjects"`
	Error    string   `json:"error,omitempty"`
}

type SaveAllSubjectsRequest struct {
	StudentID   string        `json:"studentId"`
	Subjects    []SubjectData `json:"subjects"`
	TotalPoints int           `json:"totalPoints"`
	TotalMarks  float64       `json:"totalMarks"`
	PointsRemaining int  `json:"pointsRemaining"`
}

type SubjectData struct {
	SubjectCode string  `json:"subjectCode"`
	SubjectName string  `json:"subjectName"`
	Points      int     `json:"points"`
	Marks       float64 `json:"marks"`
}

type SubjectCalculationRequest struct {
	Points int `json:"points"`
}

type SubjectCalculationResponse struct {
	Success       bool    `json:"success"`
	Points        int     `json:"points"`
	SectionPoints [4]int  `json:"sectionPoints"`
	Marks         float64 `json:"marks"`
	Error         string  `json:"error,omitempty"`
}

// Update your AllocationResult struct to include department info
type AllocationResult struct {
	TotalPoints      int                 `json:"totalPoints"`
	TotalMarks       float64             `json:"totalMarks"`
	PointsRemaining  int                 `json:"pointsRemaining"`
	Subjects         []SubjectAllocation `json:"subjects"`
	Department       string              `json:"department,omitempty"`
	Year             int                 `json:"year,omitempty"`
}