package tutorial

import (
	"database/sql"
	"fmt"
	"html"
	"strconv"
	"strings"

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
		payload:    params.payload,
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

func (params createParams) validate() (createArgs, error) {

	var args createArgs
	params = createParams{
		name:        html.EscapeString(params.name),
		description: html.EscapeString(params.description),
		fileID:      params.fileID,
		scheduleID:  params.scheduleID,
	}

	if helper.IsEmpty(params.name) {
		return args, fmt.Errorf("Name cannot be empty")
	}

	if len(params.name) > 100 {
		return args, fmt.Errorf("Name maximum consist of 100 character")
	}

	name := strings.Title(params.name)

	var desc sql.NullString
	if !helper.IsEmpty(params.description) {
		desc = sql.NullString{
			Valid:  true,
			String: params.description,
		}
	}

	if !helper.IsValidFileID(params.fileID) {
		return args, fmt.Errorf("Invalid Request")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	return createArgs{
		name:        name,
		description: desc,
		fileID:      params.fileID,
		scheduleID:  scheduleID,
	}, nil
}

func (params deleteParams) validate() (deleteArgs, error) {
	var args deleteArgs
	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	return deleteArgs{id: id}, nil
}

func (params updateParams) validate() (updateArgs, error) {
	var args updateArgs
	params = updateParams{
		id:          params.id,
		name:        html.EscapeString(params.name),
		description: html.EscapeString(params.description),
		fileID:      params.fileID,
		scheduleID:  params.scheduleID,
	}

	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	if helper.IsEmpty(params.name) {
		return args, fmt.Errorf("Name cannot be empty")
	}

	if len(params.name) > 100 {
		return args, fmt.Errorf("Name maximum consist of 100 character")
	}

	name := strings.Title(params.name)

	var desc sql.NullString
	if !helper.IsEmpty(params.description) {
		desc = sql.NullString{
			Valid:  true,
			String: params.description,
		}
	}

	if !helper.IsValidFileID(params.fileID) {
		return args, fmt.Errorf("Invalid Request")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	return updateArgs{
		id:          id,
		name:        name,
		description: desc,
		fileID:      params.fileID,
		scheduleID:  scheduleID,
	}, nil
}
