package assignment

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	as "github.com/melodiez14/meiko/src/module/assignment"
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
	TableID, err := as.Insert(
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
	tableName := "assignments"
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
		err = fs.UpdateRelation(fileID, tableName, TableID, tx)
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
	u, err := as.GetByAssignementID(args.IdentityCode)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}

	var status string
	switch u.Assignment.Status {
	case 0:
		status = "inactive"
	case 1:
		status = "active"
	}

	res := detailResponse{
		ID:               u.Assignment.ID,
		Status:           status,
		Name:             u.Assignment.Name,
		GradeParameterID: u.Assignment.GradeParameterID,
		Description:      u.Assignment.Description,
		DueDate:          u.Assignment.DueDate,
		FilesName:        u.File.Name,
		Mime:             u.File.Mime,
		Percentage:       u.GradeParameter.Percentage,
		Type:             u.GradeParameter.Type,
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
	var tableID = strconv.FormatInt(args.ID, 10)
	// Get All relations with
	filesIDDB, err := fs.GetByStatus(fs.StatusExist, args.ID)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	// Add new file
	for _, fileID := range filesIDUser {
		if !fs.IsExistID(fileID) {
			filesIDDB = append(filesIDDB, fileID)
			// Update relation
			err := fs.UpdateRelation(fileID, TableNameAssignments, tableID, tx)
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
		SetMessage("Update Assigment Success!"))
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

	tableID := fmt.Sprintf("%d%d", args.AssignmentID, args.UserID)
	//Update Relations
	if args.FileID != nil {
		for _, fileID := range args.FileID {
			err := fs.UpdateRelation(fileID, TableNameUserAssignments, tableID, tx)
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
	key := fmt.Sprintf("%d%d", args.AssignmentID, args.UserID)
	tableID, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Can not convert to int64"))
		return
	}
	// Get All relations with
	filesIDDB, err := fs.GetByStatus(fs.StatusExist, tableID)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	// Add new file
	for _, fileID := range args.FileID {
		if !fs.IsExistID(fileID) {
			filesIDDB = append(filesIDDB, fileID)
			// Update relation
			err := fs.UpdateRelation(fileID, TableNameUserAssignments, key, tx)
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
		for _, fileIDUser := range args.FileID {
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

// GetAssignmentByScheduleHandler func ...
func GetAssignmentByScheduleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := listAssignmentsParams{
		Page:       r.FormValue("pg"),
		Total:      r.FormValue("ttl"),
		ScheduleID: ps.ByName("schedule_id"),
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
	offset := (args.Page - 1) * args.Total
	assignments, err := as.GetByGradeParametersID(gradeParamsID, args.Total, offset)
	res := []listAssignmentResponse{}
	for _, val := range assignments {
		var status int8
		status = 0
		if val.Status == 0 {
			status = 2
		} else {
			if as.IsUploaded(val.ID) {
				status = 1
			}
		}
		res = append(res, listAssignmentResponse{
			ID:          val.ID,
			Name:        val.Name,
			Status:      status,
			Description: val.Description,
			DueDate:     val.DueDate,
		})
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).SetData(res))
	return

	// get all assignment by schedule join with p_users_assignment
	// serve
}

// GetAssignmentHandler func ...
func GetAssignmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	res, err := as.GetByAssignementID(args.AssignmentID)
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
	files, err := fs.GetByUserIDTableIDName(sess.ID, tableID, TableNameUserAssignments)
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

// func GetIncompleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	u := r.Context().Value("User").(*auth.User)

// 	a, err := assignment.SelectIncompleteByUserID(u.ID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError).
// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"github.com/julienschmidt/httprouter"
// 	as "github.com/melodiez14/meiko/src/module/assignment"
// 	cs "github.com/melodiez14/meiko/src/module/course"
// 	fs "github.com/melodiez14/meiko/src/module/file"
// 	rg "github.com/melodiez14/meiko/src/module/rolegroup"
// 	usr "github.com/melodiez14/meiko/src/module/user"
// 	"github.com/melodiez14/meiko/src/util/auth"
// 	"github.com/melodiez14/meiko/src/util/conn"
// 	"github.com/melodiez14/meiko/src/webserver/template"
// )

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
// 	u, err := as.GetByAssignementID(args.IdentityCode)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusNotFound))
// 		return
// 	}

// 	var status string
// 	switch u.Assignment.Status {
// 	case 0:
// 		status = "inactive"
// 	case 1:
// 		status = "active"
// 	}

// 	res := detailResponse{
// 		ID:               u.Assignment.ID,
// 		Status:           status,
// 		Name:             u.Assignment.Name,
// 		GradeParameterID: u.Assignment.GradeParameterID,
// 		Description:      u.Assignment.Description,
// 		DueDate:          u.Assignment.DueDate,
// 		FilesName:        u.File.Name,
// 		Mime:             u.File.Mime,
// 		Percentage:       u.GradeParameter.Percentage,
// 		Type:             u.GradeParameter.Type,
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
// 		if !fs.IsIDActive(fs.StatusExist, fileID, tableID) {
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
// 		SetMessage("Update Assigment Success!"))
// 	return

// }

// // CreateHandlerByUser func ...
// func CreateHandlerByUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	sess := r.Context().Value("User").(*auth.User)
// 	params := uploadAssignmentParams{
// 		FileID:       r.FormValue("file_id"),
// 		AssignmentID: r.FormValue("assignment_id"),
// 		UserID:       sess.ID,
// 		Subject:      r.FormValue("subject"),
// 		Description:  r.FormValue("description"),
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
// 	// Insert
// 	tx := conn.DB.MustBegin()
// 	err = as.UploadAssignment(args.AssignmentID, args.UserID, args.Description, tx)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).
// 			AddError("You can only submit one time for Assignment"))
// 		return
// 	}
// 	var filesID = strings.Split(args.FileID, "~")
// 	tableID := fmt.Sprintf("%d%d", args.AssignmentID, args.UserID)
// 	//Update Relations
// 	if args.FileID != "" {
// 		for _, fileID := range filesID {
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
// 		Subject:      r.FormValue("subject"),
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
// 	var filesIDUser = strings.Split(args.FileID, "~")
// 	// Get All relations with
// 	filesIDDB, err := fs.GetByStatus(fs.StatusExist, tableID)
// 	if err != nil {
// 		tx.Rollback()
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusInternalServerError))
// 		return
// 	}
// 	// Add new file
// 	for _, fileID := range filesIDUser {
// 		if !fs.IsIDActive(fs.StatusExist, fileID, key) {
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

// // GetAssignmentByScheduleHandler func ...
// func GetAssignmentByScheduleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	params := listAssignmentsParams{
// 		Page:       r.FormValue("pg"),
// 		Total:      r.FormValue("ttl"),
// 		ScheduleID: ps.ByName("schedule_id"),
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
// 	offset := (args.Page - 1) * args.Total
// 	res, err := as.GetByGradeParametersID(gradeParamsID, args.Total, offset)
// 	template.RenderJSONResponse(w, new(template.Response).
// 		SetCode(http.StatusOK).SetData(res))
// 	return

// 	// get all assignment by schedule join with p_users_assignment
// 	// serve
// }

// // GetAssignmentHandler func ...
// func GetAssignmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// Get assignment ID
// 	params := readDetailParam{
// 		AssignmentID: ps.ByName("id"),
// 	}
// 	args, err := params.validate()
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError(err.Error()))
// 		return
// 	}
// 	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
// 	if gradeParameterID == 0 {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError("Grade Parameter ID not found"))
// 		return
// 	}
// 	scheduleID := cs.GetScheduleID(gradeParameterID)
// 	if scheduleID == 0 {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError("You do not took this course"))
// 		return
// 	}
// 	res, err := as.GetByAssignementID(args.AssignmentID)
// 	if err != nil {
// 		template.RenderJSONResponse(w, new(template.Response).
// 			SetCode(http.StatusBadRequest).AddError("Error"))
// 		return
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

// //UpdateScoreByAdminHandler func ...
// func UpdateScoreByAdminHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// // func GetIncompleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// // 	u := r.Context().Value("User").(*auth.User)

// // 	a, err := assignment.SelectIncompleteByUserID(u.ID)
// // 	if err != nil {
// // 		template.RenderJSONResponse(w, new(template.Response).
// // 			SetCode(http.StatusInternalServerError).
// // 			AddError(err.Error()))
// // 		return
// // 	}

// // 	res := []summaryResponse{}
// // 	for _, v := range a {
// // 		res = append(res, summaryResponse{
// // 			ID:   v.ID,
// // 			Name: v.Name,
// // 		})
// // 	}

// // 	template.RenderJSONResponse(w, new(template.Response).
// // 		SetCode(http.StatusOK).
// // 		SetData(res))
// // 	return
// // }

// // func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// // 	var res []profileSummaryResponse
// // 	u := r.Context().Value("User").(*auth.User)

// // 	// get all enrolled course using using userID
// // 	courses, err := course.GetByUserID(u.ID)
// // 	if err != nil {
// // 		template.RenderJSONResponse(w, new(template.Response).
// // 			SetCode(http.StatusInternalServerError).
// // 			AddError(err.Error()))
// // 		return
// // 	}

// // 	// if there is no enrolled course
// // 	if len(courses) < 1 {
// // 		template.RenderJSONResponse(w, new(template.Response).
// // 			SetCode(http.StatusOK).
// // 			SetData(res))
// // 	}

// // 	// get completed assignments have been posted in all courses
// // 	ca, err := assignment.SelectCompleteByUserID(u.ID)
// // 	if err != nil {
// // 		template.RenderJSONResponse(w, new(template.Response).
// // 			SetCode(http.StatusInternalServerError).
// // 			AddError(err.Error()))
// // 		return
// // 	}

// // 	// iterate all courses to get the summary
// // 	for _, v := range courses {

// // 		pSummary := profileSummaryResponse{
// // 			CourseName: v.Name,
// // 			Complete:   0,
// // 			Incomplete: 0,
// // 		}

// // 		// get all assignments per courses
// // 		assignments, err := assignment.SelectByCourseID(v.ID)
// // 		if err != nil {
// // 			template.RenderJSONResponse(w, new(template.Response).
// // 				SetCode(http.StatusInternalServerError).
// // 				AddError(err.Error()))
// // 			return
// // 		}

// // 		// compare course assignments with all assignments in p_users_assignments
// // 		// if course assignment id exist in p_users_assignments then increment the complete
// // 		// else increment the incomplete
// // 		for _, a := range assignments {
// // 			if helper.Int64InSlice(a.ID, ca) {
// // 				pSummary.Complete++
// // 			} else {
// // 				pSummary.Incomplete++
// // 			}
// // 		}

// // 		// append summary per courses
// // 		res = append(res, pSummary)
// // 	}

// // 	template.RenderJSONResponse(w, new(template.Response).
// // 		SetCode(http.StatusOK).
// // 		SetData(res))
// // 	return
// // }
