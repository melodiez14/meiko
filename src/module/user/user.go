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

func GetByEmail(email string, column ...string) (User, error) {
	var user User
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColEmail,
			ColGender,
			ColNote,
			ColStatus,
			ColIdentityCode,
			ColLineID,
			ColPhone,
			ColRoleGroupsID,
		}
	} else {
		for _, val := range column {
			c = append(c, val)
		}
	}
	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(queryGetByEmail, cols, email)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetByIdentityCode(identityCode int64, column ...string) (User, error) {
	var user User
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColEmail,
			ColGender,
			ColNote,
			ColStatus,
			ColIdentityCode,
			ColLineID,
			ColPhone,
			ColRoleGroupsID,
		}
	} else {
		for _, val := range column {
			c = append(c, val)
		}
	}
	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(queryGetByIdentityCode, cols, identityCode)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SignIn(email, password string) (User, error) {
	var user User
	query := fmt.Sprintf(querySignIn, email, password)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SignUp(identityCode int64, name, email, password string) error {
	query := fmt.Sprintf(querySignUp, name, email, password, identityCode)
	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

func IsPhoneExist(identityCode int64, phone string) bool {
	var user User
	query := fmt.Sprintf(`
			SELECT
				phone
			FROM
				 users
			WHERE
				phone = ('%s') AND
				identity_code != (%d)
			LIMIT 1
		`, phone, identityCode)
	err := conn.DB.Get(&user, query)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func IsLineIDExist(identityCode int64, lineID string) bool {
	var user User
	query := fmt.Sprintf(`
			SELECT
				line_id
			FROM
				 users
			WHERE
				line_id = ('%s') AND
				identity_code != (%d)
			LIMIT 1
		`, lineID, identityCode)
	err := conn.DB.Get(&user, query)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func UpdateProfile(identityCode int64, name, note string, phone, lineID sql.NullString, gender int8) error {

	if gender != GenderMale && gender != GenderFemale {
		gender = GenderUndefined
	}

	queryLineID := fmt.Sprintf("NULL")
	if lineID.Valid {
		queryLineID = fmt.Sprintf("('%s')", lineID.String)
	}

	queryPhone := fmt.Sprintf("NULL")
	if phone.Valid {
		queryPhone = fmt.Sprintf("('%s')", phone.String)
	}

	query := fmt.Sprintf(`
		UPDATE
			users
		SET
			name = ('%s'),
			phone = %s,
			line_id = %s,
			note = ('%s'),
			gender = (%d),
			updated_at = NOW()
		WHERE
			identity_code = (%d);
		`, name, queryPhone, queryLineID, note, gender, identityCode)
	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

func Get(column ...string) QueryGet {
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColEmail,
			ColGender,
			ColNote,
			ColStatus,
			ColIdentityCode,
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

func (q QuerySelect) Limit(value uint16) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s LIMIT %d", q.string, value)}
}

func (q QuerySelect) Offset(value uint16) QuerySelect {
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

func GenerateVerification(identity int64) (Verification, error) {

	rand.Seed(time.Now().UTC().UnixNano())
	v := Verification{
		Code:           uint16(rand.Intn(8999) + 1000),
		ExpireDuration: "30 Minutes",
		ExpireDate:     time.Now().Add(30 * time.Minute),
		Attempt:        0,
	}

	query := fmt.Sprintf(generateVerificationQuery, v.Code, identity)
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

func UpdateToVerified(identityCode int64) error {
	query := fmt.Sprintf(queryUpdateToVerified, StatusVerified, identityCode)
	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

func UpdateStatus(identityCode int64, status int8) error {
	query := fmt.Sprintf(queryUpdateStatus, status, identityCode)
	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

func SelectDashboard(id int64, limit, offset uint16) ([]User, error) {
	var user []User
	query := fmt.Sprintf(querySelectDashboard, StatusVerified, StatusActivated, id, limit, offset)
	err := conn.DB.Select(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func ChangePassword(identityCode int64, password, oldPassword string) error {
	query := fmt.Sprintf(`
		UPDATE
			users
		SET
			password = ('%s')
		WHERE
			identity_code = (%d) AND
			password = ('%s');
		`, password, identityCode, oldPassword)
	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

func ForgotNewPassword(email, password string) error {
	query := fmt.Sprintf(queryForgotNewPassword, password, email)
	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}
