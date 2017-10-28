package bot

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	cs "github.com/melodiez14/meiko/src/module/course"
	usr "github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/helper"
)

func handleAssistant(text string, userID int64) ([]map[string]interface{}, error) {

	var args []map[string]interface{}
	var filterDays []int8
	var filterDaysLen int
	var filterDates []time.Time
	var filterCourses []string
	var filterCoursesLen int
	var filterCoursesRgx *regexp.Regexp

	params := sEntity{
		text:   text,
		userID: userID,
	}

	// get day entity
	filterDays = params.getDay()

	// get date entity
	filterDates, err := params.getTime()
	if err != nil {
		return nil, err
	}
	// change time into days
	filterDays = append(filterDays, helper.TimeToDayInt(filterDates...)...)
	filterDaysLen = len(filterDays)
	fmt.Println(filterDates)

	// get course entity
	filterCourses = params.getCourse()
	filterCoursesLen = len(filterCourses)
	filterCoursesRgx = regexp.MustCompile(strings.Join(filterCourses, "|"))

	// select enrolled schedule by userID
	scheduleID, err := cs.SelectScheduleIDByUserID(userID)
	if err != nil {
		return args, nil
	}

	// select courses details by scheduleID
	courses, err := cs.SelectByScheduleID(scheduleID, cs.StatusScheduleActive)
	if err != nil {
		return args, nil
	}

	for _, val := range courses {

		// check if course name not match with regex
		if filterCoursesLen > 0 {
			if !filterCoursesRgx.MatchString(strings.ToLower(val.Course.Name)) {
				continue
			}
		}

		// check if day not match with entity filter
		if filterDaysLen > 0 {
			if !helper.Int8InSlice(val.Schedule.Day, filterDays) {
				continue
			}
		}

		// select assistant ID by schedule
		assistantID, err := cs.SelectAssistantID(val.Schedule.ID)
		if err != nil {
			return args, nil
		}

		// select assistant ID by schedule
		assistants, err := usr.SelectByID(assistantID, usr.ColName)
		if err != nil {
			return args, nil

		}

		// assistant
		var mapAssistant []map[string]string
		for _, v := range assistants {
			mapAssistant = append(mapAssistant, map[string]string{
				"name": v.Name,
			})
		}

		args = append(args, map[string]interface{}{
			"course_name": val.Course.Name,
			"assistant":   mapAssistant,
		})
	}

	return args, nil
}
