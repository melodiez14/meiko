package user

import (
	"html"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	search := html.EscapeString(r.FormValue("q"))
	users, err := user.Search(search)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("You don't have privilege"))
		return
	}

	response := []searchResponse{}
	for _, user := range users {
		response = append(response, searchResponse{
			IdentityCode: user.IdentityCode,
			Name:         user.Name,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(response))
	return
}
