package user

import (
	"fmt"
	"html"
	"strconv"

	"github.com/melodiez14/meiko/src/util/alias"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params signUpParams) Validate() (signUpArgs, error) {

	var args signUpArgs
	params = signUpParams{
		ID:       params.ID,
		Name:     params.Name,
		Email:    params.Email,
		Password: html.EscapeString(params.Password),
	}

	// ID validation
	id, err := helper.NormalizeUserID(params.ID)
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
		ID:       id,
		Name:     name,
		Email:    email,
		Password: params.Password,
	}
	return args, nil
}

func (params emailVerificationParams) Validate() (emailVerificationArgs, error) {

	var args emailVerificationArgs

	// Email validation
	email, err := helper.NormalizeEmail(params.Email)
	if err != nil {
		return args, err
	}

	// IsResendCode validation
	if !helper.IsEmpty(params.IsResendCode) {
		if params.IsResendCode != "true" {
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

func (params getVerifiedParams) Validate() (getVerifiedArgs, error) {

	var args getVerifiedArgs
	if helper.IsEmpty(params.Page) || helper.IsEmpty(params.Total) {
		return args, fmt.Errorf("Invalid request")
	}

	page, err := strconv.ParseInt(params.Page, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	total, err := strconv.ParseInt(params.Total, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	// should be positive number
	if page < 0 || total < 0 {
		return args, fmt.Errorf("Invalid request")
	}

	args = getVerifiedArgs{
		Page:  page,
		Total: total,
	}
	return args, nil
}

func (params activationParams) Validate() (activationArgs, error) {

	var args activationArgs
	// Check is params empty
	if helper.IsEmpty(params.ID) || helper.IsEmpty(params.Status) {
		return args, fmt.Errorf("Bad Request")
	}

	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error validation: ID should be numeric")
	}

	var status int8
	switch params.Status {
	case "active":
		status = alias.UserStatusActivated
	case "inactive":
		status = alias.UserStatusVerified
	default:
		return args, fmt.Errorf("Error validation: wrong status")
	}

	args = activationArgs{
		ID:     id,
		Status: status,
	}

	return args, nil
}

func (params signInParams) Validate() (signInArgs, error) {

	var args signInArgs
	params = signInParams{
		Email:    params.Email,
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
		Password: params.Password,
	}
	return args, nil
}

func (params updateProfileParams) Validate() (updateProfileArgs, error) {

	var args updateProfileArgs
	params = updateProfileParams{
		Name:    params.Name,
		Note:    html.EscapeString(params.Note),
		Gender:  params.Gender,
		College: params.College,
		Phone:   params.Phone,
		LineID:  html.EscapeString(params.LineID),
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
	var gender int8 = alias.UserGenderUndefined
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

	// College validation (need more validation)
	var college string
	if !helper.IsEmpty(params.College) {
		college, err = helper.NormalizeCollege(params.College)
		if err != nil {
			return args, fmt.Errorf("Error validation: Not valid college")
		}
	}

	// Phone validation (can be empty)
	if !helper.IsEmpty(params.Phone) {
		if !helper.IsPhone(params.Phone) {
			return args, fmt.Errorf("Error validation: wrong input phone")
		}
	}

	// Line verification (can be empty)
	if !helper.IsEmpty(params.LineID) {
		if len(params.LineID) > alias.UserLineIDLengthMax {
			return args, fmt.Errorf("Error validation: Line Id too long")
		}
	}

	args = updateProfileArgs{
		Name:    name,
		Gender:  gender,
		Phone:   params.Phone,
		LineID:  params.LineID,
		College: college,
		Note:    params.Note,
	}

	return args, nil
}

func (params changePasswordParams) Validate() (changePasswordArgs, error) {

	var args changePasswordArgs
	params = changePasswordParams{
		OldPassword:     html.EscapeString(params.OldPassword),
		Password:        html.EscapeString(params.Password),
		ConfirmPassword: html.EscapeString(params.ConfirmPassword),
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
		OldPassword: helper.StringToMD5(params.OldPassword),
		Password:    helper.StringToMD5(params.Password),
	}
	return args, nil
}

func (params forgotParams) Validate() (forgotArgs, error) {

	var args forgotArgs
	params = forgotParams{
		Email:      params.Email,
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
		Password:   params.Password,
	}
	return args, nil
}
