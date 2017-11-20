package course

import (
	"database/sql"
)

const (
	SheduleStatusAssistant = 2
	SheduleStatusPraktikan = 1
)

type readParams struct {
	Page  string
	Total string
}

type readArgs struct {
	Page  uint16
	Total uint16
}

type searchParams struct {
	Text string
}

type searchArgs struct {
	Text string
}

type searchResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UCU         int8   `json:"ucu"`
}

type readResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Class      string `json:"class"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Day        string `json:"day"`
	Status     string `json:"status"`
	ScheduleID int64  `json:"schedule_id"`
}

type readDetailParams struct {
	ScheduleID string
}

type readDetailArgs struct {
	ScheduleID int64
}

type readDetailResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UCU         int8   `json:"ucu"`
	Status      string `json:"status"`
	Semester    int8   `json:"semester"`
	Year        int16  `json:"year"`
	StartTime   uint16 `json:"start_time"`
	EndTime     uint16 `json:"end_time"`
	Class       string `json:"class"`
	Day         string `json:"day"`
	PlaceID     string `json:"place_id"`
	ScheduleID  int64  `json:"schedule_id"`
}

type listParameterResponse struct {
	Name string `json:"id"`
	Text string `json:"text"`
}

type gradeParameter struct {
	Type         string  `json:"type"`
	Percentage   float32 `json:"percentage"`
	StatusChange uint8   `json:"status_change"`
}

type createParams struct {
	ID             string
	Name           string
	Description    string
	UCU            string
	Semester       string
	Year           string
	StartTime      string
	EndTime        string
	Class          string
	Day            string
	PlaceID        string
	IsUpdate       string
	GradeParameter string
}

type createArgs struct {
	ID             string
	Name           string
	Description    sql.NullString
	UCU            int8
	Semester       int8
	Year           int16
	StartTime      int16
	EndTime        int16
	Class          string
	Day            int8
	PlaceID        string
	IsUpdate       bool
	GradeParameter []gradeParameter
}

type updateParams struct {
	ID             string
	Name           string
	Description    string
	UCU            string
	ScheduleID     string
	Status         string
	Semester       string
	Year           string
	StartTime      string
	EndTime        string
	Class          string
	Day            string
	PlaceID        string
	IsUpdate       string
	GradeParameter string
}

type updateArgs struct {
	ID             string
	Name           string
	Description    sql.NullString
	UCU            int8
	ScheduleID     int64
	Status         int8
	Semester       int8
	Year           int16
	StartTime      int16
	EndTime        int16
	Class          string
	Day            int8
	PlaceID        string
	IsUpdate       bool
	GradeParameter []gradeParameter
}

type summaryResponse struct {
	Status string           `json:"status"`
	Course []courseResponse `json:"courses"`
}

type getParams struct {
	Payload string
}

type getArgs struct {
	Payload string
}

type getResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Class       string `json:"class"`
	Semester    int8   `json:"semester"`
}

type getAssistantParams struct {
	ScheduleID string
}

type getAssistantArgs struct {
	ScheduleID int64
}

type getAssistantResponse struct {
	Name  string `json:"name"`
	Roles string `json:"role"`
	Phone string `json:"phone_number"`
	Email string `json:"email"`
}

type courseResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	UCU      int8   `json:"ucu"`
	Semester int8   `json:"semester"`
}

type deleteScheduleParams struct {
	ScheduleID string
}

type deleteScheduleArgs struct {
	ScheduleID int64
}

type readScheduleParameterParams struct {
	ScheduleID string
}

type readScheduleParameterArgs struct {
	ScheduleID int64
}

type readScheduleParameterResponse struct {
	Type         string  `json:"type"`
	Percentage   float32 `json:"percentage"`
	StatusChange uint8   `json:"status_change"`
}

type listStudentParams struct {
	scheduleID string
}

type listStudentArgs struct {
	scheduleID int64
}

type listStudentResponse struct {
	UserIdentityCode int64  `json:"id"`
	UserName         string `json:"name"`
}
type gradeParameterResponse struct {
	ID         int64   `json:"id"`
	Type       string  `json:"type"`
	Percentage float32 `json:"percentage"`
	Nilai      float32 `json:"nilai"`
}
type scheduleGrade struct {
	ScheduleID int64                    `json:"schedule_id"`
	Total      float32                  `json:"total"`
	Grade      []gradeParameterResponse `json:"grade"`
}
type responseGradeSummery struct {
	UsersID  int64           `json:"npm"`
	Schedule []scheduleGrade `json:"schedules"`
}
