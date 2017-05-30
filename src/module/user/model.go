package user

type User struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Gender   string `db:"gender"`
	College  string `db:"college"`
	Note     string `db:"note"`
	Status   bool   `db:"active"`
}
