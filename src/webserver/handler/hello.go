package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func HelloMeiko(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userData := r.Context().Value("User").(*user.User)
	if userData != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(userData))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Masuk Tanpa Cookie"))
}
