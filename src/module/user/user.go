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
		Code:           uint16(rand.Uint32() / uint32(9999)),
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

func IsValidUserLogin(email, password string) bool {
	user := &User{}
	query := fmt.Sprintf(getUserLoginQuery, email, password)
	err := conn.DB.Get(user, query)
	if err != nil {
		return false
	}
	return true
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
