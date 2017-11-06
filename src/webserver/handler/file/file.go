package file

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"

	"github.com/disintegration/imaging"
	"github.com/julienschmidt/httprouter"
	fl "github.com/melodiez14/meiko/src/module/file"
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
		imaging.Save(mImg, "files/var/www/meiko/data/profile/"+mImgID+".jpg")
		imaging.Save(tImg, "files/var/www/meiko/data/profile/"+tImgID+".jpg")
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

	params := getFileParams{
		payload:  ps.ByName("payload"),
		filename: ps.ByName("filename"),
	}

	args, err := params.validate()
	if err != nil {
		w.WriteHeader(404)
		return
	}

	path := fmt.Sprintf("files/var/www/meiko/data/%s/%s", args.payload, args.filename)
	file, err := os.Open(path)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	defer file.Close()

	// set header
	switch args.payload {
	case "profile":
		w.Header().Set("Content-Type", "image/jpeg")
	default:
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=2628000")

	io.Copy(w, file)
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
		w.WriteHeader(404)
		return
	}

	file, err := fl.GetByTypeUserID(sess.ID, typ, fl.ColID, fl.ColExtension)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	url := fmt.Sprintf("/api/v1/files/profile/%s.%s", file.ID, file.Extension)
	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}

// not functional yet
func UploadAssignmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
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
	params := uploadAssignmentParams{
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

	// get filename
	t := time.Now().UnixNano()
	rand.Seed(t)
	id := fmt.Sprintf("%d.%06d", t, rand.Intn(999999))

	// save file
	go func() {
		path := fmt.Sprintf("files/var/www/meiko/data/%s/%s.%s", "assignment", id, args.Extension)
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		defer f.Close()

		file.Seek(0, 0)
		io.Copy(f, file)
	}()

	err = fl.Insert(id, args.FileName, args.Mime, args.Extension, sess.ID, fl.TypAssignment, nil)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK))
	return
}
