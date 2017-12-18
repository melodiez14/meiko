package file

import (
	"database/sql"
)

const (
	ColID        = "id"
	ColName      = "name"
	ColMime      = "mime"
	ColExtension = "extension"
	ColUserID    = "users_id"
	ColType      = "type"
	ColTableName = "table_name"
	ColTableID   = "table_id"

	TypProfPict         = "PL-IMG-M"
	TypProfPictThumb    = "PL-IMG-T"
	TypAssignment       = "ASG-FILE"
	TypAssignmentUpload = "ASG-UPL"
	TypTutorial         = "TT-FILE"
	TypInfPict          = "INF-IMG-M"
	TypInfPictThumb     = "INF-IMG-T"
	TypInf              = "INF-FILE"

	TableAssignment = "assignments"
	TableTutorial   = "tutorials"

	StatusDeleted = 0
	StatusExist   = 1

	NoImgAvailable = "/static/img/noimgavailable.png"
	NotFoundURL    = "/api/v1/file/default/notfound.png"
	UsrNoPhotoURL  = "/api/v1/file/default/usrnophoto.png"
)

type File struct {
	ID        string         `db:"id"`
	Name      string         `db:"name"`
	Mime      string         `db:"mime"`
	Extension string         `db:"extension"`
	UserID    int64          `db:"users_id"`
	Type      string         `db:"type"`
	TableName sql.NullString `db:"table_name"`
	TableID   sql.NullString `db:"table_id"`
}
