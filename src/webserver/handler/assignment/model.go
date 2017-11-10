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
)

type createParams struct {
	FilesID           string
	GradeParametersID string
	Name              string
	Description       string
	Status            string
	DueDate           string
}

type createArgs struct {
	FilesID           string
	GradeParametersID int64
	Name              string
	Description       sql.NullString
	Status            string
	DueDate           string
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
	ID               int64          `json:"id"`
	Status           string         `json:"status"`
	Name             string         `json:"name"`
	GradeParameterID int32          `json:"grade_parameters_id"`
	Description      sql.NullString `json:"description"`
	DueDate          time.Time      `json:"due_date"`
	FilesID          string         `json:"files_id"`
	FilesName        sql.NullString `json:"files_name"`
	Mime             sql.NullString `json:"mime"`
	Type             string         `json:"type"`
	Percentage       float32        `json:"percentage"`
}
type uploadAssignmentParams struct {
	FileID       string
	AssignmentID string
	UserID       int64
	Subject      string
	Description  string
}
type uploadAssignmentArgs struct {
	FileID       string
	AssignmentID int64
	UserID       int64
	Subject      sql.NullString
	Description  sql.NullString
}
type readUploadedAssignmentParams struct {
	ScheudleID   string
	AssignmentID string
	Name         string
	Description  string
	Score        string
	DueDate      string
	PathFile     string
}
type readUploadedAssignmentArgs struct {
	ScheudleID   int64
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
