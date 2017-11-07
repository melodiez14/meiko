package log

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/meiko/src/util/conn"
)

// Insert used for logging the data and inserting the log into bot_logs table
func Insert(text string, userID int64, status uint8, tx *sqlx.Tx) (int64, error) {

	var result sql.Result
	var err error

	query := fmt.Sprintf(`
		INSERT INTO
			bot_logs(
				message,
				users_id,
				status,
				created_at
			) VALUES (
				('%s'),
				(%d),
				(%d),
				NOW()
			);`, text, userID, status)

	if tx != nil {
		result, err = conn.DB.Exec(query)
	} else {
		result, err = tx.Exec(query)
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
