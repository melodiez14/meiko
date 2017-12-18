package information

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/melodiez14/meiko/src/util/conn"

	"github.com/melodiez14/meiko/src/webserver/template"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/course"
	cs "github.com/melodiez14/meiko/src/module/course"
	fs "github.com/melodiez14/meiko/src/module/file"
	inf "github.com/melodiez14/meiko/src/module/information"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/auth"
)

// GetHandler ...
func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := getParams{
		total: r.FormValue("ttl"),
		page:  r.FormValue("pg"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid request"))
		return
	}

	scheduleID, err := cs.SelectScheduleIDByUserID(sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// get active course
	courses, err := cs.SelectByScheduleID(scheduleID, cs.StatusScheduleActive)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	scheduleID = []int64{}
	for _, val := range courses {
		scheduleID = append(scheduleID, val.Schedule.ID)
	}

	offset := (args.page - 1) * args.total
	informations, err := inf.SelectByPage(scheduleID, args.total, offset)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	informationsID := []string{}
	for _, val := range informations {
		informationsID = append(informationsID, strconv.FormatInt(val.ID, 10))
	}

	thumbs, err := fs.SelectByRelation(fs.TypInfPictThumb, informationsID, &sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tImg := map[string]fs.File{}
	for _, val := range thumbs {
		if val.TableID.Valid {
			tImg[val.TableID.String] = val
		}
	}

	images, err := fs.SelectByRelation(fs.TypInfPict, informationsID, &sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	mImg := map[string]fs.File{}
	for _, val := range images {
		if val.TableID.Valid {
			mImg[val.TableID.String] = val
		}
	}

	resp := []getResponse{}
	var thumb, main string
	for _, val := range informations {
		thumb = fs.NoImgAvailable
		main = fs.NoImgAvailable
		if v, ok := tImg[strconv.FormatInt(val.ID, 10)]; ok {
			thumb = fmt.Sprintf("/api/v1/file/information/%s.%s", v.ID, v.Extension)
		}
		if v, ok := mImg[strconv.FormatInt(val.ID, 10)]; ok {
			main = fmt.Sprintf("/api/v1/file/information/%s.%s", v.ID, v.Extension)
		}
		resp = append(resp, getResponse{
			ID:             val.ID,
			Title:          val.Title,
			Description:    val.Description.String,
			Date:           val.CreatedAt.Format("Monday, 2 January 2006"),
			Image:          main,
			ImageThumbnail: thumb,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// CreateHandler func ...
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleInformation, rg.RoleCreate, rg.RoleXCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := createParams{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ScheduleID:  r.FormValue("schedule_did"),
		FilesID:     r.FormValue("file_id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	tx := conn.DB.MustBegin()
	// Insert
	tableID, err := inf.Insert(args.Title, args.Description, args.ScheduleID, tx)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	if args.FilesID != nil {
		for _, fileID := range args.FilesID {
			err := fs.UpdateRelation(fileID, fs.TypInf, tableID, tx)
			if err != nil {
				tx.Rollback()
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusBadRequest).
					AddError("Wrong File ID"))
				return
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Information created successfully"))
	return

}

// UpdateHandler func ...
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleInformation, rg.RoleUpdate, rg.RoleXUpdate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := updateParams{
		ID:          ps.ByName("id"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ScheduleID:  r.FormValue("schedule_id"),
		FilesID:     r.FormValue("file_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// check is information id exist?
	if !inf.IsInformationIDExist(args.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Information ID does not exist"))
		return
	}
	// check is shedule ID exist
	if args.ScheduleID != 0 {
		if !cs.IsExistScheduleID(args.ScheduleID) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("Schedule ID does not exist"))
			return
		}
	}
	tx := conn.DB.MustBegin()
	err = inf.Update(args.Title, args.Description, args.ScheduleID, args.ID, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	// Get All relations with
	filesIDDB, err := fs.GetByStatus(fs.StatusExist, args.ID)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	var tableID = strconv.FormatInt(args.ID, 10)
	// Add new file
	for _, fileID := range args.FilesID {
		if !fs.IsExistID(fileID) {
			filesIDDB = append(filesIDDB, fileID)
			// Update relation
			err := fs.UpdateRelation(fileID, fs.TypInf, tableID, tx)
			if err != nil {
				tx.Rollback()
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError))
				return
			}
		}
	}
	for _, fileIDDB := range filesIDDB {
		isSame := 0
		for _, fileIDUser := range args.FilesID {
			if fileIDUser == fileIDDB {
				isSame = 1
			}
		}
		if isSame == 0 {
			err := fs.UpdateStatusFiles(fileIDDB, fs.StatusDeleted, tx)
			if err != nil {
				tx.Rollback()
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError))
				return
			}
		}
	}
	err = tx.Commit()
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Update information succesfully"))
	return
}

// GetDetailByAdminHandler func ...
func GetDetailByAdminHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleInformation, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := detailInfromationParams{
		ID: ps.ByName("id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// check is information id exist?
	if !inf.IsInformationIDExist(args.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Information ID does not exist"))
		return
	}
	res, err := inf.GetByID(args.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	id := res.ScheduleID.Int64
	if id != 0 {
		if !cs.IsEnrolled(sess.ID, id) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("You does not have permission"))
			return
		}
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusBadRequest).
		SetData(res))
	return

}

// ReadHandler func ...
func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)

	params := readListParams{
		total: r.FormValue("ttl"),
		page:  r.FormValue("pg"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid request"))
		return
	}

	scheduleID, err := cs.SelectScheduleIDByUserID(sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	offset := (args.page - 1) * args.total
	result, err := inf.SelectByPage(scheduleID, args.total, offset)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(result))
	return
}

// DeleteHandler func ...
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleInformation, rg.RoleDelete, rg.RoleXDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := deleteParams{
		ID: ps.ByName("id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// check is information id exist?
	if !inf.IsInformationIDExist(args.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Information ID does not exist"))
		return
	}
	// delete query
	err = inf.Delete(args.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Delete failed"))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Delete information successfully"))
	return

}

// GetDetailHandler func ..
func GetDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := detailInfromationParams{
		ID: ps.ByName("id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError(err.Error()))
		return
	}
	scheduleID := inf.GetScheduleIDByID(args.ID)
	if scheduleID != 0 {
		if !course.IsEnrolled(sess.ID, scheduleID) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("you do not have permission to this informations"))
			return
		}
	}
	res, err := inf.GetByID(args.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Information does not exist"))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}
