package course

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"

	"database/sql"

	"github.com/melodiez14/meiko/src/util/conn"
)

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

func (q QueryGet) Exec() (Course, error) {
	var course Course
	query := fmt.Sprintf("%s LIMIT 1", q.string)
	err := conn.DB.Get(&course, query)
	if err != nil {
		return course, err
	}
	return course, nil
}

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
		return QuerySelect{fmt.Sprintf("%s OR %s %s (%s)", q.string, column, operator, str)}
	default:
		return q
	}
}

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

func (q QuerySelect) Limit(value uint16) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s LIMIT %d", q.string, value)}
}

func (q QuerySelect) Offset(value uint16) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s OFFSET %d", q.string, value)}
}

func (q QuerySelect) Exec() ([]Course, error) {
	var courses []Course
	err := conn.DB.Select(&courses, q.string)
	if err != nil {
		return courses, err
	}
	return courses, nil
}

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

func Update(column map[string]interface{}) QueryUpdate {
	c := []string{"updated_at = NOW()"}
	for i, val := range column {
		switch val.(type) {
		case int, int8, int64:
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

func (q QueryUpdate) Exec() error {
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

func GetByUserID(userID int64) ([]Course, error) {
	var courses []Course
	query := fmt.Sprintf(queryGetCourseByUserID, userID)
	err := conn.DB.Select(&courses, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return courses, nil
}

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

func IsEnrolled(userID, courseID int64) bool {
	var v int64
	query := fmt.Sprintf("SELECT users_id FROM p_users_courses WHERE users_id = (%d) AND courses_id = (%d) LIMIT 1", userID, courseID)
	err := conn.DB.Get(&v, query)
	if err != nil {
		return false
	}
	return true
}

func SelectAssistantID(courseID int64) ([]int64, error) {

	userIDs := []int64{}
	query := fmt.Sprintf(`SELECT users_id FROM p_users_courses WHERE courses_id = (%d) AND status = (%d)`, courseID, PStatusAssistant)
	err := conn.DB.Select(&userIDs, query)
	if err != nil && err != sql.ErrNoRows {
		return userIDs, err
	}

	return userIDs, nil
}
