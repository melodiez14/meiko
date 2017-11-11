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
	ID            uint64         `db:"id"`
	MeetingNumber int8           `db:"number"`
	Description   sql.NullString `db:"description"`
	Date          time.Time      `db:"date"`
	ScheduleID    int64          `db:"schedules_id"`
}
