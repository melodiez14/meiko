package notification

import "time"
import "github.com/go-sql-driver/mysql"

type Notification struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"descriptions"`
	TableName   string         `db:"table_name"`
	TableID     string         `db:"table_id"`
	ReadAt      mysql.NullTime `db:"read_at"`
	CreatedAt   time.Time      `db:"created_at"`
}
