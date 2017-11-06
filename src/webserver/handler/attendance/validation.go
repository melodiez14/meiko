package attendance

import (
	"fmt"
	"strconv"

	"github.com/melodiez14/meiko/src/util/helper"
)

func (params listStudentParams) validate() (listStudentArgs, error) {
	var args listStudentArgs

	if helper.IsEmpty(params.meetingNumber) {
		return args, fmt.Errorf("ScheduleID cannot be empty")
	}

	if helper.IsEmpty(params.scheduleID) {
		return args, fmt.Errorf("ScheduleID cannot be empty")
	}

	scheduleID, err := strconv.ParseInt(params.scheduleID, 10, 64)
	if err != nil {
		return args, err
	}

	meetingNumber, err := strconv.ParseUint(params.meetingNumber, 10, 8)
	if err != nil {
		return args, err
	}

	return listStudentArgs{
		meetingNumber: uint8(meetingNumber),
		scheduleID:    scheduleID,
	}, nil
}
