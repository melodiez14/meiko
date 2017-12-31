package assignment

import (
	"database/sql"
	"time"

	fs "github.com/melodiez14/meiko/src/module/file"
)

type getParams struct {
	scheduleID string
	filter     string
}

type getArgs struct {
	scheduleID sql.NullInt64
	filter     sql.NullString
}

type getResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	IsAllowUpload bool   `json:"is_allow_upload"`
	Description   string `json:"description"`
	DueDate       string `json:"due_date"`
	Score         string `json:"score"`
	UpdatedAt     string `json:"updated_at"`
}

type getDetailParams struct {
	id string
}

type getDetailArgs struct {
	id int64
}

type getDetailResponse struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	DueDate              string `json:"due_date"`
	Score                string `json:"score"`
	Status               string `json:"status"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	AssignmentFile       []file `json:"assignment_file"`
	IsAllowUpload        bool   `json:"is_allow_upload"`
	SubmittedDescription string `json:"submitted_description"`
	SubmittedFile        []file `json:"submitted_file"`
	SubmittedDate        string `json:"submitted_date"`
}

type submitParams struct {
	id          string
	description string
	fileID      string
}

type submitArgs struct {
	id          int64
	description sql.NullString
	fileID      []string
}

type readParams struct {
	scheduleID string
	page       string
	total      string
}

type readArgs struct {
	scheduleID int64
	page       int
	total      int
}

type readResponse struct {
	Page        int    `json:"page"`
	TotalPage   int    `json:"total_page"`
	Assignments []read `json:"assignments"`
}

type read struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	DueDate   string `json:"due_date"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
}

type createParams struct {
	filesID          string
	gpID             string
	name             string
	description      string
	status           string
	dueDate          string
	maxSizeFile      string
	allowedTypesFile string
	maxFile          string
}

type createArgs struct {
	filesID          []string
	gpID             int64
	name             string
	description      sql.NullString
	status           int8
	dueDate          string
	maxSizeFile      int64
	allowedTypesFile []string
	maxFile          int64
}
type updateParams struct {
	ID               string
	filesID          string
	gpID             string
	name             string
	description      string
	status           string
	dueDate          string
	maxSizeFile      string
	allowedTypesFile string
	maxFile          string
}
type updateArgs struct {
	ID               int64
	filesID          []string
	gpID             int64
	name             string
	description      sql.NullString
	status           int8
	dueDate          time.Time
	maxSizeFile      int64
	allowedTypesFile []string
	maxFile          int64
}

type deleteParams struct {
	id string
}

type deleteArgs struct {
	id int64
}
type availableParams struct {
	id string
}

type availableArgs struct {
	id int64
}

type getReportResponse struct {
	ScheduleID int64  `json:"schedule_id"`
	CourseName string `json:"course_name"`
	Attendance string `json:"attendance"`
	Assignment string `json:"assignment"`
	Quiz       string `json:"quiz"`
	Mid        string `json:"mid"`
	Final      string `json:"final"`
	Total      string `json:"total"`
}

type getGradeResponse struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Score string `json:"grade"`
}

type detailParams struct {
	ID      string
	page    string
	total   string
	payload string
}

type detailArgs struct {
	ID      int64
	page    int
	total   int
	payload string
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
type detAsgUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UploadedAt  string `json:"uploaded_at"`
	Link        string `json:"url"`
}
type respDetAsgUser struct {
	TotalPage   int          `json:"total_page"`
	CurrentPage int          `json:"current_page"`
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	DueDate     string       `json:"due_date"`
	Status      string       `json:"status"`
	DetAsgUser  []detAsgUser `json:"users"`
}
type respDetailUpdate struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	DueDate          time.Time `json:"due_date"`
	GradeParameterID int64     `json:"grade_parameter_id"`
	Status           int8      `json:"status"`
	MaxSize          int8      `json:"max_size"`
	MaxFile          int8      `json:"max_file"`
	Type             []string  `json:"types"`
	FilesID          []file    `json:"files"`
}

type file struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	URLThumbnail string `json:"url_thumbnail"`
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
type respGP struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
