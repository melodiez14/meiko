package course

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"

	"database/sql"

	"github.com/melodiez14/meiko/src/util/conn"
)

// Get Query language to do get from course in database
func Get(column ...string) QueryGet {
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColDescription,
			ColUCU,
			ColSemester,
			ColStatus,
			ColStartTime,
			ColEndTime,
			ColClass,
			ColDay,
			ColPlaceID,
		}
	} else {
		for _, v := range column {
			c = append(c, v)
		}
	}
	columnQuery := strings.Join(c, ", ")
	return QueryGet{fmt.Sprintf(queryGet, columnQuery)}
}

// Where Query language to do get and using where from course in database
func (q QueryGet) Where(column, operator string, value interface{}) QueryGet {
	switch value.(type) {
	case int, int8, int64:
		return QueryGet{fmt.Sprintf("%s WHERE %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryGet{fmt.Sprintf("%s WHERE %s %s ('%s')", q.string, column, operator, value)}
	case []int64:
		var vals []string
		rv := reflect.ValueOf(value).Interface().([]int64)
		if len(rv) < 1 {
			return q
		}
		for _, v := range rv {
			vals = append(vals, fmt.Sprintf("%d", v))
		}
		str := strings.Join(vals, ", ")
		return QueryGet{fmt.Sprintf("%s OR %s %s (%s)", q.string, column, operator, str)}
	default:
		return q
	}
}

// AndWhere QueryGet function to use (and where) from course in database
func (q QueryGet) AndWhere(column, operator string, value interface{}) QueryGet {
	switch value.(type) {
	case int, int8, int64:
		return QueryGet{fmt.Sprintf("%s AND %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryGet{fmt.Sprintf("%s AND %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

// OrWhere QueryGet function to use (or where)  from course in database
func (q QueryGet) OrWhere(column, operator string, value interface{}) QueryGet {
	switch value.(type) {
	case int, int8, int64:
		return QueryGet{fmt.Sprintf("%s OR %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryGet{fmt.Sprintf("%s OR %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

// Exec QueryGet Function to execute query select from course in database 
func (q QueryGet) Exec() (Course, error) {
	var course Course
	query := fmt.Sprintf("%s LIMIT 1", q.string)
	err := conn.DB.Get(&course, query)
	if err != nil {
		return course, err
	}
	return course, nil
}


// Select Query language to do select from course in database
func Select(column ...string) QuerySelect {
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColDescription,
			ColUCU,
			ColSemester,
			ColStatus,
			ColStartTime,
			ColEndTime,
			ColClass,
			ColDay,
			ColPlaceID,
		}
	} else {
		for _, v := range column {
			c = append(c, v)
		}
	}
	columnQuery := strings.Join(c, ", ")
	return QuerySelect{fmt.Sprintf(querySelect, columnQuery)}
}

// Where QuerySelect function to use (where) from course in database
func (q QuerySelect) Where(column, operator string, value interface{}) QuerySelect {
	switch value.(type) {
	case int, int8, int64:
		return QuerySelect{fmt.Sprintf("%s WHERE %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QuerySelect{fmt.Sprintf("%s WHERE %s %s ('%s')", q.string, column, operator, value)}
	case []int64:
		var vals []string
		rv := reflect.ValueOf(value).Interface().([]int64)
		for _, v := range rv {
			vals = append(vals, fmt.Sprintf("%d", v))
		}
		str := strings.Join(vals, ", ")
		return QuerySelect{fmt.Sprintf("%s WHERE %s %s (%s)", q.string, column, operator, str)}
	default:
		return q
	}
}

// AndWhere QuerySelect function to use (and where) from course in database
func (q QuerySelect) AndWhere(column, operator string, value interface{}) QuerySelect {
	switch value.(type) {
	case int, int8, int64:
		return QuerySelect{fmt.Sprintf("%s AND %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QuerySelect{fmt.Sprintf("%s AND %s %s ('%s')", q.string, column, operator, value)}
	case []int64:
		var vals []string
		rv := reflect.ValueOf(value).Interface().([]int64)
		if len(rv) < 1 {
			return q
		}
		for _, v := range rv {
			vals = append(vals, fmt.Sprintf("%d", v))
		}
		str := strings.Join(vals, ", ")
		return QuerySelect{fmt.Sprintf("%s AND %s %s (%s)", q.string, column, operator, str)}
	default:
		return q
	}
}

// OrWhere QuerySelect function to use (or where)  from course in database
func (q QuerySelect) OrWhere(column, operator string, value interface{}) QuerySelect {
	switch value.(type) {
	case int, int8, int64:
		return QuerySelect{fmt.Sprintf("%s OR %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QuerySelect{fmt.Sprintf("%s OR %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

// Limit Select Query language to do select using limit from course in database
func (q QuerySelect) Limit(value uint16) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s LIMIT %d", q.string, value)}
}

// Offset Select Query language to do select using offset from course in database
func (q QuerySelect) Offset(value uint16) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s OFFSET %d", q.string, value)}
}

// Exec QuerySelect Function to execute query select from course in database 
func (q QuerySelect) Exec() ([]Course, error) {
	var courses []Course
	err := conn.DB.Select(&courses, q.string)
	if err != nil {
		return courses, err
	}
	return courses, nil
}

// Insert Query language to do insert from course in database
func Insert(column map[string]interface{}) QueryInsert {

	c := []string{"created_at", "updated_at"}
	v := []string{"NOW()", "NOW()"}
	for i, val := range column {
		switch val.(type) {
		case int, int8, int16, int64:
			c = append(c, i)
			v = append(v, fmt.Sprintf("(%d)", val))
		case string:
			c = append(c, i)
			v = append(v, fmt.Sprintf("('%s')", val))
		}
	}
	columnQuery := strings.Join(c, ", ")
	valueQuery := strings.Join(v, ", ")
	return QueryInsert{fmt.Sprintf(queryInsert, columnQuery, valueQuery)}
}

// Exec QueryInsert Function to execute query insert from course in database
func (q QueryInsert) Exec(txs ...*sqlx.Tx) error {
	// if exist tx
	if len(txs) == 1 {
		result, err := txs[0].Exec(q.string)
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if rows == 0 {
			return fmt.Errorf("No rows affected")
		}
		return nil
	}

	result, err := conn.DB.Exec(q.string)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

// Update Query language to do update from course in database
func Update(column map[string]interface{}) QueryUpdate {
	c := []string{"updated_at = NOW()"}
	for i, val := range column {
		switch val.(type) {
		case int, int16, int8, int64:
			c = append(c, fmt.Sprintf("%s = (%d)", i, val))
		case string:
			c = append(c, fmt.Sprintf("%s = ('%s')", i, val))
		case sql.NullString:
			str := reflect.ValueOf(val).Interface().(sql.NullString)
			if str.Valid {
				c = append(c, fmt.Sprintf("%s = ('%s')", i, str.String))
			} else {
				c = append(c, fmt.Sprintf("%s = NULL", i))
			}
		}
	}
	columnQuery := strings.Join(c, ", ")
	return QueryUpdate{fmt.Sprintf(queryUpdate, columnQuery)}
}

// Where QueryUpdate function to use (where) from course in database
func (q QueryUpdate) Where(column, operator string, value interface{}) QueryUpdate {
	switch value.(type) {
	case int, int8, int64:
		return QueryUpdate{fmt.Sprintf("%s WHERE %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryUpdate{fmt.Sprintf("%s WHERE %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

// AndWhere QueryUpdate function to use (and where) from course in database
func (q QueryUpdate) AndWhere(column, operator string, value interface{}) QueryUpdate {
	switch value.(type) {
	case int, int8, int64:
		return QueryUpdate{fmt.Sprintf("%s AND %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryUpdate{fmt.Sprintf("%s AND %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

// OrWhere QueryUpdate function to use (or where)  from course in database
func (q QueryUpdate) OrWhere(column, operator string, value interface{}) QueryUpdate {
	switch value.(type) {
	case int, int8, int64:
		return QueryUpdate{fmt.Sprintf("%s OR %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryUpdate{fmt.Sprintf("%s OR %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

// Exec QueryUpdate Function to execute query update from course in database
func (q QueryUpdate) Exec(txs ...*sqlx.Tx) error {
	// if exist tx
	if len(txs) == 1 {
		result, err := txs[0].Exec(q.string)
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if rows == 0 {
			return fmt.Errorf("No rows affected")
		}
		return nil
	}
	result, err := conn.DB.Exec(q.string)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

// GetByUserID Query function to get course that take user using user ID
/*
	@params:
		userID		= int64 
	@example:
		userID		= 4
	@return:
		courseID	= []All course take by user
*/
func GetByUserID(userID int64) ([]Course, error) {
	var courses []Course
	query := fmt.Sprintf(queryGetCourseByUserID, userID)
	err := conn.DB.Select(&courses, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return courses, nil
}

// SelectIDByUserID Query function to show course that take user using user ID
/*
	@params:
		userID		= int64 
		status		= bool
	@example:
		userID		= 4
		status	= yes (enrolled)
	@return:
		courseID	= 1
*/
func SelectIDByUserID(userID int64, status ...int8) ([]int64, error) {

	var st string
	if len(status) == 1 {
		st = fmt.Sprintf("AND status = (%d)", status[0])
	}

	courseIDs := []int64{}
	query := fmt.Sprintf(`SELECT courses_id FROM p_users_courses WHERE users_id = (%d) %s`, userID, st)
	err := conn.DB.Select(&courseIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return courseIDs, err
	}

	return courseIDs, nil
}

// IsEnrolled Query function check is course has enrolled or not
/*
	@params:
		userID		= int64 
		courseID	= int64
	@example:
		userID		= 4
		courseID	= 12
	@return:
		enrolled_status	= true (enrolled) false (not enrolled)
*/
func IsEnrolled(userID, courseID int64) bool {
	var v int64
	query := fmt.Sprintf("SELECT users_id FROM p_users_courses WHERE users_id = (%d) AND courses_id = (%d) LIMIT 1", userID, courseID)
	err := conn.DB.Get(&v, query)
	if err != nil {
		return false
	}
	return true
}

// SelectAssistantID Query to select Assisten by its iD
/*
	@params:
		courseID	= int64
	@example:
		courseID	= 1
	@return:
		userIDs	= userID{}
*/
func SelectAssistantID(courseID int64) ([]int64, error) {

	userIDs := []int64{}
	query := fmt.Sprintf(`SELECT users_id FROM p_users_courses WHERE courses_id = (%d) AND status = (%d)`, courseID, PStatusAssistant)
	err := conn.DB.Select(&userIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return userIDs, err
	}

	return userIDs, nil
}
