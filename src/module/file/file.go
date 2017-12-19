package file

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetByIDExt(id string, column ...string) (File, error) {

	var c []string
	var file File
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColMime,
			ColExtension,
			ColUserID,
			ColType,
			ColTableName,
			ColTableID,
		}
	} else {
		for _, val := range column {
			c = append(c, val)
		}
	}

	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(`SELECT %s FROM files WHERE id = ('%s') LIMIT 1;`, cols, id)
	err := conn.DB.Get(&file, query)
	if err != nil {
		return file, err
	}
	return file, nil
}

func GetByTypeUserID(userID int64, typ string, column ...string) (File, error) {
	var c []string
	var file File
	if len(column) < 1 {
		c = []string{
			ColID,
			ColName,
			ColMime,
			ColExtension,
			ColUserID,
			ColType,
			ColTableName,
			ColTableID,
		}
	} else {
		for _, val := range column {
			c = append(c, val)
		}
	}

	cols := strings.Join(c, ", ")
	query := fmt.Sprintf(`SELECT %s FROM files WHERE users_id = (%d) AND type = ('%s') AND status = (%d) LIMIT 1;`, cols, userID, typ, StatusExist)
	err := conn.DB.Get(&file, query)
	if err != nil {
		return file, err
	}
	return file, nil
}

func DeleteProfileImage(userID int64, tx *sqlx.Tx) error {
	query := fmt.Sprintf(`
		UPDATE
			files
		SET
			status = (%d),
			updated_at = NOW()
		WHERE
			users_id = (%d) AND
			type IN ('%s', '%s');`, StatusDeleted, userID, TypProfPict, TypProfPictThumb)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteByRelation ...
func DeleteByRelation(typ, tableID string, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		UPDATE
			files
		SET
			status = (%d),
			updated_at = NOW()
		WHERE
			table_name = ('%s') AND
			table_id = ('%s');`, StatusDeleted, typ, tableID)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func Delete(id string, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		UPDATE
			files
		SET
			status = (%d),
			updated_at = NOW()
		WHERE
			id = ('%s');`, StatusDeleted, id)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}
	return nil
}

// GetByStatus func ...
func GetByStatus(status int, tableID int64) ([]string, error) {

	var files []string
	query := fmt.Sprintf(`
			SELECT 
				id
			FROM
				files
			WHERE
				status = (%d) AND table_id = (%d) 
			;`, status, tableID)

	rows, err := conn.DB.Query(query)
	if err != nil {
		return files, err
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return files, err
		}
		files = append(files, id)
	}
	return files, nil
}

func Insert(id, name, mime, extension string, userID int64, typ string, tx *sqlx.Tx) error {

	var result sql.Result
	var err error
	query := fmt.Sprintf(`
		INSERT INTO
		files (
			id,
			name,
			mime,
			extension,
			type,
			users_id,
			created_at,
			updated_at
		) VALUES (
			('%s'),
			('%s'),
			('%s'),
			('%s'),
			('%s'),
			(%d),
			NOW(),
			NOW()
		);`, id, name, mime, extension, typ, userID)

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

func InsertImageProfile(id, name, mime, extension string, userID int64, typ string, tx *sqlx.Tx) error {

	var result sql.Result
	var err error
	query := fmt.Sprintf(`
		INSERT INTO
		files (
			id,
			name,
			mime,
			extension,
			type,
			users_id,
			table_id,
			created_at,
			updated_at
		) VALUES (
			('%s'),
			('%s'),
			('%s'),
			('%s'),
			('%s'),
			(%d),
			(%d),
			NOW(),
			NOW()
		);`, id, name, mime, extension, typ, userID, userID)

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

func UpdateRelation(id, typ, tableID string, tx *sqlx.Tx) error {

	var result sql.Result
	var err error
	query := fmt.Sprintf(`
		UPDATE
			files
		SET
			table_id = ('%s'),
			updated_at = NOW()
		WHERE
			id = ('%s') AND
			type = ('%s') AND
			table_id IS NULL;
		`, tableID, id, typ)

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

func IsHasRelation(id string) bool {
	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			files
		WHERE
			id = ('%s') AND
			table_name IS NOT NULL AND
			table_id IS NOT NULL
		LIMIT 1;	
	`, id)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// UpdateStatusFiles func ...
func UpdateStatusFiles(id string, status int, tx *sqlx.Tx) error {

	var result sql.Result
	var err error
	query := fmt.Sprintf(`
		UPDATE 
			files
		SET
			status = (%d)
		WHERE
			id = ('%s')
		;`, status, id)

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

// IsExistID func ...
func IsExistID(fileID string) bool {

	var x string
	query := fmt.Sprintf(`
		SELECT 
			'x'
		FROM
			files
		WHERE
			id = ('%s') AND
			status = (%d)
		LIMIT 1;
		`, fileID, StatusExist)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

// SelectByRelation func ...
func SelectByRelation(typ string, tablesID []string, userID *int64) ([]File, error) {
	var files []File

	if len(tablesID) < 1 {
		return files, nil
	}

	var queryUserID string
	if userID != nil {
		queryUserID = fmt.Sprintf("users_id = (%d) AND ", *userID)
	}

	queryTableID := strings.Join(tablesID, ", ")
	query := fmt.Sprintf(`
		SELECT 
			id,
			name,
			extension,
			mime,
			type,
			users_id,
			table_id
		FROM
			files
		WHERE
			%s
			status = (%d) AND
			type = ('%s') AND
			table_id IN (%s);
		`, queryUserID, StatusExist, typ, queryTableID)

	err := conn.DB.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, nil
}

func SelectIDByRelation(typ string, tableID string, userID int64) ([]string, error) {

	var filesID []string
	query := fmt.Sprintf(`
		SELECT 
			id
		FROM
			files
		WHERE
			users_id = (%d) AND
			status = (%d) AND
			type = ('%s') AND
			table_id = ('%s');
		`, userID, StatusExist, typ, tableID)

	err := conn.DB.Select(&filesID, query)
	if err != nil {
		return filesID, err
	}
	return filesID, nil
}

// UpdateStatusFilesByNameID func ...
func UpdateStatusFilesByNameID(TableName string, Status, TableID int64, tx *sqlx.Tx) error {
	query := fmt.Sprintf(`
		UPDATE
			files
		SET
			status=(%d)
		WHERE
			table_name=('%s') AND table_id=(%d)
		;`, Status, TableName, TableID)
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

// GetByRelation ...
func GetByRelation(typ, tableID string) (File, error) {

	var file File
	query := fmt.Sprintf(`
		SELECT 
			id,
			extension
		FROM
			files
		WHERE
			status = (%d) AND
			type = ('%s') AND
			table_id = ('%s')
		LIMIT 1;
		`, StatusExist, typ, tableID)

	err := conn.DB.Get(&file, query)
	if err != nil {
		return file, err
	}
	return file, nil
}

// GetByUserIDTableIDName func ...
func GetByUserIDTableIDName(UserID, TableID int64, TableName string) ([]File, error) {
	var files []File
	query := fmt.Sprintf(`
		SELECT 
			id,
			extension
		FROM
			files
		WHERE
			users_id = (%d) AND table_name=('%s') AND table_id=(%d)
		`, UserID, TableName, TableID)

	err := conn.DB.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, nil
}
