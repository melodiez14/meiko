package assignment

import (
	"database/sql"
	"time"
)

const (
	// MaximumID const
	MaximumID = 40
	// StatusAssignmentInactive const
	StatusAssignmentInactive = 0
	// StatusAssignmentActive const
	StatusAssignmentActive = 1
)

// Assignment struct ...
type Assignment struct {
	ID               int64          `db:"id"`
	Name             string         `db:"name"`
	Status           int8           `db:"status"`
	Description      sql.NullString `db:"description"`
	GradeParameterID int32          `db:"grade_parameters_id"`
	DueDate          time.Time      `db:"due_date"`
}

// File struct ...
type File struct {
	ID        string         `db:"id"`
	Name      sql.NullString `db:"name"`
	Mime      sql.NullString `db:"mime"`
	Extension string         `db:"extension"`
	UserID    int64          `db:"users_id"`
	Type      sql.NullString `db:"type"`
	TableName sql.NullString `db:"table_name"`
	TableID   sql.NullString `db:"table_id"`
}

// FileAssignment struct ...
type FileAssignment struct {
	Assignment Assignment
}

// GradeParameter struct ...
type GradeParameter struct {
	Type         string  `json:"type"`
	Percentage   float32 `json:"percentage"`
	StatusChange uint8   `json:"status_change"`
}

//DetailAssignment struct ...
type DetailAssignment struct {
	Assignment     Assignment
	File           File
	GradeParameter GradeParameter
}
