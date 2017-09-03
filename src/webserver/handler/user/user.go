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

func ForgotConfirmation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u := r.Context().Value("User").(*auth.User)
	if u != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

	param := forgotConfirmationParams{
		Email:    r.FormValue("email"),
		Code:     r.FormValue("code"),
		Password: r.FormValue("password"),
	}

	args, err := param.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError(err.Error()))
		return
	}

	v := user.IsValidConfirmationCode(args.Email, args.Code)
	if !v {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid confirmation code"))
		return
	}

	if len(args.Password) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK))
		return
	}

	go user.SetNewPassword(args.Email, args.Password)

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage("New password has been updated").
		SetCode(http.StatusOK))
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if sess != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

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

	u, err := user.GetUserLogin(args.Email, args.Password)
	fmt.Println(u)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid email or password"))
		return
	}

	s := auth.User{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Gender:  u.Gender,
		College: u.College,
		Note:    u.Note,
		Status:  u.Status,
	}

	err = s.SetSession(w)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			SetMessage("Internal server error"))
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Login success"))
	return
}
