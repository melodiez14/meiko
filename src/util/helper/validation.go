package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/melodiez14/meiko/src/util/alias"
)

func IsAlphaSpace(text string) bool {
	valid, err := regexp.MatchString(`^[a-zA-Z ]+$`, text)
	if err != nil {
		return false
	}
	return valid
}

func IsAlphaNumericSpace(text string) bool {
	valid, err := regexp.MatchString(`^[a-zA-Z\d ]+$`, text)
	if err != nil {
		return false
	}
	return valid
}

func IsPhone(phone string) bool {

	if len(phone) < 10 || len(phone) > 12 {
		return false
	}

	valid, err := regexp.MatchString(`^\d+$`, phone)
	if err != nil {
		return false
	}
	return valid
}

func IsEmail(email string) bool {
	valid, err := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if err != nil {
		return false
	}
	return valid
}

func IsPassword(password string) bool {
	regex := []string{`[a-z]`, `[A-Z]`, `[0-9]`}
	for _, val := range regex {
		v, _ := regexp.MatchString(val, password)
		if !v {
			return false
		}
	}
	return true
}

func IsEmpty(text string) bool {
	if len(text) < 1 {
		return true
	}
	return false
}

func Normalize(text string, format func(string) bool) (string, error) {

	if !format(text) {
		return "", fmt.Errorf("text is not valid")
	}

	splitted := strings.Fields(text)
	return strings.Join(splitted, " "), nil
}

func NormalizeUserID(id string) (int64, error) {
	var userID int64
	if len(id) < 1 {
		return userID, fmt.Errorf("user id can't be empty")
	}
	if len(id) != 12 {
		return userID, fmt.Errorf("invalid user id")
	}
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return userID, fmt.Errorf("user id must be numeric")
	}
	return userID, nil
}

func NormalizeName(name string) (string, error) {
	if IsEmpty(name) {
		return "", fmt.Errorf("name cant't be empty")
	}
	if len(name) > alias.UserNameLengthMax {
		return "", fmt.Errorf("name is too long")
	}
	return Normalize(name, IsAlphaSpace)
}

func NormalizeCollege(college string) (string, error) {
	return Normalize(college, IsAlphaSpace)
}

func NormalizeEmail(email string) (string, error) {

	if len(email) < 1 {
		return "", fmt.Errorf("Email can't be empty")
	}

	if len(email) > 45 {
		return "", fmt.Errorf("Email is too long")
	}

	if !IsEmail(email) {
		return "", fmt.Errorf("Not valid email format")
	}

	parts := strings.Split(email, "@")
	parts[0] = strings.ToLower(parts[0])
	parts[1] = strings.ToLower(parts[1])
	if parts[1] == "gmail.com" || parts[1] == "googlemail.com" {
		parts[1] = "gmail.com"
	}
	return strings.Join(parts, "@"), nil
}
