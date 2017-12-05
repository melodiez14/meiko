package file

import (
	"fmt"
	"html"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/melodiez14/meiko/src/util/alias"

	"github.com/melodiez14/meiko/src/util/conn"

	"github.com/disintegration/imaging"
	"github.com/julienschmidt/httprouter"
	cs "github.com/melodiez14/meiko/src/module/course"
	fl "github.com/melodiez14/meiko/src/module/file"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	tt "github.com/melodiez14/meiko/src/module/tutorial"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/helper"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func UploadProfileImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)

	// get uploaded file
	r.ParseMultipartForm(2 * MB)
	file, header, err := r.FormFile("file")
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("File is not exist"))
		return
	}
	defer file.Close()

	// extract file extension
	fn, ext, err := helper.ExtractExtension(header.Filename)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("File doesn't have an extension"))
		return
	}

	// decode file
	img, err := imaging.Decode(file)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Not valid image"))
		return
	}

	bound := img.Bounds()
	params := uploadImageParams{
		Height:    bound.Dx(),
		Width:     bound.Dy(),
		FileName:  fn,
		Extension: ext,
		Mime:      header.Header.Get("Content-Type"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// generate file id
	t := time.Now().UnixNano()
	rand.Seed(t)
	mImgID := fmt.Sprintf("%d.%06d.1", t, rand.Intn(999999))
	tImgID := fmt.Sprintf("%d.%06d.2", t, rand.Intn(999999))

	go func() {
		// resize image
		mImg := imaging.Resize(img, 300, 0, imaging.Lanczos)
		tImg := imaging.Thumbnail(img, 128, 128, imaging.Lanczos)

		// save image to storage
		imaging.Save(mImg, alias.Dir["data"]+"/profile/"+mImgID+".jpg")
		imaging.Save(tImg, alias.Dir["data"]+"/profile/"+tImgID+".jpg")
	}()

	// begin transaction to db
	tx, err := conn.DB.Beginx()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	// delete last image
	err = fl.DeleteProfileImage(sess.ID, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	// insert main image
	err = fl.Insert(mImgID, args.FileName, args.Mime, args.Extension, sess.ID, fl.TypProfPict, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	// insert thumbnail image
	err = fl.Insert(tImgID, args.FileName, args.Mime, args.Extension, sess.ID, fl.TypProfPictThumb, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Status OK"))
	return
}

func GetFileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var err error
	payload := ps.ByName("payload")
	filename := html.EscapeString(ps.ByName("filename"))

	switch payload {
	case "assignment", "tutorial":
		err = handleSingleWithMeta(payload, filename, w)
	case "profile", "default", "information":
		err = handleSingleWithoutMeta(payload, filename, w)
	case "assignment-user":
		err = handleUserAssignment(w) // change the parameter
	default:
		err = fmt.Errorf("Invalid")
	}

	if err != nil {
		http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
		return
	}
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	var typ string
	payload := ps.ByName("payload")

	switch payload {
	case "pl":
		typ = fl.TypProfPict
	case "pl_t":
		typ = fl.TypProfPictThumb
	default:
		http.Redirect(w, r, fl.UsrNoPhotoURL, http.StatusSeeOther)
		return
	}

	file, err := fl.GetByTypeUserID(sess.ID, typ, fl.ColID)
	if err != nil {
		http.Redirect(w, r, fl.UsrNoPhotoURL, http.StatusSeeOther)
		return
	}

	url := fmt.Sprintf("/api/v1/file/profile/%s.jpg", file.ID)
	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}

// UploadFileHandler ...
func UploadFileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	// access validation
	var typ string
	payload := r.FormValue("payload")
	isHasAccess := false
	switch payload {
	case "assignment":
		isHasAccess = sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXCreate, rg.RoleCreate, rg.RoleXUpdate, rg.RoleUpdate)
		typ = fl.TypAssignment
	case "tutorial":
		isHasAccess = sess.IsHasRoles(rg.ModuleTutorial, rg.RoleXCreate, rg.RoleCreate, rg.RoleXUpdate, rg.RoleUpdate)
		typ = fl.TypTutorial
	}

	if !isHasAccess {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	// logic
	r.ParseMultipartForm(2 * MB)
	file, header, err := r.FormFile("file")
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("File is not exist"))
		return
	}
	defer file.Close()

	// extract file extension
	fn, ext, err := helper.ExtractExtension(header.Filename)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("File doesn't have an extension"))
		return
	}

	// add mime validation
	params := uploadFileParams{
		fileName:  fn,
		extension: ext,
		mime:      header.Header.Get("Content-Type"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// get filename
	t := time.Now().UnixNano()
	rand.Seed(t)
	id := fmt.Sprintf("%d.%06d", t, rand.Intn(999999))

	// save file
	go func() {
		path := fmt.Sprintf("%s/%s/%s.%s", alias.Dir["data"], payload, id, args.extension)
		f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()

		file.Seek(0, 0)
		io.Copy(f, file)
	}()

	err = fl.Insert(id, args.fileName, args.mime, args.extension, sess.ID, typ, nil)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage(id))
	return
}

// RouterFileHandler ...
func RouterFileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if sess == nil {
		http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
		return
	}

	var typ string
	var isHasAccess bool

	params := routerParams{
		id:      r.FormValue("id"),
		payload: r.FormValue("payload"),
		role:    r.FormValue("role"),
	}

	args, err := params.validate()
	if err != nil {
		http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
		return
	}

	if args.payload == "tutorial" {
		tutorial, err := tt.GetByID(args.id)
		if err != nil {
			http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
			return
		}

		switch args.role {
		case "assistant":
			if !sess.IsHasRoles(rg.ModuleTutorial, rg.RoleXRead, rg.RoleRead) || !cs.IsAssistant(sess.ID, tutorial.ScheduleID) {
				http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
				return
			}
		case "student":
			if !cs.IsEnrolled(sess.ID, tutorial.ScheduleID) {
				http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
				return
			}
		}

		isHasAccess = true
		typ = fl.TypTutorial
	}

	if !isHasAccess {
		http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
		return
	}

	file, err := fl.GetByRelation(typ, strconv.FormatInt(args.id, 10))
	if err != nil {
		http.Redirect(w, r, fl.NotFoundURL, http.StatusSeeOther)
		return
	}

	url := fmt.Sprintf("/api/v1/file/%s/%s.%s", args.payload, file.ID, file.Extension)
	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}
