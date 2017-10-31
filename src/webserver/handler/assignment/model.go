package assignment

type summaryResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status int8   `json:"status,omitempty"`
}

type profileSummaryResponse struct {
	CourseName string `json:"course_name"`
	Complete   int8   `json:"complete"`
	Incomplete int8   `json:"incomplete"`
}

type createParams struct {
	Subject     string
	Description string
	DueDate     string
	Attachment  string
}

type createArgs struct {
	Subject     string
	Description string
	DueDate     string
	Attachment  string
}
