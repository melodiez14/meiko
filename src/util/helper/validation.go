package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/melodiez14/meiko/src/util/alias"
)

func IsAlpha(text string) bool {
	valid, err := regexp.MatchString(`^[a-zA-Z ]+$`, text)
	if err != nil {
		return false
	}
	return valid
}

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

func IsImageMime(mime string) bool {
	imageMime := []string{"image/jpeg", "image/png"}
	for _, v := range imageMime {
		if v == mime {
			return true
		}
	}
	return false
}

func IsImageExtension(extension string) bool {
	imageExtension := []string{"jpeg", "jpg", "png"}
	for _, val := range imageExtension {
		if val == extension {
			return true
		}
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

func NormalizeNPM(str string) (int64, error) {
	var npm int64
	if len(str) < 1 {
		return npm, fmt.Errorf("npm can't be empty")
	}
	if len(str) != 12 {
		return npm, fmt.Errorf("invalid npm")
	}
	npm, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return npm, fmt.Errorf("npm must be numeric")
	}
	return npm, nil
}

func NormalizeIdentity(str string) (int64, error) {
	var identity int64
	if len(str) < 1 {
		return identity, fmt.Errorf("identity can't be empty")
	}
	if len(str) < 10 || len(str) > 18 {
		return identity, fmt.Errorf("invalid npm, nidn, nip, ktp, or sim number")
	}
	identity, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return identity, fmt.Errorf("identity must be numeric")
	}
	return identity, nil
}

func NormalizeName(name string) (string, error) {
	if IsEmpty(name) {
		return "", fmt.Errorf("name cant't be empty")
	}
	if len(name) > alias.UserNameLengthMax {
		return "", fmt.Errorf("name is too long")
	}
	name, err := Normalize(name, IsAlphaSpace)
	if err != nil {
		return "", fmt.Errorf("name should be alphabet and space only")
	}
	return Normalize(name, IsAlphaSpace)
}

func NormalizeCollege(college string) (string, error) {

	if len(college) > alias.UserCollegeLengthMax {
		return "", fmt.Errorf("College is too long")
	}

	return Normalize(college, IsAlphaSpace)
}

func NormalizeEmail(email string) (string, error) {

	if IsEmpty(email) {
		return "", fmt.Errorf("Email can't be empty")
	}

	if len(email) > alias.UserEmailLengthMax {
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
