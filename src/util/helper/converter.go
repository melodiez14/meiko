package helper

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
	"time"
)

// StringToMD5 convert inputed text to MD5 format
/*
	@params:
		data	= string
	@example:
		data 	= me
	@return
		MD5		= ab86a1e1ef70dff97959067b723c5c24
*/
func StringToMD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// ExtractExtension extract file name given to be seperate into file name and it's extension
/*
	@params:
		fileName	= string
	@example:
		fileName 	= me.pdf
	@return
		fn			= me
		ext			= pdf
		error		= nil/false
*/
func ExtractExtension(fileName string) (string, string, error) {

	splitted := strings.Split(fileName, ".")
	length := len(splitted)
	lastChar := string(fileName[len(fileName)-1])

	if length < 2 || lastChar == "." {
		return "", "", fmt.Errorf("Doesn't have extension")
	}
	if strings.Replace(splitted[length-1], " ", "", -1) == "" {
		return "", "", fmt.Errorf("Doesn't have extension")
	} else {
		splitted[length-1] = strings.Replace(splitted[length-1], " ", "", -1)
	}

	lastIndex := length - 1
	fn := strings.Join(splitted[:lastIndex], ".")
	ext := splitted[lastIndex]

	if !IsAlpha(ext) {
		return "", "", fmt.Errorf("Doesn't have extension")
	}

	return fn, ext, nil
}

// DateToString Give time range from accesse time by user
/*
	@params:
		t1		= time reference
		t2		= time acces by user
	@example:
		t1 		= 18 October 2017, 5:25:10
		t2		= 18 October 2017, 5:25:15
	@return
		string 	= 5 seconds
*/
func DateToString(t1, t2 time.Time) string {

	ts := t2.Sub(t1)
	seconds := int64(math.Floor(ts.Seconds() + 0.5))
	minutes := int64(math.Floor(ts.Minutes() + 0.5))
	hours := int64(math.Floor(ts.Hours() + 0.5))
	days := hours / 24

	if seconds < 60 {
		return "Just now"
	} else if minutes < 60 {
		return fmt.Sprintf("%d minutes ago", minutes)
	} else if hours < 24 {
		return fmt.Sprintf("%d hours ago", hours)
	} else if days < 4 {
		return fmt.Sprintf("%d days ago", days)
	} else {
		return t2.Format("Monday, 1 January 2006")
	}

}

// IntDayToString Change day format from time format to string
/*
	@params:
		day		= int8
	@example:
		day 	= 0
	@return
		string 	= Moday
*/
func IntDayToString(day int8) string {

	if day < 0 || day > 6 {
		return ""
	}

	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	return days[day]
}

// MinutesToTimeString Change minutes format from time format to string
/*
	@params:
		minutes	= uint16
	@example:
		minutes = 0
	@return
		string 	= 00:00
*/
func MinutesToTimeString(minutes uint16) string {
	if minutes > 1440 {
		return ""
	}
	var h, m string
	hh := minutes / 60
	mm := minutes % 60

	if hh < 10 {
		h = fmt.Sprintf("0%d", hh)
	} else {
		h = fmt.Sprintf("%d", hh)
	}

	if mm < 10 {
		m = fmt.Sprintf("0%d", mm)
	} else {
		m = fmt.Sprintf("%d", mm)
	}
	return fmt.Sprintf("%s:%s", h, m)
}

// DayStringToInt Change day format fromSstring format to Int
/*
	@params:
		day		= String
	@example:
		day 	= Monday
	@return
		int8 	= 0
		error	= nil/true
*/
func DayStringToInt(day string) (int8, error) {
	day = strings.ToLower(day)
	days := []string{
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
	}

	for i, val := range days {
		if day == val {
			return int8(i), nil
		}
	}

	return 0, fmt.Errorf("Not valid day")
}
