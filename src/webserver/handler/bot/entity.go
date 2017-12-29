package bot

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	nw "github.com/jinzhu/now"
)

func (params *sEntity) getTime() []time.Time {

	date := []time.Time{}
	now := time.Now()

	// days 2 time pattern
	patternDays2Time := []string{
		`\d\s*hari\s(yang|\s)*lalu`,
		`\d\s*hari\s(yang|\s)*kemarin`,
		`\d\s*hari\s(yang|\s)*kemaren`,
	}

	for _, val := range patternDays2Time {
		rgx := regexp.MustCompile(val)
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
				sub := digit * ((-1) * int64(time.Hour) * 24 * 1)
				// substract datetime now with time duration
				date = append(date, now.Add(time.Duration(sub)))
				// replace the selected string from with (date)
				params.text = strings.Replace(params.text, vall, "(date)", -1)
			}
		}
	}

	// week 2 time pattern
	patternWeek2Time := []string{
		`\d\s*minggu\s(yang|\s)*kemarin`,
		`\d\s*minggu\s(yang|\s)*lalu`,
		`\d\s*minggu\s(yang|\s)*kemaren`,
	}

	for _, val := range patternWeek2Time {
		rgx := regexp.MustCompile(val)
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
				sub := digit * ((-1) * int64(time.Hour) * 24 * 7)
				// get time of last week
				lastWeekTime := now.Add(time.Duration(sub))
				// convert get isoweek starttime
				startTime := nw.New(lastWeekTime).BeginningOfWeek()
				endTime := nw.New(startTime).EndOfWeek()
				// substract datetime now with time duration
				date = append(date, startTime)
				date = append(date, endTime)
				// replace the selected string from with (date)
				params.text = strings.Replace(params.text, vall, "(date)", -1)
			}
		}
	}

	patternMonth2Time := []string{
		`\d\s*bulan\s(yang|\s)*lalu`,
		`\d\s*bulan\s(yang|\s)*kemarin`,
		`\d\s*bulan\s(yang|\s)*kemaren`,
	}

	for _, val := range patternMonth2Time {
		rgx := regexp.MustCompile(val)
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
				sub := digit * ((-1) * int64(time.Hour) * 24 * 30)
				// get time of last week
				lastMonthTime := now.Add(time.Duration(sub))
				// convert get lastmonthtime starttime
				startTime := nw.New(lastMonthTime).BeginningOfMonth()
				endTime := nw.New(startTime).EndOfMonth()
				// substract datetime now with time duration
				date = append(date, startTime)
				date = append(date, endTime)
				// replace the selected string from with (date)
				params.text = strings.Replace(params.text, vall, "(date)", -1)
			}
		}
	}

	// word expression
	patternWord := map[string]time.Time{
		`kemaren\s*lusa`:  now.AddDate(0, 0, -2),
		`pekan\s*kemarin`: now.AddDate(0, 0, -7),
		`kemarin`:         now.AddDate(0, 0, -1),
		`kemaren`:         now.AddDate(0, 0, -1),
		`hari\s*ini`:      now.AddDate(0, 0, 0),
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

	if len(date) > 1 {
		startDate := date[0]
		endDate := date[0]
		for _, val := range date {
			if val.Before(startDate) {
				startDate = val
			} else if val.After(endDate) {
				endDate = val
			}
		}
		date = []time.Time{startDate, endDate}
	}

	return date
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
