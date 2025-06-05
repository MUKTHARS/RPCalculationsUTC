package models

type YearlyConfig struct {
    RedemptionRatio     int       `json:"redemptionRatio"`
    SubjectsCount       int       `json:"subjectsCount"`
    MaxPointsPerSubject int       `json:"maxPointsPerSubject"`
    RedemptionVariance  []int     `json:"redemptionVariance"`
    SectionPoints       []int     `json:"sectionPoints"`
    SectionMaxMarks     []float64 `json:"sectionMaxMarks"`
    PointsPerMark       []float64 `json:"pointsPerMark"`
    MinPointsForHalf    []float64 `json:"minPointsForHalf"`
}

type YearlyStudentAllocation struct {
    Year                int                 `json:"year"`
    TotalPoints         int                 `json:"totalPoints"`
    Subjects            []SubjectAllocation `json:"subjects"`
    TotalMarks          float64             `json:"totalMarks"`
    PointsRemaining     int                 `json:"pointsRemaining"`
    SubmissionCount     int                 `json:"submissionCount"`
}