package user

import (
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetUserByID(id int64) (*User, error) {
	user := &User{}
	query := fmt.Sprintf(getUserByIDQuery, id)
	err := conn.DB.Get(&user, query)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func IsValidUserLogin(email, password string) error {
	user := &User{}

	query := fmt.Sprintf(getUserByEmailQuery, email, password)
	err := conn.DB.Get(user, query)
	if err != nil {
		return err
	}
	return nil
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
