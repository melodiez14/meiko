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
	if helper.IsEmpty(params.page) || helper.IsEmpty(params.total) {
		return args, fmt.Errorf("page or total is empty")
	}

	page, err := strconv.ParseInt(params.page, 10, 64)
	if err != nil {
		return args, fmt.Errorf("page must be numeric")
	}

	total, err := strconv.ParseInt(params.total, 10, 64)
	if err != nil {
		return args, fmt.Errorf("total must be numeric")
	}

	// should be positive number
	if page < 1 || total < 1 {
		return args, fmt.Errorf("page or total must be positive number")
	}

	args = readArgs{
		page:  int(page),
		total: int(total),
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
	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	args = getAssistantArgs{
		payload:    params.payload,
		scheduleID: scheduleID,
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
			// validate type
			if !helper.IsStringInSlice(val.Type, gpType) {
				return args, fmt.Errorf("Invalid grade parameter")
			}
			if helper.IsStringInSlice(val.Type, gpChoosen) {
				return args, fmt.Errorf("Invalid grade parameter should be unique")
			}
			gpChoosen = append(gpChoosen, val.Type)
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

func (params addAssistantParams) validate() (addAssistantArgs, error) {

	var args addAssistantArgs
	idInt64 := []int64{}

	if !helper.IsEmpty(params.assistentIdentityCodes) {
		idString := strings.Split(params.assistentIdentityCodes, "~")
		for _, val := range idString {
			id, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return args, err
			}
			idInt64 = append(idInt64, id)
		}
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	return addAssistantArgs{
		assistentIdentityCodes: idInt64,
		scheduleID:             scheduleID,
	}, nil
}

func (params getTodayParams) validate() (getTodayArgs, error) {
	var args getTodayArgs

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	return getTodayArgs{
		scheduleID: scheduleID,
	}, nil
}

func (params getDetailParams) validate() (getDetailArgs, error) {
	var args getDetailArgs

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	return getDetailArgs{
		scheduleID: scheduleID,
	}, nil
}

func (params enrollRequestParams) validate() (enrollRequestArgs, error) {
	var args enrollRequestArgs

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	return enrollRequestArgs{
		payload:    params.payload,
		scheduleID: scheduleID,
	}, nil
}

func (params addInvolvedParams) validate() (addInvolvedArgs, error) {
	var args addInvolvedArgs

	identityCode, err := strconv.ParseInt(params.identityCode, 10, 64)
	if err != nil {
		return args, err
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	role := 0
	switch params.role {
	case "assistant":
		role = cs.PStatusAssistant
	case "student":
		role = cs.PStatusStudent
	default:
		return args, fmt.Errorf("Invalid Role")
	}

	// student can be add and activate if there is a request
	if role == cs.PStatusStudent && params.status != "add" && params.status != "active" {
		return args, fmt.Errorf("Student status must be add or active")
		// admin can only add
	} else if role == cs.PStatusAssistant && params.status != "add" {
		return args, fmt.Errorf("Assistant status must be add")
	}

	return addInvolvedArgs{
		identityCode: identityCode,
		scheduleID:   scheduleID,
		role:         role,
		status:       params.status,
	}, nil
}

func (params getInvolvedParams) validate() (getInvolvedArgs, error) {
	var args getInvolvedArgs

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	return getInvolvedArgs{
		role:       params.role,
		scheduleID: scheduleID,
	}, nil
}

func (params searchUninvolvedParams) validate() (searchUninvolvedArgs, error) {
	var args searchUninvolvedArgs

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	text := html.EscapeString(params.text)

	return searchUninvolvedArgs{
		scheduleID: scheduleID,
		text:       text,
	}, nil
}
