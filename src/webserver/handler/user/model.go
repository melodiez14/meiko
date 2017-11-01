package user

import (
	"database/sql"
)

// signUpParams Parameter that needed in sign up
/*
	@params:
		IdentityCode= string
		Name		= string
		Email		= string
		Password	= string
	@example:
		IdentityCode= 140810140060
		Name		= khairil azmi ashari
		Email		= khairil_azmi_ashari@yahoo.com
		Password	= Khairil14001
	@return
*/
type signUpParams struct {
	IdentityCode string
	Name         string
	Email        string
	Password     string
}

// signUpArgs Parameter that will use to sign up
/*
	@params:
		IdentityCode= string
		Name		= string
		Email		= string
		Password	= string
	@example:
		IdentityCode= 140810140060
		Name		= khairil azmi ashari
		Email		= khairil_azmi_ashari@yahoo.com
		Password	= Khairil14001
	@return
*/
type signUpArgs struct {
	IdentityCode int64
	Name         string
	Email        string
	Password     string
}

// emailVerificationParams Parameter that needed in email verification
/*
	@params:
		Email			= string
		IsResendCode	= string
		Code			= string
	@example:
		Email			= khairil_azmi_ashari@yahoo.com
		IsresendCode	= true
		Code			= 123456
	@return
*/
type emailVerificationParams struct {
	Email        string
	IsResendCode string
	Code         string
}

// emailVerificationArgs Parameter that will use to doing email verification
/*
	@params:
		Email			= string
		IsResendCode	= string
		Code			= string
	@example:
		Email			= khairil_azmi_ashari@yahoo.com
		IsresendCode	= true
		Code			= 123456
	@return
*/
type emailVerificationArgs struct {
	Email        string
	IsResendCode string
	Code         uint16
}

// signInParams Parameter that needed in sign in
/*
	@params:
		Email		= string
		Password	= string
	@example:
		Email		= khairil_azmi_ashari@yahoo.com
		Password	= Khairil14001
	@return
*/
type signInParams struct {
	Email    string
	Password string
}

// signInArgs Parameter that will use to sign in
/*
	@params:
		Email		= string
		Password	= string
	@example:
		Email		= khairil_azmi_ashari@yahoo.com
		Password	= Khairil14001
	@return
*/
type signInArgs struct {
	Email    string
	Password string
}

// signInResponse Variable that will be send to server when sign in
/*
	@params:
		IsLoggedIn	= bool
		Modules		= string
	@example:
		IsLoggedIn	= true
		Modules		= assignment
	@return
*/
type signInResponse struct {
	IsLoggedIn bool                `json:"is_logged_in"`
	Modules    map[string][]string `json:"modules"`
}

// forgotResponse Variable that will be send to server when forgot password clicked
/*
	@params:
		Email			= string
		ExpireDuration	= string
		MaAttempt		= uint8
	@example:
		Email			= khairil_azmi_ashari@yahoo.com
		ExpireDuration	= 1 hour
		MaAttempt		= uint8
	@return
*/
type forgotResponse struct {
	Email          string `json:"email"`
	ExpireDuration string `json:"expire_duration"`
	MaxAttempt     uint8  `json:"max_attempt"`
}

// forgotParams Parameter that needed in forgot password. to generate new password to email
/*
	@params:
		Email			= string
		IsSendCode		= string
		Password		= string
		Code			= string
	@example:
		Email			= khairil_azmi_ashari@yahoo.com
		ExpireDuration	= true
		Password		= Khairil14001
		Code			= 123456
	@return
*/
type forgotParams struct {
	Email      string
	IsSendCode string
	Password   string
	Code       string
}

// forgotArgs Parameter that will be use to forgot password. to generate new password to email
/*
	@params:
		Email			= string
		IsSendCode		= string
		Password		= string
		Code			= string
	@example:
		Email			= khairil_azmi_ashari@yahoo.com
		IsSendCode	= true
		Password		= Khairil14001
		Code			= 123456
	@return
*/
type forgotArgs struct {
	Email      string
	IsSendCode string
	Password   string
	Code       uint16
}

// getVeriviedParams Parameter that needed to get verified item to show in list.
/*
	@params:
		Page			= string
		Total			= string
	@example:
		Page			= 35
		Total			= 60
	@return
*/
type getVerifiedParams struct {
	Page  string
	Total string
}

// getVeriviedArgs Parameter that will be use to get verified item to show in list.
/*
	@params:
		Page			= uint16
		Total			= uint16
	@example:
		Page			= 35
		Total			= 60
	@return
*/
type getVerifiedArgs struct {
	Page  uint16
	Total uint16
}

