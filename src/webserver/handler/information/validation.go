package information

import (
	"fmt"
	"strconv"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params detailInfromationParams) validate() (detailInfromationArgs, error) {
	var args detailInfromationArgs
	// Information ID validation
	if helper.IsEmpty(params.ID) {
		return args, fmt.Errorf("Information ID can not be empty")
	}
	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Error to convert Information string ID to int64")
	}
	return detailInfromationArgs{
		ID: id,
	}, nil
}
