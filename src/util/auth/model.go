package auth

type User struct {
	ID      int64               `json:"id"`
	Name    string              `json:"name"`
	Email   string              `json:"email"`
	Gender  int8                `json:"gender"`
	College string              `json:"college"`
	Note    string              `json:"note"`
	Roles   map[string][]string `json:"roles"`
	LineID  string              `json:"line_id"`
	Phone   string              `json:"phone"`
	Status  int8                `json:"active"`
}
