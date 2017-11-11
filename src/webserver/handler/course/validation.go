package course

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"strings"

	cs "github.com/melodiez14/meiko/src/module/course"
	"github.com/melodiez14/meiko/src/util/helper"
)

func (params readParams) validate() (readArgs, error) {

	var args readArgs
	if helper.IsEmpty(params.Page) || helper.IsEmpty(params.Total) {
		return args, fmt.Errorf("page or total is empty")
	}

	page, err := strconv.ParseInt(params.Page, 10, 64)
	if err != nil {
		return args, fmt.Errorf("page must be numeric")
	}

	total, err := strconv.ParseInt(params.Total, 10, 64)
	if err != nil {
		return args, fmt.Errorf("total must be numeric")
	}

	// should be positive number
	if page < 0 || total < 0 {
		return args, fmt.Errorf("page or total must be positive number")
	}

	args = readArgs{
		Page:  uint16(page),
		Total: uint16(total),
	}
	return args, nil
}

func (params createParams) validate() (createArgs, error) {

	var args createArgs
	params = createParams{
		ID:             html.EscapeString(helper.Trim(params.ID)),
		Name:           helper.Trim(params.Name),
		Description:    html.EscapeString(helper.Trim(params.Description)),
		UCU:            params.UCU,
		Semester:       params.Semester,
		Year:           params.Year,
		StartTime:      params.StartTime,
		EndTime:        params.EndTime,
		Class:          strings.ToUpper(helper.Trim(params.Class)),
		Day:            strings.ToLower(params.Day),
		PlaceID:        html.EscapeString(strings.ToUpper(helper.Trim(params.PlaceID))),
		IsUpdate:       params.IsUpdate,
		GradeParameter: params.GradeParameter,
	}

	// Course Validation
	if helper.IsEmpty(params.ID) {
		return args, fmt.Errorf("course_id cannot be empty")
	}
	if len(params.ID) > cs.MaximumID {
		return args, fmt.Errorf("course_id can have only maximum 45 characters length")
	}

	// Name validation
	name, err := helper.NormalizeName(params.Name)
	if err != nil {
		return args, err
	}

	// Description validation
	var description sql.NullString
	if !helper.IsEmpty(params.Description) {
		description = sql.NullString{Valid: true, String: params.Description}
	}

	// UCU validation
	if helper.IsEmpty(params.UCU) {
		return args, fmt.Errorf("UCU can't be empty")
	}
	ucu, err := strconv.ParseInt(params.UCU, 10, 8)
	if err != nil {
		return args, fmt.Errorf("UCU must be numeric")
	}
	if ucu < 1 || ucu > 5 {
		return args, fmt.Errorf("Invalid UCU")
	}

	// Semester validation
	if helper.IsEmpty(params.Semester) {
		return args, fmt.Errorf("Semester can't be empty")
	}
	semester, err := strconv.ParseInt(params.Semester, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Semester must be numeric")
	}
	if semester < 1 || semester > 7 {
		return args, fmt.Errorf("Invalid semester")
	}

	// Year validation
	if helper.IsEmpty(params.Year) {
		return args, fmt.Errorf("Year can't be empty")
	}
	year, err := strconv.ParseInt(params.Year, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Year must be numeric")
	}
	if year < 2017 || year > 2020 {
		return args, fmt.Errorf("Invalid year")
	}

	// Start time validation
	if helper.IsEmpty(params.StartTime) {
		return args, fmt.Errorf("Start time can't be empty")
	}
	startTime, err := strconv.ParseInt(params.StartTime, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Invalid start time")
	}
	if startTime < 0 || startTime >= 1440 {
		return args, fmt.Errorf("Invalid start time")
	}

	// End time validation
	if helper.IsEmpty(params.EndTime) {
		return args, fmt.Errorf("End time can't be empty")
	}
	endTime, err := strconv.ParseInt(params.EndTime, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Invalid end time")
	}
	if endTime < 0 || endTime >= 1440 {
		return args, fmt.Errorf("Invalid end time")
	}

	if startTime > endTime {
		return args, fmt.Errorf("Start time more than end time")
	}

	// Class validation
	if helper.IsEmpty(params.Class) {
		return args, fmt.Errorf("Class can't be empty")
	}
	if len(params.Class) != 1 || !helper.IsAlpha(params.Class) {
		return args, fmt.Errorf("Invalid class")
	}
	class := strings.ToUpper(params.Class)

	// Day validation
	if helper.IsEmpty(params.Day) {
		return args, fmt.Errorf("Day can't be empty")
	}
	day, err := helper.DayStringToInt(params.Day)
	if err != nil {
		return args, err
	}

	// Place ID validation
	if helper.IsEmpty(params.PlaceID) {
		return args, fmt.Errorf("Place ID can't be empty")
	}
	if len(params.PlaceID) > 30 {
		return args, fmt.Errorf("Invalid place id")
	}

	isUpdate := false
	// Is Update Course
	if params.IsUpdate == "true" {
		isUpdate = true
	}

	var gps []gradeParameter
	if !helper.IsEmpty(params.GradeParameter) {
		var gp []gradeParameter
		// convert json into gradeParameter struct
		err = json.Unmarshal([]byte(params.GradeParameter), &gp)
		if err != nil {
			return args, fmt.Errorf("Invalid grade parameter")
		}

		var percentage float32
		gpType := []string{
			cs.GradeParameterAssignment,
			cs.GradeParameterAttendance,
			cs.GradeParameterFinal,
			cs.GradeParameterMid,
			cs.GradeParameterQuiz,
		}

		for _, val := range gp {
			percentage += val.Percentage
			// check status change
			if val.StatusChange != cs.GradeParameterStatusChange && val.StatusChange != cs.GradeParameterStatusUnchange {
				return args, fmt.Errorf("Invalid grade parameter")
			}
			// validate type
			if !helper.IsStringInSlice(val.Type, gpType) {
				return args, fmt.Errorf("Invalid grade parameter")
			}
			gps = append(gps, val)
		}

		// validate total percentage should be 100
		if helper.Float32Round(percentage) != 100 {
			return args, fmt.Errorf("Total percentage must be 100%%")
		}
	}

	return createArgs{
		ID:             params.ID,
		Name:           name,
		Description:    description,
		UCU:            int8(ucu),
		Semester:       int8(semester),
		Year:           int16(year),
		StartTime:      int16(startTime),
		EndTime:        int16(endTime),
		Class:          class,
		Day:            day,
		PlaceID:        params.PlaceID,
		IsUpdate:       isUpdate,
		GradeParameter: gps,
	}, nil
}

