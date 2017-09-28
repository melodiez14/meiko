package user

import (
	"fmt"
	"net/http"

	"github.com/melodiez14/meiko/src/util/helper"

	"database/sql"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/email"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/alias"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// SignUpHandler handles the http request for first registration process
/*
	@params:
		identity= required, numeric, characters=12
		name	= required, alphaspace, 0<characters<50
		email	= required, email format, 0<characters<45
		password= required, minimum 1 uppercase, lowercase, numeric, characters>=6
	@example:
		identity=140810140016
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
		IdentityCode: r.FormValue("id"),
		Name:         r.FormValue("name"),
		Email:        r.FormValue("email"),
		Password:     r.FormValue("password"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// check is email registered
	_, err = user.Get(user.ColEmail).
		Where(user.ColEmail, user.OperatorEquals, args.Email).
		Exec()
	if err == nil || (err != nil && err != sql.ErrNoRows) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%s has been registered", args.Email)))
		return
	}

	// check is id registered
	_, err = user.Get(user.ColIdentityCode).
		Where(user.ColIdentityCode, user.OperatorEquals, args.IdentityCode).
		Exec()
	if err == nil || (err != nil && err != sql.ErrNoRows) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%d has been registered!", args.IdentityCode)))
		return
	}

	// insert new user
	err = user.Insert(map[string]interface{}{
		user.ColIdentityCode: args.IdentityCode,
		user.ColName:         args.Name,
		user.ColEmail:        args.Email,
		user.ColPassword:     args.Password,
	}).Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Internal server error"))
		return
	}

	// send code activation to email
	verification, err := user.GenerateVerification(args.IdentityCode)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Internal server error"))
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

// EmailVerificationHandler handles the http request for resend activation code or activate the email
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

	u, err := user.Get(user.ColIdentityCode).
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
		verification, err := user.GenerateVerification(u.IdentityCode)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError).
				AddError("Internal server error"))
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

	go user.Update(map[string]interface{}{
		user.ColStatus: alias.UserStatusVerified,
	}).Where(user.ColEmail, user.OperatorEquals, args.Email).Exec()

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage(fmt.Sprintf("Your account %s is being activated by admin", args.Email)).
		SetCode(http.StatusOK))
	return
}

// ReadUserHandler handles the http request for listing all verified and activated users. Accessing this handler needs READ or XREAD ability
/*
	@params:
		pg	= required, positive numeric
		ttl	= required, positive numeric
	@example:
		pg=1
		ttl=10
	@return
		[]{name, email, status, identity}
*/
func ReadUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	if !sess.IsHasRoles(rg.ModuleUser, rg.RoleRead, rg.RoleXRead) {
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
	u, _ := user.Select(user.ColIdentityCode, user.ColName, user.ColEmail, user.ColStatus).
		Where(user.ColStatus, user.OperatorEquals, alias.UserStatusVerified).
		OrWhere(user.ColStatus, user.OperatorEquals, alias.UserStatusActivated).
		AndWhere(user.ColID, user.OperatorUnquals, sess.ID).
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
			Name:         val.Name,
			Email:        val.Email,
			IdentityCode: val.IdentityCode,
			Status:       status,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

// ActivationHandler handles the http request for changing user status to activated or verified. Accessing this handler needs UPDATE or XUPDATE ability
/*
	@params:
		identity	= required, numeric, characters=12
		status		= required, string
	@example:
		identity	= 140810140016
		status		= active or inactive
	@return
*/
func ActivationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	if !sess.IsHasRoles(rg.ModuleUser, rg.RoleUpdate, rg.RoleXUpdate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := activationParams{
		IdentityCode: r.FormValue("id"),
		Status:       r.FormValue("status"),
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
		Where(user.ColIdentityCode, user.OperatorEquals, args.IdentityCode).
		AndWhere(user.ColStatus, user.OperatorEquals, oldStatus).
		AndWhere(user.ColID, user.OperatorUnquals, sess.ID).
		Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	go func() {
		// change if args.Status == activated update redis
		// if args.Status == Verified delete redis
		_ = user.Update(map[string]interface{}{
			user.ColStatus: args.Status,
		}).Where(user.ColEmail, user.OperatorEquals, u.Email).Exec()

		roles := make(map[string][]string)
		if u.RoleGroupsID.Valid {
			roles = rg.GetModuleAccess(u.RoleGroupsID.Int64)
		}

		s := auth.User{
			ID:           u.ID,
			Name:         u.Name,
			Email:        u.Email,
			Gender:       u.Gender,
			Note:         u.Note,
			Status:       u.Status,
			IdentityCode: u.IdentityCode,
			LineID:       u.LineID.String,
			Phone:        u.Phone.String,
			Roles:        roles,
		}

		s.UpdateSession(w)
	}()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Status Updated"))
	return
}

// SignInHandler handles the http request for create a new session of user
/*
	@params:
		email	= required, email format, 0<characters<45
		password= required, minimum 1 uppercase, lowercase, numeric, characters>=6
	@example:
		email	= risal.falah@gmail.com
		password= Qwerty123
	@return
*/
func SignInHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if sess != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

	params := signInParams{
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

	u, err := user.Get().
		Where(user.ColEmail, user.OperatorEquals, args.Email).
		AndWhere(user.ColPassword, user.OperatorEquals, args.Password).
		Exec()
	if err != nil {
		fmt.Println(err.Error())
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
		roles = rg.GetModuleAccess(u.RoleGroupsID.Int64)
	}

	s := auth.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Gender:       u.Gender,
		Note:         u.Note,
		Status:       u.Status,
		IdentityCode: u.IdentityCode,
		LineID:       u.LineID.String,
		Phone:        u.Phone.String,
		Roles:        roles,
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

// ForgotHandler handles the http request for create a new session of user
// If resend is true so only email and resend is used. It used for requesting the code to send to email
// If resend is empty so code can't be empty
// If resend is empty, code is not empty, and password is not empty, it will set the new password
/*
	@params:
		email	= required, email format, 0<characters<45
		resend	= optional, value=true or empty
		code	= required if resend is empty, numeric, characters=4
		password= optional if code is empty, minimum 1 uppercase, lowercase, numeric, characters>=6
	@example:
		email=risal.falah@gmail.com
		resend=true
		code=1234 or empty if resend is true
		password= Qwerty123
	@return
*/
func ForgotHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if sess != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError("You have already logged in"))
		return
	}

	params := forgotParams{
		Email:      r.FormValue("email"),
		IsSendCode: r.FormValue("resend"),
		Code:       r.FormValue("code"),
		Password:   r.FormValue("password"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusFound).
			AddError(err.Error()))
		return
	}

	// if send code to email then return
	if args.IsSendCode {
		u, err := user.Get(user.ColIdentityCode).
			Where(user.ColEmail, user.OperatorEquals, args.Email).
			Exec()
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("Email is not registered"))
			return
		}

		// generate verification code
		verification, err := user.GenerateVerification(u.IdentityCode)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError).
				AddError("Internal server error"))
			return
		}

		// change to email template
		email.NewRequest(args.Email, fmt.Sprintf("Reset Password %d", verification.Code)).Deliver()

		// for debugging purpose
		fmt.Println(verification.Code)

		res := forgotResponse{
			Email:          args.Email,
			ExpireDuration: verification.ExpireDuration,
			MaxAttempt:     3,
		}

		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(res))
		return
	}

	v := user.IsValidConfirmationCode(args.Email, args.Code)
	if !v {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid confirmation code"))
		return
	}

	if helper.IsEmpty(args.Password) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK))
		return
	}

	go user.UpdateNewPassword(args.Email, args.Password)

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage("New password has been updated").
		SetCode(http.StatusOK))
	return
}

