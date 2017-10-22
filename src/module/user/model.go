package user

import (
	"database/sql"
	"time"
)

// const const for specify variable in user
/*
	@params:
		ID				= int64
		Name			= string
		Email			= string
		Gender			= int8
		Note			= string
		Status			= int8
		IdentityCode	= int64
		LineID			= sql.string
		Phone			= sql.string
		RoleGroupsID	= sql.int64
	@example:
		ID				= 140810140060
		Name			= kharil azmi ashari
		Email			= khairil_azmi_ashari@yahoo.com
		Gender			= 1
		Note			= just doing nothing
		Status			= 1
		IdentityCode	= 140810140060
		LineID			= khaazas
		Phone			= 082214467300
		RoleGroupsID	= 0
	@return
*/
const (
	ColID           = "id"
	ColName         = "name"
	ColEmail        = "email"
	ColGender       = "gender"
	ColNote         = "note"
	ColStatus       = "status"
	ColIdentityCode = "identity_code"
	ColLineID       = "line_id"
	ColPhone        = "phone"
	ColRoleGroupsID = "rolegroups_id"
	ColPassword     = "password"

	StatusUnverified = 0
	StatusVerified   = 1
	StatusActivated  = 2

	GenderUndefined = 0
	GenderMale      = 1
	GenderFemale    = 2

	OperatorEquals  = "="
	OperatorUnquals = "!="
	OperatorIn      = "IN"
	OperatorMore    = ">"
	OperatorLess    = "<"
)

type (
	// QueryGet struct for get query data from database
	QueryGet    struct{ string }
	// QuerySelect struct for show query data from database
	QuerySelect struct{ string }
	// QueryInsert struct for send query data to database
	QueryInsert struct{ string }
	// QueryUpdate struct for send update query data to database
	QueryUpdate struct{ string }
)

// User struct for save user information
/*
	@params:
		ID				= int64
		Name			= string
		Email			= string
		Gender			= int8
		Note			= string
		Status			= int8
		IdentityCode	= int64
		LineID			= sql.string
		Phone			= sql.string
		RoleGroupsID	= sql.int64
	@example:
		ID				= 140810140060
		Name			= kharil azmi ashari
		Email			= khairil_azmi_ashari@yahoo.com
		Gender			= 1
		Note			= just doing nothing
		Status			= 1
		IdentityCode	= 140810140060
		LineID			= khaazas
		Phone			= 082214467300
		RoleGroupsID	= 0
	@return
*/
type User struct {
	ID           int64          `db:"id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	Gender       int8           `db:"gender"`
	Note         string         `db:"note"`
	Status       int8           `db:"status"`
	IdentityCode int64          `db:"identity_code"`
	LineID       sql.NullString `db:"line_id"`
	Phone        sql.NullString `db:"phone"`
	RoleGroupsID sql.NullInt64  `db:"rolegroups_id"`
}

// Verification struct for verify account confirmation through email
/*
	@params:
		Code			= uint16
		ExpireDuration	= string
		ExpireDate		= time.Time
		Attempt			= uint8
	@example:
		Code			= 140810
		ExpireDuration	= 1 hour
		ExpireDate		= 21 October 2017, 3:30:10
		Attempt			= 140810
	@return
*/
type Verification struct {
	Code           uint16 `db:"email_verification_code"`
	ExpireDuration string
	ExpireDate     time.Time `db:"email_verification_expire_date"`
	Attempt        uint8     `db:"email_verification_attempt"`
}

// Confirmation struct for confirmation the account through email conformation
/*
	@params:
		ID		= int64
		Code	= sql.int64
		Attempt	= sql.int64
	@example:
		ID		= 140810140060
		Code	= 140810
		Attempt	= 140810
	@return
*/
type Confirmation struct {
	ID      int64         `db:"id"`
	Code    sql.NullInt64 `db:"email_verification_code"`
	Attempt sql.NullInt64 `db:"email_verification_attempt"`
}
