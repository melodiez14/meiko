package course

import (
	"fmt"

	"database/sql"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetCourseByUserID(userID string) ([]Course, error) {
	var course []Course
	query := fmt.Sprintf(queryGetCourseByUserID, userID)
	err := conn.DB.Select(&course, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return course, nil
}
