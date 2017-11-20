package handler

import (
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Hello from meiko. Have a nice day! :)"))
	return
}

// IsValidFileID ...
func IsValidFileID(str string) bool {
	valid, _ := regexp.MatchString(`^[0-9.]{26,32}$`, str)
	return valid
}
