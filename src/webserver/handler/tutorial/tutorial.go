package tutorial

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	cs "github.com/melodiez14/meiko/src/module/course"
	fl "github.com/melodiez14/meiko/src/module/file"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	tt "github.com/melodiez14/meiko/src/module/tutorial"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// ReadHandler ...
func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := readParams{
		payload:    r.FormValue("payload"),
		scheduleID: r.FormValue("schedule_id"),
		page:       r.FormValue("pg"),
		total:      r.FormValue("ttl"),
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

	isHasAccess := false
	switch args.payload {
	case "assistant":
		isHasAccess = cs.IsAssistant(sess.ID, args.scheduleID) &&
			sess.IsHasRoles(rg.ModuleTutorial, rg.RoleXRead, rg.RoleRead)
	case "student":
		isHasAccess = cs.IsEnrolled(sess.ID, args.scheduleID)
	}

	if !isHasAccess {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	offset := (args.page - 1) * args.total
	tutorials, count, err := tt.SelectByPage(args.scheduleID, args.total, offset, true)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	respTutorial := []readTutorial{}
	for _, val := range tutorials {
		desc := "-"
		if val.Description.Valid {
			desc = val.Description.String
		}

		respTutorial = append(respTutorial, readTutorial{
			ID:          val.ID,
			Name:        val.Name,
			Description: desc,
			Time:        val.CreatedAt.Unix(),
			URL:         fmt.Sprintf("/api/v1/filerouter/?id=%d&payload=tutorial", val.ID),
		})
	}

	totalPage := count / args.total
	if count%args.total > 0 {
		totalPage++
	}

	resp := readResponse{
		Page:      args.page,
		TotalPage: totalPage,
		Tutorials: respTutorial,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// ReadDetailHandler ...
func ReadDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleTutorial, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readDetailParams{
		id: ps.ByName("tutorial_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	tutorial, err := tt.GetByID(args.id)
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

	resp := readDetailResponse{
		ID:          tutorial.ID,
		Name:        tutorial.Name,
		Description: tutorial.Description.String,
		Time:        tutorial.CreatedAt.Unix(),
		URL:         fmt.Sprintf("/api/v1/filerouter/?id=%d&payload=tutorial", tutorial.ID),
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// CreateHandler ...
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleTutorial, rg.RoleXCreate, rg.RoleCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := createParams{
		name:        r.FormValue("name"),
		description: r.FormValue("description"),
		fileID:      r.FormValue("file_id"),
		scheduleID:  r.FormValue("schedule_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if !cs.IsAssistant(sess.ID, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	if tt.IsExistName(args.name, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			AddError("Tutorial name is already used"))
		return
	}

	if fl.IsHasRelation(args.fileID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid File"))
		return
	}

	tx := conn.DB.MustBegin()
	lastInsertID, err := tt.Insert(args.name, args.description, args.scheduleID, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tutorialID := strconv.FormatInt(lastInsertID, 10)

	err = fl.UpdateRelation(args.fileID, fl.TableTutorial, tutorialID, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tx.Commit()
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Success"))
	return
}

// DeleteHandler ...
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleTutorial, rg.RoleXDelete, rg.RoleDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := deleteParams{
		id: ps.ByName("tutorial_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if !tt.IsExistID(args.id) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	id := strconv.FormatInt(args.id, 10)

	tx := conn.DB.MustBegin()
	err = tt.Delete(args.id, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	err = fl.DeleteByRelation(fl.TableTutorial, id, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tx.Commit()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Success"))
	return
}

// UpdateHandler ...
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleTutorial, rg.RoleXDelete, rg.RoleDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := updateParams{
		id:          ps.ByName("tutorial_id"),
		name:        r.FormValue("name"),
		description: r.FormValue("description"),
		fileID:      r.FormValue("file_id"),
		scheduleID:  r.FormValue("schedule_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	tutorial, err := tt.GetByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if tutorial.ScheduleID != args.scheduleID {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if tt.IsExistName(args.name, tutorial.ScheduleID, tutorial.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			AddError("Tutorial name is already used"))
		return
	}

	id := strconv.FormatInt(args.id, 10)
	file, err := fl.GetByRelation(fl.TableTutorial, id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tx := conn.DB.MustBegin()
	err = tt.Update(args.id, args.name, args.description, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if file.ID != args.fileID {

		if fl.IsHasRelation(args.fileID) {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("Invalid File"))
			return
		}

		err = fl.Delete(file.ID, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}

		err = fl.UpdateRelation(args.fileID, fl.TableTutorial, id, tx)
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
		SetMessage("Success"))
	return
}
