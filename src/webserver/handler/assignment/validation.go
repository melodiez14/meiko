package assignment

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"strings"
	"time"

	asg "github.com/melodiez14/meiko/src/module/assignment"
	"github.com/melodiez14/meiko/src/util/helper"
)

func (params getParams) validate() (getArgs, error) {
	var args getArgs
	if helper.IsEmpty(params.scheduleID) {
		return args, nil
	}
	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}
	return getArgs{
		scheduleID: sql.NullInt64{Valid: true, Int64: scheduleID},
	}, nil
}

func (params getDetailParams) validate() (getDetailArgs, error) {
	var args getDetailArgs
	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}
	return getDetailArgs{
		id: id,
	}, nil
}

func (params submitParams) validate() (submitArgs, error) {

	var args submitArgs
	params = submitParams{
		id:          params.id,
		fileID:      params.fileID,
		description: html.EscapeString(params.description),
	}

	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	var desc sql.NullString
	if !helper.IsEmpty(params.description) {
		if len(params.description) > asg.MaxDesc {
			return args, fmt.Errorf("Description reach maximum of 1000 character")
		}
		desc = sql.NullString{
			Valid:  true,
			String: params.description,
		}
	}

	fileID := strings.Split(params.fileID, "~")
	for _, val := range fileID {
		if !helper.IsValidFileID(val) {
			return args, fmt.Errorf("Invalid fileID format")
		}
	}

	return submitArgs{
		id:          id,
		description: desc,
		fileID:      fileID,
	}, nil
}

