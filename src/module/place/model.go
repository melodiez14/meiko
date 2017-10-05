package place

import (
	"database/sql"
)

type Place struct {
	ID          string         `db:"id"`
	Description sql.NullString `db:"description"`
}
