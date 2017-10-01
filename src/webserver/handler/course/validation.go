package course

import (
	"database/sql"
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params readParams) Validate() (readArgs, error) {

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

func (params createParams) Validate() (createArgs, error) {

	var args createArgs
	params = createParams{
		Name:        params.Name,
		Description: html.EscapeString(params.Description),
		UCU:         params.UCU,
		Semester:    params.Semester,
		StartTime:   params.StartTime,
		EndTime:     params.EndTime,
		Class:       strings.ToUpper(params.Class),
		Day:         strings.ToLower(params.Day),
		PlaceID:     html.EscapeString(params.PlaceID),
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

	// Start time validation
	if helper.IsEmpty(params.StartTime) {
		return args, fmt.Errorf("Start time can't be empty")
	}
	startTime, err := strconv.ParseInt(params.StartTime, 10, 16)
	if err != nil {
		return args, fmt.Errorf("Invalid end time")
	}
	if startTime < 0 || startTime >= 1440 {
		return args, fmt.Errorf("Invalid start time")
	}

	// End time validation
	if helper.IsEmpty(params.EndTime) {
		return args, fmt.Errorf("End time can't be empty")
	}
	endTime, err := strconv.ParseInt(params.EndTime, 10, 16)
	if helper.IsEmpty(params.EndTime) {
		return args, fmt.Errorf("Invalid end time")
	}
	if endTime < 0 || endTime >= 1440 {
		return args, fmt.Errorf("Invalid end time")
	}

	// Class validation
	if helper.IsEmpty(params.Class) {
		return args, fmt.Errorf("Class can't be empty")
	}
	if len(params.Class) != 1 || !helper.IsAlphaSpace(params.Class) {
		return args, fmt.Errorf("Invalid class")
	}

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
		return args, fmt.Errorf("Day can't be empty")
	}
	if len(params.PlaceID) > 30 {
		return args, fmt.Errorf("Invalid place id")
	}

	return createArgs{
		Name:        name,
		Description: description,
		UCU:         int8(ucu),
		Semester:    int8(semester),
		StartTime:   int16(startTime),
		EndTime:     int16(endTime),
		Class:       params.Class,
		Day:         day,
		PlaceID:     params.PlaceID,
	}, nil
}

func (params getParams) Validate() (getArgs, error) {

	var args getArgs
	if params.Payload != "last" && params.Payload != "current" && params.Payload != "all" {
		return args, fmt.Errorf("Invalid payload")
	}

	args = getArgs{
		Payload: params.Payload,
	}

	return args, nil
}

func (params getAssistantParams) Validate() (getAssistantArgs, error) {

	var args getAssistantArgs
	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Bad request")
	}

	args = getAssistantArgs{
		ID: id,
	}
	return args, nil
}
