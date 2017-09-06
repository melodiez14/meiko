package user

import (
	"fmt"
	"math/rand"
	"time"

	"log"

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
	log.Println(query)
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
func UpdateCodeUser(email string, status int) {
	query := fmt.Sprintf(setStatusUserQuery, email, status)
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
