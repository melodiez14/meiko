package tutorial

import "database/sql"

type readParams struct {
	payload    string
	scheduleID string
	page       string
	total      string
}

type readArgs struct {
	payload    string
	scheduleID int64
	page       uint64
	total      uint64
}

type readResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"file_url"`
	Time        int64  `json:"time"`
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
	Time        int64  `json:"time"`
}

type createParams struct {
	name        string
	description string
	fileID      string
	scheduleID  string
}

type createArgs struct {
	name        string
	description sql.NullString
	fileID      string
	scheduleID  int64
}

type deleteParams struct {
	id string
}

type deleteArgs struct {
	id int64
}

type updateParams struct {
	id          string
	name        string
	description string
	fileID      string
	scheduleID  string
}

type updateArgs struct {
	id          int64
	name        string
	description sql.NullString
	fileID      string
	scheduleID  int64
}
