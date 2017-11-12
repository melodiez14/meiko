package bot

import (
	"database/sql"
	"fmt"
	"html"
	"strconv"
	"strings"

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
	var id sql.NullInt64

	if !helper.IsEmpty(params.id) {
		temp, err := strconv.ParseInt(params.id, 10, 64)
		if err != nil {
			return args, err
		}
		id = sql.NullInt64{Valid: true, Int64: temp}
	}

	return loadHistoryArgs{id: id}, nil
}
