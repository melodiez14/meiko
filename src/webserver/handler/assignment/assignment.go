package assignment

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"

	"github.com/julienschmidt/httprouter"
	asg "github.com/melodiez14/meiko/src/module/assignment"
	cs "github.com/melodiez14/meiko/src/module/course"
	fl "github.com/melodiez14/meiko/src/module/file"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/helper"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// GetHandler ...
func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	resp := []getResponse{}
	sess := r.Context().Value("User").(*auth.User)

	params := getParams{
		scheduleID: r.FormValue("schedule_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	var schedulesID []int64
	switch args.scheduleID.Valid {
	case true:
		// specific scheduleID
		if !cs.IsEnrolled(sess.ID, args.scheduleID.Int64) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusNoContent))
			return
		}
		schedulesID = []int64{args.scheduleID.Int64}
	case false:
		// all enrolled schedule
		schedulesID, err = cs.SelectScheduleIDByUserID(sess.ID, cs.PStatusStudent)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		// return if empty
		if len(schedulesID) < 1 {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusOK).
				SetData(resp))
			return
		}
	}

	gps, err := cs.SelectGPBySchedule(schedulesID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if len(gps) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
		return
	}

	var gpsID []int64
	for _, val := range gps {
		gpsID = append(gpsID, val.ID)
	}

	assignments, err := asg.SelectByGP(gpsID, true)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var asgID []int64
	for _, val := range assignments {
		asgID = append(asgID, val.ID)
	}

	if len(asgID) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
		return
	}

	submitted, err := asg.SelectSubmittedByUser(asgID, sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	submitMap := map[int64]asg.UserAssignment{}
	for _, val := range submitted {
		submitMap[val.AssignmentID] = val
	}

	for _, assignment := range assignments {
		submit, exist := submitMap[assignment.ID]
		desc := "-"
		score := "-"
		status := "unsubmitted"

		if exist {
			if assignment.DueDate.Before(time.Now()) {
				status = "overdue"
			} else {
				status = "submitted"
				if submit.Score.Valid {
					score = fmt.Sprintf("%.3g", submit.Score.Float64)
				}
			}
		}

		if assignment.Description.Valid {
			desc = assignment.Description.String
		}

		resp = append(resp, getResponse{
			ID:          assignment.ID,
			DueDate:     assignment.DueDate.Format("Monday, 2 January 2006 15:04:05"),
			Name:        assignment.Name,
			Score:       score,
			Status:      status,
			Description: desc,
			UpdatedAt:   assignment.UpdatedAt.Format("Monday, 2 January 2006 15:04:05"),
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// GetDetailHandler ...
func GetDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := getDetailParams{
		id: ps.ByName("id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	assignment, err := asg.GetByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}

	scheduleID, err := cs.GetScheduleIDByGP(assignment.GradeParameterID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if !cs.IsEnrolled(sess.ID, scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}

	submitted, err := asg.GetSubmittedByUser(args.id, sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	// response data preparation
	status := "unsubmitted"
	score := "-"
	submittedDate := "-"
	if assignment.DueDate.Before(time.Now()) {
		status = "overdue"
	}
	if submitted != nil {
		status = "submitted"
		submittedDate = submitted.UpdatedAt.Format("Monday, 2 January 2006 15:04:05")
		if submitted.Score.Valid {
			score = fmt.Sprintf("%.3g", submitted.Score.Float64)
		}
	}
	if assignment.Status == asg.StatusUploadNotRequired {
		status = "notrequired"
	}

	tableID := []string{strconv.FormatInt(args.id, 10)}
	asgFile, err := fl.SelectByRelation(fl.TypAssignment, tableID, nil)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	submittedFile, err := fl.SelectByRelation(fl.TypAssignmentUpload, tableID, &sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	// file from assistant
	rAsgFile := []file{}
	for _, val := range asgFile {
		rAsgFile = append(rAsgFile, file{
			Name:         fmt.Sprintf("%s.%s", val.Name, val.Extension),
			URL:          fmt.Sprintf("/api/v1/file/assignment/%s.%s", val.ID, val.Extension),
			URLThumbnail: helper.MimeToThumbnail(val.Mime),
		})
	}

	// file from student
	rSubmittedFile := []file{}
	for _, val := range submittedFile {
		rSubmittedFile = append(rSubmittedFile, file{
			Name:         fmt.Sprintf("%s.%s", val.Name, val.Extension),
			URL:          fmt.Sprintf("/api/v1/file/assignment/%s.%s", val.ID, val.Extension),
			URLThumbnail: helper.MimeToThumbnail(val.Mime),
		})
	}

	resp := getDetailResponse{
		ID:             assignment.ID,
		Name:           assignment.Name,
		Status:         status,
		Description:    assignment.Description.String,
		DueDate:        assignment.DueDate.Format("Monday, 2 January 2006 15:04:05"),
		Score:          score,
		CreatedAt:      assignment.CreatedAt.Format("Monday, 2 January 2006"),
		UpdatedAt:      assignment.UpdatedAt.Format("Monday, 2 January 2006"),
		AssignmentFile: rAsgFile,
		SubmittedFile:  rSubmittedFile,
		SubmittedDate:  submittedDate,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// SubmitHandler ...
func SubmitHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := submitParams{
		id:          ps.ByName("id"),
		description: r.FormValue("description"),
		fileID:      r.FormValue("file_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	assignment, err := asg.GetByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}

	if assignment.Status == asg.StatusUploadNotRequired {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid Request"))
		return
	}

	if assignment.DueDate.Before(time.Now()) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Sorry, you can't upload this assignment because of overdue."))
		return
	}

	scheduleID, err := cs.GetScheduleIDByGP(assignment.GradeParameterID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if !cs.IsEnrolled(sess.ID, scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	upload, err := asg.GetSubmittedByUser(args.id, sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	// update
	if upload != nil {
		err = handleSubmitUpdate(assignment.ID, sess.ID, args.description, args.fileID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}

		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetMessage("Success"))
		return
	}

	// insert
	err = handleSubmitInsert(assignment.ID, sess.ID, args.description, args.fileID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Success"))
	return
}

func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXCreate, rg.RoleCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := createParams{
		name:        r.FormValue("name"),
		description: r.FormValue("description"),
		gpID:        r.FormValue("gpid"),
		dueDate:     r.FormValue("due_date"),
		status:      r.FormValue("status"),
		filesID:     r.FormValue("file_id"),
		fileSize:    r.FormValue("file_size"),
		fileType:    r.FormValue("file_type"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	scheduleID, err := cs.GetScheduleIDByGP(args.gpID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if !cs.IsAssistant(sess.ID, scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	// validate file_id more ui friendly
	// slow performance

	tx := conn.DB.MustBegin()
	id, err := asg.Insert(args.name, args.description, args.gpID, args.fileSize, args.dueDate, args.status, tx)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	idStr := strconv.FormatInt(id, 10)
	for _, fileID := range args.filesID {
		if fl.UpdateRelation(fileID, fl.TypAssignment, idStr, tx) != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	tx.Commit()
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Success"))
	return
}

// // CreateHandler function is
// func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	sess := r.Context().Value("User").(*auth.User)
// 	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleCreate, rg.RoleXCreate) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusForbidden).
// 			AddError("You don't have privilege"))
// 		return
// 	}
// 	params := createParams{
// 		FilesID:           r.FormValue("file_id"),
// 		GradeParametersID: r.FormValue("grade_parameter_id"),
// 		Name:              r.FormValue("name"),
// 		Description:       r.FormValue("description"),
// 		Status:            r.FormValue("status"),
// 		DueDate:           r.FormValue("due_date"),
// 		Type:              r.FormValue("type"),
// 		Size:              r.FormValue("size"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}

// 	// is grade_parameter exist
// 	if !as.IsExistByGradeParameterID(args.GradeParametersID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Grade parameters id does not exist!"))
// 		return
// 	}
// 	// Insert to table assignments
// 	tx := conn.DB.MustBegin()
// 	TableID, err := as.Insert(
// 		args.GradeParametersID,
// 		args.Name,
// 		args.Status,
// 		args.DueDate,
// 		args.Description,
// 		tx,
// 	)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	// Files null
// 	if args.FilesID == "" {
// 		tx.Commit()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusOK).
// 			SetMessage("Success Without files"))
// 		return
// 	}
// 	// Split files id if possible
// 	filesID := strings.Split(args.FilesID, "~")
// 	tableName := "assignments"
// 	for _, fileID := range filesID {
// 		// Wrong file code
// 		if !as.IsFileIDExist(fileID) {
// 			tx.Rollback()
// 			template.RenderJSONResponse(w, new(template.Response).
// 				SetCode(http.StatusBadRequest).
// 				SetMessage("Wrong file code!"))
// 			return
// 		}
// 		// Update files
// 		err = fs.UpdateRelation(fileID, tableName, TableID, tx)
// 		if err != nil {
// 			tx.Rollback()
// 			template.RenderJSONResponse(w, new(template.Response).
// 				SetCode(http.StatusInternalServerError))
// 			return
// 		}
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetMessage("Success!"))
// 	return

// }

// // GetAllAssignmentHandler func is ...
// func GetAllAssignmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	sess := r.Context().Value("User").(*auth.User)
// 	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusForbidden).
// 			AddError("You don't have privilege"))
// 		return
// 	}
// 	params := readParams{
// 		Page:  r.FormValue("pg"),
// 		Total: r.FormValue("ttl"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusForbidden).
// 			AddError("Invalid request"))
// 		return
// 	}
// 	offset := (args.Page - 1) * args.Total
// 	assignments, err := as.SelectByPage(args.Total, offset)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	var status string
// 	var res []readResponse
// 	for _, val := range assignments {

// 		if val.Status == as.StatusAssignmentActive {
// 			status = "active"
// 		} else {
// 			status = "inactive"
// 		}

// 		res = append(res, readResponse{
// 			Name:             val.Name,
// 			Description:      val.Description,
// 			Status:           status,
// 			DueDate:          val.DueDate,
// 			GradeParameterID: val.GradeParameterID,
// 		})
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return

// }

// // DetailHandler func is ...
// func DetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	sess := r.Context().Value("User").(*auth.User)
// 	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusForbidden).
// 			AddError("You don't have privilege"))
// 		return
// 	}

// 	params := detailParams{
// 		IdentityCode: ps.ByName("id"),
// 	}

// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}

// 	assignment, err := as.GetByAssignementID(args.IdentityCode)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusNotFound))
// 		return
// 	}

// 	desc := "-"
// 	if assignment.Description.Valid {
// 		desc = assignment.Description.String
// 	}

// 	res := detailResponse{
// 		ID:               assignment.ID,
// 		Status:           assignment.Status,
// 		Name:             assignment.Name,
// 		GradeParameterID: assignment.GradeParameterID,
// 		Description:      desc,
// 		DueDate:          assignment.DueDate.Format("Monday, 2 January 2006 15:04:05"),
// 	}

// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return
// }

// // UpdateHandler func is ...
// func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	sess := r.Context().Value("User").(*auth.User)
// 	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleUpdate, rg.RoleXRead) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusForbidden).
// 			AddError("You don't have privilege"))
// 		return
// 	}

// 	params := updatePrams{
// 		ID:                ps.ByName("id"),
// 		FilesID:           r.FormValue("file_id"),
// 		GradeParametersID: r.FormValue("grade_parameter_id"),
// 		Name:              r.FormValue("name"),
// 		Description:       r.FormValue("description"),
// 		Status:            r.FormValue("status"),
// 		DueDate:           r.FormValue("due_date"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}
// 	// Params ID is exist
// 	if !as.IsAssignmentExist(args.ID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusNotFound))
// 		return
// 	}
// 	// is grade_parameter exist
// 	if !as.IsExistByGradeParameterID(args.GradeParametersID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Grade parameters id does not exist!"))
// 		return
// 	}

// 	// Insert to table assignments
// 	tx := conn.DB.MustBegin()
// 	err = as.Update(
// 		args.GradeParametersID,
// 		args.ID,
// 		args.Name,
// 		args.Status,
// 		args.DueDate,
// 		args.Description,
// 		tx,
// 	)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}

// 	var filesIDUser = strings.Split(args.FilesID, "~")
// 	var tableID = strconv.FormatInt(args.ID, 10)
// 	// Get All relations with
// 	filesIDDB, err := fs.GetByStatus(fs.StatusExist, args.ID)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	// Add new file
// 	for _, fileID := range filesIDUser {
// 		if !fs.IsExistID(fileID) {
// 			filesIDDB = append(filesIDDB, fileID)
// 			// Update relation
// 			err := fs.UpdateRelation(fileID, TableNameAssignments, tableID, tx)
// 			if err != nil {
// 				tx.Rollback()
// 				template.RenderJSONResponse(w, new(template.Response).
// 					SetCode(http.StatusInternalServerError))
// 				return
// 			}
// 		}
// 	}
// 	for _, fileIDDB := range filesIDDB {
// 		isSame := 0
// 		for _, fileIDUser := range filesIDUser {
// 			if fileIDUser == fileIDDB {
// 				isSame = 1
// 			}
// 		}
// 		if isSame == 0 {
// 			err := fs.UpdateStatusFiles(fileIDDB, fs.StatusDeleted, tx)
// 			if err != nil {
// 				tx.Rollback()
// 				template.RenderJSONResponse(w, new(template.Response).
// 					SetCode(http.StatusInternalServerError))
// 				return
// 			}
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetMessage("Update Assignment Success!"))
// 	return

// }

// // CreateHandlerByUser func ...
// func CreateHandlerByUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	sess := r.Context().Value("User").(*auth.User)
// 	params := uploadAssignmentParams{
// 		UserID:       sess.ID,
// 		AssignmentID: r.FormValue("assignment_id"),
// 		Description:  r.FormValue("description"),
// 		FileID:       r.FormValue("file_id"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}

// 	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
// 	scheduleID := cs.GetScheduleID(gradeParameterID)
// 	isValidAssignment := usr.IsUserTakeSchedule(args.UserID, scheduleID)
// 	if !isValidAssignment {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Invalid Assignment ID"))
// 		return
// 	}
// 	// due date checking
// 	dueDate, err := as.GetDueDateAssignment(args.AssignmentID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	overdue := !time.Now().Before(dueDate)
// 	if overdue {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusOK).
// 			SetMessage("Assignment has been overdue"))
// 		return
// 	}
// 	// Insert
// 	tx := conn.DB.MustBegin()
// 	err = as.UploadAssignment(args.AssignmentID, args.UserID, args.Description, tx)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("You can only submit once for this assignment"))
// 		return
// 	}

// 	tableID := fmt.Sprintf("%d%d", args.AssignmentID, args.UserID)
// 	//Update Relations
// 	if args.FileID != nil {
// 		for _, fileID := range args.FileID {
// 			err := fs.UpdateRelation(fileID, TableNameUserAssignments, tableID, tx)
// 			if err != nil {
// 				tx.Rollback()
// 				template.RenderJSONResponse(w, new(template.Response).
// 					SetCode(http.StatusBadRequest).
// 					AddError("Wrong File ID"))
// 				return
// 			}
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetMessage("Insert Assignment Success!"))
// 	return
// }

// // UpdateHandlerByUser func ...
// func UpdateHandlerByUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	sess := r.Context().Value("User").(*auth.User)
// 	params := uploadAssignmentParams{
// 		FileID:       r.FormValue("file_id"),
// 		AssignmentID: ps.ByName("id"),
// 		UserID:       sess.ID,
// 		Description:  r.FormValue("description"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}
// 	// Get grade parameters id
// 	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
// 	scheduleID := cs.GetScheduleID(gradeParameterID)
// 	isValidAssignment := usr.IsUserTakeSchedule(args.UserID, scheduleID)
// 	if !isValidAssignment {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Invalid Assignment ID"))
// 		return
// 	}
// 	// Update
// 	tx := conn.DB.MustBegin()
// 	err = as.UpdateUploadAssignment(args.AssignmentID, args.UserID, args.Description, tx)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Update Failed"))
// 		return
// 	}
// 	key := fmt.Sprintf("%d%d", args.AssignmentID, args.UserID)
// 	tableID, err := strconv.ParseInt(key, 10, 64)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Can not convert to int64"))
// 		return
// 	}
// 	// Get All relations with
// 	filesIDDB, err := fs.GetByStatus(fs.StatusExist, tableID)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	// Add new file
// 	for _, fileID := range args.FileID {
// 		if !fs.IsExistID(fileID) {
// 			filesIDDB = append(filesIDDB, fileID)
// 			// Update relation
// 			err := fs.UpdateRelation(fileID, TableNameUserAssignments, key, tx)
// 			if err != nil {
// 				tx.Rollback()
// 				template.RenderJSONResponse(w, new(template.Response).
// 					SetCode(http.StatusInternalServerError))
// 				return
// 			}
// 		}
// 	}
// 	for _, fileIDDB := range filesIDDB {
// 		isSame := 0
// 		for _, fileIDUser := range args.FileID {
// 			if fileIDUser == fileIDDB {
// 				isSame = 1
// 			}
// 		}
// 		if isSame == 0 {
// 			err := fs.UpdateStatusFiles(fileIDDB, fs.StatusDeleted, tx)
// 			if err != nil {
// 				tx.Rollback()
// 				template.RenderJSONResponse(w, new(template.Response).
// 					SetCode(http.StatusInternalServerError))
// 				return
// 			}
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetMessage("Update Assignment Success!"))
// 	return
// }

// // GetAssignmentByScheduleHandler func ...
// func GetAssignmentByScheduleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	params := listAssignmentsParams{
// 		ScheduleID: r.FormValue("id"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Bad Request"))
// 		return
// 	}
// 	// is correct schedule id
// 	if !cs.IsExistScheduleID(args.ScheduleID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("You do not took this course"))
// 		return
// 	}
// 	gradeParamsID, err := cs.GetGradeParametersIDByScheduleID(args.ScheduleID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Error"))
// 		return
// 	}
// 	assignments := as.SelectByGradeParametersID(gradeParamsID)
// 	res := []listAssignmentResponse{}
// 	for _, val := range assignments {
// 		status := "must_not_upload"
// 		submited := false
// 		var description string
// 		if val.Description.Valid {
// 			description = fmt.Sprintf(val.Description.String)
// 		}
// 		if val.Status == 1 {
// 			status = "must_upload"
// 			if as.IsUploaded(val.ID) {
// 				submited = true
// 			}
// 		}
// 		res = append(res, listAssignmentResponse{
// 			Submitted:   submited,
// 			ID:          val.ID,
// 			Name:        val.Name,
// 			Status:      status,
// 			Description: description,
// 			DueDate:     val.DueDate.Format("Monday, 2 January 2006 15:04:05"),
// 		})
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).SetData(res))
// 	return
// }

// // GetUploadedDetailHandler func ...
// func GetUploadedDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	sess := r.Context().Value("User").(*auth.User)
// 	params := readUploadedDetailParams{
// 		UserID:       ps.ByName("id"),
// 		ScheduleID:   ps.ByName("schedule_id"),
// 		AssignmentID: ps.ByName("assignment_id"),
// 	}

// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError(err.Error()))
// 		return
// 	}
// 	// Is Valid User ID
// 	if sess.ID != args.UserID {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError("Wrong User ID"))
// 		return
// 	}
// 	if !usr.IsUserTakeSchedule(sess.ID, args.ScheduleID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError("Wrong Schedule ID"))
// 		return
// 	}
// 	// Get Assignments Detail
// 	assignment, err := as.GetUploadedAssignmentByID(args.AssignmentID, sess.ID)
// 	key := fmt.Sprintf("%d%d", args.AssignmentID, sess.ID)
// 	tableID, err := strconv.ParseInt(key, 10, 64)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError(err.Error()))
// 		return
// 	}

// 	// Get File
// 	files, err := fs.GetByUserIDTableIDName(sess.ID, tableID, TableNameUserAssignments)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}

// 	res := readUploadedDetailArgs{
// 		UserID:       args.UserID,
// 		ScheduleID:   args.ScheduleID,
// 		AssignmentID: args.AssignmentID,
// 		Name:         assignment.Name,
// 		Description:  assignment.DescriptionAssignment,
// 		PathFile:     files,
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return

// }

// // CreateScoreHandler func ...
// func CreateScoreHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	sess := r.Context().Value("User").(*auth.User)
// 	// get schedule, assignment
// 	params := createScoreParams{
// 		ScheduleID:   ps.ByName("schedule_id"),
// 		AssignmentID: ps.ByName("assignment_id"),
// 		Users:        r.FormValue("users"),
// 	}
// 	// validate
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}
// 	// is schedule_id match with assignments_id?
// 	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
// 	if !as.IsAssignmentExistByGradeParameterID(args.AssignmentID, gradeParameterID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusNotFound))
// 		return
// 	}

// 	// is user (admin) took that course?
// 	if !cs.IsAssistant(sess.ID, args.ScheduleID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("you do not took this course!"))
// 		return
// 	}
// 	usersID, err := usr.SelectIDByIdentityCode(args.IdentityCode)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	if len(usersID) != len(args.IdentityCode) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("There is invalid identity code"))
// 		return
// 	}
// 	if !cs.IsAllUsersEnrolled(args.ScheduleID, usersID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("There is a invalid identity code"))
// 		return
// 	}
// 	tx := conn.DB.MustBegin()
// 	// Insert
// 	err = as.CreateScore(args.AssignmentID, usersID, args.Score, tx)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetMessage("Add score assignment successfully"))
// 	return

// }

// // GetDetailAssignmentByAdmin func ...
// func GetDetailAssignmentByAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	sess := r.Context().Value("User").(*auth.User)
// 	// get schedule, assignment
// 	params := detailAssignmentParams{
// 		ScheduleID:   ps.ByName("schedule_id"),
// 		AssignmentID: ps.ByName("assignment_id"),
// 	}
// 	// validate
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}
// 	// is schedule_id match with assignments_id?
// 	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
// 	if !as.IsAssignmentExistByGradeParameterID(args.AssignmentID, gradeParameterID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusNotFound))
// 		return
// 	}

// 	// is user (admin) took that course?
// 	if !cs.IsAssistant(sess.ID, args.ScheduleID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("you do not took this course!"))
// 		return
// 	}

// 	// get assignment detail
// 	assignments, err := as.GetAssignmentByID(args.AssignmentID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	res := detailAssignmentResponse{}
// 	userList := []userAssignment{}
// 	// is assingment must upload or not?
// 	if as.IsAssignmentMustUpload(args.AssignmentID) {
// 		userAssignments, err := as.SelectUserAssignmentsByStatusID(args.AssignmentID)
// 		if err != nil {
// 			template.RenderJSONResponse(w, new(template.Response).
// 				SetCode(http.StatusInternalServerError))
// 			return
// 		}
// 		for _, value := range userAssignments {
// 			userList = append(userList, userAssignment{
// 				UserID: value.UserID,
// 				Name:   value.Name,
// 				Grade:  0,
// 			})
// 		}
// 		res = detailAssignmentResponse{
// 			Name:          assignments.Name,
// 			Description:   assignments.Description,
// 			DueDate:       assignments.DueDate,
// 			IsCreateScore: false,
// 			Praktikan:     userList,
// 		}
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusOK).
// 			SetData(res))
// 		return
// 	}
// 	userAssignments, err := usr.SelectUserByScheduleID(args.ScheduleID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}

// 	for _, value := range userAssignments {
// 		userList = append(userList, userAssignment{
// 			UserID: value.UserID,
// 			Name:   value.Name,
// 			Grade:  0,
// 		})
// 	}
// 	res = detailAssignmentResponse{
// 		Name:          assignments.Name,
// 		Description:   assignments.Description,
// 		DueDate:       assignments.DueDate,
// 		IsCreateScore: true,
// 		Praktikan:     userList,
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return

// }

// //ScoreHandler func ...
// func ScoreHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	params := updateScoreParams{
// 		ScheduleID:   ps.ByName("id"),
// 		AssignmentID: ps.ByName("assignment_id"),
// 		UserID:       r.FormValue("user_id"),
// 		Score:        r.FormValue("score"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}
// 	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
// 	if !as.IsAssignmentExistByGradeParameterID(args.AssignmentID, gradeParameterID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusNotFound))
// 		return
// 	}
// 	// User (Praktikan) took that course
// 	if !cs.IsEnrolled(args.UserID, args.ScheduleID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("User do not took this course!"))
// 		return
// 	}
// 	// check schedule have assignments
// 	if !cs.IsUserHasUploadedFile(args.AssignmentID, args.UserID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("User has not uploaded yet assignment"))
// 		return
// 	}
// 	tx := conn.DB.MustBegin()
// 	err = as.UpdateScoreAssignment(args.AssignmentID, args.UserID, args.Score, tx)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError("Can not update score"))
// 		return
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetMessage("Score Updated"))
// 	return
// }

// // GetUploadedAssignmentByAdminHandler func ...
// func GetUploadedAssignmentByAdminHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	sess := r.Context().Value("User").(*auth.User)
// 	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusForbidden).
// 			AddError("You don't have privilege"))
// 		return
// 	}
// 	params := readUploadedAssignmentParams{
// 		ScheduleID:   ps.ByName("id"),
// 		AssignmentID: ps.ByName("assignment_id"),
// 		Page:         r.FormValue("pg"),
// 		Total:        r.FormValue("ttl"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError(err.Error()))
// 		return
// 	}
// 	offset := (args.Page - 1) * args.Total
// 	// Check schedule id
// 	if !cs.IsExistScheduleID(args.ScheduleID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Schdule ID does not exist"))
// 		return
// 	}
// 	// Check assignment id
// 	if !as.IsAssignmentExist(args.AssignmentID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Assignment ID does not exist"))
// 		return
// 	}
// 	// Get all data p_users_assignment
// 	assignments, err := as.GetAllUserAssignmentByAssignmentID(args.AssignmentID, args.Total, offset)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	//files, err := fs.GetByTableIDName(args.AssignmentID, TableNameUserAssignments)
// 	// get all data files relations
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).SetData(assignments))
// 	// serve json
// }

// // DeleteAssignmentHandler func ...
// func DeleteAssignmentHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

// 	sess := r.Context().Value("User").(*auth.User)
// 	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleDelete, rg.RoleXDelete) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusForbidden).
// 			AddError("You don't have privilege"))
// 		return
// 	}
// 	params := deleteParams{
// 		ID: pr.ByName("assignment_id"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}
// 	if !as.IsAssignmentExist(args.ID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Wrong Assignment ID"))
// 		return
// 	}
// 	if as.IsUserHaveUploadedAsssignment(args.ID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("Forbiden to delete this assignments"))
// 		return
// 	}
// 	tx := conn.DB.MustBegin()
// 	err = as.DeleteAssignment(args.ID, tx)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	err = fs.UpdateStatusFilesByNameID(TableNameAssignments, fs.StatusDeleted, args.ID, tx)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK))
// 	return
// }

// // GradeBySchedule func ...
// func GradeBySchedule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	sess := r.Context().Value("User").(*auth.User)
// 	params := scoreParams{
// 		ScheduleID: ps.ByName("id"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError(err.Error()))
// 		return
// 	}
// 	// Is this fucking user took a schedule
// 	if !cs.IsEnrolled(sess.ID, args.ScheduleID) {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("You don't have privilage for this course"))
// 		return
// 	}
// 	// take grade parameters
// 	allType := [5]string{"ATTENDANCE", "ASSIGNMENT", "QUIZ", "MID", "FINAL"}
// 	gradeParameters, err := cs.SelectGradeParameterByScheduleID(args.ScheduleID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	res := responseScoreSchedule{
// 		ScheduleID: args.ScheduleID,
// 	}
// 	var total float32
// 	for _, val := range allType {
// 		switch val {
// 		case "ATTENDANCE":
// 			res.Attendance = "-"
// 			for _, grade := range gradeParameters {
// 				if grade.Type == val {
// 					var present, totalMeeting int
// 					var percentage float32

// 					meetingsID, err := atd.SelectMeetingIDByScheduleID(sess.ID, args.ScheduleID)
// 					if err != nil {
// 						template.RenderJSONResponse(w, new(template.Response).
// 							SetCode(http.StatusInternalServerError))
// 						return
// 					}

// 					totalMeeting = len(meetingsID)
// 					if totalMeeting > 0 {
// 						present, err = atd.CountByUserMeeting(sess.ID, meetingsID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						percentage = float32(present) * 100 / float32(totalMeeting)
// 					}
// 					totalAttendance := (percentage * grade.Percentage) / 100
// 					res.Attendance = fmt.Sprintf("%.2f", totalAttendance)
// 					total = total + totalAttendance
// 				}
// 			}
// 		case "ASSIGNMENT":
// 			res.Assignment = "-"
// 			for _, grade := range gradeParameters {
// 				if grade.Type == val {
// 					res.Assignment = "0"
// 					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 					if err != nil {
// 						template.RenderJSONResponse(w, new(template.Response).
// 							SetCode(http.StatusInternalServerError))
// 						return
// 					}
// 					if assignmentsID != nil {
// 						scores, err := as.SelectScore(sess.ID, assignmentsID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if len(scores) > 0 {
// 							var t float32
// 							t = 0
// 							for _, score := range scores {
// 								t += score
// 							}
// 							totalAssignment := (t / float32(len(scores)) * grade.Percentage / 100)
// 							total = total + totalAssignment
// 							res.Assignment = fmt.Sprintf("%.2f", totalAssignment)
// 						}
// 					}
// 				}
// 			}
// 		case "QUIZ":
// 			res.Quiz = "-"
// 			for _, grade := range gradeParameters {
// 				if grade.Type == val {
// 					res.Quiz = "0"
// 					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 					if err != nil {
// 						template.RenderJSONResponse(w, new(template.Response).
// 							SetCode(http.StatusInternalServerError))
// 						return
// 					}
// 					if assignmentsID != nil {
// 						scores, err := as.SelectScore(sess.ID, assignmentsID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if len(scores) > 0 {
// 							var t float32
// 							t = 0
// 							for _, score := range scores {
// 								t += score
// 							}
// 							totalQuiz := (t / float32(len(scores)) * grade.Percentage / 100)
// 							total = total + totalQuiz
// 							res.Quiz = fmt.Sprintf("%.2f", totalQuiz)
// 						}
// 					}
// 				}
// 			}
// 		case "MID":
// 			res.Mid = "-"
// 			for _, grade := range gradeParameters {
// 				if grade.Type == val {
// 					res.Mid = "0"
// 					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 					if err != nil {
// 						template.RenderJSONResponse(w, new(template.Response).
// 							SetCode(http.StatusInternalServerError))
// 						return
// 					}
// 					if assignmentsID != nil {
// 						scores, err := as.SelectScore(sess.ID, assignmentsID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if len(scores) > 0 {
// 							var t float32
// 							t = 0
// 							for _, score := range scores {
// 								t += score
// 							}
// 							totalMid := (t / float32(len(scores)) * grade.Percentage / 100)
// 							total = total + totalMid
// 							res.Mid = fmt.Sprintf("%.2f", totalMid)
// 						}
// 					}
// 				}
// 			}
// 		case "FINAL":
// 			res.Final = "-"
// 			for _, grade := range gradeParameters {
// 				if grade.Type == val {
// 					res.Final = "0"
// 					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 					if err != nil {
// 						template.RenderJSONResponse(w, new(template.Response).
// 							SetCode(http.StatusInternalServerError))
// 						return
// 					}
// 					if assignmentsID != nil {
// 						scores, err := as.SelectScore(sess.ID, assignmentsID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if len(scores) > 0 {
// 							var t float32
// 							t = 0
// 							for _, score := range scores {
// 								t += score
// 							}
// 							totalFinal := (t / float32(len(scores)) * grade.Percentage / 100)
// 							total = total + totalFinal
// 							res.Final = fmt.Sprintf("%.2f", totalFinal)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	res.Total = fmt.Sprintf("%.2f", total)
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return
// }

// // GradeSummary func ...
// func GradeSummary(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	sess := r.Context().Value("User").(*auth.User)
// 	listSchedule, err := cs.SelectScheduleIDByUserID(sess.ID, 1)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}

// 	if listSchedule == nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusOK).
// 			SetData(nil))
// 		return
// 	}

// 	schedules, err := cs.SelectByScheduleID(listSchedule, cs.StatusScheduleActive)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}

// 	var res []responseScoreSchedule
// 	for _, c := range schedules {
// 		allType := [5]string{"ATTENDANCE", "ASSIGNMENT", "QUIZ", "MID", "FINAL"}
// 		gradeParameters, err := cs.SelectGradeParameterByScheduleID(c.Schedule.ID)
// 		if err != nil {
// 			template.RenderJSONResponse(w, new(template.Response).
// 				SetCode(http.StatusInternalServerError))
// 			return
// 		}
// 		re := responseScoreSchedule{
// 			ScheduleID: c.Schedule.ID,
// 			CourseName: c.Course.Name,
// 		}
// 		var total float32
// 		for _, val := range allType {
// 			switch val {
// 			case "ATTENDANCE":
// 				re.Attendance = "-"
// 				for _, grade := range gradeParameters {
// 					if grade.Type == val {
// 						re.Attendance = "0"
// 						var present, totalMeeting int
// 						var percentage float32

// 						meetingsID, err := atd.SelectMeetingIDByScheduleID(sess.ID, c.Schedule.ID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}

// 						totalMeeting = len(meetingsID)
// 						if totalMeeting > 0 {
// 							present, err = atd.CountByUserMeeting(sess.ID, meetingsID)
// 							if err != nil {
// 								template.RenderJSONResponse(w, new(template.Response).
// 									SetCode(http.StatusInternalServerError))
// 								return
// 							}
// 							percentage = float32(present) * 100 / float32(totalMeeting)
// 						}
// 						totalAttendance := (percentage * grade.Percentage) / 100
// 						re.Attendance = fmt.Sprintf("%.2f", totalAttendance)
// 						total = total + totalAttendance
// 					}
// 				}
// 			case "ASSIGNMENT":
// 				re.Assignment = "-"
// 				for _, grade := range gradeParameters {
// 					if grade.Type == val {
// 						re.Assignment = "0"
// 						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if assignmentsID != nil {
// 							scores, err := as.SelectScore(sess.ID, assignmentsID)
// 							if err != nil {
// 								template.RenderJSONResponse(w, new(template.Response).
// 									SetCode(http.StatusInternalServerError))
// 								return
// 							}
// 							if len(scores) > 0 {
// 								var t float32
// 								t = 0
// 								for _, score := range scores {
// 									t += score
// 								}
// 								totalAssignment := (t / float32(len(scores)) * grade.Percentage / 100)
// 								total = total + totalAssignment
// 								re.Assignment = fmt.Sprintf("%.2f", totalAssignment)
// 							}
// 						}
// 					}
// 				}
// 			case "QUIZ":
// 				re.Quiz = "-"
// 				for _, grade := range gradeParameters {
// 					if grade.Type == val {
// 						re.Quiz = "0"
// 						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if assignmentsID != nil {
// 							scores, err := as.SelectScore(sess.ID, assignmentsID)
// 							if err != nil {
// 								template.RenderJSONResponse(w, new(template.Response).
// 									SetCode(http.StatusInternalServerError))
// 								return
// 							}
// 							if len(scores) > 0 {
// 								var t float32
// 								t = 0
// 								for _, score := range scores {
// 									t += score
// 								}
// 								totalQuiz := (t / float32(len(scores)) * grade.Percentage / 100)
// 								total = total + totalQuiz
// 								re.Quiz = fmt.Sprintf("%.2f", totalQuiz)
// 							}
// 						}
// 					}
// 				}
// 			case "MID":
// 				re.Mid = "-"
// 				for _, grade := range gradeParameters {
// 					if grade.Type == val {
// 						re.Mid = "0"
// 						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if assignmentsID != nil {
// 							scores, err := as.SelectScore(sess.ID, assignmentsID)
// 							if err != nil {
// 								template.RenderJSONResponse(w, new(template.Response).
// 									SetCode(http.StatusInternalServerError))
// 								return
// 							}
// 							if len(scores) > 0 {
// 								var t float32
// 								t = 0
// 								for _, score := range scores {
// 									t += score
// 								}
// 								totalMid := (t / float32(len(scores)) * grade.Percentage / 100)
// 								total = total + totalMid
// 								re.Mid = fmt.Sprintf("%.2f", totalMid)
// 							}
// 						}
// 					}
// 				}
// 			case "FINAL":
// 				re.Final = "-"
// 				for _, grade := range gradeParameters {
// 					if grade.Type == val {
// 						re.Final = "0"
// 						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
// 						if err != nil {
// 							template.RenderJSONResponse(w, new(template.Response).
// 								SetCode(http.StatusInternalServerError))
// 							return
// 						}
// 						if assignmentsID != nil {
// 							scores, err := as.SelectScore(sess.ID, assignmentsID)
// 							if err != nil {
// 								template.RenderJSONResponse(w, new(template.Response).
// 									SetCode(http.StatusInternalServerError))
// 								return
// 							}
// 							if len(scores) > 0 {
// 								var t float32
// 								t = 0
// 								for _, score := range scores {
// 									t += score
// 								}
// 								totalFinal := (t / float32(len(scores)) * grade.Percentage / 100)
// 								total = total + totalFinal
// 								re.Final = fmt.Sprintf("%.2f", totalFinal)
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 		re.Total = fmt.Sprintf("%.2f", total)
// 		res = append(res, re)
// 	}
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).
// 		SetData(res))
// 	return
// }
