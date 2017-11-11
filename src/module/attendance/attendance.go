package attendance

import (
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

// func GetByUserCourseID(userID, courseID int64) ([]Attendance, error) {
// 	var attendances []Attendance
// 	query := fmt.Sprintf(queryGetByUserCourse, userID, courseID)
// 	err := conn.DB.Select(&attendances, query)
// 	if err != nil && err != sql.ErrNoRows {
// 		return nil, err
// 	}
// 	return attendances, nil
// }

func GetMeeting(meetingNumber uint8, scheduleID int64) (Meeting, error) {

	var meeting Meeting
	query := fmt.Sprintf(`
		SELECT
			id,
			description,
			date
		FROM
			meetings
		WHERE
			number = (%d) AND
			schedules_id = (%d)
		LIMIT 1;
		`, meetingNumber, scheduleID)
	err := conn.DB.Get(&meeting, query)
	if err != nil {
		return meeting, err
	}

	return meeting, nil
}
