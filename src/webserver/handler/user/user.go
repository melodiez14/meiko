package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := signInParams{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	args, err := param.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	err = user.IsValidUserLogin(args.Email, args.Password)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(err.Error()))
		return
	}

	// then set session to redis
	// ===== here ====

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Login success"))
	return
}
