package attendance

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"time"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params listStudentParams) validate() (listStudentArgs, error) {
	var args listStudentArgs

	if helper.IsEmpty(params.meetingNumber) {
		return args, fmt.Errorf("Meeting number cannot be empty")
	}

	if helper.IsEmpty(params.scheduleID) {
		return args, fmt.Errorf("ScheduleID cannot be empty")
	}

	meetingNumber, err := strconv.ParseUint(params.meetingNumber, 10, 8)
	if err != nil {
		return args, err
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	return listStudentArgs{
		meetingNumber: uint8(meetingNumber),
		scheduleID:    scheduleID,
	}, nil
}

func (params readMeetingParams) validate() (readMeetingArgs, error) {

	var args readMeetingArgs
	if helper.IsEmpty(params.scheduleID) || helper.IsEmpty(params.page) || helper.IsEmpty(params.total) {
		return args, fmt.Errorf("Invalid request")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	page, err := strconv.ParseUint(params.page, 10, 8)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	total, err := strconv.ParseUint(params.total, 10, 8)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	// should be positive number
	if page < 1 || total < 1 {
		return args, fmt.Errorf("Invalid request")
	}

	return readMeetingArgs{
		scheduleID: scheduleID,
		page:       uint8(page),
		total:      uint8(total),
	}, nil
}

func (params createMeetingParams) validate() (createMeetingArgs, error) {

	var args createMeetingArgs
	params = createMeetingParams{
		subject:       html.EscapeString(params.subject),
		meetingNumber: params.meetingNumber,
		scheduleID:    params.scheduleID,
		description:   html.EscapeString(params.description),
		date:          params.date,
		users:         params.users,
	}

	if helper.IsEmpty(params.subject) {
		return args, fmt.Errorf("Subject cannot be empty")
	}

	subject := helper.Trim(params.subject)
	if len(subject) > 255 {
		return args, fmt.Errorf("Subject is too long")
	}

	if helper.IsEmpty(params.meetingNumber) {
		return args, fmt.Errorf("Meeting number cannot be empty")
	}

	meetingNumber, err := strconv.ParseUint(params.meetingNumber, 10, 8)
	if err != nil {
		return args, err
	}

	if helper.IsEmpty(params.scheduleID) {
		return args, fmt.Errorf("Invalid Request")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	var description sql.NullString
	if !helper.IsEmpty(params.description) {
		if len(params.description) > 1000 {
			return args, fmt.Errorf("Description maximum consist of 1000 character")
		}
		description = sql.NullString{Valid: true, String: params.description}
	}

	if helper.IsEmpty(params.date) {
		return args, fmt.Errorf("Date cannot be empty")
	}

	timeInt, err := strconv.ParseInt(params.date, 10, 64)
	if err != nil {
		return args, err
	}

	t := time.Unix(0, timeInt)

	var users []int64
	if !helper.IsEmpty(params.users) {
		var std []student
		err = json.Unmarshal([]byte(params.users), &std)
		if err != nil {
			return args, err
		}

		for _, val := range std {
			if val.Status == "present" {
				users = append(users, val.IdentityCode)
			}
		}
	}

	return createMeetingArgs{
		subject:           subject,
		meetingNumber:     uint8(meetingNumber),
		description:       description,
		scheduleID:        scheduleID,
		date:              t,
		userIdentityCodes: users,
	}, nil
}

func (params updateMeetingParams) validate() (updateMeetingArgs, error) {

	var args updateMeetingArgs
	params = updateMeetingParams{
		id:            params.id,
		subject:       html.EscapeString(params.subject),
		meetingNumber: params.meetingNumber,
		scheduleID:    params.scheduleID,
		description:   html.EscapeString(params.description),
		date:          params.date,
		isForceUpdate: params.isForceUpdate,
		users:         params.users,
	}

	if helper.IsEmpty(params.id) {
		return args, fmt.Errorf("Invalid Request")
	}

	id, err := strconv.ParseUint(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	if helper.IsEmpty(params.meetingNumber) {
		return args, fmt.Errorf("Meeting number cannot be empty")
	}

	if helper.IsEmpty(params.subject) {
		return args, fmt.Errorf("Subject cannot be empty")
	}

	subject := helper.Trim(params.subject)
	if len(subject) > 255 {
		return args, fmt.Errorf("Subject is too long")
	}

	if helper.IsEmpty(params.meetingNumber) {
		return args, fmt.Errorf("Meeting number cannot be empty")
	}

	meetingNumber, err := strconv.ParseUint(params.meetingNumber, 10, 8)
	if err != nil {
		return args, err
	}

	if helper.IsEmpty(params.scheduleID) {
		return args, fmt.Errorf("Invalid Request")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	var description sql.NullString
	if !helper.IsEmpty(params.description) {
		if len(params.description) > 1000 {
			return args, fmt.Errorf("Description maximum consist of 1000 character")
		}
		description = sql.NullString{Valid: true, String: params.description}
	}

	if helper.IsEmpty(params.date) {
		return args, fmt.Errorf("Date cannot be empty")
	}

	timeInt, err := strconv.ParseInt(params.date, 10, 64)
	if err != nil {
		return args, err
	}

	t := time.Unix(0, timeInt)

	isForceUpdate := false
	if params.isForceUpdate == "true" {
		isForceUpdate = true
	}

	var users []int64
	if !helper.IsEmpty(params.users) {
		var std []student

		err = json.Unmarshal([]byte(params.users), &std)
		if err != nil {
			return args, fmt.Errorf("Invalid student")
		}

		for _, val := range std {
			if val.Status == "present" {
				users = append(users, val.IdentityCode)
			}
		}
	}

	return updateMeetingArgs{
		id:                id,
		subject:           subject,
		meetingNumber:     uint8(meetingNumber),
		description:       description,
		scheduleID:        scheduleID,
		date:              t,
		isForceUpdate:     isForceUpdate,
		userIdentityCodes: users,
	}, nil
}

func (params deleteMeetingParams) validate() (deleteMeetingArgs, error) {
	var args deleteMeetingArgs
	if helper.IsEmpty(params.id) {
		return args, fmt.Errorf("Invalid Request")
	}

	id, err := strconv.ParseUint(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	isForceDelete := false
	if params.isForceDelete == "true" {
		isForceDelete = true
	}

	return deleteMeetingArgs{
		id:            id,
		isForceDelete: isForceDelete,
	}, nil
}

func (params readMeetingDetailParams) validate() (readMeetingDetailArgs, error) {
	var args readMeetingDetailArgs
	if helper.IsEmpty(params.meetingID) {
		return args, fmt.Errorf("Meeting ID cannot be nil")
	}

	meetingID, err := strconv.ParseUint(params.meetingID, 10, 64)
	if err != nil {
		return args, err
	}

	return readMeetingDetailArgs{
		meetingID: meetingID,
	}, nil
}
