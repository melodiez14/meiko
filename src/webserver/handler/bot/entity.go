package bot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (params *sEntity) getTime() ([]time.Time, error) {

	date := []time.Time{}
	now := time.Now()
	// a go pattern
	pattern := map[string]int64{
		`\d\s*hari\s(yang|\s)*lalu`:      (-1) * int64(time.Hour) * 24,
		`\d\s*minggu\s(yang|\s)*lalu`:    (-1) * int64(time.Hour) * 24 * 7,
		`\d\s*bulan\s(yang|\s)*lalu`:     (-1) * int64(time.Hour) * 24 * 30,
		`\d\s*hari\s(yang|\s)*kemarin`:   (-1) * int64(time.Hour) * 24,
		`\d\s*minggu\s(yang|\s)*kemarin`: (-1) * int64(time.Hour) * 24 * 7,
		`\d\s*bulan\s(yang|\s)*kemarin`:  (-1) * int64(time.Hour) * 24 * 30,
		`\d\s*hari\s(yang|\s)*kemaren`:   (-1) * int64(time.Hour) * 24,
		`\d\s*minggu\s(yang|\s)*kemaren`: (-1) * int64(time.Hour) * 24 * 7,
		`\d\s*bulan\s(yang|\s)*kemaren`:  (-1) * int64(time.Hour) * 24 * 30,
	}

	for i, val := range pattern {
		rgx := regexp.MustCompile(i)
		if rgx.MatchString(params.text) {
			// get the string
			str := rgx.FindAllString(params.text, -1)
			for _, vall := range str {
				// change regex into \d
				rgx = regexp.MustCompile(`\d`)
				// find numeric in the selected pattern
				dt := rgx.FindString(vall)
				// convert the string into integer
				digit, _ := strconv.ParseInt(dt, 10, 64)
				// times the matched integer with negative time duration
				sub := digit * val
				// substract datetime now with time duration
				date = append(date, now.Add(time.Duration(sub)))
				// replace the selected string from with (date)
				params.text = strings.Replace(params.text, vall, "(date)", -1)
			}
		}
	}

	// word expression
	patternWord := map[string]time.Time{
		`kemaren\s*lusa`:  now.AddDate(0, 0, -1),
		`pekan\s*kemarin`: now.AddDate(0, 0, -7),
		`kemarin`:         now.AddDate(0, 0, -2),
		`kemaren`:         now.AddDate(0, 0, -2),
		`hari\s*ini`:      now.AddDate(0, 0, -1),
	}

	for i, val := range patternWord {
		rgx := regexp.MustCompile(i)
		if rgx.MatchString(params.text) {
			// get the string
			str := rgx.FindString(params.text)
			// append the selected date
			date = append(date, val)
			// replace the selected string from with (date)
			params.text = strings.Replace(params.text, str, "(date)", -1)
		}
	}

	for i, val := range patternWord {
		rgx := regexp.MustCompile(i)
		if rgx.MatchString(params.text) {
			// get the string
			str := rgx.FindString(params.text)
			// append the selected date
			date = append(date, val)
			// replace the selected string from with (date)
			params.text = strings.Replace(params.text, str, "(date)", -1)
		}
	}

	if len(date) > 2 {
		return date, fmt.Errorf("Sorry this bot can only detect at least 2 date")
	}

	if len(date) == 2 {
		if date[0].After(date[1]) {
			temp := date[0]
			date[0] = date[1]
			date[1] = temp
		}
	}

	return date, nil
}

func (params *sEntity) getDay() []int8 {

	var days []int8

	// monday
	str := regexp.MustCompile(rgxMonday).FindAllString(params.text, -1)
	if len(str) > 0 {
		days = append(days, int8(time.Monday))
		for _, val := range str {
			params.text = strings.Replace(params.text, val, ("(day)"), -1)
		}
	}

	// tuesday
	str = regexp.MustCompile(rgxTuesday).FindAllString(params.text, -1)
	if len(str) > 0 {
		days = append(days, int8(time.Tuesday))
		for _, val := range str {
			params.text = strings.Replace(params.text, val, ("(day)"), -1)
		}
	}

	// wednesday
	str = regexp.MustCompile(rgxWednesday).FindAllString(params.text, -1)
	if len(str) > 0 {
		days = append(days, int8(time.Wednesday))
		for _, val := range str {
			params.text = strings.Replace(params.text, val, ("(day)"), -1)
		}
	}

	// thursday
	str = regexp.MustCompile(rgxThursday).FindAllString(params.text, -1)
	if len(str) > 0 {
		days = append(days, int8(time.Thursday))
		for _, val := range str {
			params.text = strings.Replace(params.text, val, ("(day)"), -1)
		}
	}

	// friday
	str = regexp.MustCompile(rgxFriday).FindAllString(params.text, -1)
	if len(str) > 0 {
		days = append(days, int8(time.Friday))
		for _, val := range str {
			params.text = strings.Replace(params.text, val, ("(day)"), -1)
		}
	}

	// saturday
	str = regexp.MustCompile(rgxSaturday).FindAllString(params.text, -1)
	if len(str) > 0 {
		days = append(days, int8(time.Saturday))
		for _, val := range str {
			params.text = strings.Replace(params.text, val, ("(day)"), -1)
		}
	}

	// sunday
	str = regexp.MustCompile(rgxSunday).FindAllString(params.text, -1)
	if len(str) > 0 {
		days = append(days, int8(time.Sunday))
		for _, val := range str {
			params.text = strings.Replace(params.text, val, ("(day)"), -1)
		}
	}

	return days
}

func (params *sEntity) getAssistant() []string {
	rgx := regexp.MustCompile(rgxAssistant)
	str := rgx.FindAllString(params.text, -1)
	for _, val := range str {
		params.text = strings.Replace(params.text, val, ("(assistant)"), -1)
	}
	return str
}

func (params *sEntity) getCourse() []string {
	rgx := regexp.MustCompile(rgxCourse)
	str := rgx.FindAllString(params.text, -1)
	for _, val := range str {
		params.text = strings.Replace(params.text, val, ("(course)"), -1)
	}
	return str
}
