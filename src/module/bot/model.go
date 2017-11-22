package bot

import (
	"time"
)

const (
	StatusUser = 0
	StatusBot  = 1
)

type Log struct {
	ID        uint64    `db:"id"`
	Message   string    `db:"message"`
	UserID    int64     `db:"users_id"`
	Status    uint8     `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

type Assistant struct {
	IdentityCode int64  `db:"identity_code"`
	Name         string `db:"name"`
	Phone        string `db:"phone"`
	LineID       string `db:"line_id"`
	CourseID     string `db:"courses_id"`
	CourseName   string `db:"courses_name"`
}

type Schedule struct {
	CourseName string `db:"name"`
	Day        int8   `db:"day"`
	Place      string `db:"places_id"`
	StartTime  uint16 `db:"start_time"`
	EndTime    uint16 `db:"end_time"`
}

type Assignment struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	DueDate     time.Time `db:"due_date"`
	CourseName  string    `db:"course_name"`
}

type Grade struct {
	AssignmentID string    `db:"id"`
	Name         string    `db:"name"`
	Score        float32   `db:"score"`
	CourseName   string    `db:"course_name"`
	UpdatedAt    time.Time `db:"updated_at"`
}
