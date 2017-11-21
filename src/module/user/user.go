package user

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	as "github.com/melodiez14/meiko/src/module/assignment"
	"github.com/melodiez14/meiko/src/util/helper"

	"github.com/melodiez14/meiko/src/util/conn"
)

// SelectByID function for check in database to get user information from database using user ID
/*
	@params
		id				= []int64
	@example
		id				= [140810140060,140810140061]
	@return:
		ID				= 140810140060
		Name			= kharil azmi ashari
		Email			= khairil_azmi_ashari@yahoo.com
		Gender			= 1
		Note			= just doing nothing
		Status			= 1
		IdentityCode	= 140810140060
		LineID			= khaazas
		Phone			= 082214467300
		RoleGroupsID	= 0
*/
func SelectByID(id []int64, isSort bool, column ...string) ([]User, error) {
	var user []User
	var c []string
	var sortQuery string

	d := helper.Int64ToStringSlice(id)

	if isSort {
		sortQuery = "ORDER BY identity_code ASC"
	}

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
	ids := strings.Join(d, ", ")
	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(`
		SELECT
			%s
		FROM
			users
		WHERE
			id IN (%s) %s;`, cols, ids, sortQuery)
	err := conn.DB.Select(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail function for check in database to get user information from database using user email
/*
	@params
		email			= string
	@example
		email			= khairil_azmi_ashari@yahoo.com
	@return:
		ID				= 140810140060
		Name			= kharil azmi ashari
		Email			= khairil_azmi_ashari@yahoo.com
		Gender			= 1
		Note			= just doing nothing
		Status			= 1
		IdentityCode	= 140810140060
		LineID			= khaazas
		Phone			= 082214467300
		RoleGroupsID	= 0
*/
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

// GetByIdentityCode function for check in database to get user information from database using user idetity code
/*
	@params
	@example
	@return:
		ID				= 140810140060
		Name			= kharil azmi ashari
		Email			= khairil_azmi_ashari@yahoo.com
		Gender			= 1
		Note			= just doing nothing
		Status			= 1
		IdentityCode	= 140810140060
		LineID			= khaazas
		Phone			= 082214467300
		RoleGroupsID	= 0
*/
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

// SignIn function to sign using email and password then check in database is this valid account
/*
	@params:
		email			= required, string
		password		= required, string
	@example
		email			= khairil_azmi_ashari@yahoo.com
		password		= Khairil14001
	@return
		ID				= 140810140060
		Name			= kharil azmi ashari
		Email			= khairil_azmi_ashari@yahoo.com
		Gender			= 1
		Note			= just doing nothing
		Status			= 1
		IdentityCode	= 140810140060
		LineID			= khaazas
		Phone			= 082214467300
		RoleGroupsID	= 0
*/
func SignIn(email, password string) (User, error) {
	var user User
	query := fmt.Sprintf(querySignIn, email, password)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

// SignUp function to sign up an account by fill with identity code, name, email, and password.
//then check to database and save the information in database if there is valid new account
/*
	@params:
		identityCode	= required, int64
		name			= required, string
		email			= required, string
		password		= required, string
	@example
		identityCode	= 140810140060
		name			= khairil azmi ashari
		email			= khairil_azmi_ashari@yahoo.com
		password		= Khairil14001
	@return
*/
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

// IsPhoneExist check is phone number exist in database
// SignIn function to sign using email and password then check in database is this valid account
/*
	@params:
		identityCode	= required, int64
		phone			= required, string
	@example
		identityCode	= 140810140060
		phone			= 08214467300
	@return
*/
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

// IsLineIDExist check is Line ID exist in database
// SignIn function to sign using email and password then check in database is this valid account
/*
	@params:
		identityCode	= required, int64
		lineID			= required, string
	@example
		identityCode	= 140810140060
		lineID			= khaazass
	@return
*/
func IsLineIDExist(identityCode int64, lineID string) bool {
	var x string
	query := fmt.Sprintf(`
			SELECT
				'x'
			FROM
				 users
			WHERE
				line_id = ('%s') AND
				identity_code != (%d)
			LIMIT 1;
		`, lineID, identityCode)
	err := conn.DB.Get(&x, query)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

// UpdateProfile function to update user profile information in database
/*
	@params:
		identityCode	= int64
		name			= string
		note			= string
		phone			= string
		lineID			= string
		gender			= int8
	@example:
		identityCode	= 140810140060
		name			= kharil azmi ashari
		note			= just doing nothing
		phone			= 140810140060
		lineID			= khaazas
		gender			= 1
	@return
*/
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

// GenerateVerification ganarate verifocation code to get confirmation from valid email
/*
	@params:
		identity	= int64
	@example:
		identity	= 140810140060
	@return
*/
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

// IsValidConfirmationCode function to check is confirmation code is valid code used
// form right user email account
/*
	@params:
		email	= string
		code	= uint16
	@example:
		email	= khairil_azmi_ashari@yahoo.com
		code	= 140810
	@return
*/
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

// UpdateToVerified function to change account status to be verified after do email verification
/*
	@params:
		identityCode	= string
	@example:
		identityCode	= 140810140060
	@return
*/
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

// UpdateStatus function to update status and send update to database
/*
	@params:
		identityCode	= string
		status			= string
	@example:
		identityCode	= 140810140060
		status			= i'm single, thanks you
	@return
*/
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

// SelectDashboard function to show dashboard and send user information
/*
	@params:
		id		= string
		limit	= uint16
		offset	= uint16
	@example:
		id		= 140810140060
		limit	= 1
		offset	= 1
	@return:
		ID				= 140810140060
		Name			= kharil azmi ashari
		Email			= khairil_azmi_ashari@yahoo.com
		Gender			= 1
		Note			= just doing nothing
		Status			= 1
		IdentityCode	= 140810140060
		LineID			= khaazas
		Phone			= 082214467300
		RoleGroupsID	= 0
*/
func SelectDashboard(id int64, limit, offset uint16) ([]User, error) {
	var user []User
	query := fmt.Sprintf(querySelectDashboard, StatusVerified, StatusActivated, id, limit, offset)
	err := conn.DB.Select(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

// ChangePassword function to change user password account
/*
	@params:
		idnetityCode	= int64
		password		= string
		oldpassword		= string
	@example:
		idnetityCode	= 140810140060
		password		= Old001
		oldpassword		= New001
	@return
*/
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

// ForgotNewPassword function to renew password when user forget their own password
// and password will sent to verified email
/*
	@params:
		email		= string
		paswword	= string
	@example:
		email		= khairil_azmi_ashari@yahoo.com
		paswword	= Khairil14001
	@return
*/
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

// Update function to update user information from inputed data
/*
	@params:
		identityCode	= int64
		name			= string
		note			= string
		lineID			= sql.String
		gender			= int8
		status			= int8
	@example:
		identityCode	= 140810140060
		name			= kharil azmi ashari
		note			= just doing nothing
		lineID			= khaazas
		gender			= 1
		status			= 1
	@return
*/
func Update(identityCode int64, name, note string, phone, lineID sql.NullString, gender, status int8) error {

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
				status = (%d),
				updated_at = NOW()
			WHERE
				identity_code = (%d);
			`, name, queryPhone, queryLineID, note, gender, status, identityCode)
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

// Delete function to delete user using identity code
/*
	@params:
		identityCode	= int64
	@example:
		identityCode	= 140810140060
	@return
*/
func Delete(identityCode int64) error {
	query := fmt.Sprintf(`
		DELETE FROM
			users
		WHERE
			identity_code = (%d);
		`, identityCode)

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

// Create function to create user to database from valid singup process
/*
	@params:
		identityCode	= int64
		name			= string
		email			= khairil_azmi_ashari@yahoo.com
	@example:
		identityCode	= 140810140060
		name			= kharil azmi ashari
		email			= khairil_azmi_ashari@yahoo.com
	@return
*/
func Create(identityCode int64, name, email string) error {
	query := fmt.Sprintf(`
		INSERT INTO
		users (
			name,
			email,
			password,
			identity_code,
			status,
			created_at,
			updated_at
		) VALUES (
			('%s'),
			('%s'),
			('x'),
			(%d),
			(%d),
			NOW(),
			NOW()
		);
		`, name, email, identityCode, StatusActivated)
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

// IsUserExist func ...
func IsUserExist(identityCode int64) bool {

	var x string
	query := fmt.Sprintf(`
			SELECT
				'x'
			FROM
				users
			WHERE
				identity_code = (%d)
			LIMIT 1;`, identityCode)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// IsUserTakeSchedule func ...
func IsUserTakeSchedule(UserID, ScheduleID int64) bool {
	var x string
	query := fmt.Sprintf(`
			SELECT
				'x'
			FROM
				p_users_schedules pus
			WHERE
				pus.users_id = (%d) AND schedules_id = (%d)
			LIMIT 1;`, UserID, ScheduleID)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// SelectUserByScheduleID func ...
func SelectUserByScheduleID(scheduleID int64) ([]as.UserAssignmentDetail, error) {
	var user []as.UserAssignmentDetail
	query := fmt.Sprintf(`
			SELECT
				usr.identity_code,
				usr.name
			FROM
				p_users_schedules pus
			INNER JOIN
				users usr
			ON
				usr.id=pus.users_id
			WHERE
				pus.schedules_id=(%d);
		`, scheduleID)

	err := conn.DB.Select(&user, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

// SelectIDByIdentityCode ...
func SelectIDByIdentityCode(identityCode []int64) ([]int64, error) {
	var ids []int64
	queryIdentity := strings.Join(helper.Int64ToStringSlice(identityCode), ", ")
	query := fmt.Sprintf(`
		SELECT
			id
		FROM
			users
		WHERE
			identity_code IN (%s)
		;`, queryIdentity)
	err := conn.DB.Select(&ids, query)
	if err != nil {
		return ids, err
	}
	return ids, nil

}

// IsExistRolegroupID ...
func IsExistRolegroupID(rolegroupID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			users
		WHERE
			rolegroups_id = (%d)
		LIMIT 1;
	`, rolegroupID)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}

	return true
}

// SelectDistinctRolegroupID ...
func SelectDistinctRolegroupID() ([]int64, error) {
	var rolegroupsID []int64
	query := fmt.Sprintf(`
		SELECT
			DISTINCT rolegroups_id
		FROM
			users
		WHERE rolegroups_id IS NOT NULL;
	`)
	err := conn.DB.Select(&rolegroupsID, query)
	if err != nil {
		return rolegroupsID, err
	}

	return rolegroupsID, nil
}
