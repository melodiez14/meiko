package course

import (
	"fmt"
	"strings"

	"github.com/melodiez14/meiko/src/util/helper"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/meiko/src/util/conn"
)

func SelectIDByUserID(userID int64, status ...int8) ([]int64, error) {
	var st string
	if len(status) == 1 {
		st = fmt.Sprintf("AND status = (%d)", status[0])
	}

	var scheduleID []int64
	query := fmt.Sprintf(`SELECT schedules_id FROM p_users_schedules WHERE users_id = (%d) %s`, userID, st)
	err := conn.DB.Select(&scheduleID, query)
	if err != nil && err != sql.ErrNoRows {
		return scheduleID, err
	}

	return scheduleID, nil
}

func SelectScheduleIDByUserID(userID int64, status ...int8) ([]int64, error) {

	var st string
	if len(status) == 1 {
		st = fmt.Sprintf("AND status = (%d)", status[0])
	}

	var scheduleIDs []int64
	query := fmt.Sprintf(`SELECT schedules_id FROM p_users_schedules WHERE users_id = (%d) %s;`, userID, st)
	err := conn.DB.Select(&scheduleIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return scheduleIDs, err
	}
	return scheduleIDs, nil
}

func IsEnrolled(userID, scheduleID int64) bool {
	var x string
	query := fmt.Sprintf("SELECT 'x' FROM p_users_schedules WHERE users_id = (%d) AND schedules_id = (%d) LIMIT 1", userID, scheduleID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func SelectAssistantID(scheduleID int64) ([]int64, error) {

	userIDs := []int64{}
	query := fmt.Sprintf(`SELECT
		p.users_id
	FROM
		p_users_schedules p
	WHERE 
		p.status = (%d) AND
		p.schedules_id = (%d);`, PStatusAssistant, scheduleID)
	err := conn.DB.Select(&userIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return userIDs, err
	}

	return userIDs, nil
}

func IsExist(courseID string) bool {

	var x string
	query := fmt.Sprintf(`SELECT 'x' FROM courses WHERE id = ('%s') LIMIT 1`, courseID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true

}

func Update(courseID, name string, description sql.NullString, ucu int8, tx ...*sqlx.Tx) error {

	queryDescription := fmt.Sprintf("NULL")
	if description.Valid {
		queryDescription = fmt.Sprintf("('%s')", description.String)
	}

	query := fmt.Sprintf(`
		UPDATE
			courses
		SET
			name = ('%s'),
			description = %s,
			ucu = (%d),
			updated_at = NOW()
		WHERE
			id = ('%s');
		`, name, queryDescription, ucu, courseID)

	var result sql.Result
	var err error
	if len(tx) == 1 {
		result, err = tx[0].Exec(query)
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

func Insert(courseID, name string, description sql.NullString, ucu int8, tx ...*sqlx.Tx) error {

	queryDescription := fmt.Sprintf("NULL")
	if description.Valid {
		queryDescription = fmt.Sprintf("('%s')", description.String)
	}

	query := fmt.Sprintf(`
		INSERT INTO
			courses (
				id,
				name,
				description,
				ucu,
				created_at,
				updated_at
			)
		VALUES (
			('%s'),
			('%s'),
			%s,
			(%d),
			NOW(),
			NOW()
		);`, courseID, name, queryDescription, ucu)

	var result sql.Result
	var err error
	if len(tx) == 1 {
		result, err = tx[0].Exec(query)
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

func IsExistSchedule(semester int8, year int16, courseID, class string, scheduleID ...int64) bool {

	var sc string
	if len(scheduleID) == 1 {
		sc = fmt.Sprintf(" AND id != (%d) ", scheduleID[0])
	}

	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			schedules
		WHERE
			semester = (%d) AND
			year = (%d) AND
			courses_id = ('%s') AND
			class = ('%s') %s
		LIMIT 1;`, semester, year, courseID, class, sc)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func InsertSchedule(userID int64, startTime, endTime, year int16, semester, day, status int8, class, courseID, placeID string, tx ...*sqlx.Tx) error {

	query := fmt.Sprintf(`
		INSERT INTO
			schedules (
				status,
				start_time,
				end_time,
				day,
				class,
				semester,
				year,
				courses_id,
				places_id,
				created_by,
				created_at,
				updated_at
			)
		VALUES (
			(%d),
			(%d),
			(%d),
			(%d),
			('%s'),
			(%d),
			(%d),
			('%s'),
			('%s'),
			(%d),
			NOW(),
			NOW()
		)`, status, startTime, endTime, day, class, semester, year, courseID, placeID, userID)

	var result sql.Result
	var err error
	switch len(tx) {
	case 1:
		result, err = tx[0].Exec(query)
	default:
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

func SelectByPage(limit, offset uint16) ([]CourseSchedule, error) {

	var course []CourseSchedule
	query := fmt.Sprintf(`
		SELECT
			cs.id,
			cs.name,
			cs.description,
			cs.ucu,
			sc.id,
			sc.status,
			sc.start_time,
			sc.end_time,
			sc.day,
			sc.class,
			sc.semester,
			sc.year,
			sc.places_id,
			sc.created_by
		FROM
			courses cs
		RIGHT JOIN
			schedules sc
		ON
			cs.id = sc.courses_id
		LIMIT %d OFFSET %d;`, limit, offset)

	rows, err := conn.DB.Queryx(query)
	defer rows.Close()
	if err != nil {
		return course, err
	}

	for rows.Next() {
		var id, name, class, placeID string
		var description sql.NullString
		var ucu, status, day, semester int8
		var startTime, endTime uint16
		var year int16
		var scheduleID, createdBy int64

		err := rows.Scan(&id, &name, &description, &ucu, &scheduleID, &status, &startTime, &endTime, &day, &class, &semester, &year, &placeID, &createdBy)
		if err != nil {
			return course, err
		}

		course = append(course, CourseSchedule{
			Course: Course{
				ID:          id,
				Name:        name,
				Description: description,
				UCU:         ucu,
			},
			Schedule: Schedule{
				ID:        scheduleID,
				Status:    status,
				StartTime: startTime,
				EndTime:   endTime,
				Day:       day,
				Class:     class,
				Semester:  semester,
				Year:      year,
				PlaceID:   placeID,
				CreatedBy: createdBy,
			},
		})
	}

	return course, nil
}

func GetByScheduleID(scheduleID int64) (CourseSchedule, error) {

	var course CourseSchedule
	query := fmt.Sprintf(`
		SELECT
			cs.id,
			cs.name,
			cs.description,
			cs.ucu,
			sc.id,
			sc.status,
			sc.start_time,
			sc.end_time,
			sc.day,
			sc.class,
			sc.semester,
			sc.year,
			sc.places_id,
			sc.created_by
		FROM
			courses cs
		RIGHT JOIN
			schedules sc
		ON
			cs.id = sc.courses_id
		WHERE
			sc.id = (%d)
		LIMIT 1;`, scheduleID)

	rows := conn.DB.QueryRowx(query)

	// scan data to variable
	var id, name, class, placeID string
	var description sql.NullString
	var ucu, status, day, semester int8
	var startTime, endTime uint16
	var year int16
	var createdBy int64

	err := rows.Scan(&id, &name, &description, &ucu, &scheduleID, &status, &startTime, &endTime, &day, &class, &semester, &year, &placeID, &createdBy)
	if err != nil {
		return course, err
	}

	return CourseSchedule{
		Course: Course{
			ID:          id,
			Name:        name,
			Description: description,
			UCU:         ucu,
		},
		Schedule: Schedule{
			ID:        scheduleID,
			Status:    status,
			StartTime: startTime,
			EndTime:   endTime,
			Day:       day,
			Class:     class,
			Semester:  semester,
			Year:      year,
			PlaceID:   placeID,
			CreatedBy: createdBy,
		},
	}, nil
}

func IsExistScheduleID(scheduleID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			schedules
		WHERE
			id = (%d)
		LIMIT 1;`, scheduleID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func UpdateSchedule(scheduleID int64, startTime, endTime, year int16, semester, day, status int8, class, courseID, placeID string, tx ...*sqlx.Tx) error {

	query := fmt.Sprintf(`
		UPDATE 
			schedules
		SET
			status = (%d),
			start_time = (%d),
			end_time = (%d),
			day = (%d),
			class = ('%s'),
			semester = (%d),
			year = (%d),
			courses_id = ('%s'),
			places_id = ('%s'),
			updated_at = NOW()
		WHERE
			id = (%d);`, status, startTime, endTime, day, class, semester, year, courseID, placeID, scheduleID)

	var result sql.Result
	var err error
	switch len(tx) {
	case 1:
		result, err = tx[0].Exec(query)
	default:
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

func SelectByScheduleID(scheduleID []int64, status int8) ([]CourseSchedule, error) {

	var course []CourseSchedule
	if len(scheduleID) < 1 {
		return course, nil
	}

	d := helper.Int64ToStringSlice(scheduleID)
	ids := strings.Join(d, ", ")

	query := fmt.Sprintf(`
		SELECT
			cs.id,
			cs.name,
			cs.description,
			cs.ucu,
			sc.id,
			sc.status,
			sc.start_time,
			sc.end_time,
			sc.day,
			sc.class,
			sc.semester,
			sc.year,
			sc.places_id,
			sc.created_by
		FROM
			courses cs
		RIGHT JOIN
			schedules sc
		ON
			cs.id = sc.courses_id
		WHERE
			sc.id IN (%s) AND
			sc.status = (%d)`, ids, status)

	rows, err := conn.DB.Queryx(query)
	defer rows.Close()
	if err != nil {
		return course, err
	}

	for rows.Next() {
		var id, name, class, placeID string
		var description sql.NullString
		var ucu, status, day, semester int8
		var startTime, endTime uint16
		var year int16
		var scheduleID, createdBy int64

		err := rows.Scan(&id, &name, &description, &ucu, &scheduleID, &status, &startTime, &endTime, &day, &class, &semester, &year, &placeID, &createdBy)
		if err != nil {
			return course, err
		}

		course = append(course, CourseSchedule{
			Course: Course{
				ID:          id,
				Name:        name,
				Description: description,
				UCU:         ucu,
			},
			Schedule: Schedule{
				ID:        scheduleID,
				Status:    status,
				StartTime: startTime,
				EndTime:   endTime,
				Day:       day,
				Class:     class,
				Semester:  semester,
				Year:      year,
				PlaceID:   placeID,
				CreatedBy: createdBy,
			},
		})
	}

	return course, nil
}

func SelectByStatus(status int8) ([]CourseSchedule, error) {

	var course []CourseSchedule
	query := fmt.Sprintf(`
		SELECT
			cs.id,
			cs.name,
			cs.description,
			cs.ucu,
			sc.id,
			sc.status,
			sc.start_time,
			sc.end_time,
			sc.day,
			sc.class,
			sc.semester,
			sc.year,
			sc.places_id,
			sc.created_by
		FROM
			courses cs
		RIGHT JOIN
			schedules sc
		ON
			cs.id = sc.courses_id
		WHERE
			sc.status = (%d)`, status)

	rows, err := conn.DB.Queryx(query)
	defer rows.Close()
	if err != nil {
		return course, err
	}

	for rows.Next() {
		var id, name, class, placeID string
		var description sql.NullString
		var ucu, status, day, semester int8
		var startTime, endTime uint16
		var year int16
		var scheduleID, createdBy int64

		err := rows.Scan(&id, &name, &description, &ucu, &scheduleID, &status, &startTime, &endTime, &day, &class, &semester, &year, &placeID, &createdBy)
		if err != nil {
			return course, err
		}

		course = append(course, CourseSchedule{
			Course: Course{
				ID:          id,
				Name:        name,
				Description: description,
				UCU:         ucu,
			},
			Schedule: Schedule{
				ID:        scheduleID,
				Status:    status,
				StartTime: startTime,
				EndTime:   endTime,
				Day:       day,
				Class:     class,
				Semester:  semester,
				Year:      year,
				PlaceID:   placeID,
				CreatedBy: createdBy,
			},
		})
	}

	return course, nil
}

func DeleteSchedule(scheduleID int64, tx ...*sqlx.Tx) error {

	query := fmt.Sprintf(`
		DELETE FROM
			schedules
		WHERE
			id = (%d);
		`, scheduleID)

	var result sql.Result
	var err error
	if len(tx) == 1 {
		result, err = tx[0].Exec(query)
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
