package course

type Course struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	UCU      int8   `db:"ucu"`
	Semester int8   `db:"semester"`
	Status   int8   `db:"status"`
}
