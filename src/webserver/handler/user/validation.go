package user

import (
	"database/sql"
	"fmt"
	"html"
	"strconv"

	"github.com/melodiez14/meiko/src/module/user"

	"github.com/melodiez14/meiko/src/util/alias"
	"github.com/melodiez14/meiko/src/util/helper"
)

// Function to validate signUp parameter to be used
/*
	@params:
		IdentityCode= required, numeric, characters=12
		Name	= required, alphaspace, 0<characters<50
		Email	= required, email format, 0<characters<45
		Password= required, minimum 1 uppercase, lowercase, numeric, characters>=6
	@example:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
		Password		= Khairil14001
	@return:
		IdentityCode	= 140810140060
		Name			= khairil azmi ashari
		Email			= khairilazmiashari@gmail.com
		Password		= Khairil14001
*/
func (params signUpParams) validate() (signUpArgs, error) {

	var args signUpArgs
	params = signUpParams{
		IdentityCode: helper.Trim(params.IdentityCode),
		Name:         helper.Trim(params.Name),
		Email:        helper.Trim(params.Email),
		Password:     html.EscapeString(params.Password),
	}

	// ID validation
	id, err := helper.NormalizeNPM(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Error validation: %s", err.Error())
	}

	// Name validation
	name, err := helper.NormalizeName(params.Name)
	if err != nil {
		return args, fmt.Errorf("Error validation: %s", err.Error())
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, fmt.Errorf("Error validation: %s", err.Error())
	}

	// Password validation
	if helper.IsEmpty(params.Password) {
		return args, fmt.Errorf("Error validation: password can't be empty")
	}
	if len(params.Password) < alias.UserPasswordLengthMin {
		return args, fmt.Errorf("Error validation: password at least consist of 6 characters")
	}
	if !helper.IsPassword(params.Password) {
		return args, fmt.Errorf("Error validation: password should contains at least uppercase, lowercase, and numeric")
	}

	args = signUpArgs{
		IdentityCode: id,
		Name:         name,
		Email:        email,
		Password:     helper.StringToMD5(params.Password),
	}
	return args, nil
}

// Function to validate email verification parameter to be used
/*
	@params:
		Code			= required if resend is empty, numeric, characters=4
		Email			= required, email format, 0<characters<45
		IsResendCode	= optional, value=true or empty
	@example:
		Code			= 123456
		Email			= khairil_azmi_ashari@yahoo.com
		IsresendCode	= true
	@return:
		Code			= 123456
		Email			= khairil_azmi_ashari@yahoo.com
		IsresendCode	= true
*/
func (params emailVerificationParams) validate() (emailVerificationArgs, error) {

	var args emailVerificationArgs
	params = emailVerificationParams{
		Code:         params.Code,
		Email:        helper.Trim(params.Email),
		IsResendCode: params.IsResendCode,
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, err
	}

	// IsResendCode validation
	if !helper.IsEmpty(params.IsResendCode) {
		if params.IsResendCode == "true" {
			return emailVerificationArgs{
				Email:        email,
				IsResendCode: true,
				Code:         0,
			}, nil
		}
	}

	// Code validation: if isResendCode is true, pass the Code validation
	if helper.IsEmpty(params.Code) {
		return args, fmt.Errorf("Error validation: Code can't be empty")
	} else if len(params.Code) != alias.UserCodeLength {
		return args, fmt.Errorf("Error validation: Wrong code")
	}
	code, err := strconv.ParseInt(params.Code, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Error validation: Wrong code")
	}

	args = emailVerificationArgs{
		Email:        email,
		IsResendCode: false,
		Code:         uint16(code),
	}
	return args, nil
}

