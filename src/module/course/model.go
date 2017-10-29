package course

import (
	"database/sql"
)

type (
	QueryGet    struct{ string }
	QuerySelect struct{ string }
	QueryInsert struct{ string }
	QueryUpdate struct{ string }
)

const (
	StatusScheduleInactive = 0
	StatusScheduleActive   = 1
	StatusScheduleDeleted  = 2

	MaximumID = 40

	PStatusUnapproved = 0
	PStatusStudent    = 1
	PStatusAssistant  = 2

	GradeParameterFinal      = "FINAL"
	GradeParameterMid        = "MID"
	GradeParameterAssignment = "ASSIGNMENT"
	GradeParameterAttendance = "ATTENDANCE"
	GradeParameterQuiz       = "KUIS"
)

type Course struct {
	ID          string         `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	UCU         int8           `db:"ucu"`
}

type Schedule struct {
	ID        int64  `db:"id"`
	Status    int8   `db:"status"`
	StartTime uint16 `db:"start_time"`
	EndTime   uint16 `db:"end_time"`
	Day       int8   `db:"day"`
	Class     string `db:"class"`
	Semester  int8   `db:"semester"`
	Year      int16  `db:"year"`
	CourseID  string `db:"courses_id"`
	PlaceID   string `db:"places_id"`
	CreatedBy int64  `db:"created_by"`
}

type CourseSchedule struct {
	Course   Course
	Schedule Schedule
}
