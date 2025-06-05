package models

type Subject struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
}

// SubjectPoints represents the points and marks allocated to a student for a subject
type SubjectPoints struct {
	StudentID   string  `json:"studentId"`
	SubjectCode string  `json:"subjectCode"`
	SubjectName string  `json:"subjectName"`
	Points      int     `json:"points"`
	Marks       float64 `json:"marks"`
}



// SubjectResponse represents the response for a single subject calculation
type SubjectResponse struct {
	Success      bool    `json:"success"`
	Marks        float64 `json:"marks"`
	SectionPoints []int  `json:"sectionPoints"`
	Error        string  `json:"error,omitempty"`
}