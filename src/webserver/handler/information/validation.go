package information

import (
	"database/sql"
	"fmt"
	"html"
	"strconv"
	//"strings"

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
		Title:			params.Title,
		Description:	html.EscapeString(params.Description),
		CourseID:		params.CourseID,
		Tipe:			params.Tipe,
	}

	// Title validation
	title, err := helper.NormalizeName(params.Title)
	if err != nil {
		return args, err
	}

	// Description validation
	var description sql.NullString
	if !helper.IsEmpty(params.Description) {
		description = sql.NullString{Valid: true, String: params.Description}
	}

	// CourseID validation
	if helper.IsEmpty(params.CourseID) {
		return args, fmt.Errorf("Course can't be empty")
	}
	courseID, err := strconv.ParseInt(params.CourseID, 10, 8)
	if len(params.CourseID) < 1 || len(params.CourseID) > 12 {
		return args, fmt.Errorf("Invalid CourseID")
	}

	// Tipe validation
	if helper.IsEmpty(params.Tipe) {
		return args, fmt.Errorf("Type can't be empty")
	}
	tipe, err := strconv.ParseInt(params.Tipe, 10, 8)
	if len(params.Tipe) < 1 || len(params.Tipe) > 5 {
		return args, fmt.Errorf("Invalid tipe")
	}

	return createArgs{
		Title:			title,
		Description:	description,
		CourseID:		courseID,
		Tipe:			tipe,
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

/*func (params getAssistantParams) Validate() (getAssistantArgs, error) {

	var args getAssistantArgs
	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Bad request")
	}

	args = getAssistantArgs{
		ID: id,
	}
	return args, nil
}*/

func (params updateParams) Validate() (updateArgs, error) {

	var args updateArgs
	params = updateParams{
		ID:				params.ID,
		Title:			params.Title,
		Description:	html.EscapeString(params.Description),
		CourseID:		params.CourseID,	
		Type:			params.Type,
	}

	// ID validation
	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Bad Request")
	}

	// Title validation
	title, err := helper.NormalizeName(params.Title)
	if err != nil {
		return args, err
	}

	// Description validation
	var description sql.NullString
	if !helper.IsEmpty(params.Description) {
		description = sql.NullString{Valid: true, String: params.Description}
	}

	// course_id validation
	courseID, err := strconv.ParseInt(params.CourseID, 10, 8)
	if err != nil {
		return args, fmt.Errorf("Course ID must be numeric")
	}

	// Type validation
	tipe, err := helper.NormalizeName(params.Type)
	if err != nil {
		return args, err
	}

	return updateArgs{
		ID:				id,
		Title:			title,
		Description:	description,
		CourseID:		courseID,	
		Tipe:			tipe,
	}, nil
}