func (params createParams) validate() (createArgs, error) {
	var args createArgs
	params = createParams{
		filesID:     params.filesID,
		gpID:        params.gpID,
		name:        html.EscapeString(helper.Trim(params.name)),
		description: html.EscapeString(helper.Trim(params.description)),
		status:      params.status,
		dueDate:     params.dueDate,
		fileType:    params.fileType,
		fileSize:    params.fileSize,
	}

	filesID := strings.Split(params.filesID, "~")
	for _, val := range filesID {
		if !helper.IsValidFileID(val) {
			return args, fmt.Errorf("Invalid fileID format")
		}
	}

	gpID, err := strconv.ParseInt(params.gpID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	// name
	if helper.IsEmpty(params.name) {
		return args, fmt.Errorf("Name can not be empty")
	}

	if !helper.IsAlphaNumericSpace(params.name) {
		return args, fmt.Errorf("Name can only contain of alphabet and numeric")
	}

	name := strings.Title(params.name)

	// description validation
	var desc sql.NullString
	if !helper.IsEmpty(params.description) {
		desc = sql.NullString{Valid: true, String: params.description}
	}

	// status
	var status int8
	switch params.status {
	case "required":
		status = asg.StatusUploadRequired
	case "notrequired":
		status = asg.StatusUploadNotRequired
	default:
		return args, fmt.Errorf("Invalid request")
	}

	// dueDate validation
	if helper.IsEmpty(params.dueDate) {
		return args, fmt.Errorf("due date can not be empty")
	}

	dueDateInt64, err := strconv.ParseInt(params.dueDate, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}
	dueDate := time.Unix(dueDateInt64, 0)

	size, err := strconv.ParseInt(params.fileSize, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	if size >= 100 {
		return args, fmt.Errorf("Max size too large. It should be below 100MB")
	}

	typ := strings.Split(params.fileType, "~")
	for _, val := range typ {
		if !helper.IsValidFileID(val) {
			return args, fmt.Errorf("Invalid fileID format")
		}
	}

	return createArgs{
		filesID:     filesID,
		gpID:        gpID,
		name:        name,
		description: desc,
		status:      status,
		dueDate:     dueDate,
		fileSize:    size,
		fileType:    typ,
	}, nil
}

// old

func (params updatePrams) validate() (updateArgs, error) {
	var args updateArgs
	params = updatePrams{
		ID:                params.ID,
		FilesID:           params.FilesID,
		GradeParametersID: params.GradeParametersID,
		Name:              html.EscapeString(helper.Trim(params.Name)),
		Description:       html.EscapeString(helper.Trim(params.Description)),
		Status:            params.Status,
		DueDate:           params.DueDate,
	}
	// ID
	if helper.IsEmpty(params.ID) {
		return args, fmt.Errorf("ID can not be empty")
	}
	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("id error parsing")
	}
	// GradeParameter validation
	if helper.IsEmpty(params.GradeParametersID) {
		return args, fmt.Errorf("grade_parameters can not be empty")
	}
	GradeParametersID, err := strconv.ParseInt(params.GradeParametersID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("grade_parameters error parsing")
	}
	// Name
	if helper.IsEmpty(params.Name) {
		return args, fmt.Errorf("Name can not be empty")
	}

	// Status validation
	if helper.IsEmpty(params.Status) {
		return args, fmt.Errorf("status can not be empty")
	}
	// Description validation
	var description sql.NullString
	if !helper.IsEmpty(params.Description) {
		description = sql.NullString{Valid: true, String: params.Description}
	}
	//DueDate validation
	if helper.IsEmpty(params.DueDate) {
		return args, fmt.Errorf("due_date can not be empty")
	}
	return updateArgs{
		ID:                id,
		FilesID:           params.FilesID,
		GradeParametersID: GradeParametersID,
		Name:              params.Name,
		Description:       description,
		Status:            params.Status,
		DueDate:           params.DueDate,
	}, nil

}
func (params readParams) validate() (readArgs, error) {
	var args readArgs
	// if helper.IsEmpty(params.Page) || helper.IsEmpty(params.Total) {
	// 	return args, fmt.Errorf("page or total is empty")
	// }

	// page, err := strconv.ParseInt(params.Page, 10, 64)
	// if err != nil {
	// 	return args, fmt.Errorf("page must be numeric")
	// }

	// total, err := strconv.ParseInt(params.Total, 10, 64)
	// if err != nil {
	// 	return args, fmt.Errorf("total must be numeric")
	// }

	// // should be positive number
	// if page < 0 || total < 0 {
	// 	return args, fmt.Errorf("page or total must be positive number")
	// }

	// args = readArgs{
	// 	Page:  uint16(page),
	// 	Total: uint16(total),
	// }
	return args, nil

}

func (params detailParams) validate() (detailArgs, error) {
	var args detailArgs
	identityCode, err := strconv.ParseInt(params.IdentityCode, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error validation: ID should be numeric")
	}

	args = detailArgs{
		IdentityCode: identityCode,
	}
	return args, nil
}
func (params uploadAssignmentParams) validate() (uploadAssignmentArgs, error) {

	var args uploadAssignmentArgs
	params = uploadAssignmentParams{
		FileID:       params.FileID,
		AssignmentID: params.AssignmentID,
		UserID:       params.UserID,
		Description:  html.EscapeString(helper.Trim(params.Description)),
	}
	// AssigmentID validation
	if helper.IsEmpty(params.AssignmentID) {
		return args, fmt.Errorf("Assignment ID can not be empty")
	}
	assignmentID, err := strconv.ParseInt(params.AssignmentID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error convert AssigmentID")
	}
	// Description validation
	var description sql.NullString
	if !helper.IsEmpty(params.Description) {
		description = sql.NullString{Valid: true, String: params.Description}
	}
	var filesID []string
	// FilesID validation
	if !helper.IsEmpty(params.FileID) {
		filesID = strings.Split(params.FileID, "~")
		for _, value := range filesID {
			if !helper.IsValidFileID(value) {
				return args, fmt.Errorf("Wrong Files ID")
			}
		}
	}
	return uploadAssignmentArgs{
		FileID:       filesID,
		AssignmentID: assignmentID,
		UserID:       params.UserID,
		Description:  description,
	}, nil

}

func (params readUploadedAssignmentParams) validate() (readUploadedAssignmentArgs, error) {

	var args readUploadedAssignmentArgs
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
	// User ID validation
	if helper.IsEmpty(params.UserID) {
		return args, fmt.Errorf("User ID cannot be empty")
	}
	// Schedule ID validation
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("Schedule ID cannot be empty")
	}
	userID, err := strconv.ParseInt(params.UserID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert user ID to int64")
	}
	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert schedule ID to int64")
	}
	// Assignment ID validation
	if helper.IsEmpty(params.AssignmentID) {
		return args, fmt.Errorf("Assignment ID cannto be empty")
	}
	assignmentID, err := strconv.ParseInt(params.AssignmentID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert assignment ID to int64")
	}
	return readUploadedAssignmentArgs{
		UserID:       userID,
		ScheduleID:   scheduleID,
		AssignmentID: assignmentID,
		Total:        total,
		Page:         page,
	}, nil

}
func (params readUploadedDetailParams) validate() (readUploadedDetailArgs, error) {
	var args readUploadedDetailArgs
	// User ID validation
	if helper.IsEmpty(params.UserID) {
		return args, fmt.Errorf("User ID cannot be empty")
	}
	// Schedule ID validation
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("Schedule ID cannot be empty")
	}
	userID, err := strconv.ParseInt(params.UserID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert user ID to int64")
	}
	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert schedule ID to int64")
	}
	// Assignment ID validation
	if helper.IsEmpty(params.AssignmentID) {
		return args, fmt.Errorf("Assignment ID cannto be empty")
	}
	assignmentID, err := strconv.ParseInt(params.AssignmentID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert assignment ID to int64")
	}
	return readUploadedDetailArgs{
		UserID:       userID,
		ScheduleID:   scheduleID,
		AssignmentID: assignmentID,
	}, nil

}
func (params deleteParams) validate() (deleteArgs, error) {
	var args deleteArgs
	params = deleteParams{
		ID: params.ID,
	}
	if helper.IsEmpty(params.ID) {
		return args, fmt.Errorf("Assignment ID can not be empty")
	}
	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error convert to int64")
	}
	return deleteArgs{
		ID: id,
	}, nil
}
func (params listAssignmentsParams) validate() (listAssignmentsArgs, error) {
	var args listAssignmentsArgs
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("Schedule ID can not be emrpty")
	}
	id, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, err
	}
	return listAssignmentsArgs{
		ScheduleID: id,
	}, nil
}

