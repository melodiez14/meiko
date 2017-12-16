package assignment

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/melodiez14/meiko/src/util/helper"

	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/meiko/src/util/conn"
)

// GetByCourseID func ...
func GetByCourseID(courseID int64) ([]Assignment, error) {
	var assignments []Assignment
	query := fmt.Sprintf(queryGetByCourseID, courseID)
	err := conn.DB.Select(&assignments, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignments, nil
}

// GetIncompleteByUserID func ...
func GetIncompleteByUserID(userID int64) ([]Assignment, error) {
	var assignments []Assignment
	query := fmt.Sprintf(queryGetIncompleteByUserID, userID, userID)
	err := conn.DB.Select(&assignments, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignments, nil
}

// GetCompleteByUserID func ...
func GetCompleteByUserID(userID int64) ([]int64, error) {
	var assignmentsID []int64
	query := fmt.Sprintf(queryGetCompleteByUserID, userID)
	err := conn.DB.Select(&assignmentsID, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignmentsID, nil
}

// IsExistByGradeParameterID func ...
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

// Insert function is ...
func Insert(name string, desc sql.NullString, gps, maxSize int64, duedate time.Time, status int8, tx ...*sqlx.Tx) (int64, error) {

	var id int64
	queryDesc := fmt.Sprintf("(NULL)")
	if desc.Valid {
		queryDesc = fmt.Sprintf("('%s')", desc.String)
	}

	query := fmt.Sprintf(`
		INSERT INTO
			assignments(
				name,
				description,
				status,
				due_date,
				grade_parameters_id,
				max_size,
				created_at,
				updated_at
			)
		VALUES(
			('%s'),
			%s,
			('%d'),
			('%s'),
			(%d),
			(%d),
			NOW(),
			NOW()
		);
		`, name, queryDesc, status, duedate, gps, maxSize)

	var result sql.Result
	var err error
	if len(tx) == 1 {
		result, err = tx[0].Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return id, err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return id, fmt.Errorf("No rows affected")
	}
	id, err = result.LastInsertId()
	if err != nil {
		return id, fmt.Errorf("Error getting inserted ID")
	}
	return id, nil
}

// Update func ...
func Update(gradeParameters, id int64, name, status, dueDate string, description sql.NullString, tx *sqlx.Tx) error {

	var result sql.Result
	var err error
	queryDescription := fmt.Sprintf("NULL")
	if description.Valid {
		queryDescription = fmt.Sprintf(description.String)
	}
	query := fmt.Sprintf(`
		UPDATE 
			assignments
		SET
			name = ('%s'),
			status = ('%s'),
			due_date = ('%s'),
			grade_parameters_id = (%d),
			description = ('%s'),
			updated_at = NOW()
		WHERE
			id = (%d);
		`, name, status, dueDate, gradeParameters, queryDescription, id)
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

// IsFileIDExist func ...
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

// SelectByPage func ...
func SelectByPage(gpID []int64, limit, offset int, isCount bool) ([]Assignment, int, error) {
	var assignment []Assignment
	var count int
	queryGP := strings.Join(helper.Int64ToStringSlice(gpID), ", ")
	query := fmt.Sprintf(`
		SELECT
			id,
			grade_parameters_id,
			name,
			description,
			status,
			due_date,
			created_at,
			updated_at
		FROM
			assignments
		WHERE
			grade_parameters_id IN (%s)
		LIMIT %d
		OFFSET %d;`, queryGP, limit, offset)

	err := conn.DB.Select(&assignment, query)
	if err != nil {
		return assignment, count, err
	}

	if !isCount {
		return assignment, count, nil
	}

	query = fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM
			assignments
		WHERE
			grade_parameters_id IN (%s)
		LIMIT %d
		OFFSET %d;`, queryGP, limit, offset)

	err = conn.DB.Get(&count, query)
	if err != nil {
		return assignment, count, err
	}

	return assignment, count, nil
}

// GetByGradeParametersID func ...
func GetByGradeParametersID(gradeParametersID []int64, limit, offset uint16) ([]Assignment, error) {
	var gradeParametersQuery string
	for i, value := range gradeParametersID {
		if i+1 == len(gradeParametersID) {
			gradeParametersQuery = fmt.Sprintf("%s%d", gradeParametersQuery, value)
		} else {
			gradeParametersQuery = fmt.Sprintf("%s%d, ", gradeParametersQuery, value)
		}
	}
	query := fmt.Sprintf(`
		SELECT
			id,
			due_date,
			name,
			description,
			status
		FROM
			assignments
		WHERE
			grade_parameters_id IN (%s)
		LIMIT %d OFFSET %d		
		;`, gradeParametersQuery, limit, offset)

	var result []Assignment
	err := conn.DB.Select(&result, query)
	if err != nil {
		return result, err
	}
	return result, nil

}

// SelectByGradeParametersID func ...
func SelectByGradeParametersID(gradeParametersID []int64) []Assignment {
	var gradeParametersQuery string
	for i, value := range gradeParametersID {
		if i+1 == len(gradeParametersID) {
			gradeParametersQuery = fmt.Sprintf("%s%d", gradeParametersQuery, value)
		} else {
			gradeParametersQuery = fmt.Sprintf("%s%d, ", gradeParametersQuery, value)
		}
	}
	query := fmt.Sprintf(`
		SELECT
			id,
			due_date,
			name,
			description,
			status
		FROM
			assignments
		WHERE
			grade_parameters_id IN (%s)	
		;`, gradeParametersQuery)

	var result []Assignment
	err := conn.DB.Select(&result, query)
	if err != nil {
		return result
	}
	return result

}

// GetByID ...
func GetByID(id int64) (Assignment, error) {
	var assignment Assignment
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			status,
			description,
			grade_parameters_id,
			due_date,
			created_at,
			updated_at
		FROM
			assignments
		WHERE
			id = (%d)
		LIMIT 1;`, id)

	err := conn.DB.Get(&assignment, query)
	if err != nil {
		return assignment, err
	}
	return assignment, nil
}

// IsAssignmentExist func ...
func IsAssignmentExist(AssignmentID int64) bool {

	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			assignments
		WHERE
			id = (%d)
		LIMIT 1;`, AssignmentID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// IsAssignmentExistByGradeParameterID func ...
func IsAssignmentExistByGradeParameterID(assignmentID, gradeParameterID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			assignments
		WHERE
			id = (%d) AND grade_parameters_id =(%d)
		LIMIT 1;`, assignmentID, gradeParameterID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// UpdateSubmit ...
func UpdateSubmit(id, userID int64, desc sql.NullString, tx *sqlx.Tx) error {
	var result sql.Result
	var err error

	queryDesc := fmt.Sprintf("(NULL)")
	if desc.Valid {
		queryDesc = fmt.Sprintf("('%s')", desc.String)
	}

	query := fmt.Sprintf(`
		UPDATE  
			p_users_assignments 
		SET
			description = %s,
			updated_at = NOW()
		WHERE
			assignments_id = (%d) AND
			users_id = (%d);
		`, queryDesc, id, userID)

	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

// InsertSubmit ...
func InsertSubmit(id, userID int64, desc sql.NullString, tx *sqlx.Tx) error {

	var result sql.Result
	var err error

	queryDesc := fmt.Sprintf("(NULL)")
	if desc.Valid {
		queryDesc = fmt.Sprintf("('%s')", desc.String)
	}

	query := fmt.Sprintf(`
		INSERT INTO 
			p_users_assignments (
				assignments_id,
				users_id,
				description,
				created_at,
				updated_at
			)
		VALUES(
				(%d),
				(%d),
				%s,
				NOW(),
				NOW()
			);
	`, id, userID, queryDesc)

	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}

	return nil
}

// SelectByGP ...
func SelectByGP(gpsID []int64, isSort bool) ([]Assignment, error) {
	var assignments []Assignment
	if len(gpsID) < 1 {
		return assignments, nil
	}
	queryGP := strings.Join(helper.Int64ToStringSlice(gpsID), ", ")
	querySort := ""
	if isSort {
		querySort = "ORDER BY due_date ASC"
	}
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			status,
			description,
			grade_parameters_id,
			due_date,
			created_at,
			updated_at
		FROM
			assignments
		WHERE
			grade_parameters_id
		IN
			(%s)
		%s;
		`, queryGP, querySort)
	err := conn.DB.Select(&assignments, query)
	if err != nil {
		return assignments, err
	}
	return assignments, nil
}

// GetSubmittedByUser ...
func GetSubmittedByUser(id int64, userID int64) (*UserAssignment, error) {
	assignment := &UserAssignment{}
	query := fmt.Sprintf(`
		SELECT
			assignments_id,
			users_id,
			score,
			description,
			created_at,
			updated_at
		FROM
			p_users_assignments
		WHERE
			assignments_id = (%d) AND
			users_id = (%d)
		LIMIT 1;`, id, userID)

	err := conn.DB.Get(assignment, query)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, nil
	}
	return assignment, nil
}

// SelectSubmittedByUser ...
func SelectSubmittedByUser(id []int64, userID int64) ([]UserAssignment, error) {
	assignment := []UserAssignment{}
	if len(id) < 1 {
		return assignment, nil
	}

	queryID := strings.Join(helper.Int64ToStringSlice(id), ", ")
	query := fmt.Sprintf(`
		SELECT
			assignments_id,
			users_id,
			score,
			description,
			created_at,
			updated_at
		FROM
			p_users_assignments
		WHERE
			assignments_id IN (%s) AND
			users_id = (%d)`, queryID, userID)

	err := conn.DB.Select(&assignment, query)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

// IsUserHaveUploadedAsssignment func ...
func IsUserHaveUploadedAsssignment(AssignmentID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT 
			'x'
		FROM
			p_users_assignments
		WHERE
			assignments_id = (%d)
		LIMIT 1;
		`, AssignmentID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// DeleteAssignment func ...
func DeleteAssignment(AssignmentID int64, tx *sqlx.Tx) error {
	query := fmt.Sprintf(`
		DELETE FROM
			assignments
		WHERE
			id=(%d)
		;`, AssignmentID)

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

// GetAllUserAssignmentByAssignmentID func ...
func GetAllUserAssignmentByAssignmentID(AssignmentID, limit, offset int64) ([]DetailUploadedAssignment, error) {
	query := fmt.Sprintf(`
		SELECT 
			pus.assignments_id,
			pus.score,
			pus.description,
			asg.name,
			asg.description,
			asg.due_date
		FROM
			p_users_assignments pus
		INNER JOIN
			assignments asg
		ON
			asg.id=pus.assignments_id
		WHERE
			pus.assignments_id = (%d)
		LIMIT %d OFFSET %d;
			`, AssignmentID, limit, offset)

	var assignment []DetailUploadedAssignment
	rows, err := conn.DB.Query(query)
	if err != nil {
		return assignment, err
	}
	defer rows.Close()
	for rows.Next() {
		var assignmentID int64
		var name, dueDate string
		var descriptionAssignnment, score, descriptionUser sql.NullString

		err := rows.Scan(&assignmentID, &score, &descriptionUser, &name, &descriptionAssignnment, &dueDate)
		if err != nil {
			return assignment, err
		}
		assignment = append(assignment, DetailUploadedAssignment{
			AssignmentID: assignmentID,
			Name:         name,
			DescriptionAssignment: descriptionAssignnment,
			DescriptionUser:       descriptionUser,
			Score:                 score,
			DueDate:               dueDate,
		})
	}

	return assignment, nil

}

// UpdateScoreAssignment func ...
func UpdateScoreAssignment(assignmentID, userID int64, score float32, tx *sqlx.Tx) error {
	var result sql.Result
	var err error
	query := fmt.Sprintf(`
		UPDATE  
			p_users_assignments 
		SET
				score = (%g),
				updated_at = NOW()
		WHERE
			assignments_id = (%d) AND users_id = (%d)
		;`, score, assignmentID, userID)
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

// IsAssignmentMustUpload func ...
func IsAssignmentMustUpload(assingmentID int64) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			assignments
		WHERE
			id=(%d) AND status=1
		LIMIT 1;
		`, assingmentID)

	err := conn.DB.Get(&x, query)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

// // SelectUserAssignmentsByStatusID func ..
// func SelectUserAssignmentsByStatusID(assignmentID int64) ([]UserAssignmentDetail, error) {
// 	var assignment []UserAssignmentDetail
// 	query := fmt.Sprintf(`
// 			SELECT
// 				usr.identity_code,
// 				usr.name,
// 				pas.score,
// 				pas.description,
// 				pas.updated_at
// 			FROM
// 				p_users_assignments pas
// 			INNER JOIN
// 				users usr
// 			ON
// 				pas.users_id=usr.id
// 			WHERE
// 				pas.assignments_id = (%d)
// 				`, assignmentID)
// 	err := conn.DB.Select(&assignment, query)
// 	if err != nil {
// 		return assignment, err
// 	}
// 	return assignment, nil
// }

// // CreateScore func ...
// func CreateScore(assignmentID int64, usersID []int64, score []float32, tx *sqlx.Tx) error {

// 	var value []string
// 	length := len(usersID)
// 	for i := 0; i < length; i++ {
// 		value = append(value, fmt.Sprintf("(%d, %d, %v, NOW(), NOW())", assignmentID, usersID[i], score[i]))
// 	}
// 	queryValue := strings.Join(value, ", ")
// 	query := fmt.Sprintf(`
// 		INSERT INTO
// 			p_users_assignments
// 			(
// 				assignments_id,
// 				users_id,
// 				score,
// 				created_at,
// 				updated_at
// 			)
// 		VALUES %s
// 		ON DUPLICATE KEY UPDATE
// 		score=VALUES(score), updated_at=VALUES(updated_at)
// 		;`, queryValue)
// 	var err error
// 	if tx != nil {
// 		_, err = tx.Exec(query)
// 	} else {
// 		_, err = conn.DB.Exec(query)
// 	}
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // GetDueDateAssignment func ...
// func GetDueDateAssignment(assignmentID int64) (time.Time, error) {
// 	query := fmt.Sprintf(`
// 			SELECT
// 				due_date
// 			FROM
// 				assignments
// 			WHERE
// 				id=(%d)
// 			LIMIT 1;
// 		`, assignmentID)
// 	var dueDate time.Time
// 	err := conn.DB.Get(&dueDate, query)
// 	if err != nil {
// 		return dueDate, err
// 	}
// 	return dueDate, nil
// }

// // SelectAssignmentIDByGradeParameter func ..
// func SelectAssignmentIDByGradeParameter(gradeParameterID int64) ([]int64, error) {
// 	query := fmt.Sprintf(`
// 		SELECT DISTINCT
// 			id
// 		FROM
// 			assignments
// 		WHERE
// 			grade_parameters_id=(%d);
// 		`, gradeParameterID)
// 	var id []int64
// 	err := conn.DB.Select(&id, query)
// 	if err != nil {
// 		return id, err
// 	}
// 	return id, nil
// }

// // SelectAssignmentIDByGradeParameterIN func ...
// func SelectAssignmentIDByGradeParameterIN(gradeParameterID []int64) ([]int64, error) {
// 	var gradeQuery []string
// 	for _, value := range gradeParameterID {
// 		gradeQuery = append(gradeQuery, fmt.Sprintf("%d", value))
// 	}
// 	queryGradeList := strings.Join(gradeQuery, ",")
// 	query := fmt.Sprintf(`
// 		SELECT
// 			id
// 		FROM
// 			assignments
// 		WHERE
// 			grade_parameters_id
// 		IN
// 			(%s) &&  NOW() < due_date;
// 		`, queryGradeList)
// 	var id []int64
// 	err := conn.DB.Select(&id, query)
// 	if err != nil {
// 		return id, err
// 	}
// 	return id, nil
// }

// // SelectScore func ...
// func SelectScore(userID int64, assignmentsID []int64) ([]float32, error) {
// 	var id []string
// 	for _, value := range assignmentsID {
// 		id = append(id, fmt.Sprintf("%d", value))
// 	}
// 	queryAssignmentsID := strings.Join(id, ",")
// 	query := fmt.Sprintf(`
// 		SELECT
// 			score
// 		FROM
// 			p_users_assignments
// 		WHERE
// 			users_id=(%d) AND assignments_id IN (%s)
// 		`, userID, queryAssignmentsID)
// 	var scores []float64
// 	rows, err := conn.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id sql.NullFloat64
// 		if err := rows.Scan(&id); err != nil {
// 			return nil, err
// 		}
// 		scores = append(scores, id.Float64)
// 	}

// 	var result []float32
// 	for _, value := range scores {
// 		result = append(result, float32(value))
// 	}

// 	return result, nil

// }

// // IsUploaded func ...
// func IsUploaded(assignmentID int64) bool {
// 	var x string
// 	query := fmt.Sprintf(`
// 		SELECT
// 			'x'
// 		FROM
// 			p_users_assignments
// 		WHERE
// 			assignments_id = (%d)
// 		LIMIT 1;
// 		`, assignmentID)
// 	err := conn.DB.Get(&x, query)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// // GetScoreByIDUser func ..
// func GetScoreByIDUser(assignmentID, userID int64) float32 {
// 	query := fmt.Sprintf(`
// 		SELECT
// 			score
// 		FROM
// 			p_users_assignments
// 		WHERE
// 			users_id=(%d) AND assignments_id=(%d)
// 		LIMIT 1;
// 		`, userID, assignmentID)
// 	var score float32
// 	err := conn.DB.Get(&score, query)
// 	if err != nil {
// 		return score
// 	}
// 	return score
// }

// // SelectUnsubmittedAssignment func ..
// func SelectSubmittedAssignment(assignmentID []int64, userID int64) []int64 {
// 	var assignment []string
// 	for _, value := range assignmentID {
// 		assignment = append(assignment, fmt.Sprintf("%d", value))
// 	}
// 	assignmentQuery := strings.Join(assignment, ",")
// 	query := fmt.Sprintf(`
// 		SELECT
// 			assignments_id
// 		FROM
// 			p_users_assignments
// 		WHERE
// 			assignments_id IN (%s) AND users_id = (%d)
// 		;`, assignmentQuery, userID)
// 	var res []int64
// 	err := conn.DB.Select(&res, query)
// 	if err != nil {
// 		return res
// 	}
// 	return res
// }

// // SelectAssignmentByID func ..
// func SelectAssignmentByID(assignmentID []int64) []Assignment {
// 	var assignment []string
// 	for _, value := range assignmentID {
// 		assignment = append(assignment, fmt.Sprintf("%d", value))
// 	}
// 	assignmentQuery := strings.Join(assignment, ",")
// 	query := fmt.Sprintf(`
// 		SELECT
// 			id,
// 			due_date,
// 			name,
// 			description,
// 			status
// 		FROM
// 			assignments
// 		WHERE
// 			id
// 		IN
// 			(%s)
// 		ORDER BY due_date ASC
// 		`, assignmentQuery)
// 	var res []Assignment
// 	err := conn.DB.Select(&res, query)
// 	if err != nil {
// 		return res
// 	}
// 	return res

// }
