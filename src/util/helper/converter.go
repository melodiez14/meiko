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

	fmt.Println(t1)
	fmt.Println(t2)

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
