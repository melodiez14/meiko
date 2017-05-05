package handler

import (
	"fmt"
	"regexp"

	f "github.com/melodiez14/meiko/src/webserver/function"
)

func (s signInParams) Validate() (*f.SignInArgs, error) {
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

	args := &f.SignInArgs{
		Email:    s.Email,
		Password: s.Password,
	}
	return args, nil
}
