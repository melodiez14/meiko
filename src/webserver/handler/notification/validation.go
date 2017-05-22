package notification

import (
	"fmt"
	"strconv"
)

func (n getNotificationParam) validate() (*getNotificationArgs, error) {

	if len(n.page) < 1 {
		return nil, fmt.Errorf("Error validation : page cant't be empty")
	}

	pg, err := strconv.ParseUint(n.page, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Error validation : error parsing page")
	}

	return &getNotificationArgs{
		page: uint16(pg),
	}, nil
}
