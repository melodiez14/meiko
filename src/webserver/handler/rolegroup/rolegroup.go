package rolegroup

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	usr "github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/conn"
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

	var resp getPrivilegeResponse
	sess := r.Context().Value("User").(*auth.User)
	if sess == nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
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

	resp = getPrivilegeResponse{
		IsLoggedIn: true,
		Modules:    roles,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// CreateHandler ...
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleRole, rg.RoleXCreate, rg.RoleCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := createParams{
		name:    r.FormValue("name"),
		modules: r.FormValue("modules"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if rg.IsExistName(args.name) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			AddError("Name already exist"))
		return
	}

	tx := conn.DB.MustBegin()

	rolegroupID, err := rg.Insert(args.name, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if len(args.modules) > 1 {
		err = rg.InsertModuleAccess(rolegroupID, args.modules, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	tx.Commit()
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Roles sucessfully inserted"))
	return
}

// ReadHandler ...
func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleRole, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readParams{
		page:  r.FormValue("pg"),
		total: r.FormValue("ttl"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if args.total > 100 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Max total should be less than or equal to 100"))
		return
	}

	offset := (args.page - 1) * args.total
	roles, count, err := rg.SelectByPage(args.total, offset, true)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	rolegroupID, err := usr.SelectDistinctRolegroupID()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	respRoles := []readRoles{}
	for _, val := range roles {
		isDeleteAllow := true
		if helper.Int64InSlice(val.ID, rolegroupID) {
			isDeleteAllow = false
		}

		respRoles = append(respRoles, readRoles{
			ID:            val.ID,
			Name:          val.Name,
			IsDeleteAllow: isDeleteAllow,
		})
	}

	total := count / args.total
	if count%args.total > 0 {
		total++
	}

	resp := readResponse{
		TotalPage: total,
		Page:      args.page,
		Roles:     respRoles,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// ReadDetailHandler ...
func ReadDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleRole, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readDetailParams{
		id: ps.ByName("rolegroup_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	role, err := rg.GetByID(args.id)
	if err != nil {
		if err == sql.ErrNoRows {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusNoContent))
			return
		}

		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	moduleAccess, err := rg.SelectModuleAccess(role.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	resp := readDetailResponse{
		ID:      role.ID,
		Name:    role.Name,
		Modules: moduleAccess,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// DeleteHandler ...
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleRole, rg.RoleXDelete, rg.RoleDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readDetailParams{
		id: ps.ByName("rolegroup_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if !rg.IsExist(args.id) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}

	if usr.IsExistRolegroupID(args.id) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			AddError("Error: this rolegroups has already used by user"))
		return
	}

	tx := conn.DB.MustBegin()
	err = rg.DeleteModuleAccess(args.id, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	err = rg.Delete(args.id, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tx.Commit()
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Rolegroup successfully deleted"))
	return
}

// UpdateHandler ...
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleRole, rg.RoleXUpdate, rg.RoleUpdate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := updateParams{
		id:      ps.ByName("rolegroup_id"),
		name:    r.FormValue("name"),
		modules: r.FormValue("modules"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if !rg.IsExist(args.id) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}

	tx := conn.DB.MustBegin()
	err = rg.Update(args.id, args.name, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	err = rg.DeleteModuleAccess(args.id, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if len(args.modules) > 1 {
		err = rg.InsertModuleAccess(args.id, args.modules, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	tx.Commit()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Rolegroup has been updated"))
	return
}
