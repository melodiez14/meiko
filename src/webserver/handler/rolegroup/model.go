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
	page  uint8
	total uint8
}

type readResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
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
