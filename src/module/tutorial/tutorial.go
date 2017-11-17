package tutorial

import (
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

// SelectByPage ...
func SelectByPage(scheduleID int64, limit, offset uint64) ([]Tutorial, error) {
	var tutorials []Tutorial
	query := fmt.Sprintf(`
		SELECT 
			id,
			name,
			description,
			created_at
		FROM
			tutorials
		WHERE
			schedules_id = (%d)
		LIMIT (%d)
		OFFSET (%d)
		`, scheduleID, limit, offset)

	err := conn.DB.Select(&tutorials, query)
	if err != nil {
		return tutorials, err
	}
	return tutorials, nil
}

// GetByID ...
func GetByID(id int64) (Tutorial, error) {
	var tutorial Tutorial
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			description,
			created_at
		FROM
			tutorials
		WHERE
			id = (%d)
		LIMIT 1;	
	`, id)

	err := conn.DB.Get(&tutorial, query)
	if err != nil {
		return tutorial, err
	}
	return tutorial, nil
}
