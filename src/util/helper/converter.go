package helper

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
	"time"
)

func StringToMD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func ExtractExtension(fileName string) (string, string, error) {

	splitted := strings.Split(fileName, ".")
	length := len(splitted)

	if length < 2 {
		return "", "", fmt.Errorf("Doesn't have extension")
	}

	lastIndex := length - 1
	fn := strings.Join(splitted[:lastIndex], ".")
	ext := splitted[lastIndex]

	return fn, ext, nil
}

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
		return t1.Format("Monday, 1 January 2006")
	}

}

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
