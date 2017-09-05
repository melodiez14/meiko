package helper

import (
	"strings"
)

func Int64InSlice(val int64, arr []int64) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func IsStringInSlice(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func NormalizeEmail(str string) (string, error) {

	parts := strings.Split(str, "@")
	parts[0] = strings.ToLower(parts[0])
	parts[1] = strings.ToLower(parts[1])
	if parts[1] == "gmail.com" || parts[1] == "googlemail.com" {
		parts[1] = "gmail.com"
	}
	return strings.Join(parts, "@"), nil
}
