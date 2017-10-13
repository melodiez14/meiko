package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Hello from meiko Hello from meiko Hello from meikoHello from meiko Hello from meiko Hello from meiko Hello from meiko Hello from meiko Hello from meiko Hello from meiko"))
	return
}
