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

// SignUpHandler handles the http request for first registration process
/*
	@params:
		user_id	= required, numeric, characters=12
		name	= required, alphaspace, 0<characters<50
		email	= required, email format, 0<characters<45
		password= required, minimum 1 uppercase, lowercase, numeric, characters>=6
	@example:
		id=140810140016
		name=Risal Falah
		email=risal.falah@gmail.com
		password=Qwerty1
	@return
*/
func SignUpHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if sess != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

	params := signUpParams{
		ID:       r.FormValue("user_id"),
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	_, err = user.Get(user.ColEmail).
		Where(user.ColEmail, user.OperatorEquals, args.Email).
		Exec()
	if err == nil || (err != nil && err != sql.ErrNoRows) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%s has been registered", args.Email)))
		return
	}

	_, err = user.Get(user.ColID).
		Where(user.ColEmail, user.OperatorEquals, args.ID).
		Exec()
	if err == nil || (err != nil && err != sql.ErrNoRows) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%d has been registered!", args.ID)))
		return
	}

	user.InsertNewUser(args.ID, args.Name, args.Email, args.Password)

	// send code activation to email
	verification, err := user.GenerateVerification(args.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Server error"))
		return
	}

	// change to email template
	email.NewRequest(args.Email, fmt.Sprintf("Reset Password %d", verification.Code)).Deliver()

	// for debugging purpose
	fmt.Println(verification.Code)

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("SignUp success"))
	return

}

// EmailActivationHandler handles the http request for resend activation code or activate the email
/*
	@params:
		email	= required, email format, 0<characters<45
		resend	= optional, value=true or empty
		code	= required if resend is empty, numeric, characters=4
	@example:
		email=risal.falah@gmail.com
		resend=true
		code=1234 or empty if resend is true
	@return
*/
func EmailVerificationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if sess != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

	params := emailVerificationParams{
		Email:        r.FormValue("email"),
		IsResendCode: r.FormValue("resend"),
		Code:         r.FormValue("code"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	u, err := user.Get(user.ColID).
		Where(user.ColEmail, user.OperatorEquals, args.Email).
		AndWhere(user.ColStatus, user.OperatorEquals, alias.UserStatusUnverified).
		Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid email or has been verified"))
		return
	}

	if args.IsResendCode {
		// generate verification code
		verification, err := user.GenerateVerification(u.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError).
				AddError("Server error"))
			return
		}

		// change to email template
		email.NewRequest(args.Email, fmt.Sprintf("Reset Password %d", verification.Code)).Deliver()

		// for debugging purpose
		fmt.Println(verification.Code)

		template.RenderJSONResponse(w, new(template.Response).
			SetMessage(fmt.Sprintf("Code has been sent to email")).
			SetCode(http.StatusOK))
		return
	}

	valid := user.IsValidConfirmationCode(args.Email, args.Code)
	if !valid {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid confirmation code"))
		return
	}

	go user.SetStatus(args.Email, alias.UserStatusVerified)

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage(fmt.Sprintf("Your account %s is being activated by admin", args.Email)).
		SetCode(http.StatusOK))
	return
}

// ReadUserHandler handles the http request for listing all verified and activated users. Accessing this handler needs XREAD ability
/*
	@params:
		pg	= required, positive numeric
		ttl	= required, positive numeric
	@example:
		pg=1
		ttl=10
	@return
		[]{name, email, status, user_id}
*/
func ReadUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	if !sess.IsHasRoles(alias.ModuleUser, alias.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := getVerifiedParams{
		Page:  r.FormValue("pg"),
		Total: r.FormValue("ttl"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(err.Error()))
		return
	}

	// get verified user by page
	offset := (args.Page - 1) * args.Total
	u, _ := user.Select(user.ColID, user.ColName, user.ColEmail, user.ColStatus).
		Where(user.ColStatus, user.OperatorEquals, alias.UserStatusVerified).
		OrWhere(user.ColStatus, user.OperatorEquals, alias.UserStatusActivated).
		Limit(args.Total).
		Offset(offset).
		Exec()

	var status string
	res := []getVerifiedResponse{}
	for _, val := range u {
		if val.Status == alias.UserStatusActivated {
			status = "active"
		} else {
			status = "inactive"
		}
		res = append(res, getVerifiedResponse{
			Name:   val.Name,
			Email:  val.Email,
			ID:     val.ID,
			Status: status,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

// ActivationHandler handles the http request for change user status to activated or verified. Accessing this handler needs XUPDATE ability
/*
	@params:
		user_id	= required, numeric, characters=12
		status	= required, string
	@example:
		user_id = 140810140016
		status	= active or inactive
	@return
*/
func ActivationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	if !sess.IsHasRoles(alias.ModuleUser, alias.RoleXUpdate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := activationParams{
		ID:     r.FormValue("user_id"),
		Status: r.FormValue("status"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	var oldStatus int8
	switch args.Status {
	case alias.UserStatusVerified:
		oldStatus = alias.UserStatusActivated
	case alias.UserStatusActivated:
		oldStatus = alias.UserStatusVerified
	}

	u, err := user.Get().
		Where(user.ColID, user.OperatorEquals, args.ID).
		AndWhere(user.ColStatus, user.OperatorEquals, oldStatus).
		Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	go func() {
		user.SetStatus(u.Email, args.Status)

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
			LineID:  u.LineID.String,
			Phone:   u.Phone.String,
			Roles:   roles,
		}

		s.UpdateSession(w)
	}()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Status Updated"))
	return
}

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

func GetUserAccountHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	val, err := user.GetUserByID(sess.ID)

	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%d has been registered!", sess.ID)))
		return
	}

	s := setUserAccoutArgs{
		Name:    val.Name,
		Gender:  val.Gender,
		Phone:   val.Phone.String,
		LineID:  val.LineID.String,
		College: val.College,
		Note:    val.College,
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(s))
	return

}

// UpdateUserAccountHandler
func UpdateUserAccountHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	param := setUserAccoutParams{
		Name:    r.FormValue("name"),
		Gender:  r.FormValue("gender"),
		Phone:   r.FormValue("phone"),
		LineID:  r.FormValue("line_id"),
		College: r.FormValue("College"),
		Note:    r.FormValue("note"),
	}

	args, err := param.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	user.SetUpdateUserAccount(args.Name, args.Phone, args.LineID, args.College, args.Note, args.Gender, sess.ID)

	u, err := user.GetUserByID(sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Internal error"))
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

	s.UpdateSession(w)

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage("Data updated").
		SetCode(http.StatusOK))
	return
}
func ChangePasswordHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := r.Context().Value("User").(*auth.User).ID
	param := setChangePasswordParams{
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirmPassword"),
	}
	args, err := param.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	go user.SetChangePassword(args.Password, id)
	template.RenderJSONResponse(w, new(template.Response).
		SetMessage("Password has changed").
		SetCode(http.StatusOK))
	return
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	err := sess.DestroySession(r, w)

	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			SetMessage("Internal server error"))
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Logout success"))
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
			AddError("Internal server error"))
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
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Login success"))
	return
}
