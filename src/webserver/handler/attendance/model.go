package attendance

import (
	"database/sql"
	"time"
)

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

type readMeetingParams struct {
	scheduleID string
	page       string
	total      string
}

type readMeetingArgs struct {
	scheduleID int64
	page       int
	total      int
}

type readMeetings struct {
	ID             uint64 `json:"id"`
	Subject        string `json:"subject"`
	MeetingNumber  uint8  `json:"number"`
	Date           int64  `json:"date"`
	TotalAttendant uint16 `json:"total_attendant"`
	TotalStudent   uint16 `json:"total_student"`
}

type readMeetingResponse struct {
	TotalPage int            `json:"total_page"`
	Page      int            `json:"page"`
	Meetings  []readMeetings `json:"meetings"`
}

type readMeetingDetailParams struct {
	meetingID string
}

type readMeetingDetailArgs struct {
	meetingID uint64
}

type student struct {
	IdentityCode int64  `json:"identity_code"`
	Name         string `json:"name"`
	Status       string `json:"status"`
}

type readMeetingDetailResponse struct {
	ID            uint64    `json:"meeting_id"`
	Subject       string    `json:"subject"`
	MeetingNumber uint8     `json:"meeting_number"`
	Description   string    `json:"desciption"`
	Date          int64     `json:"date"`
	Student       []student `json:"student"`
}

type createMeetingParams struct {
	subject       string
	meetingNumber string
	scheduleID    string
	description   string
	date          string
	users         string
}

type createMeetingArgs struct {
	subject           string
	meetingNumber     uint8
	scheduleID        int64
	description       sql.NullString
	date              time.Time
	userIdentityCodes []int64
}

type updateMeetingParams struct {
	id            string
	subject       string
	meetingNumber string
	scheduleID    string
	description   string
	date          string
	isForceUpdate string
	users         string
}

type updateMeetingArgs struct {
	id                uint64
	subject           string
	meetingNumber     uint8
	scheduleID        int64
	description       sql.NullString
	date              time.Time
	isForceUpdate     bool
	userIdentityCodes []int64
}

type deleteMeetingParams struct {
	id            string
	isForceDelete string
}

type deleteMeetingArgs struct {
	id            uint64
	isForceDelete bool
}

type getAttendanceParams struct {
	scheduleID string
}

type getAttendanceArgs struct {
	scheduleID int64
}

type getAttendanceResponse struct {
	Percentage   string `json:"percentage"`
	Absent       int    `json:"absent"`
	Present      int    `json:"present"`
	TotalMeeting int    `json:"total_meeting"`
}
