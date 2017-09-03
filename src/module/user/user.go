package user

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetUserByID(id int64) (*User, error) {
	user := &User{}
	query := fmt.Sprintf(getUserByIDQuery, id)
	err := conn.DB.Get(user, query)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := fmt.Sprintf(getUserEmailQuery, email)
	err := conn.DB.Get(user, query)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GenerateVerification(id int64) (*Verification, error) {

	v := &Verification{
		Code:           uint16(rand.Intn(8999) + 1000),
		ExpireDuration: "30 Minutes",
		ExpireDate:     time.Now().Add(30 * time.Minute),
		Attempt:        0,
	}

	query := fmt.Sprintf(generateVerificationQuery, v.Code, id)
	result := conn.DB.MustExec(query)
	count, _ := result.RowsAffected()
	if count < 1 {
		return nil, fmt.Errorf("Error executing query")
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

func GetUserLogin(email, password string) (*User, error) {
	user := &User{}
	query := fmt.Sprintf(getUserLoginQuery, email, password)
	err := conn.DB.Get(user, query)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func InsertUser(name, email, password, gender, college, note string, rolegroupID int64, status bool) (*User, error) {

// 	u := &User{
// 		Name:        name,
// 		Email:       email,
// 		Password:    password,
// 		Gender:      gender,
// 		College:     college,
// 		Note:        note,
// 		RolegroupID: rolegroupID,
// 		Status:      status,
// 	}
// 	_ = u
// 	return nil, nil
// }
