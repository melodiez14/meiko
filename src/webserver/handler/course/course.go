package course

import (
	"database/sql"
	"net/http"

	"github.com/melodiez14/meiko/src/util/conn"

	"github.com/melodiez14/meiko/src/util/helper"

	"github.com/julienschmidt/httprouter"
	cs "github.com/melodiez14/meiko/src/module/course"
	pl "github.com/melodiez14/meiko/src/module/place"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/alias"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// CreateHandler handles the http request for creating the course. Accessing this handler needs CREATE or XCREATE ability
/*
	@params:
		name		= required, alphabet and space only
		description	= optional
		ucu			= required, positive numeric
		semester	= required, positive numeric
		start_time	= required, positive numeric, minutes
		end_time	= required, positive numeric, minutes
		class		= required, character=1
		place		= required
	@example:
		name		= Sistem Informasi Multimedia
		description	= Praktikum ini membahas mengenai Sistem Informasi Multimedia
		ucu			= 3
		semester	= 1
		start_time	= 600
		end_time	= 800
		class		= A
		place		= UDJT-102
	@return
*/
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleCourse, rg.RoleCreate, rg.RoleXCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := createParams{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		UCU:         r.FormValue("ucu"),
		Semester:    r.FormValue("semester"),
		StartTime:   r.FormValue("start_time"),
		EndTime:     r.FormValue("end_time"),
		Class:       r.FormValue("class"),
		Day:         r.FormValue("day"),
		PlaceID:     r.FormValue("place"),
	}

	args, err := params.Validation()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	_, err = cs.Get(cs.ColID).
		Where(cs.ColDay, cs.OperatorEquals, args.Day).
		AndWhere(cs.ColClass, cs.OperatorEquals, args.Class).
		AndWhere(cs.ColSemester, cs.OperatorEquals, args.Semester).
		AndWhere(cs.ColStartTime, cs.OperatorEquals, args.StartTime).
		Exec()
	if err == nil || (err != nil && err != sql.ErrNoRows) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Course already exists"))
		return
	}

	tx := conn.DB.MustBegin()
	// validate place, create place if not exist
	if !pl.IsExistID(args.PlaceID) {
		err = pl.Place{
			ID: args.PlaceID,
		}.Insert(tx)
		if err != nil {
			_ = tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	err = cs.Insert(map[string]interface{}{
		cs.ColName:        args.Name,
		cs.ColDescription: args.Description,
		cs.ColUCU:         args.UCU,
		cs.ColSemester:    args.Semester,
		cs.ColStartTime:   args.StartTime,
		cs.ColEndTime:     args.EndTime,
		cs.ColStatus:      cs.StatusActive,
		cs.ColClass:       args.Class,
		cs.ColDay:         args.Day,
		cs.ColPlaceID:     args.PlaceID,
		cs.ColCreatedBy:   sess.ID,
	}).Exec(tx)
	if err != nil {
		_ = tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	err = tx.Commit()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Success"))
	return
}

// ReadHandler handles the http request for creating the course. Accessing this handler needs CREATE or XCREATE ability
/*
	@params:
		pg	= required, positive numeric
		ttl	= required, positive numeric
	@example:
		pg=1
		ttl=10
	@return
		[]{name, email, status, identity}
*/
func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleCourse, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readParams{
		Page:  r.FormValue("pg"),
		Total: r.FormValue("ttl"),
	}

	args, err := params.Validation()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid request"))
		return
	}

	offset := (args.Page - 1) * args.Total
	courses, err := cs.Select(cs.ColID, cs.ColName, cs.ColClass, cs.ColStartTime, cs.ColEndTime, cs.ColDay, cs.ColStatus, cs.ColPlaceID).
		Where(cs.ColStatus, cs.OperatorUnquals, cs.StatusDeleted).
		Limit(args.Total).
		Offset(offset).
		Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Internal server error"))
		return
	}

	var status string
	var res []readResponse
	for _, val := range courses {

		if val.Status == cs.StatusActive {
			status = "active"
		} else {
			status = "inactive"
		}

		res = append(res, readResponse{
			ID:        val.ID,
			Name:      val.Name,
			Class:     val.Class,
			StartTime: helper.MinutesToTimeString(val.StartTime),
			EndTime:   helper.MinutesToTimeString(val.EndTime),
			Day:       helper.IntDayToString(val.Day),
			Status:    status,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u := r.Context().Value("User").(*auth.User)
	c, err := cs.GetByUserID(u.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError(err.Error()))
		return
	}

	activeCourse := []courseResponse{}
	inactiveCourse := []courseResponse{}

	for _, v := range c {
		cres := courseResponse{
			ID:       v.ID,
			Name:     v.Name,
			UCU:      v.UCU,
			Semester: v.Semester,
		}

		if v.Status == alias.CourseActive {
			activeCourse = append(activeCourse, cres)
		} else {
			inactiveCourse = append(inactiveCourse, cres)
		}
	}

	sres := []summaryResponse{
		summaryResponse{
			Status: "Active",
			Course: activeCourse,
		},
		summaryResponse{
			Status: "Inactive",
			Course: inactiveCourse,
		},
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(sres))
	return
}
