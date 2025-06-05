package models

type Student struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type StudentReward struct {
	ID              int     `json:"id"`
	StudentID       string  `json:"studentId"`
	Points          int     `json:"points"`
	TotalMarks      float64 `json:"totalMarks"`
	PointsRemaining int     `json:"pointsRemaining"`
	CreatedAt       string  `json:"createdAt"`
}

type StudentTotal struct {
	StudentID   string  `json:"studentId"`
	TotalPoints int     `json:"totalPoints"`
	TotalMarks  float64 `json:"totalMarks"`
}
