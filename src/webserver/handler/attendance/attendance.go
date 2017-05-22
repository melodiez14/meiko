package attendance

import (
	"net/http"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/attendance"
	"github.com/melodiez14/meiko/src/module/course"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u := r.Context().Value("User").(*user.User)

	courses, err := course.GetByUserID(u.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError(err.Error()))
		return
	}

	var res []summaryResponse
	var percentage float32
	for _, c := range courses {
		a, err := attendance.GetByUserCourseID(u.ID, c.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError).
				AddError(err.Error()))
			return
		}

		if len(a) > 0 {
			percentage = (float32(len(a)) * 100) / float32(len(a))
		} else {
			percentage = 0
		}

		res = append(res, summaryResponse{
			Course:     c.Name,
			Percentage: fmt.Sprintf("%.4g%%", percentage),
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}
