package helper

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
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
	}

	splitted[length-1] = strings.Replace(splitted[length-1], " ", "", -1)

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
		string 	= Sunday
*/
func IntDayToString(day int8) string {

	if day < 0 || day > 6 {
		return ""
	}

	days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
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
		day 	= Sunday
	@return
		int8 	= 0
		error	= nil/true
*/
func DayStringToInt(day string) (int8, error) {
	day = strings.ToLower(day)

	switch day {
	case "sunday":
		return int8(time.Sunday), nil
	case "monday":
		return int8(time.Monday), nil
	case "tuesday":
		return int8(time.Tuesday), nil
	case "wednesday":
		return int8(time.Wednesday), nil
	case "thursday":
		return int8(time.Thursday), nil
	case "friday":
		return int8(time.Friday), nil
	case "saturday":
		return int8(time.Saturday), nil
	default:
		return 0, fmt.Errorf("not valid day")
	}
}

// Int64ToStringSlice ...
func Int64ToStringSlice(value []int64) []string {
	var str []string
	for _, val := range value {
		str = append(str, strconv.FormatInt(val, 10))
	}
	return str
}

// TimeToDayInt converts time.Time slice into days int8
/*
	@params:
		day		= time.Time
	@example:
		day 	= 2017-10-26 10:09:00.349054 +0700 WIB
	@return
		int8 	= 4
*/
func TimeToDayInt(time ...time.Time) []int8 {
	var days []int8
	for _, val := range time {
		days = append(days, int8(val.Weekday()))
	}
	return days
}

// Float64Round ...
func Float64Round(value float64) float64 {
	return math.Ceil(value - 0.5)
}

// Float32Round ...
func Float32Round(value float32) float32 {
	v := float64(value)
	return float32(math.Ceil(v - 0.5))
}
