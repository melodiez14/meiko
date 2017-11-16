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
