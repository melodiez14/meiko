package attendance

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	atd "github.com/melodiez14/meiko/src/module/attendance"
	cs "github.com/melodiez14/meiko/src/module/course"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	usr "github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// u := r.Context().Value("User").(*auth.User)

	// courses, err := course.GetByUserID(u.ID)
	// if err != nil {
	// 	template.RenderJSONResponse(w, new(template.Response).
	// 		SetCode(http.StatusInternalServerError).
	// 		AddError(err.Error()))
	// 	return
	// }

	// var res []summaryResponse
	// var percentage float32
	// for _, c := range courses {
	// 	a, err := attendance.GetByUserCourseID(u.ID, c.ID)
	// 	if err != nil {
	// 		template.RenderJSONResponse(w, new(template.Response).
	// 			SetCode(http.StatusInternalServerError).
	// 			AddError(err.Error()))
	// 		return
	// 	}

	// 	if len(a) > 0 {
	// 		percentage = (float32(len(a)) * 100) / float32(len(a))
	// 	} else {
	// 		percentage = 0
	// 	}

	// 	res = append(res, summaryResponse{
	// 		Course:     c.Name,
	// 		Percentage: fmt.Sprintf("%.4g%%", percentage),
	// 	})
	// }

	// template.RenderJSONResponse(w, new(template.Response).
	// 	SetCode(http.StatusOK).
	// 	SetData(res))
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK))
	return
}

// not functional yet
func ListStudentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAttendance, rg.RoleXRead, rg.RoleRead) && !sess.IsHasRoles(rg.ModuleUser, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := listStudentParams{
		meetingNumber: r.FormValue("meeting_number"),
		scheduleID:    r.FormValue("schedule_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	// check is valid meeting number and schedule id
	_, err = atd.GetMeeting(args.meetingNumber, args.scheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	studentID, err := cs.SelectEnrolledStudentID(args.scheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var resp []listStudentResponse
	if len(studentID) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
		return
	}

	users, err := usr.SelectByID(studentID, true, usr.ColID, usr.ColName, usr.ColIdentityCode)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	// need to get attendances table per meetings number
	for _, val := range users {
		resp = append(resp, listStudentResponse{
			IdentityCode: val.IdentityCode,
			StudentName:  val.Name,
			Status:       "absent",
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}
