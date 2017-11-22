package bot

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

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
			s.status = 1 AND
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

func SelectScheduleWithCourse(userID int64, rgxCourse sql.NullString, days []int8) ([]Schedule, error) {

	var schedules []Schedule
	var queryCourse string
	var queryDay string

	// course query regex
	if rgxCourse.Valid {
		queryCourse = fmt.Sprintf(`LOWER(c.name) REGEXP '%s' AND`, rgxCourse.String)
	}

	// days query
	if len(days) > 0 {
		daysString := []string{}
		for _, val := range days {
			daysString = append(daysString, strconv.FormatInt(int64(val), 10))
		}
		queryDay = fmt.Sprintf(`s.day IN (%s) AND`, strings.Join(daysString, ", "))
	}

	query := fmt.Sprintf(`
		SELECT
			c.name,
			s.day,
			s.places_id,
			s.start_time,
			s.end_time
		FROM
			schedules s
		INNER JOIN courses c ON s.courses_id = c.id
		WHERE
			s.status = 1 AND
			%s %s
			s.id IN (
				SELECT
					schedules_id
				FROM
					p_users_schedules
				WHERE
					users_id = (%d) AND
					status = 1
			);
		`, queryDay, queryCourse, userID)

	err := conn.DB.Select(&schedules, query)
	if err != nil {
		return schedules, err
	}

	return schedules, nil
}

func SelectAssignmentWithCourse(userID int64, rgxCourse sql.NullString, t []time.Time) ([]Assignment, error) {

	var assignments []Assignment
	var queryCourse string
	var queryTime string

	// course query regex
	if rgxCourse.Valid {
		queryCourse = fmt.Sprintf(`LOWER(c.name) REGEXP '%s' AND`, rgxCourse.String)
	}

	// days query
	switch len(t) {
	case 1:
		queryTime = fmt.Sprintf("date(a.due_date) = ('%s') AND", t[0].Format("2006-01-02"))
	case 2:
		queryTime = fmt.Sprintf("date(a.due_date) BETWEEN ('%s') AND ('%s') AND", t[0].Format("2006-01-02"), t[1].Format("2006-01-02"))
	}

	query := fmt.Sprintf(`
		SELECT
			a.id,
			a.name,
			COALESCE(a.description, '-') as description,
			a.due_date,
			c.name as course_name
		FROM
			assignments a
		INNER JOIN grade_parameters g ON g.id = a.grade_parameters_id
		INNER JOIN schedules s ON s.id = g.schedules_id
		INNER JOIN courses c ON c.id = s.courses_id
		WHERE
			%s %s
			a.id NOT IN (
				SELECT
					assignments_id
				FROM
					p_users_assignments
				WHERE
					users_id = (%d)
			) AND
			s.id IN (
				SELECT
					schedules_id
				FROM
					p_users_schedules
				WHERE
					users_id = (%d) AND
					status = 1
				)
		ORDER BY a.due_date ASC
		LIMIT 5;
		`, queryTime, queryCourse, userID, userID)

	err := conn.DB.Select(&assignments, query)
	if err != nil {
		return assignments, err
	}

	return assignments, nil
}

func SelectGradeWithCourse(userID int64, rgxCourse sql.NullString, t []time.Time) ([]Grade, error) {

	var grades []Grade
	var queryCourse string
	var queryTime string

	// course query regex
	if rgxCourse.Valid {
		queryCourse = fmt.Sprintf(`LOWER(c.name) REGEXP '%s' AND`, rgxCourse.String)
	}

	// days query
	switch len(t) {
	case 1:
		queryTime = fmt.Sprintf("date(p.updated_at) = ('%s') AND", t[0].Format("2006-01-02"))
	case 2:
		queryTime = fmt.Sprintf("date(p.updated_at) BETWEEN ('%s') AND ('%s') AND", t[0].Format("2006-01-02"), t[1].Format("2006-01-02"))
	}

	query := fmt.Sprintf(`
		SELECT
			a.id,
			a.name,
			c.name as course_name,
			p.score,
			p.updated_at
		FROM
			assignments a
		INNER JOIN grade_parameters g ON g.id = a.grade_parameters_id
		INNER JOIN schedules s ON s.id = g.schedules_id
		INNER JOIN courses c ON c.id = s.courses_id
		INNER JOIN p_users_assignments p ON p.assignments_id = a.id
		WHERE
			%s %s
			p.users_id = (%d) AND
			p.score IS NOT NULL AND
			s.id IN (
				SELECT
					schedules_id
				FROM
					p_users_schedules
				WHERE
					users_id = (%d) AND
					status = 1
				)
		ORDER BY p.updated_at ASC
		LIMIT 5;
		`, queryTime, queryCourse, userID, userID)

	err := conn.DB.Select(&grades, query)
	if err != nil {
		return grades, err
	}

	return grades, nil
}
