package information

import (
	"fmt"
	"strings"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/util/helper"
)

func SelectByScheduleID(scheduleID []int64, column ...string) ([]Information, error) {

	var info []Information
	var c []string
	d := helper.Int64ToStringSlice(scheduleID)

	if len(scheduleID) < 1 {
		return info, nil
	}

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

func SelectByScheduleIDAndTime(scheduleID []int64, t []time.Time, column ...string) ([]Information, error) {

	var info []Information
	var c []string
	d := helper.Int64ToStringSlice(scheduleID)

	if len(scheduleID) < 1 {
		return info, nil
	}

	var queryTime string
	if len(t) == 1 {
		queryTime = fmt.Sprintf("AND date(created_at) = ('%s')", t[0].Format("2006-01-02"))
	} else if len(t) == 2 {
		queryTime = fmt.Sprintf("AND date(created_at) BETWEEN ('%s') AND ('%s')", t[0].Format("2006-01-02"), t[1].Format("2006-01-02"))
	} else if len(t) > 2 {
		return info, fmt.Errorf("date more than two")
	}

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
			WHERE (
					schedules_id IS NULL
				OR
					schedules_id IN (%s)
			) %s
			ORDER BY created_at DESC
			LIMIT 5`, cols, ids, queryTime)
	err := conn.DB.Select(&info, query)
	if err != nil {
		return info, err
	}
	return info, nil
}
