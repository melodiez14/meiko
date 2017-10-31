package assignment

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

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

func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleCreate, rg.RoleXCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

}
