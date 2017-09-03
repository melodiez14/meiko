package user

import (
	"fmt"
	"html"
	"strconv"

	validator "gopkg.in/asaskevich/govalidator.v4"
)

func (s signInParams) Validate() (*signInArgs, error) {

	// Email Validation
	if len(s.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	email, err := validator.NormalizeEmail(html.EscapeString(s.Email))
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

	email, err := validator.NormalizeEmail(html.EscapeString(f.Email))
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

	email, err := validator.NormalizeEmail(html.EscapeString(f.Email))
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
