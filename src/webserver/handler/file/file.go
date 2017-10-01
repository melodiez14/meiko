package file

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/julienschmidt/httprouter"
	flmodule "github.com/melodiez14/meiko/src/module/file"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/alias"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/helper"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	// resize image
	mImg := imaging.Resize(img, 300, 0, imaging.NearestNeighbor)
	tImg := imaging.Thumbnail(img, 128, 128, imaging.NearestNeighbor)

	// declare mapper
	var mapper uploadImageMapper

	switch args.Payload {
	case "profile":
		mapper = uploadImageMapper{
			fn:        fn,
			multiple:  false,
			mImg:      mImg,
			tImg:      tImg,
			mPath:     filepath.Join(alias.Dir.Profile, fmt.Sprintf("%d.jpg", sess.ID)),
			tPath:     filepath.Join(alias.Dir.Profile, fmt.Sprintf("t_%d.jpg", sess.ID)),
			mType:     alias.TypeProfile,
			tType:     alias.TypeProfileThumbnail,
			tableID:   sess.ID,
			tableName: rg.ModuleUser,
		}
	default:
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	err = mapper.save(sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError("Internal server error"))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("File Uploaded"))
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
		Where(flmodule.ColTableName, flmodule.OperatorEquals, rg.ModuleUser).
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

func (mapper uploadImageMapper) save(userID int64) error {

	// save file
	err := imaging.Save(mapper.mImg, mapper.mPath)
	if err != nil {
		return fmt.Errorf("failed to save main photo")
	}

	err = imaging.Save(mapper.tImg, mapper.tPath)
	if err != nil {
		return fmt.Errorf("failed to save thumbnail photo")
	}

	if !mapper.multiple {
		_, err := flmodule.Get(flmodule.ColID).
			Where(flmodule.ColType, flmodule.OperatorEquals, mapper.mType).
			OrWhere(flmodule.ColType, flmodule.OperatorEquals, mapper.tType).
			Exec()
		if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("error checking old metadata")
		}

		// if record exists in database
		if err == nil {
			// update main image metadata
			err = flmodule.Update(map[string]interface{}{
				flmodule.ColName:      mapper.fn,
				flmodule.ColPath:      mapper.mPath,
				flmodule.ColMime:      MimeJPEG,
				flmodule.ColExtension: ExtJPEG,
				flmodule.ColUserID:    userID,
				flmodule.ColType:      mapper.mType,
				flmodule.ColTableName: mapper.tableName,
				flmodule.ColTableID:   mapper.tableID,
			}).Where(flmodule.ColType, flmodule.OperatorEquals, mapper.mType).
				Exec()
			if err != nil {
				return fmt.Errorf("error updating main metadata")
			}

			// update thumbnail image metadata
			flmodule.Update(map[string]interface{}{
				flmodule.ColName:      mapper.fn,
				flmodule.ColPath:      mapper.tPath,
				flmodule.ColMime:      MimeJPEG,
				flmodule.ColExtension: ExtJPEG,
				flmodule.ColUserID:    userID,
				flmodule.ColType:      mapper.tType,
				flmodule.ColTableName: mapper.tableName,
				flmodule.ColTableID:   mapper.tableID,
			}).Where(flmodule.ColType, flmodule.OperatorEquals, mapper.tType).
				Exec()
			if err != nil {
				return fmt.Errorf("error updating thumbnail metadata")
			}

			return nil
		}
	}

	// insert main image metadata
	err = flmodule.Insert(map[string]interface{}{
		flmodule.ColName:      mapper.fn,
		flmodule.ColPath:      mapper.mPath,
		flmodule.ColMime:      MimeJPEG,
		flmodule.ColExtension: ExtJPEG,
		flmodule.ColUserID:    userID,
		flmodule.ColType:      mapper.mType,
		flmodule.ColTableName: mapper.tableName,
		flmodule.ColTableID:   mapper.tableID,
	}).Exec()
	if err != nil {
		return fmt.Errorf("error inserting main metadata")
	}

	// insert thumbnail image metadata
	flmodule.Insert(map[string]interface{}{
		flmodule.ColName:      mapper.fn,
		flmodule.ColPath:      mapper.tPath,
		flmodule.ColMime:      MimeJPEG,
		flmodule.ColExtension: ExtJPEG,
		flmodule.ColUserID:    userID,
		flmodule.ColType:      mapper.tType,
		flmodule.ColTableName: mapper.tableName,
		flmodule.ColTableID:   mapper.tableID,
	}).Exec()
	if err != nil {
		return fmt.Errorf("error inserting thumbnail metadata")
	}

	return nil
}
