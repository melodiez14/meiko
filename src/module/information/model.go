package information

import "time"
import "database/sql"

const (
	ColID          = "id"
	ColTitle       = "title"
	ColDescription = "description"
	ColScheduleID  = "schedules_id"
	CreatedAt      = "created_at"
	UpdatedAt      = "updated_at"
)

type Information struct {
	ID          int64          `db:"id"`
	Title       string         `db:"title"`
	Description sql.NullString `db:"description"`
	ScheduleID  sql.NullInt64  `db:"schedules_id"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}
