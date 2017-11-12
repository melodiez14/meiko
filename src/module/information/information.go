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

// Insert func ...
func Insert(title, description string, scheduleID int64) error {
	var c []string
	var data string
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

// Update func ...
func Update(title, description string, scheduleID, informationID int64) error {
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
