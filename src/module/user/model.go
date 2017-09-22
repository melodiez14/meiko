package user

import (
	"database/sql"
	"time"
)

const (
	ColID           = "id"
	ColName         = "name"
	ColEmail        = "email"
	ColGender       = "gender"
	ColCollege      = "college"
	ColNote         = "note"
	ColStatus       = "status"
	ColLineID       = "line_id"
	ColPhone        = "phone"
	ColRoleGroupsID = "rolegroups_id"
	ColPassword     = "password"

	OperatorEquals = "="
	OperatorIn     = "IN"
	OperatorMore   = ">"
	OperatorLess   = "<"
)

type User struct {
	ID           int64          `db:"id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	Gender       int8           `db:"gender"`
	College      string         `db:"college"`
	Note         string         `db:"note"`
	Status       int8           `db:"status"`
	LineID       sql.NullString `db:"line_id"`
	Phone        sql.NullString `db:"phone"`
	RoleGroupsID sql.NullInt64  `db:"rolegroups_id"`
}

type Verification struct {
	Code           uint16 `db:"email_verification_code"`
	ExpireDuration string
	ExpireDate     time.Time `db:"email_verification_expire_date"`
	Attempt        uint8     `db:"email_verification_attempt"`
}

type Confirmation struct {
	ID      int64         `db:"id"`
	Code    sql.NullInt64 `db:"email_verification_code"`
	Attempt sql.NullInt64 `db:"email_verification_attempt"`
}
