package bot

import (
	"fmt"
	"html"
	"strconv"
	"strings"
	"time"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params messageParams) validate() (messageArgs, error) {

	var args messageArgs
	text := html.EscapeString(params.Text)
	text = helper.Trim(text)

	normalized := strings.ToLower(text)

	if helper.IsEmpty(normalized) {
		return args, fmt.Errorf("Text cannot be empty")
	}

	return messageArgs{
		Text:           text,
		NormalizedText: normalized,
	}, nil
}

func (params loadHistoryParams) validate() (loadHistoryArgs, error) {

	var args loadHistoryArgs

	timeInt, err := strconv.ParseInt(params.Time, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Time must be unix format")
	}
	t := time.Unix(timeInt, 0)

	isAfter := false
	if params.Position == "after" {
		isAfter = true
	}

	return loadHistoryArgs{
		Time:    t,
		IsAfter: isAfter,
	}, nil
}
