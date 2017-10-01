package place

import (
	"html"
)

func (params searchParams) Validate() (searchArgs, error) {
	return searchArgs{
		Query: html.EscapeString(params.Query),
	}, nil
}
