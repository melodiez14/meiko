package assignment

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	as "github.com/melodiez14/meiko/src/module/assignment"
	atd "github.com/melodiez14/meiko/src/module/attendance"
	cs "github.com/melodiez14/meiko/src/module/course"
	fs "github.com/melodiez14/meiko/src/module/file"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	usr "github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// CreateHandler function is
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleCreate, rg.RoleXCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := createParams{
		FilesID:           r.FormValue("file_id"),
		GradeParametersID: r.FormValue("grade_parameter_id"),
		Name:              r.FormValue("name"),
		Description:       r.FormValue("description"),
		Status:            r.FormValue("status"),
		DueDate:           r.FormValue("due_date"),
		Type:              r.FormValue("type"),
		Size:              r.FormValue("size"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// is grade_parameter exist
	if !as.IsExistByGradeParameterID(args.GradeParametersID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Grade parameters id does not exist!"))
		return
	}
	// Insert to table assignments
	tx := conn.DB.MustBegin()
	tableID, err := as.Insert(
		args.GradeParametersID,
		args.Name,
		args.Status,
		args.DueDate,
		args.Description,
		tx,
	)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	// Files null
	if args.FilesID == "" {
		tx.Commit()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetMessage("Success Without files"))
		return
	}

	// Split files id if possible
	filesID := strings.Split(args.FilesID, "~")
	for _, fileID := range filesID {
		// Wrong file code
		if !as.IsFileIDExist(fileID) {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				SetMessage("Wrong file code!"))
			return
		}
		// Update files
		err = fs.UpdateRelation(fileID, fs.TypAssignment, tableID, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
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
		SetMessage("Success!"))
	return

}

