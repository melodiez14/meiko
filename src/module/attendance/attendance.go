package attendance

import (
	"database/sql"
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetByUserCourseID(userID, courseID int64) ([]Attendance, error) {
	var attendances []Attendance
	query := fmt.Sprintf(queryGetByUserCourse, userID, courseID)
	err := conn.DB.Select(&attendances, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return attendances, nil
}