func (params getParams) validate() (getArgs, error) {

	var args getArgs
	if params.Payload != "last" && params.Payload != "current" && params.Payload != "all" {
		return args, fmt.Errorf("Invalid payload")
	}

	args = getArgs{
		Payload: params.Payload,
	}

	return args, nil
}

func (params getAssistantParams) validate() (getAssistantArgs, error) {

	var args getAssistantArgs
	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Bad request")
	}

	args = getAssistantArgs{
		ScheduleID: scheduleID,
	}
	return args, nil
}

func (params updateParams) validate() (updateArgs, error) {

	var args updateArgs
	params = updateParams{
		ID:             html.EscapeString(helper.Trim(params.ID)),
		Name:           helper.Trim(params.Name),
		Description:    html.EscapeString(helper.Trim(params.Description)),
		UCU:            params.UCU,
		ScheduleID:     params.ScheduleID,
		Status:         params.Status,
		Semester:       params.Semester,
		Year:           params.Year,
		StartTime:      params.StartTime,
		EndTime:        params.EndTime,
		Class:          strings.ToUpper(helper.Trim(params.Class)),
		Day:            strings.ToLower(params.Day),
		PlaceID:        html.EscapeString(strings.ToUpper(helper.Trim(params.PlaceID))),
		IsUpdate:       params.IsUpdate,
		GradeParameter: params.GradeParameter,
	}

	// Course Validation
	if helper.IsEmpty(params.ID) {
		return args, fmt.Errorf("course_id cannot be empty")
	}
	if len(params.ID) > cs.MaximumID {
		return args, fmt.Errorf("course_id can have only maximum 45 characters length")
	}

	// Name validation
	name, err := helper.NormalizeName(params.Name)
	if err != nil {
		return args, err
	}

	// Description validation
	var description sql.NullString
	if !helper.IsEmpty(params.Description) {
		description = sql.NullString{Valid: true, String: params.Description}
	}

	// UCU validation
	if helper.IsEmpty(params.UCU) {
		return args, fmt.Errorf("UCU can't be empty")
	}
	ucu, err := strconv.ParseInt(params.UCU, 10, 8)
	if err != nil {
		return args, fmt.Errorf("UCU must be numeric")
	}
	if ucu < 1 || ucu > 5 {
		return args, fmt.Errorf("Invalid UCU")
	}

	// Semester validation
	if helper.IsEmpty(params.Semester) {
		return args, fmt.Errorf("Semester can't be empty")
	}
	semester, err := strconv.ParseInt(params.Semester, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Semester must be numeric")
	}
	if semester < 1 || semester > 7 {
		return args, fmt.Errorf("Invalid semester")
	}

	// Year validation
	if helper.IsEmpty(params.Year) {
		return args, fmt.Errorf("Year can't be empty")
	}
	year, err := strconv.ParseInt(params.Year, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Year must be numeric")
	}
	if year < 2017 || year > 2020 {
		return args, fmt.Errorf("Invalid year")
	}

	// Start time validation
	if helper.IsEmpty(params.StartTime) {
		return args, fmt.Errorf("Start time can't be empty")
	}
	startTime, err := strconv.ParseInt(params.StartTime, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Invalid start time")
	}
	if startTime < 0 || startTime >= 1440 {
		return args, fmt.Errorf("Invalid start time")
	}

	// End time validation
	if helper.IsEmpty(params.EndTime) {
		return args, fmt.Errorf("End time can't be empty")
	}
	endTime, err := strconv.ParseInt(params.EndTime, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Invalid end time")
	}
	if endTime < 0 || endTime >= 1440 {
		return args, fmt.Errorf("Invalid end time")
	}

	if startTime > endTime {
		return args, fmt.Errorf("Start time more than end time")
	}

	// Class validation
	if helper.IsEmpty(params.Class) {
		return args, fmt.Errorf("Class can't be empty")
	}
	if len(params.Class) != 1 || !helper.IsAlpha(params.Class) {
		return args, fmt.Errorf("Invalid class")
	}
	class := strings.ToUpper(params.Class)

	// Day validation
	if helper.IsEmpty(params.Day) {
		return args, fmt.Errorf("Day can't be empty")
	}
	day, err := helper.DayStringToInt(params.Day)
	if err != nil {
		return args, err
	}

	// Place ID validation
	if helper.IsEmpty(params.PlaceID) {
		return args, fmt.Errorf("Place ID can't be empty")
	}
	if len(params.PlaceID) > 30 {
		return args, fmt.Errorf("Invalid place id")
	}

	// IsUpdate Course validation
	isUpdate := false
	if params.IsUpdate == "true" {
		isUpdate = true
	}

	// validate is schedule active
	var status int8
	switch params.Status {
	case "active":
		status = cs.StatusScheduleActive
	case "inactive":
		status = cs.StatusScheduleInactive
	default:
		return args, fmt.Errorf("Error status")
	}
	if params.Status == "active" {
		status = cs.StatusScheduleActive
	}

	// validate schedule_id
	var scheduleID int64
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("invalid request")
	}
	scheduleID, err = strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("invalid request")
	}

	// validate grade parameter and convert to []gradeParameter struct
	var gps []gradeParameter
	if !helper.IsEmpty(params.GradeParameter) {
		var gp []gradeParameter
		// convert json into gradeParameter struct
		err = json.Unmarshal([]byte(params.GradeParameter), &gp)
		if err != nil {
			return args, fmt.Errorf("Invalid grade parameter")
		}

		var percentage float32
		gpType := []string{
			cs.GradeParameterAssignment,
			cs.GradeParameterAttendance,
			cs.GradeParameterFinal,
			cs.GradeParameterMid,
			cs.GradeParameterQuiz,
		}

		var gpChoosen []string
		for _, val := range gp {
			percentage += val.Percentage
			// check status change
			if val.StatusChange != cs.GradeParameterStatusChange && val.StatusChange != cs.GradeParameterStatusUnchange {
				return args, fmt.Errorf("Invalid grade parameter")
			}
			// validate type
			if !helper.IsStringInSlice(val.Type, gpType) {
				return args, fmt.Errorf("Invalid grade parameter")
			}
			if !helper.IsStringInSlice(val.Type, gpChoosen) {
				return args, fmt.Errorf("Invalid grade parameter should be unique")
			}
			gps = append(gps, val)
		}

		// validate total percentage should be 100
		if helper.Float32Round(percentage) != 100 {
			return args, fmt.Errorf("Total percentage must be 100%%")
		}
	}

	return updateArgs{
		ID:             params.ID,
		Name:           name,
		Description:    description,
		UCU:            int8(ucu),
		ScheduleID:     scheduleID,
		Status:         status,
		Semester:       int8(semester),
		Year:           int16(year),
		StartTime:      int16(startTime),
		EndTime:        int16(endTime),
		Class:          class,
		Day:            day,
		PlaceID:        params.PlaceID,
		IsUpdate:       isUpdate,
		GradeParameter: gps,
	}, nil
}

func (params deleteScheduleParams) validate() (deleteScheduleArgs, error) {

	var args deleteScheduleArgs
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("ID cannot be empty")
	}

	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("ID must be numeric")
	}

	return deleteScheduleArgs{
		ScheduleID: scheduleID,
	}, nil
}

func (params readDetailParams) validate() (readDetailArgs, error) {
	var args readDetailArgs
	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("schedule id must be numeric")
	}
	return readDetailArgs{
		ScheduleID: scheduleID,
	}, nil
}

func (params searchParams) validate() (searchArgs, error) {

	text := html.EscapeString(params.Text)
	text = helper.Trim(text)

	return searchArgs{
		Text: text,
	}, nil
}

func (params readScheduleParameterParams) validate() (readScheduleParameterArgs, error) {
	var args readScheduleParameterArgs
	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("schedule id must be numeric")
	}
	return readScheduleParameterArgs{
		ScheduleID: scheduleID,
	}, nil
}

func (params listStudentParams) validate() (listStudentArgs, error) {
	var args listStudentArgs
	if helper.IsEmpty(params.scheduleID) {
		return args, fmt.Errorf("Schedule id cannot be empty")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	return listStudentArgs{scheduleID: scheduleID}, nil
}
