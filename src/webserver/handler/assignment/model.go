package assignment

import (
	"database/sql"
	"time"

	fs "github.com/melodiez14/meiko/src/module/file"
)

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
const (
	TableNameAssignments     = "assignments"
	TableNameUserAssignments = "p_users_assignments"
	UserMustUpload           = 0
	UserMustNotUpload        = 1
)

type createParams struct {
	FilesID           string
	GradeParametersID string
	Name              string
	Description       string
	Status            string
	DueDate           string
	Size              string
	Type              string
}

type createArgs struct {
	FilesID           string
	GradeParametersID int64
	Name              string
	Description       sql.NullString
	Status            string
	DueDate           string
	Size              int64
	Type              string
}
type updatePrams struct {
	ID                string
	FilesID           string
	GradeParametersID string
	Name              string
	Description       string
	Status            string
	DueDate           string
}
type updateArgs struct {
	ID                int64
	FilesID           string
	GradeParametersID int64
	Name              string
	Description       sql.NullString
	Status            string
	DueDate           string
	TableID           int64
}

type readParams struct {
	Page  string
	Total string
}

type readArgs struct {
	Page  uint16
	Total uint16
}

type readResponse struct {
	Name             string         `db:"name"`
	Description      sql.NullString `db:"description"`
	Status           string         `db:"status"`
	GradeParameterID int32          `db:"grade_parameters_id"`
	DueDate          time.Time      `db:"due_date"`
}

type detailParams struct {
	IdentityCode string
}

type detailArgs struct {
	IdentityCode int64
}

type detailResponse struct {
	ID               int64  `json:"id"`
	Status           int8   `json:"status"`
	Name             string `json:"name"`
	GradeParameterID int32  `json:"grade_parameters_id"`
	Description      string `json:"description"`
	DueDate          string `json:"due_date"`
	FilesID          string `json:"files_id"`
	FilesName        string `json:"files_name"`
	Mime             string `json:"mime"`
	Type             string `json:"type"`
}
type detailResponseUser struct {
	ID               int64  `json:"id"`
	Status           string `json:"status"`
	Name             string `json:"name"`
	GradeParameterID int32  `json:"grade_parameters_id"`
	Description      string `json:"description"`
	DueDate          string `json:"due_date"`
	Score            string `json:"score"`
	FilesName        string `json:"file_name"`
	UploadedStatus   bool   `json:"uploaded_status"`
	MustUploadStatus int8   `json:"must_upload"`
	ButtonType       string `json:"button_type"`
	AssistantFile    []file `json:"assignment_file"`
	StudentFile      []file `json:"uploaded_file"`
}

type file struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type uploadAssignmentParams struct {
	UserID       int64
	AssignmentID string
	Description  string
	FileID       string
}
type uploadAssignmentArgs struct {
	UserID       int64
	AssignmentID int64
	Description  sql.NullString
	FileID       []string
}
type readUploadedAssignmentParams struct {
	UserID       string
	Page         string
	Total        string
	ScheduleID   string
	AssignmentID string
	Name         string
	Description  string
	Score        string
	DueDate      string
	PathFile     string
}
type readUploadedAssignmentArgs struct {
	UserID       int64
	Page         int64
	Total        int64
	ScheduleID   int64
	AssignmentID int64
	Name         string
	Description  sql.NullString
	Score        string
	DueDate      string
	PathFile     []fs.File
}
type readUploadedDetailParams struct {
	UserID       string
	ScheduleID   string
	AssignmentID string
	Name         string
	Description  string
	Score        string
	DueDate      string
	PathFile     string
}
type readUploadedDetailArgs struct {
	UserID       int64
	ScheduleID   int64
	AssignmentID int64
	Name         string
	Description  sql.NullString
	Score        string
	DueDate      string
	PathFile     []fs.File
}

type deleteParams struct {
	ID string
}
type deleteArgs struct {
	ID int64
}
type listAssignmentsParams struct {
	Page         string
	Total        string
	ScheduleID   string
	AssignmentID string
	DueDate      string
	Name         string
	Description  string
}
type listAssignmentsArgs struct {
	Page         uint16
	Total        uint16
	ScheduleID   int64
	AssignmentID int64
	DueDate      string
	Name         string
	Description  string
}
type readDetailParam struct {
	AssignmentID string
}
type readDetailArgs struct {
	AssignmentID int64
}

type updateScoreParams struct {
	Score        string
	UserID       string
	ScheduleID   string
	AssignmentID string
}
type updateScoreArgs struct {
	Score        float32
	UserID       int64
	ScheduleID   int64
	AssignmentID int64
}
type detailAssignmentParams struct {
	ScheduleID   string
	AssignmentID string
}
type detailAssignmentArgs struct {
	ScheduleID   int64
	AssignmentID int64
}

type userAssignment struct {
	UserID int64   `json:"user_id"`
	Name   string  `json:"name"`
	Grade  float32 `json:"grade"`
}

type detailAssignmentResponse struct {
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	DueDate       time.Time      `json:"due_date"`
	IsCreateScore bool           `json:"is_create_score"`
	Praktikan     []userAssignment
}
type createScoreParams struct {
	ScheduleID   string
	AssignmentID string
	Name         string
	Description  string
	Users        string
}
type createScoreArgs struct {
	ScheduleID   int64
	AssignmentID int64
	Name         string
	Description  sql.NullString
	IdentityCode []int64
	Score        []float32
}
type student struct {
	IdentityCode int64   `json:"identity_code"`
	Name         string  `json:"name"`
	Score        float32 `json:"score"`
}
type listAssignmentResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Submitted   bool   `json:"submitted"`
}
type scoreParams struct {
	ScheduleID string
}
type scoreArgs struct {
	ScheduleID int64
}
type responseScoreSchedule struct {
	ScheduleID int64  `json:"schedule_id"`
	CourseName string `json:"course_name"`
	Attendance string `json:"attendance"`
	Assignment string `json:"assignment"`
	Quiz       string `json:"quiz"`
	Mid        string `json:"mid"`
	Final      string `json:"final"`
	Total      string `json:"total"`
}
