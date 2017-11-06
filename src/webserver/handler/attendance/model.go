package attendance

type summaryResponse struct {
	Course     string `json:"status"`
	Percentage string `json:"percentage"`
}

type listStudentParams struct {
	meetingNumber string
	scheduleID    string
}

type listStudentArgs struct {
	meetingNumber uint8
	scheduleID    int64
}

type listStudentResponse struct {
	IdentityCode int64  `json:"id"`
	StudentName  string `json:"name"`
	Status       string `json:"status"`
}
