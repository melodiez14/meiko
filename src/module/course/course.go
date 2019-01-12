package course

import (
	"fmt"
	"strconv"
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

// CountEnrolled ...
func CountEnrolled(usersID []int64, scheduleID int64) (int, error) {
	var count int
	ids := helper.Int64ToStringSlice(usersID)
	queryID := strings.Join(ids, ", ")
	query := fmt.Sprintf("SELECT COUNT(*) FROM p_users_schedules WHERE users_id IN (%s) AND schedules_id = (%d) AND status = (%d) LIMIT 1", queryID, scheduleID, PStatusStudent)
	err := conn.DB.Get(&count, query)
	if err != nil {
		return count, err
	}
	return count, nil
}

func IsEnrolled(userID, scheduleID int64) bool {
	var x string
	query := fmt.Sprintf("SELECT 'x' FROM p_users_schedules WHERE users_id = (%d) AND schedules_id = (%d) AND status = (%d) LIMIT 1", userID, scheduleID, PStatusStudent)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func IsAssistant(userID, scheduleID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			p_users_schedules
		WHERE
			users_id = (%d) AND
			schedules_id = (%d) AND
			status = (%d)
		LIMIT 1;
	`, userID, scheduleID, PStatusAssistant)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func IsUnapproved(userID, scheduleID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			p_users_schedules
		WHERE
			users_id = (%d) AND
			schedules_id = (%d) AND
			status = (%d)
		LIMIT 1;
	`, userID, scheduleID, PStatusUnapproved)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func IsCreator(userID, scheduleID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			schedules
		WHERE
			id = (%d) AND
			created_by = (%d)
		LIMIT 1;
	`, scheduleID, userID)
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

func SelectAllAssistantID() ([]int64, error) {

	userIDs := []int64{}
	query := fmt.Sprintf(`SELECT
		users_id
	FROM
		p_users_schedules
	WHERE 
		status = (%d);`, PStatusAssistant)
	err := conn.DB.Select(&userIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return userIDs, err
	}

	return userIDs, nil
}

func SelectEnrolledStudentID(scheduleID int64) ([]int64, error) {
	userIDs := []int64{}
	query := fmt.Sprintf(`SELECT
			users_id
		FROM
			p_users_schedules
		WHERE 
			status = (%d) AND
			schedules_id = (%d);`, PStatusStudent, scheduleID)
	err := conn.DB.Select(&userIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return userIDs, err
	}

	return userIDs, nil
}

func SelectAllName() ([]string, error) {

	var names []string
	query := fmt.Sprintf(`SELECT name FROM courses;`)
	err := conn.DB.Select(&names, query)
	if err != nil && err != sql.ErrNoRows {
		return names, err
	}

	return names, nil
}

func IsExist(courseID string) bool {

	var x string
	query := fmt.Sprintf(`SELECT 'x' FROM courses WHERE id = ('%s') LIMIT 1;`, courseID)
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

func InsertSchedule(userID int64, startTime, endTime, year int16, semester, day, status int8, class, courseID, placeID string, tx ...*sqlx.Tx) (int64, error) {

	var id int64
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
		return id, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return id, fmt.Errorf("Cannot get last insert id")
	}

	return id, nil
}

func SelectByPage(limit, offset int, isCount bool, scheduleID []int64) ([]CourseSchedule, int, error) {

	var course []CourseSchedule
	var count int

	if len(scheduleID) < 1 {
		return course, count, nil
	}

	scid := strings.Join(helper.Int64ToStringSlice(scheduleID), ", ")
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
		RIGHT JOIN schedules sc ON cs.id = sc.courses_id
		WHERE sc.id IN (%s)
		LIMIT %d OFFSET %d;`, scid, limit, offset)
	rows, err := conn.DB.Queryx(query)
	defer rows.Close()
	if err != nil {
		return course, count, err
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
			return course, count, err
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

	if !isCount {
		return course, count, nil
	}

	query = fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM
			schedules
		WHERE id IN (%s)`, scid)
	err = conn.DB.Get(&count, query)
	if err != nil {
		return course, count, err
	}

	return course, count, nil
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
			sc.status = (%d)
		ORDER BY day ASC;`, ids, status)

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
		var scID, createdBy int64

		err := rows.Scan(&id, &name, &description, &ucu, &scID, &status, &startTime, &endTime, &day, &class, &semester, &year, &placeID, &createdBy)
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
				ID:        scID,
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

func SelectByName(name string) ([]Course, error) {
	var courses []Course
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			description,
			ucu
		FROM
			courses
		WHERE
			name LIKE ('%%%s%%')
		LIMIT 5;
	`, name)
	err := conn.DB.Select(&courses, query)
	if err != nil && err != sql.ErrNoRows {
		return courses, err
	}

	return courses, nil
}

func InsertGradeParameter(typ string, percentage float32, scheduleID int64, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		INSERT INTO
		grade_parameters (
			type,
			percentage,
			schedules_id,
			created_at,
			updated_at
		)
		VALUES (
			('%s'),
			(%f),
			(%d),
			NOW(),
			NOW()
		);
		`, typ, percentage, scheduleID)

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

func SelectGPBySchedule(scheduleID []int64) ([]GradeParameter, error) {
	var gps []GradeParameter

	if len(scheduleID) < 1 {
		return gps, nil
	}

	querySchID := strings.Join(helper.Int64ToStringSlice(scheduleID), ", ")
	query := fmt.Sprintf(`
		SELECT
			id,
			type,
			percentage,
			schedules_id
		FROM
			grade_parameters
		WHERE
			schedules_id IN (%s);
		`, querySchID)
	err := conn.DB.Select(&gps, query)
	if err != nil && err != sql.ErrNoRows {
		return gps, err
	}
	return gps, nil
}

// SelectGradeParameterByScheduleIDIN func
func SelectGradeParameterByScheduleIDIN(scheduleID []int64) ([]int64, error) {
	var gps []int64
	var gradeQuery []string
	for _, value := range scheduleID {
		gradeQuery = append(gradeQuery, fmt.Sprintf("%d", value))
	}
	queryGradeList := strings.Join(gradeQuery, ",")
	query := fmt.Sprintf(`
		SELECT
			id
		FROM
			grade_parameters
		WHERE
			schedules_id
		IN
			 (%s);
		`, queryGradeList)
	err := conn.DB.Select(&gps, query)
	if err != nil && err != sql.ErrNoRows {
		return gps, err
	}
	return gps, nil
}

func DeleteGradeParameter(id int64, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
			DELETE FROM
				grade_parameters
			WHERE
				id = (%d);
			`, id)

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

func UpdateGradeParameter(typ string, percentage float32, scheduleID int64, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		UPDATE
			grade_parameters
		SET
			percentage = (%f),
			updated_at = NOW()
		WHERE
			type = ('%s') AND
			schedules_id = (%d);
		`, percentage, typ, scheduleID)

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

// GetScheduleIDByGP ...
func GetScheduleIDByGP(gpID int64) (int64, error) {
	var scheduleID int64
	query := fmt.Sprintf(`
		SELECT 
			schedules_id
		FROM
			grade_parameters
		WHERE
			id = (%d)
		`, gpID)
	err := conn.DB.Get(&scheduleID, query)
	if err != nil {
		return scheduleID, err
	}

	return scheduleID, nil
}

// GetGradeParametersID func ...
func GetGradeParametersID(AssignmentID int64) int64 {
	query := fmt.Sprintf(`
		SELECT 
			grade_parameters_id
		FROM
			assignments
		WHERE
			id = (%d)
		`, AssignmentID)
	var assignmentID string
	err := conn.DB.QueryRow(query).Scan(&assignmentID)
	if err != nil {
		return 0
	}
	assignmentid, err := strconv.ParseInt(assignmentID, 10, 64)
	if err != nil {
		return 0
	}

	return assignmentid
}

// GetGradeParametersIDByScheduleID func ...
func GetGradeParametersIDByScheduleID(ScheduleID int64) ([]int64, error) {
	query := fmt.Sprintf(`
		SELECT
			id
		FROM
			grade_parameters
		WHERE
			schedules_id = (%d)
		;`, ScheduleID)

	rows, err := conn.DB.Query(query)
	var gradeParamsID []int64
	if err != nil {
		return gradeParamsID, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return gradeParamsID, err
		}
		gradeParamsID = append(gradeParamsID, id)
	}
	return gradeParamsID, nil
}

// IsUserHasUploadedFile func ...
func IsUserHasUploadedFile(assignmentID, userID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			p_users_assignments
		WHERE
			assignments_id = (%d) AND users_id =(%d)
		LIMIT 1;`, assignmentID, userID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// IsAllUsersEnrolled func ...
func IsAllUsersEnrolled(scheduleID int64, usersID []int64) bool {
	userIDs := []int64{}
	var userList []string
	for _, value := range usersID {
		userList = append(userList, fmt.Sprintf("%d", value))
	}
	queryUserList := strings.Join(userList, ",")
	query := fmt.Sprintf(`SELECT
			users_id
		FROM
			p_users_schedules
		WHERE 
			status = (%d) AND
			schedules_id = (%d) AND users_id IN(%s);`, PStatusStudent, scheduleID, queryUserList)

	err := conn.DB.Select(&userIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return false
	}

	if len(userIDs) != len(usersID) {
		return false
	}
	return true
}

// GetCourseID func ...
func GetCourseID(scheduleID int64) (string, error) {
	query := fmt.Sprintf(`
		SELECT
			courses_id
		FROM
			schedules
		WHERE
			id=(%d)
		LIMIT 1;
		`, scheduleID)
	var res string
	err := conn.DB.Get(&res, query)
	if err != nil {
		return res, err
	}
	return res, nil
}

// GetName func ...
func GetName(courseID string) (string, error) {
	query := fmt.Sprintf(`
		SELECT
			name
		FROM
			courses
		WHERE
			id=('%s')
		LIMIT 1;
		`, courseID)
	var res string
	err := conn.DB.Get(&res, query)
	if err != nil {
		return res, err
	}
	return res, nil
}
func SelectJoinScheduleCourse(scheduleID []int64) ([]CourseConcise, error) {
	var res []CourseConcise
	id := helper.Int64ToStringSlice(scheduleID)
	queryID := strings.Join(id, ", ")
	query := fmt.Sprintf(`
		SELECT
			sc.id,
			cs.name
		FROM
			schedules sc
		INNER JOIN
			courses cs
		ON
			sc.courses_id = cs.id
		WHERE 
			sc.id IN (%s)
		;
		`, queryID)
	err := conn.DB.Select(&res, query)
	if err != nil {
		return res, err
	}
	return res, nil

}

// InsertAssistant ...
func InsertAssistant(usersID []int64, scheduleID int64, tx *sqlx.Tx) error {

	var values []string
	for _, val := range usersID {
		value := fmt.Sprintf("(%d, %d, %d, NOW(), NOW())", val, scheduleID, PStatusAssistant)
		values = append(values, value)
	}

	queryValue := strings.Join(values, ", ")
	query := fmt.Sprintf(`
		INSERT INTO
			p_users_schedules (
				users_id,
				schedules_id,
				status,
				created_at,
				updated_at
			) VALUES %s;
	`, queryValue)

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

func DeleteAssistant(usersID []int64, scheduleID int64, tx *sqlx.Tx) error {

	usersIDString := helper.Int64ToStringSlice(usersID)
	queryUsersID := strings.Join(usersIDString, ", ")
	query := fmt.Sprintf(`
		DELETE FROM
			p_users_schedules
		WHERE
			status = (%d) AND
			schedules_id = (%d) AND
			users_id IN (%s);
	`, PStatusAssistant, scheduleID, queryUsersID)

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

func SelectByDayScheduleID(day int8, schedulesID []int64) ([]CourseSchedule, error) {
	var course []CourseSchedule
	if len(schedulesID) < 1 {
		return course, nil
	}

	querySchedulesID := strings.Join(helper.Int64ToStringSlice(schedulesID), ", ")
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
			sc.day = (%d) AND
			sc.id IN (%s)
		;`, day, querySchedulesID)
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

func InsertUnapproved(userID, scheduleID int64) error {
	query := fmt.Sprintf(`
		INSERT INTO
			p_users_schedules (
				users_id,
				schedules_id,
				status,
				created_at,
				updated_at
			) VALUES (
				(%d),
				(%d),
				(%d),
				NOW(),
				NOW()
			);
	`, userID, scheduleID, PStatusUnapproved)

	_, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserRelation(userID, scheduleID int64) error {

	query := fmt.Sprintf(`
		DELETE FROM
			p_users_schedules
		WHERE
			users_id = (%d) AND
			schedules_id = (%d);
	`, userID, scheduleID)

	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}

	if valid, err := result.RowsAffected(); err != nil || valid < 1 {
		return err
	}

	return nil
}

// InsertInvolved ...
func InsertInvolved(userID, scheduleID int64, role int, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		INSERT INTO
			p_users_schedules (
				users_id,
				schedules_id,
				status,
				created_at,
				updated_at
			) VALUES (
				(%d),
				(%d),
				(%d),
				NOW(),
				NOW()
			);
	`, userID, scheduleID, role)

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

// ActivateStudent ...
func ActivateStudent(userID, scheduleID int64, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		UPDATE
			p_users_schedules
		SET
			status = (%d),
			updated_at = NOW()
		WHERE
			users_id = (%d) AND
			schedules_id = (%d);
	`, PStatusStudent, userID, scheduleID)

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

func SelectUnapproved(scheduleID int64) ([]int64, error) {
	var usersID []int64
	query := fmt.Sprintf(`
		SELECT
			users_id
		FROM
			p_users_schedules
		WHERE
			schedules_id = (%d) AND
			status = (%d);
		`, scheduleID, PStatusStudent)
	err := conn.DB.Select(&usersID, query)
	if err != nil {
		return usersID, err
	}
	return usersID, nil
}

func SelectInvolved(scheduleID int64) ([]int64, error) {
	var usersID []int64
	query := fmt.Sprintf(`
		SELECT
			users_id
		FROM
			p_users_schedules
		WHERE
			schedules_id = (%d)
		`, scheduleID)
	err := conn.DB.Select(&usersID, query)
	if err != nil {
		return usersID, err
	}
	return usersID, nil
}

// SelectIDBySchedule ..
func SelectIDBySchedule(scheduleID int64) ([]int64, error) {
	var ids []int64
	query := fmt.Sprintf(`
		SELECT
			users_id
		FROM
			p_users_schedules
		WHERE
			schedules_id=(%d) AND status = 1
		`, scheduleID)
	err := conn.DB.Select(&ids, query)
	if err != nil {
		return ids, err
	}
	return ids, nil
}
