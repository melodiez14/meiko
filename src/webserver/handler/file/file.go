package file

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	flmodule "github.com/melodiez14/meiko/src/module/file"
	"github.com/melodiez14/meiko/src/util/alias"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/helper"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func UploadImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	// get uploaded file
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

	params := uploadImageParams{
		Payload:   r.FormValue("payload"),
		FileName:  fn,
		Extension: ext,
		Mime:      header.Header.Get("Content-Type"),
	}

	// validate request
	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError(err.Error()))
		return
	}

	// map the image to specific table
	var (
		multipleFile bool
		path         string
		pathThumb    string
		mainType     sql.NullString
		thumbType    sql.NullString
		tableID      int64
		tableName    string
	)

	switch args.Payload {
	case "profile":
		multipleFile = false
		path = filepath.Join("private/profile", fmt.Sprintf("%d", sess.ID))
		pathThumb = filepath.Join("private/profile", fmt.Sprintf("t_%d", sess.ID))
		mainType.Valid = true
		mainType.String = alias.TypeProfile
		thumbType.Valid = true
		thumbType.String = alias.TypeProfileThumbnail
		tableID = sess.ID
		tableName = alias.ModuleUser
	default:
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	if !multipleFile {
		fexist, err := flmodule.Get(flmodule.ColID).
			Where(flmodule.ColType, flmodule.OperatorEquals, mainType).
			Exec()
		if err != nil && err != sql.ErrNoRows {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError).
				AddError("Cannot create a file"))
			return
		}

		if err != sql.ErrNoRows {

			// remove main image
			err = os.Remove(path)
			if err != nil {
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError).
					AddError("Cannot remove image"))
				return
			}

			// remove thumbnail image
			err = os.Remove(pathThumb)
			if err != nil {
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError).
					AddError("Cannot remove image thumbnail"))
				return
			}

			// create new main image
			output, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError).
					AddError("Cannot create a file"))
				return
			}
			defer output.Close()
			io.Copy(output, file)

			// create new thumbnail image
			thumbnail, err := os.OpenFile(pathThumb, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError).
					AddError("Cannot create a file"))
				return
			}
			defer thumbnail.Close()
			file.Seek(0, 0)
			io.Copy(thumbnail, file)

			// update main image metadata
			flmodule.Update(map[string]interface{}{
				flmodule.ColName:      args.FileName,
				flmodule.ColPath:      path,
				flmodule.ColMime:      args.Mime,
				flmodule.ColExtension: args.Extension,
				flmodule.ColUserID:    sess.ID,
				flmodule.ColType:      alias.TypeProfile,
				flmodule.ColTableName: tableName,
				flmodule.ColTableID:   tableID,
			}).Where(flmodule.ColID, flmodule.OperatorEquals, fexist.ID).
				AndWhere(flmodule.ColType, flmodule.OperatorEquals, mainType).
				Exec()

			// update thumbnail image metadata
			flmodule.Update(map[string]interface{}{
				flmodule.ColName:      args.FileName,
				flmodule.ColPath:      pathThumb,
				flmodule.ColMime:      args.Mime,
				flmodule.ColExtension: args.Extension,
				flmodule.ColUserID:    sess.ID,
				flmodule.ColType:      thumbType,
				flmodule.ColTableName: tableName,
				flmodule.ColTableID:   tableID,
			}).Where(flmodule.ColID, flmodule.OperatorEquals, fexist.ID).
				AndWhere(flmodule.ColType, flmodule.OperatorEquals, thumbType).
				Exec()

			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusOK).
				SetMessage("Image has been changed"))
			return
		}
	}

	// create main image
	output, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Cannot create a file"))
		return
	}
	defer output.Close()
	io.Copy(output, file)

	// create thumbnail image
	thumbnail, err := os.OpenFile(pathThumb, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Cannot create a file"))
		return
	}
	defer thumbnail.Close()
	file.Seek(0, 0)
	io.Copy(thumbnail, file)

	// insert main image metadata
	flmodule.Insert(map[string]interface{}{
		flmodule.ColName:      args.FileName,
		flmodule.ColPath:      path,
		flmodule.ColMime:      args.Mime,
		flmodule.ColExtension: args.Extension,
		flmodule.ColUserID:    sess.ID,
		flmodule.ColType:      mainType,
		flmodule.ColTableName: tableName,
		flmodule.ColTableID:   tableID,
	}).Exec()

	// insert thumbnail image metadata
	flmodule.Insert(map[string]interface{}{
		flmodule.ColName:      args.FileName,
		flmodule.ColPath:      pathThumb,
		flmodule.ColMime:      args.Mime,
		flmodule.ColExtension: args.Extension,
		flmodule.ColUserID:    sess.ID,
		flmodule.ColType:      thumbType,
		flmodule.ColTableName: tableName,
		flmodule.ColTableID:   tableID,
	}).Exec()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Image has been uploaded"))
	return
}

func GetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := getProfileParams{
		Payload: ps.ByName("payload"),
	}

	args, err := params.Validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	f, err := flmodule.Get().
		Where(flmodule.ColTableName, flmodule.OperatorEquals, alias.ModuleUser).
		AndWhere(flmodule.ColTableID, flmodule.OperatorEquals, sess.ID).
		AndWhere(flmodule.ColType, flmodule.OperatorEquals, args.Payload).
		Exec()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound).
			AddError("File is not exist"))
		return
	}

	file, err := ioutil.ReadFile(f.Path)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("File is not found"))
		return
	}

	w.Header().Set("Content-Type", f.Mime)
	w.Write(file)
	return
}
