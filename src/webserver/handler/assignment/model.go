package assignment

import "time"
import "database/sql"

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
	Description       string
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
