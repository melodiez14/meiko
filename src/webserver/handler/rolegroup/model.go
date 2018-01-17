package rolegroup

type getPrivilegeResponse struct {
	IsLoggedIn bool                `json:"is_logged_in"`
	Modules    map[string][]string `json:"modules"`
}

type createParams struct {
	name    string
	modules string
}

type createArgs struct {
	name    string
	modules map[string][]string
}

type readParams struct {
	page  string
	total string
}

type readArgs struct {
	page  int
	total int
}

type readRoles struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	IsDeleteAllow bool   `json:"is_delete_allow"`
	CreatedAt     string `json:"created_at"`
}

type readResponse struct {
	Page      int         `json:"page"`
	TotalPage int         `json:"total_page"`
	Roles     []readRoles `json:"roles"`
}

type readDetailParams struct {
	id string
}

type readDetailArgs struct {
	id int64
}

type readDetailResponse struct {
	ID      int64               `json:"id"`
	Name    string              `json:"name"`
	Modules map[string][]string `json:"modules"`
}

type deleteParams struct {
	id string
}

type deleteArgs struct {
	id int64
}

type updateParams struct {
	id      string
	name    string
	modules string
}

type updateArgs struct {
	id      int64
	name    string
	modules map[string][]string
}

type searchParams struct {
	Text string
}

type searchArgs struct {
	Text string
}
type searchResponse struct {
	ID        int64  `json:"id"`
	Rolegroup string `json:"role"`
}
