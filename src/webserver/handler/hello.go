package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func HelloMeiko(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	test := "Hello Meiko"
	template.RenderJSONResponse(w, new(template.Response).SetCode(200).SetData(test).SetMessage("Ini Pesan"))
}