// GetAllAssignmentHandler func is ...
func GetAllAssignmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := readParams{
		Page:  r.FormValue("pg"),
		Total: r.FormValue("ttl"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid request"))
		return
	}
	offset := (args.Page - 1) * args.Total
	assignments, err := as.SelectByPage(args.Total, offset)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	var status string
	var res []readResponse
	for _, val := range assignments {

		if val.Status == as.StatusAssignmentActive {
			status = "active"
		} else {
			status = "inactive"
		}

		res = append(res, readResponse{
			Name:             val.Name,
			Description:      val.Description,
			Status:           status,
			DueDate:          val.DueDate,
			GradeParameterID: val.GradeParameterID,
		})
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return

}

// DetailHandler func is ...
func DetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := detailParams{
		IdentityCode: ps.ByName("id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	assignment, err := as.GetByAssignementID(args.IdentityCode)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}

	desc := "-"
	if assignment.Description.Valid {
		desc = assignment.Description.String
	}

	res := detailResponse{
		ID:               assignment.ID,
		Status:           assignment.Status,
		Name:             assignment.Name,
		GradeParameterID: assignment.GradeParameterID,
		Description:      desc,
		DueDate:          assignment.DueDate.Format("Monday, 2 January 2006 15:04:05"),
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

// UpdateHandler func is ...
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleUpdate, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := updatePrams{
		ID:                ps.ByName("id"),
		FilesID:           r.FormValue("file_id"),
		GradeParametersID: r.FormValue("grade_parameter_id"),
		Name:              r.FormValue("name"),
		Description:       r.FormValue("description"),
		Status:            r.FormValue("status"),
		DueDate:           r.FormValue("due_date"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// Params ID is exist
	if !as.IsAssignmentExist(args.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}
	// is grade_parameter exist
	if !as.IsExistByGradeParameterID(args.GradeParametersID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Grade parameters id does not exist!"))
		return
	}

	// Insert to table assignments
	tx := conn.DB.MustBegin()
	err = as.Update(
		args.GradeParametersID,
		args.ID,
		args.Name,
		args.Status,
		args.DueDate,
		args.Description,
		tx,
	)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var filesIDUser = strings.Split(args.FilesID, "~")
	// Get All relations with
	filesIDDB, err := fs.GetByStatus(fs.StatusExist, args.ID)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	// Add new file
	var tableID = strconv.FormatInt(args.ID, 10)
	for _, fileID := range filesIDUser {
		if !fs.IsExistID(fileID) {
			filesIDDB = append(filesIDDB, fileID)
			// Update relation
			err := fs.UpdateRelation(fileID, fs.TypAssignment, tableID, tx)
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
		for _, fileIDUser := range filesIDUser {
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
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Update Assignment Success!"))
	return

}

// CreateHandlerByUser func ...
func CreateHandlerByUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := uploadAssignmentParams{
		UserID:       sess.ID,
		AssignmentID: r.FormValue("assignment_id"),
		Description:  r.FormValue("description"),
		FileID:       r.FormValue("file_id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	scheduleID := cs.GetScheduleID(gradeParameterID)
	isValidAssignment := usr.IsUserTakeSchedule(args.UserID, scheduleID)
	if !isValidAssignment {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Assignment ID"))
		return
	}
	// due date checking
	dueDate, err := as.GetDueDateAssignment(args.AssignmentID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	overdue := !time.Now().Before(dueDate)
	if overdue {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetMessage("Assignment has been overdue"))
		return
	}
	// Insert
	tx := conn.DB.MustBegin()
	err = as.UploadAssignment(args.AssignmentID, args.UserID, args.Description, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("You can only submit once for this assignment"))
		return
	}

	assignmentID := strconv.FormatInt(args.AssignmentID, 10)

	//Update Relations
	if args.FileID != nil {
		for _, fileID := range args.FileID {
			err := fs.UpdateRelation(fileID, fs.TypAssignmentUpload, assignmentID, tx)
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
		SetMessage("Insert Assignment Success!"))
	return
}

// UpdateHandlerByUser func ...
func UpdateHandlerByUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := uploadAssignmentParams{
		FileID:       r.FormValue("file_id"),
		AssignmentID: ps.ByName("id"),
		UserID:       sess.ID,
		Description:  r.FormValue("description"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// Get grade parameters id
	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	scheduleID := cs.GetScheduleID(gradeParameterID)
	isValidAssignment := usr.IsUserTakeSchedule(args.UserID, scheduleID)
	if !isValidAssignment {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Assignment ID"))
		return
	}
	// Update
	tx := conn.DB.MustBegin()
	err = as.UpdateUploadAssignment(args.AssignmentID, args.UserID, args.Description, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Update Failed"))
		return
	}

	tableID := strconv.FormatInt(args.AssignmentID, 10)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Can not convert to int64"))
		return
	}

	// Get All relations with
	filesIDDB, err := fs.SelectByRelation(fs.TypAssignmentUpload, []string{tableID}, &sess.ID)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	_ = filesIDDB

	// Add new file
	// for _, fileID := range args.FileID {
	// 	if !fs.IsExistID(fileID) {
	// 		filesIDDB = append(filesIDDB, fileID)
	// 		// Update relation
	// 		err := fs.UpdateRelation(fileID, fs.TypAssignmentUpload, tableID, tx)
	// 		if err != nil {
	// 			tx.Rollback()
	// 			template.RenderJSONResponse(w, new(template.Response).
	// 				SetCode(http.StatusInternalServerError))
	// 			return
	// 		}
	// 	}
	// }
	// for _, fileIDDB := range filesIDDB {
	// 	isSame := 0
	// 	for _, fileIDUser := range args.FileID {
	// 		if fileIDUser == fileIDDB {
	// 			isSame = 1
	// 		}
	// 	}
	// 	if isSame == 0 {
	// 		err := fs.UpdateStatusFiles(fileIDDB, fs.StatusDeleted, tx)
	// 		if err != nil {
	// 			tx.Rollback()
	// 			template.RenderJSONResponse(w, new(template.Response).
	// 				SetCode(http.StatusInternalServerError))
	// 			return
	// 		}
	// 	}
	// }

	err = tx.Commit()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Update Assignment Success!"))
	return
}

// GetAssignmentHandler func ...
func GetAssignmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	schedulesIDs, err := cs.SelectScheduleIDByUserID(sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	// select Grade parameter
	gradeList, err := cs.SelectGradeParameterByScheduleIDIN(schedulesIDs)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	assignmentID, err := as.SelectAssignmentIDByGradeParameterIN(gradeList)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	submittedAssignmetID := as.SelectSubmittedAssignment(assignmentID, sess.ID)
	var unsubmittedAssignmentID []int64
	for _, val := range assignmentID {
		for _, submitted := range submittedAssignmetID {
			if val != submitted {
				unsubmittedAssignmentID = append(unsubmittedAssignmentID, val)
			}
		}
	}
	res := []listAssignmentResponse{}
	unsubmittedAssignment := as.SelectAssignmentByID(unsubmittedAssignmentID)
	for _, val := range unsubmittedAssignment {
		r := listAssignmentResponse{}
		if val.Status == 1 {
			r.Status = "must_upload"
		} else {
			r.Status = "must_not_upload"
		}
		description := fmt.Sprintf("NULL")
		if val.Description.Valid {
			description = fmt.Sprintf(val.Description.String)
		}
		r.ID = val.ID
		r.Name = val.Name
		r.DueDate = val.DueDate.Format("Monday, 2 January 2006 15:04:05")
		r.Submitted = false
		r.Description = description
		res = append(res, r)
	}
	submittedAssignment := as.SelectAssignmentByID(submittedAssignmetID)
	for _, val := range submittedAssignment {
		r := listAssignmentResponse{}
		if val.Status == 1 {
			r.Status = "must_upload"
		} else {
			r.Status = "must_not_upload"
		}
		description := fmt.Sprintf("NULL")
		if val.Description.Valid {
			description = fmt.Sprintf(val.Description.String)
		}
		r.ID = val.ID
		r.Name = val.Name
		r.DueDate = val.DueDate.Format("Monday, 2 January 2006 15:04:05")
		r.Submitted = true
		r.Description = description
		res = append(res, r)
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return

}

// GetAssignmentByScheduleHandler func ...
func GetAssignmentByScheduleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := listAssignmentsParams{
		ScheduleID: r.FormValue("id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}
	// is correct schedule id
	if !cs.IsExistScheduleID(args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("You do not took this course"))
		return
	}
	gradeParamsID, err := cs.GetGradeParametersIDByScheduleID(args.ScheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Error"))
		return
	}
	assignments := as.SelectByGradeParametersID(gradeParamsID)
	res := []listAssignmentResponse{}
	for _, val := range assignments {
		status := "must_not_upload"
		submited := false
		var description string
		if val.Description.Valid {
			description = fmt.Sprintf(val.Description.String)
		}
		if val.Status == 1 {
			status = "must_upload"
			if as.IsUploaded(val.ID) {
				submited = true
			}
		}
		res = append(res, listAssignmentResponse{
			Submitted:   submited,
			ID:          val.ID,
			Name:        val.Name,
			Status:      status,
			Description: description,
			DueDate:     val.DueDate.Format("Monday, 2 January 2006 15:04:05"),
		})
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).SetData(res))
	return
}

// GetAssignmentHandler func ...
func GetAssignmentDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	// Get assignment ID
	params := readDetailParam{
		AssignmentID: ps.ByName("id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError(err.Error()))
		return
	}
	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	if gradeParameterID == 0 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError("Grade Parameter ID not found"))
		return
	}
	scheduleID := cs.GetScheduleID(gradeParameterID)
	if scheduleID == 0 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError("You do not took this course"))
		return
	}
	a, err := as.GetByAssignementID(args.AssignmentID)

	status := "must_not_upload"
	isUploaded := false
	score := "-"
	buttonType := "none"

	if as.IsUploaded(args.AssignmentID) {
		isUploaded = true
		buttonType = "update"
		score = fmt.Sprintf("%.3g", as.GetScoreByIDUser(args.AssignmentID, sess.ID))
	}

	if a.Status == 1 {
		status = "must_upload"
		buttonType = "add"
	} else {
		buttonType = "none"
	}

	description := fmt.Sprintf("NULL")
	if a.Description.Valid {
		description = fmt.Sprintf(a.Description.String)
	}
	prefix := a.Name
	if len(a.Name) > 9 {
		p := a.Name[0:8]
		s := strings.Fields(p)
		prefix = strings.Join(s, "_")
	}

	tableID := []string{strconv.FormatInt(args.AssignmentID, 10)}
	assistantFile, err := fs.SelectByRelation(fs.TypAssignment, tableID, nil)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	studentFile, err := fs.SelectByRelation(fs.TypAssignmentUpload, tableID, &sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	resAssistantFile := []file{}
	for _, val := range assistantFile {
		resAssistantFile = append(resAssistantFile, file{
			Name: val.Name,
			URL:  fmt.Sprintf("/api/v1/file/assignment/%s.%s", val.ID, val.Extension),
		})
	}

	resStudentFile := []file{}
	for _, val := range studentFile {
		resStudentFile = append(resStudentFile, file{
			Name: val.Name,
			URL:  fmt.Sprintf("/api/v1/file/assignment/%s.%s", val.ID, val.Extension),
		})
	}

	res := detailResponseUser{
		ID:             a.ID,
		Name:           a.Name,
		Status:         status,
		Description:    description,
		DueDate:        a.DueDate.Format("Monday, 2 January 2006 15:04:05"),
		Score:          score,
		FilesName:      a.UpdatedAt.Format("2006_01_02") + "-" + prefix,
		UploadedStatus: isUploaded,
		ButtonType:     buttonType,
		AssistantFile:  resAssistantFile,
		StudentFile:    resStudentFile,
	}
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError("Error"))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).SetData(res))
	return
}

// GetUploadedDetailHandler func ...
func GetUploadedDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := readUploadedDetailParams{
		UserID:       ps.ByName("id"),
		ScheduleID:   ps.ByName("schedule_id"),
		AssignmentID: ps.ByName("assignment_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError(err.Error()))
		return
	}
	// Is Valid User ID
	if sess.ID != args.UserID {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError("Wrong User ID"))
		return
	}
	if !usr.IsUserTakeSchedule(sess.ID, args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError("Wrong Schedule ID"))
		return
	}
	// Get Assignments Detail
	assignment, err := as.GetUploadedAssignmentByID(args.AssignmentID, sess.ID)
	key := fmt.Sprintf("%d%d", args.AssignmentID, sess.ID)
	tableID, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError(err.Error()))
		return
	}

	// Get File
	files, err := fs.GetByUserIDTableID(sess.ID, tableID, fs.TypAssignmentUpload)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	res := readUploadedDetailArgs{
		UserID:       args.UserID,
		ScheduleID:   args.ScheduleID,
		AssignmentID: args.AssignmentID,
		Name:         assignment.Name,
		Description:  assignment.DescriptionAssignment,
		PathFile:     files,
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return

}

