package assignment

import (
	"database/sql"
	"time"
)

// ...
const (
	MaxDesc                 = 1000
	StatusUploadNotRequired = 0
	StatusUploadRequired    = 1
	MaxFile                 = 5
	MinFile                 = 1
	MaxSizeFile             = 100
	MinSizeFile             = 1
	MaxPage                 = 10
)

// Assignment struct ...
type Assignment struct {
	ID               int64          `db:"id"`
	Name             string         `db:"name"`
	Status           int8           `db:"status"`
	Description      sql.NullString `db:"description"`
	GradeParameterID int64          `db:"grade_parameters_id"`
	DueDate          time.Time      `db:"due_date"`
	MaxSize          sql.NullInt64  `db:"max_size"`
	MaxFile          sql.NullInt64  `db:"max_file"`
	CreatedAt        time.Time      `db:"created_at"`
	UpdatedAt        time.Time      `db:"updated_at"`
}

// ConciseAssignment ..
type ConciseAssignment struct {
	ID      int64     `db:"id"`
	Name    string    `db:"name"`
	DueDate time.Time `db:"due_date"`
	Status  int8      `db:"status"`
}

// UserAssignment struct ...
type UserAssignment struct {
	UserID       int64           `db:"users_id"`
	AssignmentID int64           `db:"assignments_id"`
	Score        sql.NullFloat64 `db:"score"`
	Description  sql.NullString  `db:"description"`
	CreatedAt    time.Time       `db:"created_at"`
	UpdatedAt    time.Time       `db:"updated_at"`
}

// UserScore  struct ...
type UserScore struct {
	UserID    int64           `db:"users_id"`
	Score     sql.NullFloat64 `db:"score"`
	UpdatedAt time.Time       `db:"updated_at"`
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
	Type       string
	Percentage float32
}

//DetailAssignment struct ...
type DetailAssignment struct {
	Assignment     Assignment
	File           File
	GradeParameter GradeParameter
}

// DetailUploadedAssignment struct ...
type DetailUploadedAssignment struct {
	ScheduleID            int64
	AssignmentID          int64
	Name                  string
	DescriptionUser       sql.NullString
	DescriptionAssignment sql.NullString
	Score                 sql.NullString
	DueDate               string
	PathFile              sql.NullString
}
