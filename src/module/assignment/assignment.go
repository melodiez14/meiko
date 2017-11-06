package assignment

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/meiko/src/util/conn"
)

func GetByCourseID(courseID int64) ([]Assignment, error) {
	var assignments []Assignment
	query := fmt.Sprintf(queryGetByCourseID, courseID)
	err := conn.DB.Select(&assignments, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignments, nil
}

func GetIncompleteByUserID(userID int64) ([]Assignment, error) {
	var assignments []Assignment
	query := fmt.Sprintf(queryGetIncompleteByUserID, userID, userID)
	err := conn.DB.Select(&assignments, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignments, nil
}

func GetCompleteByUserID(userID int64) ([]int64, error) {
	var assignmentsID []int64
	query := fmt.Sprintf(queryGetCompleteByUserID, userID)
	err := conn.DB.Select(&assignmentsID, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignmentsID, nil
}

func IsExistByGradeParameterID(gpID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			id
		FROM
			grade_parameters
		WHERE
			id = (%d)
		LIMIT 1;
	`, gpID)
	err := conn.DB.Get(&x, query)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}
func Insert(GradeParameters int64, Status, Description, DueDate string, tx ...*sqlx.Tx) error {
	query := fmt.Sprintf(`
		INSERT INTO
			assignments(
				status,
				due_date,
				grade_parameters_id,
				description,
				created_at,
				updated_at
			)
		VALUES(
			('%s'),
			('%s'),
			('%d'),
			('%s'),
			NOW(),
			NOW()
		);
		`, Status, DueDate, GradeParameters, Description)

	var result sql.Result
	var err error
	if len(tx) == 1 {
		result, err = tx[0].Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	fmt.Println(result)
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}

	return nil
}
func IsFileIDExist(ID string) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			id
		FROM
			files
		WHERE
			id = ('%s')
		LIMIT 1;
			`, ID)
	err := conn.DB.Get(&x, query)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}