func (params updateScoreParams) validate() (updateScoreArgs, error) {
	var args updateScoreArgs
	//Schedule ID validation
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("Schedule ID can not be empty")
	}
	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, err
	}
	// Assignment ID validation
	if helper.IsEmpty(params.AssignmentID) {
		return args, fmt.Errorf("Assignment ID can not be empty")
	}
	assignmentID, err := strconv.ParseInt(params.AssignmentID, 10, 64)
	if err != nil {
		return args, err
	}
	// Users ID validation
	if helper.IsEmpty(params.UserID) {
		return args, fmt.Errorf("User ID can not be empty")
	}
	userID, err := strconv.ParseInt(params.UserID, 10, 64)
	if err != nil {
		return args, err
	}
	// Score Validation
	if helper.IsEmpty(params.Score) {
		return args, fmt.Errorf("Score can not be empty")
	}
	score, err := strconv.ParseFloat(params.Score, 64)
	if err != nil {
		return args, err
	}
	return updateScoreArgs{
		ScheduleID:   scheduleID,
		AssignmentID: assignmentID,
		UserID:       userID,
		Score:        float32(score),
	}, nil
}

func (params detailAssignmentParams) validate() (detailAssignmentArgs, error) {
	var args detailAssignmentArgs
	//Schedule ID validation
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("Schedule ID can not be empty")
	}
	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, err
	}
	// Assignment ID validation
	if helper.IsEmpty(params.AssignmentID) {
		return args, fmt.Errorf("Assignment ID can not be empty")
	}
	assignmentID, err := strconv.ParseInt(params.AssignmentID, 10, 64)
	if err != nil {
		return args, err
	}
	return detailAssignmentArgs{
		ScheduleID:   scheduleID,
		AssignmentID: assignmentID,
	}, nil
}

func (params createScoreParams) validate() (createScoreArgs, error) {
	args := createScoreArgs{}
	params = createScoreParams{
		ScheduleID:   params.ScheduleID,
		AssignmentID: params.AssignmentID,
		Users:        params.Users,
	}

	//Schedule ID validation
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("Schedule ID can not be empty")
	}

	scheduleID, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	// Assignment ID validation
	if helper.IsEmpty(params.AssignmentID) {
		return args, fmt.Errorf("Assignment ID can not be empty")
	}
	assignmentID, err := strconv.ParseInt(params.AssignmentID, 10, 64)
	if err != nil {
		return args, err
	}
	// Users validation
	var users []int64
	var score []float32
	if !helper.IsEmpty(params.Users) {
		var std []student
		err := json.Unmarshal([]byte(params.Users), &std)

		if err != nil {
			return args, err
		}
		for _, val := range std {
			users = append(users, val.IdentityCode)
			score = append(score, val.Score)
		}
	}

	return createScoreArgs{
		ScheduleID:   scheduleID,
		AssignmentID: assignmentID,
		IdentityCode: users,
		Score:        score,
	}, nil

}

func (params scoreParams) validate() (scoreArgs, error) {
	var args scoreArgs
	params = scoreParams{
		ScheduleID: params.ScheduleID,
	}
	if helper.IsEmpty(params.ScheduleID) {
		return args, fmt.Errorf("Schedule ID can not be empty")
	}
	id, err := strconv.ParseInt(params.ScheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error convert to int64")
	}
	return scoreArgs{
		ScheduleID: id,
	}, nil
}
