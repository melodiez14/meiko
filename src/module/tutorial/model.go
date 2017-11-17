package tutorial

import "database/sql"
import "time"

// Tutorial ...
type Tutorial struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	CreatedAt   time.Time      `db:"created_at"`
}
