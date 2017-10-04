package file

const (
	ColID        = "id"
	ColName      = "name"
	ColPath      = "path"
	ColMime      = "mime"
	ColExtension = "extension"
	ColUserID    = "users_id"
	ColType      = "type"
	ColTableName = "table_name"
	ColTableID   = "table_id"

	OperatorEquals  = "="
	OperatorUnquals = "!="
	OperatorIn      = "IN"
	OperatorMore    = ">"
	OperatorLess    = "<"
)

type (
	QueryGet    struct{ string }
	QuerySelect struct{ string }
	QueryInsert struct{ string }
	QueryUpdate struct{ string }
)

type File struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Path      string `db:"path"`
	Mime      string `db:"mime"`
	Extension string `db:"extension"`
	UserID    int64  `db:"users_id"`
	Type      string `db:"type"`
	TableName string `db:"table_name"`
	TableID   string `db:"table_id"`
}
