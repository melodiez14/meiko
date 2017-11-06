package assignment

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	as "github.com/melodiez14/meiko/src/module/assignment"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// CreateHandler function is
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleCreate, rg.RoleXCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := createParams{
		FilesID:           r.FormValue("id"),
		GradeParametersID: r.FormValue("grade_parameter"),
		Name:              r.FormValue("name"),
		Description:       r.FormValue("description"),
		Status:            r.FormValue("status"),
		DueDate:           r.FormValue("due_date"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// is grade_parameter exist
	//
	if !as.IsExistByGradeParameterID(args.GradeParametersID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Grade parameters id does not exist!"))
		return
	}
	// Insert to table assignments
	//
	tx := conn.DB.MustBegin()
	TableID, err := as.Insert(
		args.GradeParametersID,
		args.Name,
		args.Status,
		args.Description,
		args.DueDate,
		tx,
	)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	// Files null
	//
	if args.FilesID == "" {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetMessage("Success Without files"))
		return
	}
	// Wrong file code
	//
	if !as.IsFileIDExist(args.FilesID) {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			SetMessage("Wrong file code!"))
		return
	}
	err = as.UpdateFiles(args.FilesID, TableID, tx)
	if err != nil {
		tx.Rollback()
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
		SetMessage("Success!"))
	return

}

// GetAllAssignmentHandler func is ...
func GetAllAssignmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := readParams{
		Page:  r.FormValue("pg"),
		Total: r.FormValue("ttl"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid request"))
		return
	}
	offset := (args.Page - 1) * args.Total
	assignments, err := as.SelectByPage(args.Total, offset)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	var status string
	var res []readResponse
	for _, val := range assignments {

		if val.Assignment.Status == as.StatusAssignmentActive {
			status = "active"
		} else {
			status = "inactive"
		}

		res = append(res, readResponse{
			Name:             val.Assignment.Name,
			Description:      val.Assignment.Description,
			Status:           status,
			DueDate:          val.Assignment.DueDate,
			GradeParameterID: val.Assignment.GradeParameterID,
		})
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return

}

// DetailHandler func is ...
func DetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := detailParams{
		IdentityCode: ps.ByName("id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	u, err := as.GetByAssignementID(args.IdentityCode)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}

	var status string
	switch u.Assignment.Status {
	case 0:
		status = "inactive"
	case 1:
		status = "active"
	}

	res := detailResponse{
		ID:               u.Assignment.ID,
		Status:           status,
		Name:             u.Assignment.Name,
		GradeParameterID: u.Assignment.GradeParameterID,
		Description:      u.Assignment.Description,
		DueDate:          u.Assignment.DueDate,
		Mime:             u.File.Mime,
		Percentage:       u.GradeParameter.Percentage,
	}
	_ = res

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

// func GetIncompleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	u := r.Context().Value("User").(*auth.User)

// 	a, err := assignment.GetIncompleteByUserID(u.ID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError).
// 			AddError(err.Error()))
// 		return
// 	}

// 	res := []summaryResponse{}
// 	for _, v := range a {
// 		res = append(res, summaryResponse{
// 			ID:   v.ID,
// 			Name: v.Name,
// 		})
// 	}

// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return
// }

// func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	var res []profileSummaryResponse
// 	u := r.Context().Value("User").(*auth.User)

// 	// get all enrolled course using using userID
// 	courses, err := course.GetByUserID(u.ID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError).
// 			AddError(err.Error()))
// 		return
// 	}

// 	// if there is no enrolled course
// 	if len(courses) < 1 {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusOK).
// 			SetData(res))
// 	}

// 	// get completed assignments have been posted in all courses
// 	ca, err := assignment.GetCompleteByUserID(u.ID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError).
// 			AddError(err.Error()))
// 		return
// 	}

// 	// iterate all courses to get the summary
// 	for _, v := range courses {

// 		pSummary := profileSummaryResponse{
// 			CourseName: v.Name,
// 			Complete:   0,
// 			Incomplete: 0,
// 		}

// 		// get all assignments per courses
// 		assignments, err := assignment.GetByCourseID(v.ID)
// 		if err != nil {
// 			template.RenderJSONResponse(w, new(template.Response).
// 				SetCode(http.StatusInternalServerError).
// 				AddError(err.Error()))
// 			return
// 		}

// 		// compare course assignments with all assignments in p_users_assignments
// 		// if course assignment id exist in p_users_assignments then increment the complete
// 		// else increment the incomplete
// 		for _, a := range assignments {
// 			if helper.Int64InSlice(a.ID, ca) {
// 				pSummary.Complete++
// 			} else {
// 				pSummary.Incomplete++
// 			}
// 		}

// 		// append summary per courses
// 		res = append(res, pSummary)
// 	}

// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return
// }
