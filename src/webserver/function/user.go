package function

import (
	"fmt"

	"github.com/melodiez14/meiko/src/module/user"
)

func (s *SignInArgs) SignIn() error {
	err := user.IsValidUserLogin(s.Email, s.Password)
	if err != nil {
		return fmt.Errorf("Email/Password is incorrect")
	}
	// set session to redis
	return nil
}
