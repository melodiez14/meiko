package course

import (
	"database/sql"
)

type (
	QueryGet    struct{ string }
	QuerySelect struct{ string }
	QueryInsert struct{ string }
	QueryUpdate struct{ string }
)

const (
	ColID          = "id"
	ColName        = "name"
	ColDescription = "description"
	ColUCU         = "ucu"
	ColSemester    = "semester"
	ColStatus      = "status"
	ColStartTime   = "start_time"
	ColEndTime     = "end_time"
	ColClass       = "classes"
	ColDay         = "day"
	ColPlaceID     = "places_id"
	ColCreatedBy   = "created_by"

	StatusInactive = 0
	StatusActive   = 1
	StatusDeleted  = 2

	OperatorEquals  = "="
	OperatorUnquals = "!="
	OperatorIn      = "IN"
	OperatorMore    = ">"
	OperatorLess    = "<"
)

type Course struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	UCU         int8           `db:"ucu"`
	Semester    int8           `db:"semester"`
	Status      int8           `db:"status"`
	StartTime   uint16         `db:"start_time"`
	EndTime     uint16         `db:"end_time"`
	Class       string         `db:"classes"`
	Day         int8           `db:"day"`
	PlaceID     string         `db:"places_id"`
	CreatedBy   int64          `db:"created_by"`
}
