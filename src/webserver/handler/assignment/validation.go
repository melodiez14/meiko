package assignment

import (
	"database/sql"
	"fmt"
	"html"
	"strconv"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params createParams) validate() (createArgs, error) {
	var args createArgs
	params = createParams{
		FilesID:           params.FilesID,
		GradeParametersID: params.GradeParametersID,
		Name:              html.EscapeString(helper.Trim(params.Name)),
		Description:       html.EscapeString(helper.Trim(params.Description)),
		Status:            params.Status,
		DueDate:           params.DueDate,
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
	return createArgs{
		FilesID:           params.FilesID,
		GradeParametersID: GradeParametersID,
		Name:              params.Name,
		Description:       description,
		Status:            params.Status,
		DueDate:           params.DueDate,
	}, nil

}
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
		Subject:      html.EscapeString(helper.Trim(params.Subject)),
	}
	// AssigmentID validation
	if helper.IsEmpty(params.AssignmentID) {
		return args, fmt.Errorf("Assignment ID can not be empty")
	}
	assignmentID, err := strconv.ParseInt(params.AssignmentID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error convert AssigmentID")
	}
	// Subject validation
	var subject sql.NullString
	if !helper.IsEmpty(params.Subject) {
		subject = sql.NullString{Valid: true, String: params.Subject}
	}
	// Description validation
	var description sql.NullString
	if !helper.IsEmpty(params.Description) {
		description = sql.NullString{Valid: true, String: params.Description}
	}
	return uploadAssignmentArgs{
		FileID:       params.FileID,
		AssignmentID: assignmentID,
		UserID:       params.UserID,
		Subject:      subject,
		Description:  description,
	}, nil

}
