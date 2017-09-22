package user

import (
	"database/sql"
)

type signUpParams struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type signUpArgs struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type emailVerificationParams struct {
	Email        string
	IsResendCode string
	Code         string
}

type emailVerificationArgs struct {
	Email        string
	IsResendCode bool
	Code         uint16
}

type signInParams struct {
	Email    string
	Password string
}

type signInArgs struct {
	Email    string
	Password string
}

type forgotResponse struct {
	Email          string `json:"email"`
	ExpireDuration string `json:"expire_duration"`
	MaxAttempt     uint8  `json:"max_attempt"`
}

type forgotParams struct {
	Email      string
	IsSendCode string
	Password   string
	Code       string
}

type forgotArgs struct {
	Email      string
	IsSendCode bool
	Password   string
	Code       uint16
}

type getVerifiedParams struct {
	Page  string
	Total string
}

type getVerifiedArgs struct {
	Page  int64
	Total int64
}

type getVerifiedResponse struct {
	ID     int64  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type updateProfileParams struct {
	Name    string
	Gender  string
	Phone   string
	LineID  string
	College string
	Note    string
}

type updateProfileArgs struct {
	Name    string
	Gender  int8
	Phone   sql.NullString
	LineID  sql.NullString
	College string
	Note    string
}

type getProfileResponse struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	LineID  string `json:"line_id"`
	College string `json:"college"`
	Note    string `json:"note"`
}

type changePasswordParams struct {
	OldPassword     string
	Password        string
	ConfirmPassword string
}

type changePasswordArgs struct {
	OldPassword string
	Password    string
}

type activationParams struct {
	ID     string
	Status string
}

type activationArgs struct {
	ID     int64
	Status int8
}