// CreateScoreHandler func ...
func CreateScoreHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	// get schedule, assignment
	params := createScoreParams{
		ScheduleID:   ps.ByName("schedule_id"),
		AssignmentID: ps.ByName("assignment_id"),
		Users:        r.FormValue("users"),
	}
	// validate
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// is schedule_id match with assignments_id?
	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	if !as.IsAssignmentExistByGradeParameterID(args.AssignmentID, gradeParameterID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}

	// is user (admin) took that course?
	if !cs.IsAssistant(sess.ID, args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("you do not took this course!"))
		return
	}
	usersID, err := usr.SelectIDByIdentityCode(args.IdentityCode)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	if len(usersID) != len(args.IdentityCode) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("There is invalid identity code"))
		return
	}
	if !cs.IsAllUsersEnrolled(args.ScheduleID, usersID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("There is a invalid identity code"))
		return
	}
	tx := conn.DB.MustBegin()
	// Insert
	err = as.CreateScore(args.AssignmentID, usersID, args.Score, tx)
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
		SetMessage("Add score assignment successfully"))
	return

}

// GetDetailAssignmentByAdmin func ...
func GetDetailAssignmentByAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	// get schedule, assignment
	params := detailAssignmentParams{
		ScheduleID:   ps.ByName("schedule_id"),
		AssignmentID: ps.ByName("assignment_id"),
	}
	// validate
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// is schedule_id match with assignments_id?
	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	if !as.IsAssignmentExistByGradeParameterID(args.AssignmentID, gradeParameterID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}

	// is user (admin) took that course?
	if !cs.IsAssistant(sess.ID, args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("you do not took this course!"))
		return
	}

	// get assignment detail
	assignments, err := as.GetAssignmentByID(args.AssignmentID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	res := detailAssignmentResponse{}
	userList := []userAssignment{}
	// is assingment must upload or not?
	if as.IsAssignmentMustUpload(args.AssignmentID) {
		userAssignments, err := as.SelectUserAssignmentsByStatusID(args.AssignmentID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		for _, value := range userAssignments {
			userList = append(userList, userAssignment{
				UserID: value.UserID,
				Name:   value.Name,
				Grade:  0,
			})
		}
		res = detailAssignmentResponse{
			Name:          assignments.Name,
			Description:   assignments.Description,
			DueDate:       assignments.DueDate,
			IsCreateScore: false,
			Praktikan:     userList,
		}
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(res))
		return
	}
	userAssignments, err := usr.SelectUserByScheduleID(args.ScheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	for _, value := range userAssignments {
		userList = append(userList, userAssignment{
			UserID: value.UserID,
			Name:   value.Name,
			Grade:  0,
		})
	}
	res = detailAssignmentResponse{
		Name:          assignments.Name,
		Description:   assignments.Description,
		DueDate:       assignments.DueDate,
		IsCreateScore: true,
		Praktikan:     userList,
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return

}

//ScoreHandler func ...
func ScoreHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := updateScoreParams{
		ScheduleID:   ps.ByName("id"),
		AssignmentID: ps.ByName("assignment_id"),
		UserID:       r.FormValue("user_id"),
		Score:        r.FormValue("score"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	if !as.IsAssignmentExistByGradeParameterID(args.AssignmentID, gradeParameterID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}
	// User (Praktikan) took that course
	if !cs.IsEnrolled(args.UserID, args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("User do not took this course!"))
		return
	}
	// check schedule have assignments
	if !cs.IsUserHasUploadedFile(args.AssignmentID, args.UserID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("User has not uploaded yet assignment"))
		return
	}
	tx := conn.DB.MustBegin()
	err = as.UpdateScoreAssignment(args.AssignmentID, args.UserID, args.Score, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError("Can not update score"))
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
		SetMessage("Score Updated"))
	return
}

