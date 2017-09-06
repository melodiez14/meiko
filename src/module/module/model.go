package module

type Module struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
