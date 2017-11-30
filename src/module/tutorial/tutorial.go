package tutorial

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/melodiez14/meiko/src/util/conn"
)

// SelectByPage ...
func SelectByPage(scheduleID int64, limit, offset int, isCount bool) ([]Tutorial, int, error) {
	var tutorials []Tutorial
	var count int
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			description,
			created_at
		FROM
			tutorials
		WHERE
			schedules_id = (%d)
		LIMIT %d
		OFFSET %d;
		`, scheduleID, limit, offset)
	err := conn.DB.Select(&tutorials, query)
	if err != nil {
		return tutorials, count, err
	}

	if !isCount {
		return tutorials, count, err
	}

	query = fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM
			tutorials
		WHERE
			schedules_id = (%d);
		`, scheduleID)
	err = conn.DB.Get(&count, query)
	if err != nil {
		return tutorials, count, err
	}
	return tutorials, count, nil
}

// GetByID ...
func GetByID(id int64) (Tutorial, error) {
	var tutorial Tutorial
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			description,
			schedules_id,
			created_at
		FROM
			tutorials
		WHERE
			id = (%d)
		LIMIT 1;	
	`, id)

	err := conn.DB.Get(&tutorial, query)
	if err != nil {
		return tutorial, err
	}
	return tutorial, nil
}

// IsExistName ...
func IsExistName(name string, scheduleID int64, currentID ...int64) bool {
	var queryID string
	if len(currentID) == 1 {
		queryID = fmt.Sprintf("id != (%d) AND ", currentID[0])
	}

	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			tutorials
		WHERE
			%s
			name = ('%s') AND
			schedules_id = (%d)
		LIMIT 1;	
	`, queryID, name, scheduleID)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// Insert ...
func Insert(name string, description sql.NullString, scheduleID int64, tx *sqlx.Tx) (int64, error) {

	desc := "(NULL)"
	if description.Valid {
		desc = fmt.Sprintf("('%s')", description.String)
	}

	query := fmt.Sprintf(`
		INSERT INTO
			tutorials (
				name,
				description,
				schedules_id,
				created_at,
				updated_at
			) VALUES (
				('%s'),
				%s,
				(%d),
				NOW(),
				NOW()
			);
	`, name, desc, scheduleID)

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("No rows affected")
	}

	return lastInsertID, nil
}

// IsExistID ...
func IsExistID(id int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			tutorials
		WHERE
			id = ('%d')
		LIMIT 1;	
	`, id)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func Delete(id int64, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		DELETE FROM
			tutorials
		WHERE
			id = (%d);
	`, id)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}

	if err != nil {
		return err
	}

	return nil
}

func Update(id int64, name string, description sql.NullString, tx *sqlx.Tx) error {

	desc := "(NULL)"
	if description.Valid {
		desc = fmt.Sprintf("('%s')", description.String)
	}

	query := fmt.Sprintf(`
		UPDATE
			tutorials
		SET 
			name = ('%s'),
			description = %s,
			updated_at = NOW()
		WHERE
			id = (%d)
	`, name, desc, id)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	return nil
}
