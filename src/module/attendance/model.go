package attendance

import (
	"time"
)

type Attendance struct {
	ID            int64     `db:"id"`
	MeetingNumber int8      `db:"meeting_number"`
	Status        int8      `db:"status"`
	Date          time.Time `db:"meeting_date"`
}
