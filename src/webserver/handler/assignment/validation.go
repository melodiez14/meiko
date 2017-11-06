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
	if helper.IsEmpty(params.Description) {
		return args, fmt.Errorf("descriptin can not be empty")
	}
	//DueDate validation
	if helper.IsEmpty(params.Description) {
		return args, fmt.Errorf("due_date can not be empty")
	}
	return createArgs{
		FilesID:           params.FilesID,
		GradeParametersID: GradeParametersID,
		Name:              params.Name,
		Description:       params.Description,
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
