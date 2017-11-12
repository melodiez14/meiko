package bot

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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

// SelectAssistantWithCourse ...
func SelectAssistantWithCourse(userID int64, rgxCourse sql.NullString, days []int8) ([]Assistant, error) {

	var assistants []Assistant
	var queryCourse string
	var queryDay string

	// validate regex
	if rgxCourse.Valid {
		queryCourse = fmt.Sprintf(`LOWER(c.name) REGEXP '%s' AND`, rgxCourse.String)
	}

	if len(days) > 0 {
		daysString := []string{}
		for _, val := range days {
			daysString = append(daysString, strconv.FormatInt(int64(val), 10))
		}
		queryDay = fmt.Sprintf(`s.day IN (%s)`, strings.Join(daysString, ", "))
	}

	query := fmt.Sprintf(`
		SELECT
			u.identity_code,
			u.name,
			COALESCE(u.phone, '-') as phone,
			COALESCE(u.line_id, '-') as line_id,
			c.id as courses_id,
			c.name as courses_name
		FROM
			users u
		INNER JOIN p_users_schedules pus ON u.id = pus.users_id
		INNER JOIN schedules s ON pus.schedules_id = s.id
		INNER JOIN courses c ON s.courses_id = c.id
		WHERE
			%s %s
			pus.status = 2 AND
			u.id IN (
				SELECT
					DISTINCT(users_id)
				FROM p_users_schedules
					WHERE schedules_id
				IN (
					SELECT
						DISTINCT(schedules_id)
					FROM
						p_users_schedules
					WHERE
						users_id = (%d) AND
						status = 1
				)	
			);
	`, queryCourse, queryDay, userID)

	err := conn.DB.Select(&assistants, query)
	if err != nil {
		return assistants, err
	}

	return assistants, nil
}
