package assignment

import (
	"database/sql"
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetIncompleteAssignment(userID int64) ([]Assignment, error) {
	var assignments []Assignment
	query := fmt.Sprintf(queryGetIncompleteAssignment, userID, userID)
	err := conn.DB.Select(&assignments, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignments, nil
}
