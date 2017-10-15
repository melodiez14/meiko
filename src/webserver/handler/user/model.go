package user

import (
	"database/sql"
)

type signUpParams struct {
	IdentityCode string
	Name         string
	Email        string
	Password     string
}

type signUpArgs struct {
	IdentityCode int64
	Name         string
	Email        string
	Password     string
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
	Page  uint16
	Total uint16
}

type getVerifiedResponse struct {
	IdentityCode int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

type updateProfileParams struct {
	IdentityCode string
	Name         string
	Email        string
	Gender       string
	Phone        string
	LineID       string
	Note         string
}

type updateProfileArgs struct {
	IdentityCode int64
	Name         string
	Email        string
	Gender       int8
	Phone        sql.NullString
	LineID       sql.NullString
	Note         string
}

type getProfileResponse struct {
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	Gender                string `json:"gender"`
	Phone                 string `json:"phone"`
	IdentityCode          int64  `json:"id"`
	LineID                string `json:"line_id"`
	Note                  string `json:"about_me"`
	ImageProfile          string `json:"img"`
	ImageProfileThumbnail string `json:"img_t"`
}

type changePasswordParams struct {
	IdentityCode    string
	Email           string
	OldPassword     string
	Password        string
	ConfirmPassword string
}

type changePasswordArgs struct {
	IdentityCode int64
	Email        string
	OldPassword  string
	Password     string
}

type activationParams struct {
	IdentityCode string
	Status       string
}

type activationArgs struct {
	IdentityCode int64
	Status       int8
}

type detailParams struct {
	IdentityCode string
}

type detailArgs struct {
	IdentityCode int64
}

type detailResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone"`
	IdentityCode int64  `json:"id"`
	LineID       string `json:"line_id"`
	Note         string `json:"about_me"`
}

type updateParams struct {
	IdentityCode string
	Name         string
	Email        string
	Gender       string
	Phone        string
	LineID       string
	Note         string
	Status       string
}

type updateArgs struct {
	IdentityCode int64
	Name         string
	Email        string
	Gender       int8
	Phone        sql.NullString
	LineID       sql.NullString
	Note         string
	Status       int8
}

type deleteParams struct {
	IdentityCode string
}

type deleteArgs struct {
	IdentityCode int64
}

type createParams struct {
	IdentityCode string
	Name         string
	Email        string
}

type createArgs struct {
	IdentityCode int64
	Name         string
	Email        string
}
