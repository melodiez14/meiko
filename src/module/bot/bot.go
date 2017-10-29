package bot

import (
	"fmt"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"
)

func LoadByTime(t time.Time, isAfter bool, userID int64) ([]Log, error) {

	var log []Log

	opr := "<"
	order := "DESC"
	if isAfter {
		opr = ">"
		order = "ASC"
	}

	query := fmt.Sprintf(`
		SELECT
			message,
			status,
			created_at
		FROM
			bot_logs
		WHERE
			created_at %s ('%s') AND
			users_id = (%d)
		ORDER BY
			created_at %s
		LIMIT 10;
	`, opr, t, userID, order)

	err := conn.DB.Select(&log, query)
	if err != nil {
		return log, err
	}

	if !isAfter {
		logT := log
		log = []Log{}
		for i := len(logT) - 1; i >= 0; i-- {
			log = append(log, logT[i])
		}
	}

	return log, nil
}
