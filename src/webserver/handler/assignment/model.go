package assignment

// type summaryResponse struct {
// 	ID     int64  `json:"id"`
// 	Name   string `json:"name"`
// 	Status int8   `json:"status,omitempty"`
// }

// type profileSummaryResponse struct {
// 	CourseName string `json:"course_name"`
// 	Complete   int8   `json:"complete"`
// 	Incomplete int8   `json:"incomplete"`
// }

type createParams struct {
	ID             string
	GradeParameter string
	Status         string
	Description    string
	DueDate        string
}

type createArgs struct {
	ID             string
	GradeParameter int64
	Status         string
	Description    string
	DueDate        string
}
