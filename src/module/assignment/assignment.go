package assignment

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

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

// Insert function is ...
func Insert(GradeParameters int64, Name, Status, DueDate string, Description sql.NullString, tx ...*sqlx.Tx) (string, error) {
	queryDescription := fmt.Sprintf("NULL")
	if Description.Valid {
		queryDescription = fmt.Sprintf("('%s')", Description.String)
	}
	query := fmt.Sprintf(`
		INSERT INTO
			assignments(
				name,
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
			('%s'),
			('%d'),
			%s,
			NOW(),
			NOW()
		);
		`, Name, Status, DueDate, GradeParameters, queryDescription)

	var result sql.Result
	var err error
	if len(tx) == 1 {
		result, err = tx[0].Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return "", err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return "", fmt.Errorf("No rows affected")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("Error get LastIsertedId")
	}
	ID := strconv.FormatInt(id, 10)
	return ID, nil
}

// Update func ...
func Update(GradeParameters, id int64, Name, Status, DueDate string, Description sql.NullString, tx *sqlx.Tx) error {

	var result sql.Result
	var err error
	queryDescription := fmt.Sprintf("NULL")
	if Description.Valid {
		queryDescription = fmt.Sprintf(Description.String)
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
		`, Name, Status, DueDate, GradeParameters, queryDescription, id)
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
func SelectByPage(limit, offset uint16) ([]FileAssignment, error) {
	var assignment []FileAssignment
	query := fmt.Sprintf(`
			SELECT
				asg.grade_parameters_id,
				asg.name,
				asg.description,
				asg.status,
				asg.due_date
			FROM
				assignments asg
			LIMIT %d OFFSET %d;`, limit, offset)

	rows, err := conn.DB.Queryx(query)
	defer rows.Close()
	if err != nil {
		return assignment, err
	}

	for rows.Next() {
		var name string
		var description sql.NullString
		var status int8
		var gradeParameterID int32
		var dueDate time.Time

		err := rows.Scan(&gradeParameterID, &name, &description, &status, &dueDate)
		if err != nil {
			return assignment, err
		}
		assignment = append(assignment, FileAssignment{
			Assignment: Assignment{
				Name:             name,
				Status:           status,
				Description:      description,
				GradeParameterID: gradeParameterID,
				DueDate:          dueDate,
			},
		})
	}
	return assignment, nil
}

// GetByAssignementID func ...
func GetByAssignementID(assignmentID int64) (DetailAssignment, error) {

	var assignment DetailAssignment
	query := fmt.Sprintf(`
			SELECT
				asg.id,
				asg.status,
				asg.name,
				asg.grade_parameters_id,
				asg.description,
				asg.due_date,
				fs.name,
				fs.mime,
				gp.type,
				gp.percentage
			FROM((
				assignments asg
			LEFT JOIN
				files fs
			ON
				asg.id = fs.table_id)
			LEFT JOIN
				grade_parameters gp
			ON
				gp.id = asg.grade_parameters_id)
			WHERE
				asg.id = (%d)
			LIMIT 1;`, assignmentID)

	rows := conn.DB.QueryRowx(query)

	// scan data to variable
	var name, Type string
	var nameFile, mime, description sql.NullString
	var status int8
	var gradeParameterID int32
	var id int64
	var dueDate time.Time
	var percentage float32

	err := rows.Scan(&id, &status, &name, &gradeParameterID, &description, &dueDate, &nameFile, &mime, &Type, &percentage)
	if err != nil {
		return assignment, err
	}

	return DetailAssignment{
		Assignment: Assignment{
			ID:               id,
			Name:             name,
			Description:      description,
			Status:           status,
			GradeParameterID: gradeParameterID,
			DueDate:          dueDate,
		},
		File: File{
			Name: nameFile,
			Mime: mime,
		},
		GradeParameter: GradeParameter{
			Type:       Type,
			Percentage: percentage,
		},
	}, nil
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

// UploadAssignment func ...
func UploadAssignment(assignmentID, userID int64, description sql.NullString, tx *sqlx.Tx) error {

	var result sql.Result
	var err error

	queryDescription := fmt.Sprintf("NULL")
	if description.Valid {
		queryDescription = fmt.Sprintf(description.String)
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
				('%s'),
				NOW(),
				NOW()
			)
		;`, assignmentID, userID, queryDescription)
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
