package tutorial

import (
	"fmt"
	"strconv"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params readParams) validate() (readArgs, error) {
	var args readArgs
	if helper.IsEmpty(params.scheduleID) || helper.IsEmpty(params.page) || helper.IsEmpty(params.total) {
		return args, fmt.Errorf("Invalid request")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	page, err := strconv.ParseUint(params.page, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	total, err := strconv.ParseUint(params.total, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	// should be positive number
	if page < 1 || total < 1 {
		return args, fmt.Errorf("Invalid request")
	}

	return readArgs{
		scheduleID: scheduleID,
		page:       page,
		total:      total,
	}, nil
}

func (params readDetailParams) validate() (readDetailArgs, error) {

	var args readDetailArgs
	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	return readDetailArgs{id: id}, nil
}
