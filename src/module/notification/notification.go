package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

func Push(title, description string, playerID []string) error {

	payload := map[string]interface{}{
		"app_id":             "9f1b7d96-d6d8-4e3f-9d2b-9978f5f2b5a1",
		"include_player_ids": playerID,
		"headings": map[string]string{
			"en": title,
		},
		"contents": map[string]string{
			"en": description,
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Basic OWQ0MTFkZjYtNjdlOC00N2Y2LWFmN2YtN2IxMDdkOGNhYjEw")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
