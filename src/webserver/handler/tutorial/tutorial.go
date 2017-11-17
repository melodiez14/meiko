package tutorial

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	tt "github.com/melodiez14/meiko/src/module/tutorial"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// ReadHandler ...
func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleSchedule, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readParams{
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

	offset := (args.page - 1) * args.total
	tutorials, err := tt.SelectByPage(args.scheduleID, args.total, offset)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	resp := []readResponse{}
	for _, val := range tutorials {
		desc := "-"
		if val.Description.Valid {
			desc = val.Description.String
		}

		resp = append(resp, readResponse{
			ID:          val.ID,
			Name:        val.Name,
			Description: desc,
			Time:        val.CreatedAt.Format(time.RFC1123),
			URL:         fmt.Sprintf("/api/v1/file/tutorial/%d", val.ID),
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// ReadDetailHandler ...
func ReadDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleSchedule, rg.RoleXRead, rg.RoleRead) {
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
		Time:        tutorial.CreatedAt.Format(time.RFC1123),
		URL:         fmt.Sprintf("/api/v1/file/tutorial/%d", tutorial.ID),
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}
