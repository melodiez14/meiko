package bot

import (
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

func LoadByUserID(userID int64) ([]Log, error) {

	var log []Log

	query := fmt.Sprintf(`
			SELECT
				id,
				message,
				status,
				created_at
			FROM
				bot_logs
			WHERE
				users_id = (%d)
			ORDER BY id DESC
			LIMIT 20;
		`, userID)

	err := conn.DB.Select(&log, query)
	if err != nil {
		return log, err
	}

	logT := log
	log = []Log{}
	for i := len(logT) - 1; i >= 0; i-- {
		log = append(log, logT[i])
	}

	return log, nil
}

func LoadByID(id, userID int64) ([]Log, error) {

	var log []Log

	query := fmt.Sprintf(`
		SELECT
			id,
			message,
			status,
			created_at
		FROM
			bot_logs
		WHERE
			id < (%d) AND
			users_id = (%d)
		ORDER BY id DESC
		LIMIT 20;
	`, id, userID)

	err := conn.DB.Select(&log, query)
	if err != nil {
		return log, err
	}

	logT := log
	log = []Log{}
	for i := len(logT) - 1; i >= 0; i-- {
		log = append(log, logT[i])
	}

	return log, nil
}
