package rolegroup

type getPrivilegeResponse struct {
	IsLoggedIn bool                `json:"is_logged_in"`
	Modules    map[string][]string `json:"modules"`
}
