package user

import (
	"fmt"
	"regexp"
	"strconv"
)

func (s signInParams) Validate() (*signInArgs, error) {
	const emailRegex = `(^[a-zA-Z0-9]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$)`
	if len(s.Email) < 1 {
		return nil, fmt.Errorf("Error validation : email cant't be empty")
	}

	v, err := regexp.MatchString(`(^[a-zA-Z0-9]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$)`, s.Email)
	if err != nil || v == false {
		return nil, fmt.Errorf("Error validation : email doesn't have a valid format")
	}

	if len(s.Password) < 1 {
		return nil, fmt.Errorf("Error validation : password cant't be empty")
	}

	// v, err = regexp.MatchString(`(^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&]{8,16})`, s.Email)
	// err = nil
	// if err != nil || v == false {
	// 	return nil, fmt.Errorf("Error validation : password doesn't have a valid format")
	// }

	args := &signInArgs{
		Email:    s.Email,
		Password: s.Password,
	}
	return args, nil
}

func (f forgotRequestParams) Validate() (*forgotRequestArgs, error) {
	const emailRegex = `(^[a-zA-Z0-9]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$)`
	if len(f.Email) < 1 {
		return nil, fmt.Errorf("Error validation : email cant't be empty")
	}

	v, err := regexp.MatchString(emailRegex, f.Email)
	if err != nil || v == false {
		return nil, fmt.Errorf("Error validation : email doesn't have a valid format")
	}

	args := &forgotRequestArgs{
		Email: f.Email,
	}
	return args, nil
}

func (f forgotConfirmationParams) Validate() (*forgotConfirmationArgs, error) {

	// fix the email validation regex
	const emailRegex = `(^[a-zA-Z0-9]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$)`
	if len(f.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	v, err := regexp.MatchString(emailRegex, f.Email)
	if err != nil || v == false {
		return nil, fmt.Errorf("Error validation: email doesn't have a valid format")
	}

	// Password == nil for code confirmation
	// Password != nil for set new password
	if len(f.Password) > 0 {
		v, err = regexp.MatchString(`(^[a-zA-Z0-9]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$)`, f.Password)
		if err != nil || v == false {
			return nil, fmt.Errorf("Error validation: password must be 6-8 characters")
		}
	}

	if len(f.Code) < 1 {
		return nil, fmt.Errorf("Error validation : code cant't be empty")
	} else if len(f.Code) != 4 {
		return nil, fmt.Errorf("Error validation : code must be 4 digits")
	}

	c, err := strconv.ParseInt(f.Code, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Error validation : code should be numeric")
	}

	args := &forgotConfirmationArgs{
		Email:    f.Email,
		Code:     uint16(c),
		Password: f.Password,
	}

	return args, nil
}