// GetUploadedAssignmentByAdminHandler func ...
func GetUploadedAssignmentByAdminHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := readUploadedAssignmentParams{
		ScheduleID:   ps.ByName("id"),
		AssignmentID: ps.ByName("assignment_id"),
		Page:         r.FormValue("pg"),
		Total:        r.FormValue("ttl"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).AddError(err.Error()))
		return
	}
	offset := (args.Page - 1) * args.Total
	// Check schedule id
	if !cs.IsExistScheduleID(args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Schdule ID does not exist"))
		return
	}
	// Check assignment id
	if !as.IsAssignmentExist(args.AssignmentID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Assignment ID does not exist"))
		return
	}
	// Get all data p_users_assignment
	assignments, err := as.GetAllUserAssignmentByAssignmentID(args.AssignmentID, args.Total, offset)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	//files, err := fs.GetByTableIDName(args.AssignmentID, TableNameUserAssignments)
	// get all data files relations
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).SetData(assignments))
	// serve json
}

// DeleteAssignmentHandler func ...
func DeleteAssignmentHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleDelete, rg.RoleXDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := deleteParams{
		ID: pr.ByName("assignment_id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	if !as.IsAssignmentExist(args.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Wrong Assignment ID"))
		return
	}
	if as.IsUserHaveUploadedAsssignment(args.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Forbiden to delete this assignments"))
		return
	}
	tx := conn.DB.MustBegin()
	err = as.DeleteAssignment(args.ID, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	err = fs.UpdateStatusFilesByNameID(TableNameAssignments, fs.StatusDeleted, args.ID, tx)
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
		SetCode(http.StatusOK))
	return
}

