package assignment

import (
	"time"
)

type Assignment struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Status     int8      `db:"status"`
	UploadDate time.Time `db:"upload_date"`
	DueDate    time.Time `db:"due_date"`
}
