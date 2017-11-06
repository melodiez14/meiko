package assignment

import (
	"fmt"
	"html"
	"strconv"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params createParams) validate() (createArgs, error) {
	var args createArgs
	params = createParams{
		ID:             params.ID,
		GradeParameter: params.GradeParameter,
		Status:         params.Status,
		Description:    html.EscapeString(helper.Trim(params.Description)),
		DueDate:        params.DueDate,
	}
	// GradeParameter validation
	if helper.IsEmpty(params.GradeParameter) {
		return args, fmt.Errorf("grade_parameters can not be empty")
	}
	id, err := strconv.ParseInt(params.GradeParameter, 10, 64)
	if err != nil {
		return args, fmt.Errorf("grade_parameters error parsing")
	}

	// Status validation
	if helper.IsEmpty(params.Status) {
		return args, fmt.Errorf("status can not be empty")
	}
	// Description validation
	if helper.IsEmpty(params.Description) {
		return args, fmt.Errorf("descriptin can not be empty")
	}
	//DueDate validation
	if helper.IsEmpty(params.Description) {
		return args, fmt.Errorf("due_date can not be empty")
	}
	return createArgs{
		ID:             params.ID,
		GradeParameter: id,
		Status:         params.Status,
		Description:    params.Description,
		DueDate:        params.DueDate,
	}, nil

}
