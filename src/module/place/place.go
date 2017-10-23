package place

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/meiko/src/util/conn"
)

func Search(id string) ([]string, error) {
	places := []string{}
	query := fmt.Sprintf("SELECT id FROM places WHERE id LIKE ('%%%s%%')", id)
	err := conn.DB.Select(&places, query)
	if err != nil {
		return places, err
	}
	return places, nil
}

func IsExistID(id string) bool {
	var place string
	query := fmt.Sprintf("SELECT id FROM places WHERE id = ('%s') LIMIT 1", id)
	err := conn.DB.Get(&place, query)
	if err != nil {
		return false
	}
	return true
}

func Insert(id string, description sql.NullString, tx ...*sqlx.Tx) error {

	queryDescription := "(NULL)"
	if description.Valid {
		queryDescription = fmt.Sprintf("('%s')", description.String)
	}

	query := fmt.Sprintf(`INSERT INTO
		places (
			id,
			description,
			created_at,
			updated_at
		) VALUES (
			('%s'),
			%s,
			NOW(),
			NOW()
		)`, id, queryDescription)

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

	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}

	return nil
}
