package rolegroup

import (
	"time"
)

const (
	ModuleUser        = "users"
	ModuleCourse      = "courses"
	ModuleRole        = "roles"
	ModuleAttendance  = "attendances"
	ModuleSchedule    = "schedules"
	ModuleAssignment  = "assignments"
	ModuleInformation = "informations"
	ModuleTutorial    = "tutorials"

	RoleCreate  = "CREATE"
	RoleRead    = "READ"
	RoleUpdate  = "UPDATE"
	RoleDelete  = "DELETE"
	RoleXCreate = "XCREATE"
	RoleXRead   = "XREAD"
	RoleXUpdate = "XUPDATE"
	RoleXDelete = "XDELETE"
)

type RoleGroup struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"updated_at"`
}

type Privilege struct {
	ID      int64  `db:"id"`
	Module  string `db:"module"`
	Ability string `db:"ability"`
}
