package tutorial

import "database/sql"
import "time"

// Tutorial ...
type Tutorial struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	ScheduleID  int64          `db:"schedules_id"`
	CreatedAt   time.Time      `db:"created_at"`
}