// Function to validate verified parameter to be used.
/*
	@params:
		Page	= required, positive numeric
		Total	= required, positive numeric
	@example:
		Page	= 35
		Total	= 60
	@return:
		Page	= 35
		Total	= 60
*/
func (params getVerifiedParams) validate() (getVerifiedArgs, error) {

	var args getVerifiedArgs
	if helper.IsEmpty(params.Page) || helper.IsEmpty(params.Total) {
		return args, fmt.Errorf("Invalid request")
	}

	page, err := strconv.ParseUint(params.Page, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	total, err := strconv.ParseUint(params.Total, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	// should be positive number
	if page < 1 || total < 1 {
		return args, fmt.Errorf("Invalid request")
	}

	args = getVerifiedArgs{
		Page:  uint16(page),
		Total: uint16(total),
	}
	return args, nil
}

// Function to validate activation paramater to be used
/*
	@params:
		identity		= required, numeric, characters=12
		status			= required, string
	@example:
		IdentityCode	= 140810140016
		status			= active or inactive
	@return:
		IdentityCode	= 140810140060
		Status			= actice or inactive
*/
func (params activationParams) validate() (activationArgs, error) {

	var args activationArgs
	// Check is params empty
	if helper.IsEmpty(params.IdentityCode) || helper.IsEmpty(params.Status) {
		return args, fmt.Errorf("Invalid Request")
	}

	identityCode, err := helper.NormalizeIdentity(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Error validation: ID should be numeric")
	}

	var status int8
	switch params.Status {
	case "active":
		status = user.StatusActivated
	case "inactive":
		status = user.StatusVerified
	default:
		return args, fmt.Errorf("Error validation: wrong status")
	}

	args = activationArgs{
		IdentityCode: identityCode,
		Status:       status,
	}

	return args, nil
}

// Function to validate sign in paramater to be used
/*
	@params:
		Email	= required, email format, 0<characters<45
		Password= required, minimum 1 uppercase, lowercase, numeric, characters>=6
	@example:
		Email	= risal.falah@gmail.com
		Password= Qwerty123
	@return:
		Email	= risal.falah@gmail.com
		Password= Qwerty123
*/
func (params signInParams) validate() (signInArgs, error) {

	var args signInArgs
	params = signInParams{
		Email:    helper.Trim(params.Email),
		Password: html.EscapeString(params.Password),
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, fmt.Errorf("Error validation: %s", err.Error())
	}

	// Password validation
	if helper.IsEmpty(params.Password) {
		return args, fmt.Errorf("Error validation: password can't be empty")
	}
	if len(params.Password) < alias.UserPasswordLengthMin {
		return args, fmt.Errorf("Error validation: password at least consist of 6 characters")
	}
	if !helper.IsPassword(params.Password) {
		return args, fmt.Errorf("Error validation: password should contains at least uppercase, lowercase, and numeric")
	}

	args = signInArgs{
		Email:    email,
		Password: helper.StringToMD5(params.Password),
	}
	return args, nil
}

// Function to validate update profile parameter to be used
/*
	@params:
		IdentityCode= required, numeric, characters=12
		Email	= required, email format, 0<characters<45
		Name	= required, alphaspace, 0<characters<50
		Note	= optional, 0<character<100
		Gender	= optional, male or female
		Phone	= optional, numeric, 10<=characters<=12
		Line_id	= optional, 0<characters<=45
	@example:
		IdentityCode	= 140810140060
		Email			= khairilazmiashari@gmail.com
		Name			= khairil azmi ashari
		Note			= nothing is impossible
		Gender			= male or female
		Phone			= 082214467300
		Lide_id			= khaazas
	@return:
		IdentityCode	= 140810140060
		Email			= khairilazmiashari@gmail.com
		Name			= khairil azmi ashari
		Note			= nothing is impossible
		Gender			= male or female
		Phone			= 082214467300
		Lide_id			= khaazas
*/
func (params updateProfileParams) validate() (updateProfileArgs, error) {

	var args updateProfileArgs
	params = updateProfileParams{
		IdentityCode: params.IdentityCode,
		Email:        params.Email,
		Name:         helper.Trim(params.Name),
		Note:         helper.Trim(html.EscapeString(params.Note)),
		Gender:       params.Gender,
		Phone:        helper.Trim(params.Phone),
		LineID:       html.EscapeString(params.LineID),
	}

	// Identity code validation
	identityCode, err := helper.NormalizeIdentity(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	// Name validation
	name, err := helper.NormalizeName(params.Name)
	if err != nil {
		return args, err
	}

	// Note validation
	if !helper.IsEmpty(params.Note) {
		if len(params.Note) > alias.UserNoteLengthMax {
			return args, fmt.Errorf("Error validation: Note too long")
		}
	}

	// Gender validation
	var gender int8 = user.GenderUndefined
	if !helper.IsEmpty(params.Gender) {
		switch params.Gender {
		case "male":
			gender = alias.UserGenderMale
		case "female":
			gender = alias.UserGenderFemale
		default:
			return args, fmt.Errorf("Error validation: wrong gender")
		}
	}

	// Phone validation (can be empty)
	var phone sql.NullString
	if !helper.IsEmpty(params.Phone) {
		if !helper.IsPhone(params.Phone) {
			return args, fmt.Errorf("Error validation: wrong input phone")
		}
		phone = sql.NullString{String: params.Phone, Valid: true}
	}

	// Line vallidation (can be empty)
	var lineID sql.NullString
	if !helper.IsEmpty(params.LineID) {
		if len(params.LineID) > alias.UserLineIDLengthMax {
			return args, fmt.Errorf("Error validation: Line Id too long")
		}
		lineID = sql.NullString{String: params.LineID, Valid: true}
	}

	args = updateProfileArgs{
		IdentityCode: identityCode,
		Name:         name,
		Email:        email,
		Gender:       gender,
		Phone:        phone,
		LineID:       lineID,
		Note:         params.Note,
	}

	return args, nil
}

// Function to validate change password parameter to be used
/*
	@params:
		IdentityCode	= required, numeric, 10<=characters<=18
		Email			= required, email format, 0<characters<45
		OldPassword		= required, minimum 1 uppercase, lowercase, numeric, characters>=6
		Password		= required, minimum 1 uppercase, lowercase, numeric, characters>=6
		ConfirmPassword	= required, should be same as password
	@example:
		IdentityCode	= 140810140016
		Email			= khairilazmi@gmail.com
		OldPassword		= Qwerty123
		Password		= Qwerty321
		ConfirmPassword	= Qwerty321
	@return:
		IdentityCode	= 140810140016
		Email			= khairilazmi@gmail.com
		OldPassword		= Qwerty123
		Password		= Qwerty321
		ConfirmPassword	= Qwerty321
*/
func (params changePasswordParams) validate() (changePasswordArgs, error) {

	var args changePasswordArgs
	params = changePasswordParams{
		IdentityCode:    params.IdentityCode,
		Email:           params.Email,
		OldPassword:     html.EscapeString(params.OldPassword),
		Password:        html.EscapeString(params.Password),
		ConfirmPassword: html.EscapeString(params.ConfirmPassword),
	}

	// Identity Code validation
	identityCode, err := helper.NormalizeIdentity(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	// Old password validation
	if helper.IsEmpty(params.OldPassword) {
		return args, fmt.Errorf("Error validation: old password can't be empty")
	}
	if len(params.OldPassword) < alias.UserPasswordLengthMin {
		return args, fmt.Errorf("Error validation: old password at least consist of 6 characters")
	}
	if !helper.IsPassword(params.OldPassword) {
		return args, fmt.Errorf("Error validation: old password should contains at least uppercase, lowercase, and numeric")
	}

	// Password validation
	if helper.IsEmpty(params.Password) {
		return args, fmt.Errorf("Error validation: password can't be empty")
	}
	if len(params.Password) < alias.UserPasswordLengthMin {
		return args, fmt.Errorf("Error validation: password at least consist of 6 characters")
	}
	if !helper.IsPassword(params.Password) {
		return args, fmt.Errorf("Error validation: password should contains at least uppercase, lowercase, and numeric")
	}

	if params.Password != params.ConfirmPassword {
		return args, fmt.Errorf("password is not match")
	}

	args = changePasswordArgs{
		IdentityCode: identityCode,
		Email:        email,
		OldPassword:  helper.StringToMD5(params.OldPassword),
		Password:     helper.StringToMD5(params.Password),
	}
	return args, nil
}

// Function to validate forgot password parameter to be used
/*
	@params:
		Email		= required, email format, 0<characters<45
		IsSendCode	= optional, value=true or empty
		Code		= required if resend is empty, numeric, characters=4
		Password	= optional if code is empty, minimum 1 uppercase, lowercase, numeric, characters>=6
	@example:
		Email		= risal.falah@gmail.com
		IsSendCode	= true
		Code		= 1234 or empty if resend is true
		Password	= Qwerty123
	@return:
		Email		= risal.falah@gmail.com
		IsSendCode	= true
		Code		= 1234 or empty if resend is true
		Password	= Qwerty123
*/
func (params forgotParams) validate() (forgotArgs, error) {

	var args forgotArgs
	params = forgotParams{
		Email:      helper.Trim(params.Email),
		IsSendCode: params.IsSendCode,
		Code:       params.Code,
		Password:   html.EscapeString(params.Password),
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, fmt.Errorf("Error validation: %s", err.Error())
	}

	// IsSendCode validation
	if !helper.IsEmpty(params.IsSendCode) {
		if params.IsSendCode == "true" {
			return forgotArgs{
				Email:      email,
				IsSendCode: true,
				Code:       0,
				Password:   "",
			}, nil
		}
	}

	// Code Validation
	if helper.IsEmpty(params.Code) {
		return args, fmt.Errorf("Error validation: code cant't be empty")
	} else if len(params.Code) != alias.UserCodeLength {
		return args, fmt.Errorf("Error validation: code must be 4 digits")
	}
	code, err := strconv.ParseInt(params.Code, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Error validation: code should be numeric")
	}

	// Password Validation (Optional Field)
	if helper.IsEmpty(params.Password) {
		return forgotArgs{
			Email:      email,
			IsSendCode: false,
			Code:       uint16(code),
			Password:   "",
		}, nil
	}

	// Password validation
	if len(params.Password) < alias.UserPasswordLengthMin {
		return args, fmt.Errorf("Error validation: password at least consist of 6 characters")
	}
	if !helper.IsPassword(params.Password) {
		return args, fmt.Errorf("Error validation: password should contains at least uppercase, lowercase, and numeric")
	}

	args = forgotArgs{
		Email:      email,
		IsSendCode: false,
		Code:       uint16(code),
		Password:   helper.StringToMD5(params.Password),
	}
	return args, nil
}

// Function to validate detail information parameter to be used
/*
	@params:
		IdentityCode= required, numeric, characters=12
	@example:
		IdentityCode	= 140810140060
	@return:
		IdentityCode	= 140810140060
*/
func (params detailParams) validate() (detailArgs, error) {
	var args detailArgs
	identityCode, err := helper.NormalizeIdentity(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Error validation: ID should be numeric")
	}

	args = detailArgs{
		IdentityCode: identityCode,
	}
	return args, nil
}

// Function to validate update profile parameter to be used
/*
	@params:
		IdentityCode	= required, numeric, characters=12
		Email			= required, email format, 0<characters<45
		Name			= required, alphaspace, 0<characters<50
		Note			= optional, 0<character<100
		Gender			= optional, male or female
		Phone			= optional, numeric, 10<=characters<=12
		Line_id			= optional, 0<characters<=45
		Status			= required, actived or inactived
	@example:
		IdentityCode	= 140810140060
		Name 			= Khairil Azmi Ashari
		Note			= nothing is impossible
		Email			= khairilazmiashari@gmail.com
		Gender 			= male or female
		Phone 			= 082214467300
		Line_id 		= khaazas
		Status			= actived
	@return:
		IdentityCode	= 140810140060
		Name 			= Khairil Azmi Ashari
		Note			= nothing is impossible
		Email			= khairilazmiashari@gmail.com
		Gender 			= male or female
		Phone 			= 082214467300
		Line_id 		= khaazas
		Status			= actived
*/
func (params updateParams) validate() (updateArgs, error) {

	var args updateArgs
	params = updateParams{
		IdentityCode: helper.Trim(params.IdentityCode),
		Email:        helper.Trim(params.Email),
		Name:         helper.Trim(params.Name),
		Note:         helper.Trim(html.EscapeString(params.Note)),
		Gender:       params.Gender,
		Phone:        helper.Trim(params.Phone),
		LineID:       html.EscapeString(params.LineID),
		Status:       params.Status,
	}

	// Identity code validation
	identityCode, err := helper.NormalizeIdentity(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Invalid identity code")
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, fmt.Errorf("Invalid email")
	}

	// Name validation
	name, err := helper.NormalizeName(params.Name)
	if err != nil {
		return args, err
	}

	// Note validation
	if !helper.IsEmpty(params.Note) {
		if len(params.Note) > alias.UserNoteLengthMax {
			return args, fmt.Errorf("Error validation: Note too long")
		}
	}

	// Gender validation
	var gender int8 = user.GenderUndefined
	if !helper.IsEmpty(params.Gender) {
		switch params.Gender {
		case "male":
			gender = alias.UserGenderMale
		case "female":
			gender = alias.UserGenderFemale
		default:
			return args, fmt.Errorf("Error validation: wrong gender")
		}
	}

	// Phone validation (can be empty)
	var phone sql.NullString
	if !helper.IsEmpty(params.Phone) {
		if !helper.IsPhone(params.Phone) {
			return args, fmt.Errorf("Error validation: wrong input phone")
		}
		phone = sql.NullString{String: params.Phone, Valid: true}
	}

	// Line validation (can be empty)
	var lineID sql.NullString
	if !helper.IsEmpty(params.LineID) {
		if len(params.LineID) > alias.UserLineIDLengthMax {
			return args, fmt.Errorf("Error validation: Line Id too long")
		}
		lineID = sql.NullString{String: params.LineID, Valid: true}
	}

	// Status validation
	var status int8
	switch params.Status {
	case "active":
		status = user.StatusActivated
	case "inactive":
		status = user.StatusVerified
	default:
		return args, fmt.Errorf("Error validation: wrong status")
	}

	args = updateArgs{
		IdentityCode: identityCode,
		Name:         name,
		Email:        email,
		Gender:       gender,
		Phone:        phone,
		LineID:       lineID,
		Note:         params.Note,
		Status:       status,
	}

	return args, nil
}

// Function to validate delete profile parameter to be used
/*
	@params:
		IdentityCode= required, numeric, characters=12
	@example:
		IdentityCode	= 140810140060
	@return:
		IdentityCode	= 140810140060
*/
func (params deleteParams) validate() (deleteArgs, error) {
	var args deleteArgs
	identityCode, err := helper.NormalizeIdentity(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Error validation: ID should be numeric")
	}

	args = deleteArgs{
		IdentityCode: identityCode,
	}
	return args, nil
}

func (params createParams) validate() (createArgs, error) {
	var args createArgs
	params = createParams{
		IdentityCode: helper.Trim(params.IdentityCode),
		Name:         helper.Trim(params.Name),
		Email:        helper.Trim(params.Email),
	}

	// identity code validation
	identityCode, err := helper.NormalizeIdentity(params.IdentityCode)
	if err != nil {
		return args, fmt.Errorf("Error validation: ID should be numeric")
	}

	// Name validation
	name, err := helper.NormalizeName(params.Name)
	if err != nil {
		return args, fmt.Errorf("Error validation: %s", err.Error())
	}

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, fmt.Errorf("Error validation: %s", err.Error())
	}

	args = createArgs{
		IdentityCode: identityCode,
		Name:         name,
		Email:        email,
	}

	return args, nil
}
