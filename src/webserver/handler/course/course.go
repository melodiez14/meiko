package course

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/course"
	"github.com/melodiez14/meiko/src/util/alias"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u := r.Context().Value("User").(*auth.User)
	fmt.Println(u)
	c, err := course.GetByUserID(u.ID)
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
