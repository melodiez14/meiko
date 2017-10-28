package bot

import (
	"time"
)

const (
	StatusUser = 0
	StatusBot  = 1
)

type Log struct {
	ID        uint64    `db:"id"`
	Message   string    `db:"message"`
	UserID    int64     `db:"users_id"`
	Status    uint8     `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}
