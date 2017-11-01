package course

import (
	"database/sql"
)

// type struct for collect query insert update select and get data form database 
type (
	// QueryGet Query to get data from database
	QueryGet    struct{ string }
	// QuerySelect Query to select data from database
	QuerySelect struct{ string }
	// QueryInsert Query to select data to database
	QueryInsert struct{ string }
	// QueryUpdate Query to select data to database
	QueryUpdate struct{ string }
)

// constant to all over user information that needed to check course
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

	PStatusStudent   = 0
	PStatusAssistant = 1

	OperatorEquals  = "="
	OperatorUnquals = "!="
	OperatorIn      = "IN"
	OperatorMore    = ">"
	OperatorLess    = "<"
)

// Course struct user detail information to get course that will be send to server in database
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
