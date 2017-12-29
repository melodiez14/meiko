package course

import (
	"fmt"

	cs "github.com/melodiez14/meiko/src/module/course"
	"github.com/melodiez14/meiko/src/util/helper"
)

func getLast(userID int64) ([]getResponse, error) {
	resp := []getResponse{}
	coursesID, err := cs.SelectScheduleIDByUserID(userID, cs.PStatusStudent)
	if err != nil {
		return resp, err
	}
	courses, err := cs.SelectByScheduleID(coursesID, cs.StatusScheduleInactive)
	if err != nil {
		return resp, err
	}

	for _, val := range courses {
		startTime := helper.MinutesToTimeString(val.Schedule.StartTime)
		endTime := helper.MinutesToTimeString(val.Schedule.EndTime)
		t := fmt.Sprintf("%s - %s", startTime, endTime)

		resp = append(resp, getResponse{
			ID:          val.Schedule.ID,
			Name:        val.Course.Name,
			Description: val.Course.Description.String,
			Class:       val.Schedule.Class,
			Semester:    val.Schedule.Semester,
			Day:         helper.IntDayToString(val.Schedule.Day),
			Time:        t,
			Place:       val.Schedule.PlaceID,
			Status:      "enrolled",
		})
	}

	return resp, nil
}

func getCurrent(userID int64) ([]getResponse, error) {
	resp := []getResponse{}
	coursesID, err := cs.SelectScheduleIDByUserID(userID, cs.PStatusStudent)
	if err != nil {
		return resp, err
	}
	courses, err := cs.SelectByScheduleID(coursesID, cs.StatusScheduleActive)
	if err != nil {
		return resp, err
	}

	for _, val := range courses {
		startTime := helper.MinutesToTimeString(val.Schedule.StartTime)
		endTime := helper.MinutesToTimeString(val.Schedule.EndTime)
		t := fmt.Sprintf("%s - %s", startTime, endTime)

		resp = append(resp, getResponse{
			ID:          val.Schedule.ID,
			Name:        val.Course.Name,
			Description: val.Course.Description.String,
			Class:       val.Schedule.Class,
			Semester:    val.Schedule.Semester,
			Day:         helper.IntDayToString(val.Schedule.Day),
			Time:        t,
			Place:       val.Schedule.PlaceID,
			Status:      "enrolled",
		})
	}

	return resp, nil
}

func getAll(userID int64) ([]getResponse, error) {
	resp := []getResponse{}
	courses, err := cs.SelectByStatus(cs.StatusScheduleActive)
	if err != nil {
		return resp, err
	}

	enrolled, err := cs.SelectIDByUserID(userID, cs.PStatusStudent)
	if err != nil {
		return resp, err
	}

	unapproved, err := cs.SelectIDByUserID(userID, cs.PStatusUnapproved)
	if err != nil {
		return resp, err
	}

	enrolledResp := []getResponse{}
	unenrolledResp := []getResponse{}
	waitingResp := []getResponse{}
	for _, val := range courses {
		startTime := helper.MinutesToTimeString(val.Schedule.StartTime)
		endTime := helper.MinutesToTimeString(val.Schedule.EndTime)
		t := fmt.Sprintf("%s - %s", startTime, endTime)

		if helper.Int64InSlice(val.Schedule.ID, enrolled) {
			enrolledResp = append(enrolledResp, getResponse{
				ID:          val.Schedule.ID,
				Name:        val.Course.Name,
				Description: val.Course.Description.String,
				Class:       val.Schedule.Class,
				Semester:    val.Schedule.Semester,
				Day:         helper.IntDayToString(val.Schedule.Day),
				Time:        t,
				Place:       val.Schedule.PlaceID,
				Status:      "enrolled",
			})
		} else if helper.Int64InSlice(val.Schedule.ID, unapproved) {
			waitingResp = append(waitingResp, getResponse{
				ID:          val.Schedule.ID,
				Name:        val.Course.Name,
				Description: val.Course.Description.String,
				Class:       val.Schedule.Class,
				Semester:    val.Schedule.Semester,
				Day:         helper.IntDayToString(val.Schedule.Day),
				Time:        t,
				Place:       val.Schedule.PlaceID,
				Status:      "waiting",
			})
		} else {
			unenrolledResp = append(unenrolledResp, getResponse{
				ID:          val.Schedule.ID,
				Name:        val.Course.Name,
				Description: val.Course.Description.String,
				Class:       val.Schedule.Class,
				Semester:    val.Schedule.Semester,
				Day:         helper.IntDayToString(val.Schedule.Day),
				Time:        t,
				Place:       val.Schedule.PlaceID,
				Status:      "unenrolled",
			})
		}
	}

	resp = append(enrolledResp, waitingResp...)
	resp = append(resp, unenrolledResp...)
	return resp, nil
}

// func getInactive(userID int64) ([]cs.CourseSchedule, error) {
// 	var courses []cs.CourseSchedule

// 	coursesID, err := cs.SelectScheduleIDByUserID(userID, cs.PStatusStudent)
// 	if err != nil {
// 		return courses, err
// 	}
// }