// GradeBySchedule func ...
func GradeBySchedule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	params := scoreParams{
		ScheduleID: ps.ByName("id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	// Is this fucking user took a schedule
	if !cs.IsEnrolled(sess.ID, args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("You don't have privilage for this course"))
		return
	}
	// take grade parameters
	allType := [5]string{"ATTENDANCE", "ASSIGNMENT", "QUIZ", "MID", "FINAL"}
	gradeParameters, err := cs.SelectGradeParameterByScheduleID(args.ScheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	res := responseScoreSchedule{
		ScheduleID: args.ScheduleID,
	}
	var total float32
	for _, val := range allType {
		switch val {
		case "ATTENDANCE":
			res.Attendance = "-"
			for _, grade := range gradeParameters {
				if grade.Type == val {
					var present, totalMeeting int
					var percentage float32

					meetingsID, err := atd.SelectMeetingIDByScheduleID(sess.ID, args.ScheduleID)
					if err != nil {
						template.RenderJSONResponse(w, new(template.Response).
							SetCode(http.StatusInternalServerError))
						return
					}

					totalMeeting = len(meetingsID)
					if totalMeeting > 0 {
						present, err = atd.CountByUserMeeting(sess.ID, meetingsID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						percentage = float32(present) * 100 / float32(totalMeeting)
					}
					totalAttendance := (percentage * grade.Percentage) / 100
					res.Attendance = fmt.Sprintf("%.2f", totalAttendance)
					total = total + totalAttendance
				}
			}
		case "ASSIGNMENT":
			res.Assignment = "-"
			for _, grade := range gradeParameters {
				if grade.Type == val {
					res.Assignment = "0"
					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
					if err != nil {
						template.RenderJSONResponse(w, new(template.Response).
							SetCode(http.StatusInternalServerError))
						return
					}
					if assignmentsID != nil {
						scores, err := as.SelectScore(sess.ID, assignmentsID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if len(scores) > 0 {
							var t float32
							t = 0
							for _, score := range scores {
								t += score
							}
							totalAssignment := (t / float32(len(scores)) * grade.Percentage / 100)
							total = total + totalAssignment
							res.Assignment = fmt.Sprintf("%.2f", totalAssignment)
						}
					}
				}
			}
		case "QUIZ":
			res.Quiz = "-"
			for _, grade := range gradeParameters {
				if grade.Type == val {
					res.Quiz = "0"
					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
					if err != nil {
						template.RenderJSONResponse(w, new(template.Response).
							SetCode(http.StatusInternalServerError))
						return
					}
					if assignmentsID != nil {
						scores, err := as.SelectScore(sess.ID, assignmentsID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if len(scores) > 0 {
							var t float32
							t = 0
							for _, score := range scores {
								t += score
							}
							totalQuiz := (t / float32(len(scores)) * grade.Percentage / 100)
							total = total + totalQuiz
							res.Quiz = fmt.Sprintf("%.2f", totalQuiz)
						}
					}
				}
			}
		case "MID":
			res.Mid = "-"
			for _, grade := range gradeParameters {
				if grade.Type == val {
					res.Mid = "0"
					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
					if err != nil {
						template.RenderJSONResponse(w, new(template.Response).
							SetCode(http.StatusInternalServerError))
						return
					}
					if assignmentsID != nil {
						scores, err := as.SelectScore(sess.ID, assignmentsID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if len(scores) > 0 {
							var t float32
							t = 0
							for _, score := range scores {
								t += score
							}
							totalMid := (t / float32(len(scores)) * grade.Percentage / 100)
							total = total + totalMid
							res.Mid = fmt.Sprintf("%.2f", totalMid)
						}
					}
				}
			}
		case "FINAL":
			res.Final = "-"
			for _, grade := range gradeParameters {
				if grade.Type == val {
					res.Final = "0"
					assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
					if err != nil {
						template.RenderJSONResponse(w, new(template.Response).
							SetCode(http.StatusInternalServerError))
						return
					}
					if assignmentsID != nil {
						scores, err := as.SelectScore(sess.ID, assignmentsID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if len(scores) > 0 {
							var t float32
							t = 0
							for _, score := range scores {
								t += score
							}
							totalFinal := (t / float32(len(scores)) * grade.Percentage / 100)
							total = total + totalFinal
							res.Final = fmt.Sprintf("%.2f", totalFinal)
						}
					}
				}
			}
		}
	}
	res.Total = fmt.Sprintf("%.2f", total)
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

// GradeSummary func ...
func GradeSummary(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	listSchedule, err := cs.SelectScheduleIDByUserID(sess.ID, 1)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if listSchedule == nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(nil))
		return
	}

	schedules, err := cs.SelectByScheduleID(listSchedule, cs.StatusScheduleActive)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var res []responseScoreSchedule
	for _, c := range schedules {
		allType := [5]string{"ATTENDANCE", "ASSIGNMENT", "QUIZ", "MID", "FINAL"}
		gradeParameters, err := cs.SelectGradeParameterByScheduleID(c.Schedule.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		re := responseScoreSchedule{
			ScheduleID: c.Schedule.ID,
			CourseName: c.Course.Name,
		}
		var total float32
		for _, val := range allType {
			switch val {
			case "ATTENDANCE":
				re.Attendance = "-"
				for _, grade := range gradeParameters {
					if grade.Type == val {
						re.Attendance = "0"
						var present, totalMeeting int
						var percentage float32

						meetingsID, err := atd.SelectMeetingIDByScheduleID(sess.ID, c.Schedule.ID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}

						totalMeeting = len(meetingsID)
						if totalMeeting > 0 {
							present, err = atd.CountByUserMeeting(sess.ID, meetingsID)
							if err != nil {
								template.RenderJSONResponse(w, new(template.Response).
									SetCode(http.StatusInternalServerError))
								return
							}
							percentage = float32(present) * 100 / float32(totalMeeting)
						}
						totalAttendance := (percentage * grade.Percentage) / 100
						re.Attendance = fmt.Sprintf("%.2f", totalAttendance)
						total = total + totalAttendance
					}
				}
			case "ASSIGNMENT":
				re.Assignment = "-"
				for _, grade := range gradeParameters {
					if grade.Type == val {
						re.Assignment = "0"
						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if assignmentsID != nil {
							scores, err := as.SelectScore(sess.ID, assignmentsID)
							if err != nil {
								template.RenderJSONResponse(w, new(template.Response).
									SetCode(http.StatusInternalServerError))
								return
							}
							if len(scores) > 0 {
								var t float32
								t = 0
								for _, score := range scores {
									t += score
								}
								totalAssignment := (t / float32(len(scores)) * grade.Percentage / 100)
								total = total + totalAssignment
								re.Assignment = fmt.Sprintf("%.2f", totalAssignment)
							}
						}
					}
				}
			case "QUIZ":
				re.Quiz = "-"
				for _, grade := range gradeParameters {
					if grade.Type == val {
						re.Quiz = "0"
						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if assignmentsID != nil {
							scores, err := as.SelectScore(sess.ID, assignmentsID)
							if err != nil {
								template.RenderJSONResponse(w, new(template.Response).
									SetCode(http.StatusInternalServerError))
								return
							}
							if len(scores) > 0 {
								var t float32
								t = 0
								for _, score := range scores {
									t += score
								}
								totalQuiz := (t / float32(len(scores)) * grade.Percentage / 100)
								total = total + totalQuiz
								re.Quiz = fmt.Sprintf("%.2f", totalQuiz)
							}
						}
					}
				}
			case "MID":
				re.Mid = "-"
				for _, grade := range gradeParameters {
					if grade.Type == val {
						re.Mid = "0"
						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if assignmentsID != nil {
							scores, err := as.SelectScore(sess.ID, assignmentsID)
							if err != nil {
								template.RenderJSONResponse(w, new(template.Response).
									SetCode(http.StatusInternalServerError))
								return
							}
							if len(scores) > 0 {
								var t float32
								t = 0
								for _, score := range scores {
									t += score
								}
								totalMid := (t / float32(len(scores)) * grade.Percentage / 100)
								total = total + totalMid
								re.Mid = fmt.Sprintf("%.2f", totalMid)
							}
						}
					}
				}
			case "FINAL":
				re.Final = "-"
				for _, grade := range gradeParameters {
					if grade.Type == val {
						re.Final = "0"
						assignmentsID, err := as.SelectAssignmentIDByGradeParameter(grade.ID)
						if err != nil {
							template.RenderJSONResponse(w, new(template.Response).
								SetCode(http.StatusInternalServerError))
							return
						}
						if assignmentsID != nil {
							scores, err := as.SelectScore(sess.ID, assignmentsID)
							if err != nil {
								template.RenderJSONResponse(w, new(template.Response).
									SetCode(http.StatusInternalServerError))
								return
							}
							if len(scores) > 0 {
								var t float32
								t = 0
								for _, score := range scores {
									t += score
								}
								totalFinal := (t / float32(len(scores)) * grade.Percentage / 100)
								total = total + totalFinal
								re.Final = fmt.Sprintf("%.2f", totalFinal)
							}
						}
					}
				}
			}
		}
		re.Total = fmt.Sprintf("%.2f", total)
		res = append(res, re)
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}
