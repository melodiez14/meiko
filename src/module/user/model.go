package user

type User struct {
	ID       int64                      `json:"id"`
	Name     string                     `json:"name"`
	Email    string                     `json:"email"`
	Password string                     `json:"password"`
	Gender   string                     `json:"gender"`
	College  string                     `json:"college"`
	Note     string                     `json:"note"`
	Roles    map[string]map[string]bool `json:"roles"`
	Status   bool                       `json:"active"`
}
