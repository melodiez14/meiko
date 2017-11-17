package tutorial

type readParams struct {
	scheduleID string
	page       string
	total      string
}

type readArgs struct {
	scheduleID int64
	page       uint64
	total      uint64
}

type readResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"file_url"`
	Time        string `json:"time"`
}

type readDetailParams struct {
	id string
}

type readDetailArgs struct {
	id int64
}

type readDetailResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"file_url"`
	Time        string `json:"time"`
}
