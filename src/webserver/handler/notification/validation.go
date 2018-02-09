package notification

import (
	"fmt"
	"regexp"
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

func (params subscribeParams) validate() (subscribeArgs, error) {

	var args subscribeArgs
	valid, err := regexp.MatchString(`^[\w\d-]{36}$`, params.playerID)
	if err != nil {
		return args, err
	}

	if !valid {
		return args, fmt.Errorf("Invalid PlayerID")
	}

	return subscribeArgs{
		playerID: params.playerID,
	}, nil
}
