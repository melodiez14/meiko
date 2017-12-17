package attendance

import (
	"database/sql"
	"time"
)

type Attendance struct {
	MeetingID  uint64 `db:"meetings_id"`
	UserID     int64  `db:"p_users_schedules_users_id"`
	ScheduleID int64  `db:"p_users_schedules_schedules_id"`
}

type Meeting struct {
	ID             uint64         `db:"id"`
	Subject        string         `db:"subject"`
	Number         uint8          `db:"number"`
	Description    sql.NullString `db:"description"`
	Date           time.Time      `db:"date"`
	ScheduleID     int64          `db:"schedules_id"`
	TotalAttendant uint16         `db:"total"`
}

type AttendanceCount struct {
	MeetingID uint64 `db:"meetings_id"`
	Count     uint16 `db:"count"`
}

type AttendanceReport struct {
	MeetingTotal    int `db:"meeting_total"`
	AttendanceTotal int `db:"attendance_total"`
}
