package user

import (
	"time"
)

type User struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	Email        string `db:"email"`
	Password     string `db:"password"`
	Gender       string `db:"gender"`
	College      string `db:"college"`
	Note         string `db:"note"`
	Status       bool   `db:"status"`
	RoleGroupsID int64  `db:"rolegroups_id"`
}

type Verification struct {
	Code           uint16 `db:"email_verification_code"`
	ExpireDuration string
	ExpireDate     time.Time `db:"email_verification_expire_date"`
	Attempt        uint8     `db:"email_verification_attempt"`
}
