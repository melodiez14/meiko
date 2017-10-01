package place

import (
	"html"
)

func (params searchParams) Validation() (searchArgs, error) {
	return searchArgs{
		Query: html.EscapeString(params.Query),
	}, nil
}
