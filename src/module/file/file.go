package file

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetByIDExt(id, ext string, column ...string) (File, error) {

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
	query := fmt.Sprintf(`SELECT %s FROM files WHERE id = ('%s') AND extension = ('%s') LIMIT 1;`, cols, id, ext)
	err := conn.DB.Get(&file, query)
	if err != nil {
		fmt.Println(1, err.Error())
		return file, err
	}
	fmt.Println(2, file)
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
		fmt.Println(1, err.Error())
		return file, err
	}
	fmt.Println(2, file)
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

func UpdateRelation(id, tableName, tableID string, tx *sqlx.Tx) error {

	var result sql.Result
	var err error
	query := fmt.Sprintf(`
		UPDATE
			files
		SET
			table_name = ('%s'),
			table_id = ('%s'),
			updated_at = NOW()
		WHERE
			id = ('%s');`, tableName, tableID, id)

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
