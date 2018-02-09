package notification

import (
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

func IsExist(userID int64, onesignalID string) bool {

	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			notifications
		WHERE
			users_id = (%d) AND
			onesignal_id = ('%s')
		LIMIT 1;	
	`, userID, onesignalID)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func Insert(userID int64, onesignalID string) error {
	query := fmt.Sprintf(`
		INSERT INTO
			notifications (
				users_id,
				onesignal_id,
				created_at,
				updated_at
			) VALUES (
				(%d),
				('%s'),
				NOW(),
				NOW()
			);
	`, userID, onesignalID)

	result, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows < 1 || err != nil {
		return fmt.Errorf("No rows affected")
	}

	return nil
}
