package user

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func ForgotRequestHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u := r.Context().Value("User").(*auth.User)
	if u != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

	param := forgotRequestParams{
		Email: r.FormValue("email"),
	}

	args, err := param.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	us, err := user.GetUserByEmail(args.Email)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid email"))
		return
	}

	v, err := user.GenerateVerification(us.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Server error"))
		return
	}

	// change to email template
	// email.NewRequest(us.Email, "Reset Password").Deliver()
	// for debugging purposes
	fmt.Println(v.Code)

	res := forgotRequestResponse{
		Email:          us.Email,
		ExpireDuration: v.ExpireDuration,
		MaxAttempt:     3,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

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

	v := user.IsValidUserLogin(args.Email, args.Password)
	if !v {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid email or password"))
		return
	}

	// then set session to redis
	// ===== here ====

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Login success"))
	return
}
