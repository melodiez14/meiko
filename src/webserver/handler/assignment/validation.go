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
	fl "github.com/melodiez14/meiko/src/module/file"
	"github.com/melodiez14/meiko/src/util/helper"
)

func (params getParams) validate() (getArgs, error) {
	var args getArgs
	var scheduleID sql.NullInt64
	if !helper.IsEmpty(params.scheduleID) {
		schID, err := strconv.ParseInt(params.scheduleID, 10, 64)
		if err != nil {
			return args, err
		}
		scheduleID.Int64 = schID
		scheduleID.Valid = true
	}

	var filter sql.NullString
	if !helper.IsEmpty(params.filter) {
		switch params.filter {
		case "submitted", "unsubmitted", "overdue", "done":
			filter.String = params.filter
			filter.Valid = true
		default:
			return args, fmt.Errorf("Invalid request")
		}
	}
	return getArgs{
		scheduleID: scheduleID,
		filter:     filter,
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

	isEmptyDesc := helper.IsEmpty(params.description)
	isEmptyFile := helper.IsEmpty(params.fileID)
	if isEmptyDesc && isEmptyFile {
		return args, fmt.Errorf("File or description must be filled")
	}

	var desc sql.NullString
	if !isEmptyDesc {
		if len(params.description) > asg.MaxDesc {
			return args, fmt.Errorf("Description reach maximum of 1000 character")
		}
		desc = sql.NullString{
			Valid:  true,
			String: params.description,
		}
	}

	var fileID []string
	if !isEmptyFile {
		fileID = strings.Split(params.fileID, "~")
		for _, val := range fileID {
			if !helper.IsValidFileID(val) {
				return args, fmt.Errorf("Invalid request")
			}
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
		filesID:          params.filesID,
		gpID:             params.gpID,
		name:             html.EscapeString(helper.Trim(params.name)),
		description:      html.EscapeString(helper.Trim(params.description)),
		status:           params.status,
		dueDate:          params.dueDate,
		allowedTypesFile: params.allowedTypesFile,
		maxSizeFile:      params.maxSizeFile,
		maxFile:          params.maxFile,
	}
	var filesID []string
	if len(params.filesID) > 0 {
		filesID = strings.Split(params.filesID, "~")
		for _, val := range filesID {
			if !helper.IsValidFileID(val) {
				return args, fmt.Errorf("Invalid fileID format")
			}
		}
	}

	if helper.IsEmpty(params.gpID) {
		return args, fmt.Errorf("Grade parameters id can not be empty")
	}
	gpID, err := strconv.ParseInt(params.gpID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert grade parameter id to int64")
	}

	if helper.IsEmpty(params.name) {
		return args, fmt.Errorf("Name can not be empty")
	}
	if !helper.IsAlphaNumericSpace(params.name) {
		return args, fmt.Errorf("Name can only contain of alphabet and numeric")
	}
	name := strings.Title(params.name)

	var desc sql.NullString
	if !helper.IsEmpty(params.description) {
		desc = sql.NullString{Valid: true, String: params.description}
	}
	if helper.IsEmpty(params.status) {
		return args, fmt.Errorf("Status must upload or not can not empty")
	}
	status, err := strconv.ParseInt(params.status, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Status can convert to int64")
	}
	if status < 0 && status > 1 {
		return args, fmt.Errorf("Wrong status")
	}
	var allowedTypes []string
	var maxFile int64
	var size int64
	if status == 1 {
		if helper.IsEmpty(params.maxSizeFile) {
			return args, fmt.Errorf("Max size file can not be empty")
		}
		size, err := strconv.ParseInt(params.maxSizeFile, 10, 64)
		if err != nil {
			return args, fmt.Errorf("Can not parse from size string to size int64")
		}
		if size >= asg.MaxSizeFile {
			return args, fmt.Errorf("Max size too large. It should be below 100MB")
		}

		if helper.IsEmpty(params.allowedTypesFile) {
			return args, fmt.Errorf("types can not be empty")
		}
		allowedTypes = strings.Split(params.allowedTypesFile, "~")
		availableTypes := strings.Split(fl.AvailableTypesFile, "~")
		for _, val := range allowedTypes {
			count := 0
			for _, typeFile := range availableTypes {
				if typeFile == val {
					count++
					continue
				}
			}
			if count == 0 {
				return args, fmt.Errorf(fmt.Sprintf("%s Denied type", val))
			}
		}
		if helper.IsEmpty(params.maxFile) {
			return args, fmt.Errorf("Max file can not be empty")
		}
		maxFile, err = strconv.ParseInt(params.maxSizeFile, 10, 64)
		if err != nil {
			return args, fmt.Errorf("Can not convert max file to int64")
		}
		if maxFile > asg.MaxFile {
			return args, fmt.Errorf("Max file can not more than 5")
		}
	}
	if helper.IsEmpty(params.dueDate) {
		return args, fmt.Errorf("Due date can not be empty")
	}
	layout := `2006-01-02 15:04:05`
	dueDate, err := time.Parse(layout, params.dueDate)
	if err != nil {
		return args, fmt.Errorf(err.Error())
	}

	return createArgs{
		filesID:          filesID,
		gpID:             gpID,
		name:             name,
		description:      desc,
		status:           int8(status),
		dueDate:          dueDate,
		maxSizeFile:      size,
		allowedTypesFile: allowedTypes,
		maxFile:          maxFile,
	}, nil
}

func (params updateParams) validate() (updateArgs, error) {
	var args updateArgs
	params = updateParams{
		ID:               params.ID,
		filesID:          params.filesID,
		gpID:             params.gpID,
		name:             html.EscapeString(helper.Trim(params.name)),
		description:      html.EscapeString(helper.Trim(params.description)),
		status:           params.status,
		dueDate:          params.dueDate,
		allowedTypesFile: params.allowedTypesFile,
		maxSizeFile:      params.maxSizeFile,
		maxFile:          params.maxFile,
	}
	if helper.IsEmpty(params.ID) {
		return args, fmt.Errorf("ID can not be empty")
	}
	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert ID to int64")
	}
	var filesID []string
	if len(params.filesID) > 0 {
		filesID = strings.Split(params.filesID, "~")
		for _, val := range filesID {
			if !helper.IsValidFileID(val) {
				return args, fmt.Errorf("Invalid fileID format")
			}
		}
	}

	if helper.IsEmpty(params.gpID) {
		return args, fmt.Errorf("Grade parameters id can not be empty")
	}
	gpID, err := strconv.ParseInt(params.gpID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Can not convert grade parameter id to int64")
	}
	if !asg.IsAssignmentExistByGradeParameterID(id, gpID) {
		return args, fmt.Errorf("Wrong ID assignemnt")
	}

	if helper.IsEmpty(params.name) {
		return args, fmt.Errorf("Name can not be empty")
	}
	if !helper.IsAlphaNumericSpace(params.name) {
		return args, fmt.Errorf("Name can only contain of alphabet and numeric")
	}
	name := strings.Title(params.name)

	var desc sql.NullString
	if !helper.IsEmpty(params.description) {
		desc = sql.NullString{Valid: true, String: params.description}
	}
	if helper.IsEmpty(params.status) {
		return args, fmt.Errorf("Status must upload or not can not empty")
	}
	status, err := strconv.ParseInt(params.status, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Status can convert to int64")
	}
	if status < 0 && status > 1 {
		return args, fmt.Errorf("Wrong status")
	}
	var allowedTypes []string
	var maxFile int64
	var size int64
	if status == 1 {
		if helper.IsEmpty(params.maxSizeFile) {
			return args, fmt.Errorf("Max size file can not be empty")
		}
		size, err = strconv.ParseInt(params.maxSizeFile, 10, 64)
		if err != nil {
			return args, fmt.Errorf("Can not parse from size string to size int64")
		}
		if size >= asg.MaxSizeFile {
			return args, fmt.Errorf("Max size too large. It should be below 100MB")
		}

		if helper.IsEmpty(params.allowedTypesFile) {
			return args, fmt.Errorf("types can not be empty")
		}
		allowedTypes = strings.Split(params.allowedTypesFile, "~")
		availableTypes := strings.Split(fl.AvailableTypesFile, "~")
		for _, val := range allowedTypes {
			count := 0
			for _, typeFile := range availableTypes {
				if typeFile == val {
					count++
					continue
				}
			}
			if count == 0 {
				return args, fmt.Errorf(fmt.Sprintf("%s Denied type", val))
			}
		}
		if helper.IsEmpty(params.maxFile) {
			return args, fmt.Errorf("Max file can not be empty")
		}
		maxFile, err = strconv.ParseInt(params.maxFile, 10, 64)
		if err != nil {
			return args, fmt.Errorf("Can not convert max file to int64")
		}
		if maxFile > asg.MaxFile {
			return args, fmt.Errorf("Max file can not more than 5")
		}
	}
	if helper.IsEmpty(params.dueDate) {
		return args, fmt.Errorf("Due date can not be empty")
	}
	layout := `2006-01-02 15:04:05`
	dueDate, err := time.Parse(layout, params.dueDate)
	if err != nil {
		return args, fmt.Errorf(err.Error())
	}

	return updateArgs{
		ID:               id,
		filesID:          filesID,
		gpID:             gpID,
		name:             name,
		description:      desc,
		status:           int8(status),
		dueDate:          dueDate,
		maxSizeFile:      size,
		allowedTypesFile: allowedTypes,
		maxFile:          maxFile,
	}, nil
}

func (params readParams) validate() (readArgs, error) {

	var args readArgs
	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("scheduleID must be numeric")
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
	if page < 0 || total < 0 {
		return args, fmt.Errorf("page or total must be positive number")
	}

	args = readArgs{
		page:       int(page),
		total:      int(total),
		scheduleID: scheduleID,
	}
	return args, nil

}
func (params availableParams) validate() (availableArgs, error) {

	var args availableArgs
	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, fmt.Errorf("scheduleID must be numeric")
	}

	args = availableArgs{
		id: id,
	}
	return args, nil

}

func (params deleteParams) validate() (deleteArgs, error) {
	var args deleteArgs

	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	return deleteArgs{
		id: id,
	}, nil
}

func (params detailParams) validate() (detailArgs, error) {
	var args detailArgs
	ID, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error validation: ID should be numeric")
	}
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
	if page < 0 || total < 0 {
		return args, fmt.Errorf("page or total must be positive number")
	}

	args = detailArgs{
		ID:    ID,
		page:  int(page),
		total: int(total),
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
