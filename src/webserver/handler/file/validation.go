package file

import (
	"fmt"
	"html"
	"strconv"

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

func (params uploadFileParams) validate() (uploadFileArgs, error) {

	id, _ := strconv.ParseInt(params.id, 10, 64)

	// validate mime and extension

	return uploadFileArgs{
		id:      id,
		payload: params.payload,
		role:    params.role,
	}, nil
}

func (params metaParams) validate() (metaArgs, error) {
	var args metaArgs
	params = metaParams{
		fileName:  html.EscapeString(params.fileName),
		extension: params.extension,
		mime:      params.mime,
	}

	if helper.IsEmpty(params.fileName) {
		return args, fmt.Errorf("Filename is empty")
	}

	// validate mime and extension

	return metaArgs{
		fileName:  params.fileName,
		extension: params.extension,
		mime:      params.mime,
	}, nil
}

func (params getFileParams) validate() (getFileArgs, error) {

	// do a validation if necessary

	return getFileArgs{
		payload:  params.payload,
		filename: params.filename,
	}, nil
}

func (params routerParams) validate() (routerArgs, error) {

	var args routerArgs
	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	return routerArgs{
		payload: params.payload,
		role:    params.role,
		id:      id,
	}, nil
}
