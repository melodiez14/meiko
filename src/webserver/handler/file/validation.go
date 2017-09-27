package file

import (
	"fmt"
	"html"

	"github.com/melodiez14/meiko/src/util/alias"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params uploadImageParams) Validate() (uploadImageArgs, error) {

	var args uploadImageArgs
	params = uploadImageParams{
		Payload:   html.EscapeString(params.Payload),
		FileName:  html.EscapeString(params.FileName),
		Extension: params.Extension,
		Mime:      params.Mime,
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
		Payload:   params.Payload,
		FileName:  params.FileName,
		Extension: params.Extension,
		Mime:      params.Mime,
	}, nil
}

func (params getProfileParams) Validate() (getProfileArgs, error) {

	var args getProfileArgs

	var payload string
	switch params.Payload {
	case alias.ParamsProfile:
		payload = alias.TypeProfile
	case alias.ParamsProfileThumbnail:
		payload = alias.TypeProfileThumbnail
	default:
		return args, fmt.Errorf("Error parameter")
	}

	return getProfileArgs{
		Payload: payload,
	}, nil
}
