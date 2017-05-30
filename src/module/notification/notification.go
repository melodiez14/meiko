package notification

import (
	"database/sql"
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

func Get(userID int64, page uint16, limit uint8) ([]Notification, error) {
	var notifications []Notification

	startRow := uint32(page-1) * uint32(limit)

	query := fmt.Sprintf(queryGet, userID, startRow, limit)
	err := conn.DB.Select(&notifications, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return notifications, nil
}

func (n Notification) GetURL() string {
	return "http://URL.com"
}
