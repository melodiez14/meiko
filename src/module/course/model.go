package course

import (
	"database/sql"
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
	GradeParameterQuiz       = "QUIZ"

	GradeParameterStatusUnchange = 0
	GradeParameterStatusChange   = 1
)

// Course struct user detail information to get course that will be send to server in database
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

type GradeParameter struct {
	ID           int64   `db:"id"`
	Type         string  `db:"type"`
	Percentage   float32 `db:"percentage"`
	ScheduleID   int64   `db:"schedules_id"`
	StatusChange uint8   `db:"status_change"`
}
