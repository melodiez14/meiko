package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/melodiez14/meiko/src/util/alias"
)

// IsAlpha Check if string given alphabet only
/*
	@params:
		text	= string
	@example:
		text	= khaazas
	@return
		true/false
*/
func IsAlpha(text string) bool {
	valid, _ := regexp.MatchString(`^[a-zA-Z]+$`, text)
	return valid
}

// IsAlphaSpace Check if string given alphabet and space only
/*
	@params:
		text	= string
	@example:
		text	= kha azas
	@return
		true/false
*/
func IsAlphaSpace(text string) bool {
	valid, _ := regexp.MatchString(`^[a-zA-Z ]+$`, text)
	return valid
}

// IsAlphaNumericSpace Check if string given alphabet, numeric and space only
/*
	@params:
		text	= string
	@example:
		text	= kha azas14001
	@return
		true/false
*/
func IsAlphaNumericSpace(text string) bool {
	valid, _ := regexp.MatchString(`^[a-zA-Z\d ]+$`, text)
	return valid
}

// IsPhone Check if given string is a phone number
/*
	@params:
		phone	= string
	@example:
		phone	= 082214467300
	@return
		true/false
*/
func IsPhone(phone string) bool {

	if len(phone) < 9 || len(phone) > 11 {
		return false
	}

	valid, _ := regexp.MatchString(`^\d+$`, phone)
	return valid
}

// IsEmail Check if given string is a email address 
/*
	@params:
		email	= string
	@example:
		email	= khairil_azmi_ashari@yahoo.com
	@return
		true/false
*/
func IsEmail(email string) bool {
	valid, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return valid
}

// IsPassword Check if given string is a valid password format
/*
	@params:
		password	= string
	@example:
		password	= khairil14001
	@return
		true/false
*/
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

// IsEmpty Check if text is empty or not
/*
	@params:
		text	= string
	@example:
		text	= 
	@return
		true/false
*/
func IsEmpty(text string) bool {
	if len(text) < 1 {
		return true
	}
	return false
}

// IsImageMime Check mime in the right format jpeg or jpg
/*
	@params:
		mime	= string
	@example:
		mime	= image/jpeg
	@return
		true/false
*/
func IsImageMime(mime string) bool {
	imageMime := []string{"image/jpeg", "image/png"}
	for _, v := range imageMime {
		if v == mime {
			return true
		}
	}
	return false
}

// IsImageExtension Check image extension jpeg, jpg or png
/*
	@params:
		extension	= string
	@example:
		extension	= jpeg
	@return
		true/false
*/
func IsImageExtension(extension string) bool {
	imageExtension := []string{"jpeg", "jpg", "png"}
	for _, val := range imageExtension {
		if val == extension {
			return true
		}
	}
	return false
}

// Normalize normalize input text to appropiate text
/*
	@params:
		text	= string
	@example:
		text	= This text would be normalize
	@return
		[]{text,true/false}
*/ 
func Normalize(text string, format func(string) bool) (string, error) {

	if !format(text) {
		return "", fmt.Errorf("text is not valid")
	}

	splitted := strings.Fields(text)
	return strings.Join(splitted, " "), nil
}

// NormalizeNPM normalize inputed NPM to be valid NPM number
/*
	@params:
		str	= string
	@example:
		text	= 140810140060
	@return
		[]{npm,error}
*/ 
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

// NormalizeIdentity Normalize inputed identity number to be valid identity number
/*
	@params:
		str	= string
	@example:
		text	= 1207261801970005
	@return
		[]{identity,error}
*/ 
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

// NormalizeName Normalize inputed name to be valid name
/*
	@params:
		name	= string
	@example:
		text	= Khairil azmi ashari
	@return
		[]{name,error}
*/ 
func NormalizeName(name string) (string, error) {

	splitted := strings.Fields(name)
	name = strings.Join(splitted, " ")

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

	return name, nil
}

// NormalizeCollege Normalize inputed college name to be valid name
/*
	@params:
		college	= string
	@example:
		college	= Universitas Padjadjaran
	@return
		[]{college,error}
*/ 
func NormalizeCollege(college string) (string, error) {
	splitted := strings.Fields(college)
	college = strings.Join(splitted, " ")

	if len(college) > alias.UserCollegeLengthMax {
		return "", fmt.Errorf("College is too long")
	}
	if IsEmpty(college) {
		return "", nil
	}

	return Normalize(college, IsAlphaSpace)
}

// NormalizeEmail Normalize inputed email addres to be valid email
/*
	@params:
		email	= string
	@example:
		email	= khairil14001@gmail.com
	@return
		[]{email,error}
*/ 
func NormalizeEmail(email string) (string, error) {
	if IsEmpty(email) {
		return "", fmt.Errorf("Email can't be empty")
	}

	if len(email) > alias.UserEmailLengthMax {
		return "", fmt.Errorf("Email is too long")
	}

	email = strings.ToLower(email)
	if !IsEmail(email) {
		return "", fmt.Errorf("Not valid email format")
	}

	parts := strings.Split(email, "@")
	if parts[1] == "gmail.com" || parts[1] == "googlemail.com" {
		parts[1] = "gmail.com"
	}
	return strings.Join(parts, "@"), nil
}

func Trim(str string) string {
	splitted := strings.Fields(str)
	return strings.Join(splitted, " ")
}
