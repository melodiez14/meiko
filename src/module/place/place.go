package place

import (
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

func (args Place) Insert(txs ...*sqlx.Tx) error {

	var desc string
	if args.Description.Valid {
		desc = fmt.Sprintf("('%s')", args.Description.String)
	} else {
		desc = "(NULL)"
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
		)`, args.ID, desc)

	if len(txs) == 1 {
		_, err := txs[0].Exec(query)
		if err != nil {
			return err
		}
		return nil
	}

	_, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
