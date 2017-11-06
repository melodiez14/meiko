package assignment

import (
	"database/sql"
	"time"
)

const (
	MaximumID                = 40
	StatusAssignmentInactive = 0
	StatusAssignmentActive   = 1
)

type Assignment struct {
	ID               int64          `db:"id"`
	Name             string         `db:"name"`
	Status           int8           `db:"status"`
	Description      sql.NullString `db:"description"`
	GradeParameterID int32          `db:"grade_parameters_id"`
	DueDate          time.Time      `db:"due_date"`
}
type File struct {
	ID        string         `db:"id"`
	Name      sql.NullString `db:"name"`
	Mime      sql.NullString `db:"mime"`
	Extension string         `db:"extension"`
	UserID    int64          `db:"users_id"`
	Type      string         `db:"type"`
	TableName string         `db:"table_name"`
	TableID   string         `db:"table_id"`
}
type FileAssignment struct {
	Assignment Assignment
}
type GradeParameter struct {
	Type         string `json:"type"`
	Percentage   uint8  `json:"percentage"`
	StatusChange uint8  `json:"status_change"`
}
type DetailAssignment struct {
	Assignment     Assignment
	File           File
	GradeParameter GradeParameter
}