// getVeriviedResponse Response to server from application to get verified content.
/*
	@params:
		IdentityCode	= int64
		Name			= string
		Email			= string
		Status			= string
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
		Status			= 1
	@return
*/
type getVerifiedResponse struct {
	IdentityCode int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

// updateProfileParams Parameter that needed to update user profile.
/*
	@params:
		IdentityCode	= int64
		Name			= string
		Email			= string
		Gender			= string
		Phone			= string
		LineID			= string
		Note			= string
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
		Gender			= 1
		Phone			= 082214467300
		LineID			= khaazas
		Note			= nothing is impossible
	@return
*/
type updateProfileParams struct {
	IdentityCode string
	Name         string
	Email        string
	Gender       string
	Phone        string
	LineID       string
	Note         string
}

// updateProfileArgs Parameter that will be use to update user profile.
/*
	@params:
		IdentityCode	= int64
		Name			= string
		Email			= string
		Gender			= string
		Phone			= string
		LineID			= string
		Note			= string
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
		Gender			= 1
		Phone			= 082214467300
		LineID			= khaazas
		Note			= nothing is impossible
	@return
*/
type updateProfileArgs struct {
	IdentityCode int64
	Name         string
	Email        string
	Gender       int8
	Phone        sql.NullString
	LineID       sql.NullString
	Note         string
}

// getProfileResponse Application will send data to server to update.
/*
	@params:
		Name					= string
		Email					= string
		Gender					= string
		Phone					= string
		IdentityCode			= int64
		LineID					= string
		Note					= string
		ImageProfile			= string
		ImageProfileThumbnail	= string
	@example:
		Name					= khairil azmi ashari
		Email					= khairilazmiashari@gmail.com
		Gender					= 1
		Phone					= 082214467300
		IdentityCode			= 140810140060
		LineID					= khaazas
		Note					= nothing is impossible
		ImageProfile			= profile.jpg
		ImageProfileThumbnail	= profileThumb.jpg
	@return
*/
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

// changePasswordParams Parameter that need to change password.
/*
	@params:
		IdentityCode			= int64
		Email					= string
		OldPassword				= string
		Password				= string
		ConfirmPassword			= string
	@example:
		IdentityCode			= 140810140060
		Email					= khairilazmiashari@gmail.com
		OldPassword				= Khairil14001
		Password				= 14001Khairil
		ConfirmPassword			= 14001Khairil
	@return
*/
type changePasswordParams struct {
	IdentityCode    string
	Email           string
	OldPassword     string
	Password        string
	ConfirmPassword string
}

// changePasswordArgs Parameter that will be use to change password.
/*
	@params:
		IdentityCode			= int64
		Email					= string
		OldPassword				= string
		Password				= string
		ConfirmPassword			= string
	@example:
		IdentityCode			= 140810140060
		Email					= khairilazmiashari@gmail.com
		OldPassword				= Khairil14001
		Password				= 14001Khairil
		ConfirmPassword			= 14001Khairil
	@return
*/
type changePasswordArgs struct {
	IdentityCode int64
	Email        string
	OldPassword  string
	Password     string
}

// activationParams Parameter that needed to activation.
/*
	@params:
		IdentityCode	= string
		Status			= string
	@example:
		IdentityCode	= 140810140060
		Status			= 1
	@return
*/
type activationParams struct {
	IdentityCode string
	Status       string
}

// activationParams Parameter that will be use to activation.
/*
	@params:
		IdentityCode	= int64
		Status			= int8
	@example:
		IdentityCode	= 140810140060
		Status			= 1
	@return
*/
type activationArgs struct {
	IdentityCode int64
	Status       int8
}

// detailParams Parameter that needed to get detail information.
/*
	@params:
		IdentityCode	= string
	@example:
		IdentityCode	= 140810140060
	@return
*/
type detailParams struct {
	IdentityCode string
}

// detailParams Parameter that will be use to detail information.
/*
	@params:
		IdentityCode	= int64
	@example:
		IdentityCode	= 140810140060
	@return
*/
type detailArgs struct {
	IdentityCode int64
}

// detailResponse Parameter that will be send to server to get detail information.
/*
	@params:
		Name					= string
		Email					= string
		Gender					= string
		Phone					= string
		IdentityCode			= int64
		LineID					= string
		Note					= string
	@example:
		Name					= khairil azmi ashari
		Email					= khairilazmiashari@gmail.com
		Gender					= 1
		Phone					= 082214467300
		IdentityCode			= 140810140060
		LineID					= khaazas
		Note					= nothing is impossible
	@return
*/
type detailResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone"`
	IdentityCode int64  `json:"id"`
	LineID       string `json:"line_id"`
	Note         string `json:"about_me"`
}

// updateParams Parameter that will be needed to update user information.
/*
	@params:
		IdentityCode	= string
		Name			= string
		Email			= string
		Gender			= string
		Phone			= string
		LineID			= string
		Note			= string
		Status			= string
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
		Gender			= 1
		Phone			= 082214467300
		LineID			= khaazas
		Note			= nothing is impossible
	@return
*/
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

// updateArgs Parameter that will be use to update user information.
/*
	@params:
		IdentityCode	= int64
		Name			= string
		Email			= string
		Gender			= int8
		Phone			= string
		LineID			= string
		Note			= string
		Status			= int8
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
		Gender			= 1
		Phone			= 082214467300
		LineID			= khaazas
		Note			= nothing is impossible
		Status			= 1
	@return
*/
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

// deleteParams Parameter that needed to delete user.
/*
	@params:
		IdentityCode	= string
	@example:
		IdentityCode	= 140810140060
	@return
*/
type deleteParams struct {
	IdentityCode string
}

// deleteArgs Parameter that will be use to delete user.
/*
	@params:
		IdentityCode	= int64
	@example:
		IdentityCode	= 140810140060
	@return
*/
type deleteArgs struct {
	IdentityCode int64
}

// createParams Parameter that needed to create user.
/*
	@params:
		IdentityCode	= string
		Name			= string
		Email			= string
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
	@return
*/
type createParams struct {
	IdentityCode string
	Name         string
	Email        string
}

// createParams Parameter that will use to create user.
/*
	@params:
		IdentityCode	= int64
		Name			= string
		Email			= string
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
	@return
*/
type createArgs struct {
	IdentityCode int64
	Name         string
	Email        string
}
