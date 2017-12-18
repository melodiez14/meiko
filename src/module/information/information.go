package information

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/util/helper"
)

//SelectByScheduleID func ...
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
func GetScheduleIDByID(informationID int64) *int64 {
	var id *int64
	query := fmt.Sprintf(`
		SELECT
			schedules_id
		FROM
			informations
		WHERE
			id = (%d)
		LIMIT 1
		;`, informationID)

	err := conn.DB.Get(id, query)
	if err != nil {
		return nil
	}
	return id
}

// Insert func ...
func Insert(title, description string, scheduleID int64, tx *sqlx.Tx) (string, error) {
	var c []string
	var data string
	var result sql.Result
	var err error

	if scheduleID == 0 {
		c = []string{
			ColTitle,
			ColDescription,
			CreatedAt,
			UpdatedAt,
		}
		data = fmt.Sprintf(`'%s','%s', NOW(), NOW()`, title, description)
	} else {
		c = []string{
			ColTitle,
			ColDescription,
			ColScheduleID,
			CreatedAt,
			UpdatedAt,
		}
		data = fmt.Sprintf(`'%s','%s',(%d), NOW(), NOW()`, title, description, scheduleID)
	}
	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(`
		INSERT INTO
			informations
			(
				%s
			)
		VALUES
			(
				%s
			)
		;`, cols, data)

	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return "", err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return "", fmt.Errorf("No rows affected")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("Error get LastIsertedId")
	}
	ID := strconv.FormatInt(id, 10)
	return ID, nil
}

// Update func ...
func Update(title, description string, scheduleID, informationID int64, tx *sqlx.Tx) error {
	var data string
	if scheduleID == 0 {
		data = fmt.Sprintf(`
			title = '%s',
			description = '%s',
			schedules_id = NULL,
			updated_at = NOW()`, title, description)
	} else {
		data = fmt.Sprintf(`
			title = '%s',
			description = '%s',
			schedules_id = %d,
			updated_at = NOW()`, title, description, scheduleID)
	}
	query := fmt.Sprintf(`
		UPDATE 
			informations
		SET
			%s
		WHERE
			id = (%d)
		;`, data, informationID)
	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

// IsInformationIDExist func ...
func IsInformationIDExist(informationID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			informations
		WHERE
			id = (%d)
		LIMIT 1
		;`, informationID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// Delete func ...
func Delete(informationID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM
			informations
		WHERE
			id = (%d)
		;`, informationID)
	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if row == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

//SelectByPage func ...
func SelectByPage(scheduleID []int64, total, offset int64, column ...string) ([]Information, error) {

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
		LIMIT %d OFFSET %d`, cols, ids, total, offset)

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
