package handler

import (
	"net/http"

	"github.com/KhaAzAs/meiko/src/webserver/template"
	"github.com/julienschmidt/httprouter"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Hello from meiko Hello from meiko Hello from meikoHello from meiko Hello from meiko Hello from meiko Hello from meiko Hello from meiko Hello from meiko Hello from meiko"))
	return
}