// GetProfileHandler handles the http request for listing all profile information
/*
	@params:
	@example:
	@return
		npm			= 140810140016
		name 		= Risal Falah
		email		= risal.falah@gmail.com
		gender 		= male or female
		phone 		= 085860141146
		line_id 	= risalfa
		about_me	= hello my name is risal falah, you can call me ical
		img			= /api/v1/img/pl
		img_t		= /api/v1/img/pl_t
*/
func GetProfileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	u, err := user.Get().Where(user.ColID, user.OperatorEquals, sess.ID).Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(fmt.Sprintf("%d has been registered!", sess.ID)))
		return
	}

	var gender string
	switch u.Gender {
	case alias.UserGenderMale:
		gender = "male"
	case alias.UserGenderFemale:
		gender = "female"
	}

	res := getProfileResponse{
		Name:                  u.Name,
		Email:                 u.Email,
		Gender:                gender,
		Phone:                 u.Phone.String,
		IdentityCode:          u.IdentityCode,
		LineID:                u.LineID.String,
		Note:                  u.Note,
		ImageProfile:          alias.URLProfile,
		ImageProfileThumbnail: alias.URLProfileThumbnail,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return

}

// UpdateProfileHandler handles the http request for updating the user profile
/*
	@params:
		id		= required, numeric, 10<=characters<=18
		name	= required, alphaspace, 0<characters<=50
		email	= required, email format
		gender	= optional, male or female
		phone	= optional, numeric, 10<=characters<=12
		line_id	= optional, 0<characters<=45
		note	= optional, 0<characters<=100
	@example:
		name=Risal Falah
		gender=male
		phone=085860141146
		line_id=risalfa
		note=Hello my name is risal
	@return
*/
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := updateProfileParams{
		IdentityCode: r.FormValue("id"),
		Name:         r.FormValue("name"),
		Email:        r.FormValue("email"),
		Gender:       r.FormValue("gender"),
		Phone:        r.FormValue("phone"),
		LineID:       r.FormValue("line_id"),
		Note:         r.FormValue("about_me"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if args.IdentityCode != sess.IdentityCode || args.Email != sess.Email {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	if args.Phone.Valid {
		_, err := user.Get(user.ColPhone).
			Where(user.ColPhone, user.OperatorEquals, args.Phone.String).
			AndWhere(user.ColID, user.OperatorUnquals, sess.ID).
			Exec()
		if err == nil || (err != nil && err != sql.ErrNoRows) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusForbidden).
				AddError(fmt.Sprintf("phone %s has been registered!", args.Phone.String)))
			return
		}
	}

	if args.LineID.Valid {
		_, err := user.Get(user.ColLineID).
			Where(user.ColLineID, user.OperatorEquals, args.LineID.String).
			AndWhere(user.ColID, user.OperatorUnquals, sess.ID).
			Exec()
		if err == nil || (err != nil && err != sql.ErrNoRows) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusForbidden).
				AddError(fmt.Sprintf("line id %s has been registered!", args.LineID.String)))
			return
		}
	}

	err = user.Update(map[string]interface{}{
		user.ColName:   args.Name,
		user.ColPhone:  args.Phone,
		user.ColLineID: args.LineID,
		user.ColNote:   args.Note,
		user.ColGender: args.Gender,
	}).Where(user.ColID, user.OperatorEquals, sess.ID).Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Internal server error"))
		return
	}

	u, err := user.Get().Where(user.ColID, user.OperatorEquals, sess.ID).Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Internal server error"))
		return
	}

	roles := make(map[string][]string)
	if u.RoleGroupsID.Valid {
		roles = rg.GetModuleAccess(u.RoleGroupsID.Int64)
	}

	s := auth.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Gender:       u.Gender,
		Note:         u.Note,
		Status:       u.Status,
		IdentityCode: u.IdentityCode,
		LineID:       u.LineID.String,
		Phone:        u.Phone.String,
		Roles:        roles,
	}

	s.UpdateSession(w)

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage("Data updated").
		SetCode(http.StatusOK))
	return
}

