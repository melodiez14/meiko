package user

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"
)

type (
	QueryGet    struct{ string }
	QuerySelect struct{ string }
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

func Update(column map[string]interface{}) QueryUpdate {
	c := []string{"updated_at = NOW()"}
	for i, v := range column {
		switch v.(type) {
		case int, int8, int64:
			c = append(c, fmt.Sprintf("%s = (%d)", i, v))
		case string:
			c = append(c, fmt.Sprintf("%s = ('%s')", i, v))
		case nil:
			c = append(c, fmt.Sprintf("%s = NULL", i))
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

func GetUserByID(id int64) (User, error) {
	var user User
	query := fmt.Sprintf(getUserByIDQuery, id)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	query := fmt.Sprintf(getUserEmailQuery, email)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SetUpdateUserAccount(name, phone, lineID, College, note string, gender int8, id int64) {
	query := fmt.Sprintf(updateUserAccountQuery, name, gender, phone, lineID, College, note, id)
	_ = conn.DB.MustExec(query)
}

func SetChangePassword(password string, id int64) {
	query := fmt.Sprintf(setChangePasswordQuery, password, id)
	_ = conn.DB.MustExec(query)
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

func SetNewPassword(email, password string) {
	query := fmt.Sprintf(setNewPasswordQuery, password, email)
	_ = conn.DB.MustExec(query)
}

func SetStatus(email string, status int8) {
	query := fmt.Sprintf(setStatusUserQuery, status, email)
	_ = conn.DB.MustExec(query)
}

func GetUserLogin(email, password string) (User, error) {
	var user User
	query := fmt.Sprintf(getUserLoginQuery, email, password)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertNewUser(id int64, name, email, password string) {
	query := fmt.Sprintf(insertNewUserQuery, id, name, email, password)
	_ = conn.DB.MustExec(query)
}

func GetByStatus(status int8) ([]User, error) {
	users := []User{}
	query := fmt.Sprintf(getUserByStatusQuery, status)
	err := conn.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetByIDStatus(id int64, status int8) (User, error) {
	var user User
	query := fmt.Sprintf(getUserByIDStatusQuery, id, status)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}

	return user, nil
}
