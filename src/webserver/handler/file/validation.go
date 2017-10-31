package file

import (
	"fmt"
	"html"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params uploadImageParams) validate() (uploadImageArgs, error) {

	var args uploadImageArgs
	params = uploadImageParams{
		Height:    params.Height,
		Width:     params.Width,
		FileName:  html.EscapeString(params.FileName),
		Extension: params.Extension,
		Mime:      params.Mime,
	}

	if params.Height < 300 || params.Width < 300 {
		return args, fmt.Errorf("Height and Width should be 300px minimum")
	}

	if helper.IsEmpty(params.FileName) {
		return args, fmt.Errorf("Filename is empty")
	}

	if !helper.IsImageExtension(params.Extension) {
		return args, fmt.Errorf("Invalid file extensions")
	}

	if !helper.IsImageMime(params.Mime) {
		return args, fmt.Errorf("Invalid file type")
	}

	return uploadImageArgs{
		FileName:  params.FileName,
		Extension: params.Extension,
		Mime:      params.Mime,
	}, nil
}

func (params uploadAssignmentParams) validate() (uploadAssignmentArgs, error) {
	var args uploadAssignmentArgs
	params = uploadAssignmentParams{
		FileName:  html.EscapeString(params.FileName),
		Extension: params.Extension,
		Mime:      params.Mime,
	}

	if helper.IsEmpty(params.FileName) {
		return args, fmt.Errorf("Filename is empty")
	}

	// validate mime and extension

	return uploadAssignmentArgs{
		FileName:  params.FileName,
		Extension: params.Extension,
		Mime:      params.Mime,
	}, nil
}

func (params getFileParams) validate() (getFileArgs, error) {

	// do a validation if necessary

	return getFileArgs{
		payload:  params.payload,
		filename: params.filename,
	}, nil
}
