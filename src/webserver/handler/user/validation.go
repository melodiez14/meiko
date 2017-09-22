package user

import (
	"fmt"
	"html"
	"strconv"

	"github.com/melodiez14/meiko/src/util/alias"

	"regexp"

	"github.com/melodiez14/meiko/src/util/helper"
	validator "gopkg.in/asaskevich/govalidator.v4"
)

func (params signUpParams) Validate() (signUpArgs, error) {

	var args signUpArgs
	params = signUpParams{
		ID:       html.EscapeString(params.ID),
		Name:     html.EscapeString(params.Name),
		Email:    html.EscapeString(params.Email),
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

func (params setUserAccoutParams) Validate() (setUserAccoutArgs, error) {

	var args setUserAccoutArgs
	params = setUserAccoutParams{
		Name:    params.Name,
		Note:    html.EscapeString(params.Note),
		Gender:  params.Gender,
		College: params.College,
		Phone:   params.Phone,
		LineID:  html.EscapeString(params.LineID),
	}

	// Name validation
	if len(params.Name) < 1 {
		return args, fmt.Errorf("Error validation: name cant't be empty")
	}
	if len(params.Name) > 50 {
		return args, fmt.Errorf("Error validation: name is too long")
	}
	name, err := helper.NormalizeName(params.Name)
	if err != nil {
		return args, err
	}

	// gender validation
	var gender int8 = alias.UserGenderUndefined
	if len(params.Gender) > 0 {
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
	if len(params.Phone) > 0 {
		if !helper.IsPhone(params.Phone) {
			return args, fmt.Errorf("Error validation: wrong input gender")
		}
	}

	// Line verification (can be empty)
	if len(params.LineID) > 0 {
		if len(params.LineID) > 45 {
			return args, fmt.Errorf("Error validation: Line Id too long")
		}
	}

	// College validation (need more validation)
	var college string
	if len(params.College) > 0 {
		if len(params.College) > 100 {
			return args, fmt.Errorf("Error validation: College too long")
		}
		college, err = helper.NormalizeCollege(params.College)
		if err != nil {
			return args, fmt.Errorf("Error validation: Not valid college")
		}

	}

	// Note validation
	if len(params.Note) > 0 {
		if len(params.Note) > 100 {
			return args, fmt.Errorf("Error validation: Note too long")
		}
	}

	args = setUserAccoutArgs{
		Name:    name,
		Gender:  gender,
		Phone:   params.Phone,
		LineID:  params.LineID,
		College: college,
		Note:    params.Note,
	}

	return args, nil
}

func (s setChangePasswordParams) Validate() (*setChangePasswordArgs, error) {
	// Password
	if len(s.Password) < 1 {
		return nil, fmt.Errorf("Error validation: password can't be empty")
	} else if len(s.ConfirmPassword) < 1 {
		return nil, fmt.Errorf("Eroor validation: Confirmation password can't be empty")
	}

	if len(s.Password) < 6 {
		return nil, fmt.Errorf("Error validation: password at least consist of 6 characters")
	} else if len(s.ConfirmPassword) < 6 {
		return nil, fmt.Errorf("Error validation : Confirmation password at least consist of 6 characters")
	}

	password := html.EscapeString(s.Password)
	cp := html.EscapeString(s.ConfirmPassword)

	regexPassword := []string{`[a-z]`, `[A-Z]`, `[0-9]`}
	for _, val := range regexPassword {
		is, _ := regexp.MatchString(val, password)
		if !is {
			return nil, fmt.Errorf("Error validation: password must contains alphanumeric upper and lower case")
		}
	}

	for _, val := range regexPassword {
		is, _ := regexp.MatchString(val, cp)
		if !is {
			return nil, fmt.Errorf("Error validation: password must contains alphanumeric upper and lower case")
		}
	}

	if password != cp {
		return nil, fmt.Errorf("Error validation: Password not match")
	}
	args := &setChangePasswordArgs{
		Password: password,
	}

	return args, nil

}
func (s setStatusUserParams) Validate() (*setStatusUserArgs, error) {
	// Email Validation
	if len(s.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}
	if len(s.Email) > 45 {
		return nil, fmt.Errorf("Error validation : email too longer")
	}

	if !validator.IsEmail(s.Email) {
		return nil, fmt.Errorf("%s is not an email", s.Email)
	}
	email, err := helper.NormalizeEmail(html.EscapeString(s.Email))
	if err != nil {
		return nil, err
	}

	//Code Validation
	if len(s.Code) < 1 {
		return nil, fmt.Errorf("Error validation : Code can't be empty")
	} else if len(s.Code) != 4 {
		return nil, fmt.Errorf("Error validation : Wrong code")
	}
	val, err := regexp.MatchString(`[0-9]+$`, s.Code)
	if !val || err != nil {
		return nil, fmt.Errorf("Error validation: Wrong code")
	}

	c, err := strconv.ParseInt(s.Code, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Error validation: Wrong code")
	}

	args := &setStatusUserArgs{
		Email: email,
		Code:  uint16(c),
	}

	return args, nil

}

func (s signInParams) Validate() (*signInArgs, error) {

	// Email Validation
	if len(s.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	email, err := helper.NormalizeEmail(html.EscapeString(s.Email))
	if err != nil {
		return nil, err
	}

	// Password Validation
	password := html.EscapeString(s.Password)
	if len(password) < 6 {
		return nil, fmt.Errorf("Error validation: password at least consist of 6 characters")
	}

	args := &signInArgs{
		Email:    email,
		Password: password,
	}
	return args, nil
}

func (f forgotRequestParams) Validate() (*forgotRequestArgs, error) {

	// Email Validation
	if len(f.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	email, err := helper.NormalizeEmail(html.EscapeString(f.Email))
	if err != nil {
		return nil, err
	}

	args := &forgotRequestArgs{
		Email: email,
	}
	return args, nil
}

func (f forgotConfirmationParams) Validate() (*forgotConfirmationArgs, error) {

	// Email Validation
	if len(f.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	email, err := helper.NormalizeEmail(html.EscapeString(f.Email))
	if err != nil {
		return nil, err
	}

	// Password Validation (Optional Field)
	if len(f.Password) > 0 {
		f.Password = html.EscapeString(f.Password)
		if len(f.Password) < 6 {
			return nil, fmt.Errorf("Error validation: password at least consist of 6 characters")
		}
	}

	// Code Validation
	if len(f.Code) < 1 {
		return nil, fmt.Errorf("Error validation: code cant't be empty")
	} else if len(f.Code) != 4 {
		return nil, fmt.Errorf("Error validation: code must be 4 digits")
	}

	c, err := strconv.ParseInt(f.Code, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Error validation: code should be numeric")
	}

	args := &forgotConfirmationArgs{
		Email:    email,
		Code:     uint16(c),
		Password: f.Password,
	}

	return args, nil
}

func (param activationParams) Validate() (*activationArgs, error) {

	// Check is param empty
	if len(param.ID) < 1 || len(param.Status) < 1 {
		return nil, fmt.Errorf("Bad Request")
	}

	id, err := strconv.ParseInt(param.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error validation: ID should be numeric")
	}

	var status int8
	switch param.Status {
	case "active":
		status = alias.UserStatusActivated
	case "inactive":
		status = alias.UserStatusVerified
	default:
		return nil, fmt.Errorf("Error validation: wrong status")
	}

	args := &activationArgs{
		ID:     id,
		Status: status,
	}

	return args, nil
}
