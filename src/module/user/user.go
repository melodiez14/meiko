package user

import (
	"database/sql"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"
)

type (
	QueryGet    struct{ string }
	QuerySelect struct{ string }
	QueryInsert struct{ string }
	QueryUpdate struct{ string }
)

func Get(column ...string) QueryGet {
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColEmail,
			ColGender,
			ColCollege,
			ColNote,
			ColStatus,
			ColLineID,
			ColPhone,
			ColRoleGroupsID,
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

func (q QueryGet) Exec() (User, error) {
	var user User
	query := fmt.Sprintf("%s LIMIT 1", q.string)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func Select(column ...string) QuerySelect {
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColEmail,
			ColGender,
			ColCollege,
			ColNote,
			ColStatus,
			ColLineID,
			ColPhone,
			ColRoleGroupsID,
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

func (q QuerySelect) Limit(value int64) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s LIMIT %d", q.string, value)}
}

func (q QuerySelect) Offset(value int64) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s OFFSET %d", q.string, value)}
}

func (q QuerySelect) Exec() ([]User, error) {
	var user []User
	err := conn.DB.Select(&user, q.string)
	if err != nil {
		return user, err
	}
	return user, nil
}

func Insert(column map[string]interface{}) QueryInsert {

	c := []string{"created_at", "updated_at"}
	v := []string{"NOW()", "NOW()"}
	for i, val := range column {
		switch val.(type) {
		case int, int8, int64:
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

func (q QueryInsert) Exec() error {
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
	print(q.string)
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

func GenerateVerification(id int64) (Verification, error) {

	v := Verification{
		Code:           uint16(rand.Intn(8999) + 1000),
		ExpireDuration: "30 Minutes",
		ExpireDate:     time.Now().Add(30 * time.Minute),
		Attempt:        0,
	}

	query := fmt.Sprintf(generateVerificationQuery, v.Code, id)
	result := conn.DB.MustExec(query)
	count, _ := result.RowsAffected()
	if count < 1 {
		return v, fmt.Errorf("Error executing query")
	}

	return v, nil
}

func IsValidConfirmationCode(email string, code uint16) bool {
	var c Confirmation
	query := fmt.Sprintf(getConfirmationQuery, email)
	err := conn.DB.Get(&c, query)
	if err != nil {
		return false
	}

	if !c.Attempt.Valid || c.Attempt.Int64 >= 3 {
		return false
	}

	if !c.Code.Valid || c.Code.Int64 != int64(code) {
		query = fmt.Sprintf(attemptIncrementQuery, c.ID)
		_ = conn.DB.MustExec(query)
		return false
	}

	return true
}

func UpdateNewPassword(email, password string) {
	query := fmt.Sprintf(updateNewPasswordQuery, password, email)
	_ = conn.DB.MustExec(query)
}
