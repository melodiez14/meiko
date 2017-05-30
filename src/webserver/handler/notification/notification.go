package notification

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/notification"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	param := getNotificationParam{
		page: r.FormValue("_pg"),
	}

	args, err := param.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	const limit = uint8(10)
	u := r.Context().Value("User").(*auth.User)

	n, err := notification.Get(u.ID, args.page, limit)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError(err.Error()))
		return
	}

	var res []Notification
	for _, v := range n {
		res = append(res, Notification{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			URL:         v.GetURL(),
			CreatedAt:   v.CreatedAt.Unix(),
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}
