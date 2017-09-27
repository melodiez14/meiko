package information

import "time"
import "database/sql"

const (
	ColID          = "id"
	ColTitle       = "title"
	ColDescription = "description"
	ColCourseID    = "courses_id"
	ColType        = "type"
	ColCreatedAt   = "created_at"
	ColUpdatedAt   = "updated_at"

	OperatorEquals  = "="
	OperatorUnquals = "!="
	OperatorIn      = "IN"
	OperatorMore    = ">"
	OperatorLess    = "<"

	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

type (
	QueryGet    struct{ string }
	QuerySelect struct{ string }
	QueryInsert struct{ string }
	QueryUpdate struct{ string }
)

type Information struct {
	ID          int64         `db:"id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
	Type        int8          `db:"type"`
	CourseID    sql.NullInt64 `db:"courses_id"`
	CreatedAt   time.Time     `db:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at"`
}
