package rolegroup

const (
	ModuleUser        = "users"
	ModuleCourse      = "courses"
	ModuleRole        = "roles"
	ModuleAttendance  = "attendances"
	ModuleSchedule    = "schedules"
	ModuleAssignment  = "assignments"
	ModuleInformation = "informations"

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
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
