package rolegroup

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/helper"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// GetPrivilege handles the http request for getting the list of user privilege and access to listed modules
/*
	@params:
	@example:
	@return
		is_logged_id = true
		modules = {
			users = ["create", "read", "update"],
			course = ["read"]
		}
*/
func GetPrivilege(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var res getPrivilegeResponse
	sess := r.Context().Value("User").(*auth.User)
	if sess == nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(res))
		return
	}

	// set response data
	var role string
	roles := map[string][]string{}
	for i, val := range sess.Roles {
		for _, v := range val {
			switch v {
			case rg.RoleCreate, rg.RoleXCreate:
				role = "CREATE"
			case rg.RoleRead, rg.RoleXRead:
				role = "READ"
			case rg.RoleUpdate, rg.RoleXUpdate:
				role = "UPDATE"
			case rg.RoleDelete, rg.RoleXDelete:
				role = "DELETE"
			}
			if !helper.IsStringInSlice(role, roles[i]) {
				roles[i] = append(roles[i], role)
			}
		}
	}

	res = getPrivilegeResponse{
		IsLoggedIn: true,
		Modules:    roles,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}