// ChangePasswordHandler handles the http request for updating user password
/*
	@params:
		id						= required, numeric, characters=12
		email					= required, email format, 0<characters<45
		old_password			= required, minimum 1 uppercase, lowercase, numeric, characters>=6
		password				= required, minimum 1 uppercase, lowercase, numeric, characters>=6
		password_confirmation	= required, should be same as password
	@example:
		npm=140810140016
		email=risal.falah@gmail.com
		old_password			= Qwerty123
		password				= Qwerty321
		password_confirmation	= Qwerty321
	@return
*/
func ChangePasswordHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := changePasswordParams{
		IdentityCode:    r.FormValue("id"),
		Email:           r.FormValue("email"),
		OldPassword:     r.FormValue("old_password"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("password_confirmation"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if args.IdentityCode != sess.IdentityCode || args.Email != sess.Email {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	// Update new password
	err = user.Update(map[string]interface{}{user.ColPassword: args.Password}).
		Where(user.ColID, user.OperatorEquals, sess.ID).
		AndWhere(user.ColPassword, user.OperatorEquals, args.OldPassword).
		Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Incorect old password!"))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetMessage("Password has been changed").
		SetCode(http.StatusOK))
	return
}

// SignOutHandler handles the http request for destroying the session
/*
	@params:
	@example:
	@return
*/
func SignOutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
