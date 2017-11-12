package information

import (
	"fmt"
	"strings"

	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/util/helper"
)

//SelectByScheduleID func ...
func SelectByScheduleID(scheduleID []int64, column ...string) ([]Information, error) {

	var info []Information
	var c []string
	d := helper.Int64ToStringSlice(scheduleID)

	if len(column) < 1 {
		c = []string{
			ColID,
			ColTitle,
			ColDescription,
			ColScheduleID,
			CreatedAt,
			UpdatedAt,
		}
	} else {
		for _, val := range column {
			c = append(c, val)
		}
	}
	ids := strings.Join(d, ", ")
	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(`
		SELECT
			%s
		FROM
			informations
		WHERE
			schedules_id IS NULL
		OR
			schedules_id IN (%s)
		ORDER BY created_at DESC
		LIMIT 100`, cols, ids)
	err := conn.DB.Select(&info, query)
	if err != nil {
		return info, err
	}
	return info, nil
}

// GetByID func ...
func GetByID(informationID int64, column ...string) (Information, error) {
	var info Information
	var c []string

	if len(column) < 1 {
		c = []string{
			ColID,
			ColTitle,
			ColDescription,
			ColScheduleID,
			CreatedAt,
			UpdatedAt,
		}
	} else {
		for _, val := range column {
			c = append(c, val)
		}
	}
	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(`
		SELECT
			%s
		FROM
			informations
		WHERE
			id = (%d)
		LIMIT 1
		;`, cols, informationID)
	err := conn.DB.Get(&info, query)
	if err != nil {
		return info, err
	}
	return info, nil
}

// GetScheduleIDByID func ...
func GetScheduleIDByID(informationID int64) int64 {
	query := fmt.Sprintf(`
		SELECT
			schedules_id
		FROM
			informations
		WHERE
			id = (%d)
		LIMIT 1
		;`, informationID)

	var id int64
	err := conn.DB.Get(&id, query)
	if err != nil {
		return 0
	}
	return id
}
