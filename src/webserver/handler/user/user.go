package user

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/email"
	"github.com/melodiez14/meiko/src/module/module"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/alias"
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
	email.NewRequest(us.Email, "Reset Password").Deliver()
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

func SignUpHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u := r.Context().Value("User").(*auth.User)
	if u != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

	param := signUpParams{
		ID:       r.FormValue("id"),
		Name:     r.FormValue("name"),
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

	us, err := user.GetUserByEmail(args.Email)
	if (err != nil || us != nil) && err != sql.ErrNoRows {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%s has been registered", args.Email)))
		return
	}

	id, err := user.GetUserByID(args.ID)
	if (err != nil || id != nil) && err != sql.ErrNoRows {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%d has been registered!", args.ID)))
		return
	}

	user.InsertNewUser(args.ID, args.Name, args.Email, args.Password)

	// send code activation  to email
	g, err := user.GenerateVerification(args.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Server error"))
		return
	}

	// change to email template
	email.NewRequest(args.Email, fmt.Sprintf("Reset Password %d", g.Code)).Deliver()

	// for debugging purpose
	fmt.Println(g.Code)

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Sign Up success"))
	return

}

func GetValidatedUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	// need confirmation
	if !sess.IsHasRoles(alias.ModuleUser, alias.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	u, _ := user.GetByStatus(alias.UserStatusVerified)

	res := []getVerifiedUserResponse{}
	for _, val := range u {
		res = append(res, getVerifiedUserResponse{
			Name:  val.Name,
			Email: val.Email,
			ID:    val.ID,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

func RequestVerifiedUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if sess != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}
	param := setStatusUserParams{
		Email: r.FormValue("email"),
		Code:  r.FormValue("code"),
	}
	args, err := param.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	_, err = user.GetUserByEmail(args.Email)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid email"))
		return
	}
	v := user.IsValidConfirmationCode(args.Email, args.Code)
	if !v {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid confirmation code"))
		return
	}
	go user.UpdateCodeUser(args.Email, alias.UserStatusVerified)

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage(fmt.Sprintf("Your account with this %s is being Verified ", args.Email)).
		SetCode(http.StatusOK))
	return

}
func LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := auth.DestroySession(r, w)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			SetMessage("Internal server error"))
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Logout Sukses"))
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
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid email or password"))
		return
	}

	// check whether user activated
	switch u.Status {
	case alias.UserStatusUnverified:
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			SetMessage("Email unactivated"))
		return
	case alias.UserStatusVerified:
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Waiting for admin approval"))
		return
	case alias.UserStatusActivated:
		break
	default:
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Internal Error"))
		return
	}

	roles := make(map[string][]string)
	if u.RoleGroupsID.Valid {
		roles = module.GetPriviegeByRoleGroupID(u.RoleGroupsID.Int64)
	}

	s := auth.User{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Gender:  u.Gender,
		College: u.College,
		Note:    u.Note,
		Status:  u.Status,
		Roles:   roles,
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
