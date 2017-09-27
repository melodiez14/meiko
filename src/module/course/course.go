package course

import (
	"fmt"

	"database/sql"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetByUserID(userID int64) ([]Course, error) {
	var courses []Course
	query := fmt.Sprintf(queryGetCourseByUserID, userID)
	err := conn.DB.Select(&courses, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return courses, nil
}

func SelectIDByUserID(userID int64) ([]int64, error) {
	courseIDs := []int64{}
	query := fmt.Sprintf(queryGetIDByUserID, userID)
	err := conn.DB.Select(&courseIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return courseIDs, err
	}
	return courseIDs, nil
}
