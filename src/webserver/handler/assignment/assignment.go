package assignment

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"

	"github.com/julienschmidt/httprouter"
	asg "github.com/melodiez14/meiko/src/module/assignment"
	att "github.com/melodiez14/meiko/src/module/attendance"
	cs "github.com/melodiez14/meiko/src/module/course"
	fl "github.com/melodiez14/meiko/src/module/file"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	usr "github.com/melodiez14/meiko/src/module/user"
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
		filter:     r.FormValue("filter"),
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

	var desc, score, status string
	var isAllowUpload bool
	for _, assignment := range assignments {
		submit, exist := submitMap[assignment.ID]
		desc = "-"
		score = "-"
		status = "unsubmitted"
		isAllowUpload = true
		if assignment.DueDate.Before(time.Now()) {
			status = "overdue"
			isAllowUpload = false
		}
		if exist {
			status = "submitted"
			if submit.Score.Valid {
				status = "done"
				isAllowUpload = false
				score = fmt.Sprintf("%.3g", submit.Score.Float64)
			}
		}
		if assignment.Status == asg.StatusUploadNotRequired {
			isAllowUpload = false
		}
		if assignment.Description.Valid {
			desc = assignment.Description.String
		}
		if args.filter.Valid {
			if args.filter.String != status {
				continue
			}
		}
		resp = append(resp, getResponse{
			ID:            assignment.ID,
			DueDate:       assignment.DueDate.Format("Monday, 2 January 2006 15:04:05"),
			Name:          assignment.Name,
			Score:         score,
			Status:        status,
			IsAllowUpload: isAllowUpload,
			Description:   desc,
			UpdatedAt:     assignment.UpdatedAt.Format("Monday, 2 January 2006 15:04:05"),
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
	submittedDesc := ""
	isAllowUpload := true
	if assignment.DueDate.Before(time.Now()) {
		status = "overdue"
		isAllowUpload = false
	}
	if submitted != nil {
		status = "submitted"
		submittedDesc = submitted.Description.String
		submittedDate = submitted.UpdatedAt.Format("Monday, 2 January 2006 15:04:05")
		if submitted.Score.Valid {
			isAllowUpload = false
			status = "done"
			score = fmt.Sprintf("%.3g", submitted.Score.Float64)
		}
	}
	if assignment.Status == asg.StatusUploadNotRequired {
		status = "notrequired"
		isAllowUpload = false
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
			ID:           val.ID,
			Name:         fmt.Sprintf("%s.%s", val.Name, val.Extension),
			URL:          fmt.Sprintf("/api/v1/file/assignment/%s.%s", val.ID, val.Extension),
			URLThumbnail: helper.MimeToThumbnail(val.Mime),
		})
	}

	// file from student
	rSubmittedFile := []file{}
	for _, val := range submittedFile {
		rSubmittedFile = append(rSubmittedFile, file{
			ID:           val.ID,
			Name:         fmt.Sprintf("%s.%s", val.Name, val.Extension),
			URL:          fmt.Sprintf("/api/v1/file/assignment/%s.%s", val.ID, val.Extension),
			URLThumbnail: helper.MimeToThumbnail(val.Mime),
		})
	}

	resp := getDetailResponse{
		ID:                   assignment.ID,
		Name:                 assignment.Name,
		Status:               status,
		Description:          assignment.Description.String,
		DueDate:              assignment.DueDate.Format("Monday, 2 January 2006 15:04:05"),
		Score:                score,
		CreatedAt:            assignment.CreatedAt.Format("Monday, 2 January 2006"),
		UpdatedAt:            assignment.UpdatedAt.Format("Monday, 2 January 2006"),
		AssignmentFile:       rAsgFile,
		IsAllowUpload:        isAllowUpload,
		SubmittedDescription: submittedDesc,
		SubmittedFile:        rSubmittedFile,
		SubmittedDate:        submittedDate,
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

// GetAvailableGP ..
func GetAvailableGP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXCreate, rg.RoleCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := availableParams{
		id: ps.ByName("id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if !cs.IsAssistant(sess.ID, args.id) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	gps, err := cs.SelectGPBySchedule([]int64{args.id})
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	res := []respGP{}
	for _, gp := range gps {
		res = append(res, respGP{
			ID:   gp.ID,
			Name: gp.Type,
		})
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).SetData(res))
	return

}

// ReadHandler ..
func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	resp := readResponse{Assignments: []read{}}
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXCreate, rg.RoleCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readParams{
		scheduleID: r.FormValue("schedule_id"),
		page:       r.FormValue("pg"),
		total:      r.FormValue("ttl"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if !cs.IsAssistant(sess.ID, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	gps, err := cs.SelectGPBySchedule([]int64{args.scheduleID})
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

	offset := (args.page - 1) * args.total
	assignments, count, err := asg.SelectByPage(gpsID, args.total, offset, true)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	rAssignment := []read{}
	for _, val := range assignments {
		rAssignment = append(rAssignment, read{
			ID:        val.ID,
			Name:      val.Name,
			URL:       fmt.Sprintf("/api/v1/filerouter/?id=%d&payload=assignment&role=assistant", val.ID),
			DueDate:   val.DueDate.Format("Monday, 2 January 2006 15:04:05"),
			UpdatedAt: val.UpdatedAt.Format("Monday, 2 January 2006 15:04:05"),
		})
	}

	totalPage := count / args.total
	if count%args.total > 0 {
		totalPage++
	}

	resp = readResponse{
		Assignments: rAssignment,
		Page:        args.page,
		TotalPage:   totalPage,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// CreateHandler ..
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXCreate, rg.RoleCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := createParams{
		name:             r.FormValue("name"),
		description:      r.FormValue("description"),
		dueDate:          r.FormValue("due_date"),
		filesID:          r.FormValue("files_id"),
		gpID:             r.FormValue("grade_parameter_id"),
		status:           r.FormValue("status"),
		allowedTypesFile: r.FormValue("allowed_types"),
		maxFile:          r.FormValue("max_file"),
		maxSizeFile:      r.FormValue("max_size"),
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
	if len(args.filesID) > 0 {
		filesID, err := fl.SelectIDStatusByID(args.filesID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		for _, inFile := range args.filesID {
			for _, dbFile := range filesID {
				if inFile == dbFile.ID && dbFile.Status == fl.StatusDeleted {
					template.RenderJSONResponse(w, new(template.Response).
						SetCode(http.StatusBadRequest).
						AddError("ID file has been deleted, you can not use it again"))
					return
				}
			}
		}
	}
	tx := conn.DB.MustBegin()
	id, err := asg.Insert(args.name, args.description, args.gpID, args.maxSizeFile, args.maxFile, args.dueDate, args.status, tx)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	idStr := strconv.FormatInt(id, 10)
	if len(args.filesID) > 0 {
		for _, fileID := range args.filesID {
			if fl.UpdateRelation(fileID, fl.TypAssignment, idStr, tx) != nil {
				tx.Rollback()
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusBadRequest).
					AddError("Wrong File ID"))
				return
			}
		}
	}
	if len(args.allowedTypesFile) > 0 {
		err := fl.InsertType(args.allowedTypesFile, id, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("Wrong File ID"))
			return
		}
	}
	tx.Commit()
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Assignment created successfully"))
	return
}

// UpdateHandler ..
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleUpdate, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := updateParams{
		ID:               ps.ByName("id"),
		name:             r.FormValue("name"),
		description:      r.FormValue("description"),
		dueDate:          r.FormValue("due_date"),
		filesID:          r.FormValue("files_id"),
		gpID:             r.FormValue("grade_parameter_id"),
		status:           r.FormValue("status"),
		allowedTypesFile: r.FormValue("allowed_types"),
		maxFile:          r.FormValue("max_file"),
		maxSizeFile:      r.FormValue("max_size"),
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
	tx := conn.DB.MustBegin()
	if len(args.filesID) > 0 {
		filesID, err := fl.SelectIDStatusByID(args.filesID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		for _, inFile := range args.filesID {
			for _, dbFile := range filesID {
				if inFile == dbFile.ID && dbFile.Status == fl.StatusDeleted {
					template.RenderJSONResponse(w, new(template.Response).
						SetCode(http.StatusBadRequest).
						AddError("ID file has been deleted, you can not use it again"))
					return
				}
			}
		}
		activefilesID, err := fl.SelectIDByRelation(fl.TypAssignment, params.ID, sess.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		var uptFiles []string
		var dltFiles []string
		for _, val := range args.filesID {
			isInActive := 0
			for _, actFile := range activefilesID {
				if actFile == val {
					isInActive++
					break
				}
			}
			if isInActive == 0 {
				uptFiles = append(uptFiles, val)
			}
		}
		for _, actFile := range activefilesID {
			isInActive := 0
			for _, val := range args.filesID {
				if actFile == val {
					isInActive++
					break
				}
			}
			if isInActive == 0 {
				dltFiles = append(dltFiles, actFile)
			}
		}
		if len(uptFiles) > 0 {
			for _, val := range uptFiles {
				err := fl.UpdateRelation(val, fl.TypAssignment, params.ID, tx)
				if err != nil {
					tx.Rollback()
					template.RenderJSONResponse(w, new(template.Response).
						SetCode(http.StatusInternalServerError))
					return
				}
			}
		}
		if len(dltFiles) > 0 {
			for _, val := range dltFiles {
				err := fl.Delete(val, tx)
				if err != nil {
					tx.Rollback()
					template.RenderJSONResponse(w, new(template.Response).
						SetCode(http.StatusInternalServerError))
					return
				}
			}
		}
	}
	if args.status == 0 {
		addedTypes, err := fl.SelectTypeByID(args.ID)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		if len(addedTypes) > 0 {
			err = fl.DeleteTypeByID(addedTypes, args.ID, tx)
			if err != nil {
				tx.Rollback()
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError))
				return
			}
		}
	}
	if len(args.allowedTypesFile) > 0 && args.status > 0 {
		addedTypes, err := fl.SelectTypeByID(args.ID)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		var dltType []string
		var uptType []string
		if len(addedTypes) > 0 {
			for _, addedType := range addedTypes {
				count := 0
				for _, alwdType := range args.allowedTypesFile {
					if alwdType == addedType {
						count++
						break
					}
				}
				if count == 0 {
					dltType = append(dltType, addedType)
				}
			}
		}
		for _, alwdType := range args.allowedTypesFile {
			count := 0
			for _, addedType := range addedTypes {
				if alwdType == addedType {
					count++
					break
				}
			}
			if count == 0 {
				uptType = append(uptType, alwdType)
			}
		}
		if len(dltType) > 0 {
			err := fl.DeleteTypeByID(dltType, args.ID, tx)
			if err != nil {
				tx.Rollback()
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError))
				return
			}
		}
		if len(uptType) > 0 {
			err := fl.InsertType(uptType, args.ID, tx)
			if err != nil {
				tx.Rollback()
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError))
				return
			}
		}
	}

	err = asg.Update(args.name, args.description, args.ID, args.gpID, args.maxSizeFile, args.maxFile, args.dueDate, args.status, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	err = tx.Commit()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Assignment updated successfully"))
	return

}

// DetailHandler ..
func DetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := detailParams{
		ID:      ps.ByName("id"),
		total:   r.FormValue("ttl"),
		page:    r.FormValue("pg"),
		payload: r.FormValue("payload"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if !asg.IsAssignmentExist(args.ID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Assignment does not exist"))
		return
	}
	gp := asg.GetGradeParameterID(args.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}
	scheduleID, err := cs.GetScheduleIDByGP(gp)
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
	if args.payload == "update" {
		typs := []string{}
		assignment, err := asg.GetByID(args.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		typs, err = fl.SelectTypeByID(args.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		tableID := []string{strconv.FormatInt(args.ID, 10)}
		asgFile, err := fl.SelectByRelation(fl.TypAssignment, tableID, nil)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}

		// file from assistant
		rAsgFile := []file{}
		for _, val := range asgFile {
			rAsgFile = append(rAsgFile, file{
				ID:           val.ID,
				Name:         fmt.Sprintf("%s.%s", val.Name, val.Extension),
				URL:          fmt.Sprintf("/api/v1/file/assignment/%s.%s", val.ID, val.Extension),
				URLThumbnail: helper.MimeToThumbnail(val.Mime),
			})
		}
		desc := "-"
		if assignment.Description.Valid {
			desc = assignment.Description.String
		}
		res := respDetailUpdate{}
		if assignment.Status == 1 {
			var size int8
			var max int8
			if assignment.MaxFile.Valid {
				max = int8(assignment.MaxFile.Int64)
			}
			if assignment.MaxSize.Valid {
				size = int8(assignment.MaxSize.Int64)
			}
			res = respDetailUpdate{
				ID:               assignment.ID,
				Name:             assignment.Name,
				Description:      desc,
				DueDate:          assignment.DueDate,
				GradeParameterID: assignment.GradeParameterID,
				Status:           assignment.Status,
				MaxFile:          max,
				MaxSize:          size,
				Type:             typs,
				FilesID:          rAsgFile,
			}
		} else {
			res = respDetailUpdate{
				ID:               assignment.ID,
				Name:             assignment.Name,
				Description:      desc,
				DueDate:          assignment.DueDate,
				GradeParameterID: assignment.GradeParameterID,
				Status:           assignment.Status,
				FilesID:          rAsgFile,
			}
		}
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(res))
		return

	}
	offset := (args.page - 1) * args.total
	assignment := asg.GetAssignmentByID(args.ID)
	var res respDetAsgUser
	asgUser := []detAsgUser{}
	var totalPg int
	if assignment.Status == 0 {
		total, err := usr.SelectCountByScheduleID(scheduleID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		totalPg = total / asg.MaxPage
		if total%asg.MaxPage > 0 {
			totalPg++
		}
		ids, err := usr.SelectIDByScheduleID(scheduleID, args.total, offset)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		users, err := usr.SelectConciseUserByID(ids)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		if len(users) > 0 {
			for _, val := range users {
				asgUser = append(asgUser, detAsgUser{
					ID:          val.IdentityCode,
					Name:        val.Name,
					Description: "-",
					UploadedAt:  "-",
					Link:        "-",
				})
			}
		}
	} else {
		total, err := asg.SelectCountUsrAsgByID(args.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		totalPg = total / asg.MaxPage
		if total%asg.MaxPage > 0 {
			totalPg++
		}
		sbmtdAsg, err := asg.SelectUserAssignmentByID(args.ID, args.total, offset)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		if len(sbmtdAsg) > 0 {
			var usersID []int64
			for _, val := range sbmtdAsg {
				usersID = append(usersID, val.UserID)
			}
			users, err := usr.SelectConciseUserByID(usersID)
			if err != nil {
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError))
				return
			}
			for _, val := range sbmtdAsg {
				var desc string
				if val.Description.Valid {
					desc = val.Description.String
				}
				asgUser = append(asgUser, detAsgUser{
					ID:          val.UserID,
					Description: desc,
					UploadedAt:  val.UpdatedAt.Format("Monday, 2 January 2006 15:04:05"),
					Link:        "-",
				})
			}
			for _, a := range asgUser {
				for i, u := range users {
					if a.ID == u.ID {
						asgUser[i].ID = u.IdentityCode
						asgUser[i].Name = u.Name
					}
				}
			}
		}
	}
	status := "must_upload"
	if assignment.Status == 0 {
		status = "upload_not_required"
	}
	res = respDetAsgUser{
		TotalPage:   totalPg,
		CurrentPage: args.page,
		ID:          assignment.ID,
		Name:        assignment.Name,
		Status:      status,
		DueDate:     assignment.DueDate.Format("Monday, 2 January 2006 15:04:05"),
		DetAsgUser:  asgUser,
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

// DeleteHandler ..
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXDelete, rg.RoleDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := deleteParams{
		id: ps.ByName("id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}
	countAsg, err := asg.SelectCountByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}
	if countAsg == 0 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Assignment does not exist"))
		return
	}
	assignment, err := asg.GetByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}
	sbmtdAsg, err := asg.SelectCountUsrAsgByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	if sbmtdAsg > 0 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("not allowed to delete this assignments"))
		return
	}
	scheduleID, err := cs.GetScheduleIDByGP(assignment.GradeParameterID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if !cs.IsAssistant(sess.ID, scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	tx := conn.DB.MustBegin()
	typs, err := fl.SelectCountTypeByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	if typs > 0 {
		err = fl.DeleteAllTypeByID(args.id, tx)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}
	id := strconv.FormatInt(args.id, 10)
	fls, err := fl.SelectCountIDByRelation(fl.TypAssignment, id, sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	if fls > 0 {
		err := fl.DeleteByRelation(fl.TypAssignment, id, tx)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	err = asg.DeleteAssignment(args.id, tx)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	err = tx.Commit()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Assignment deleted successfully"))
	return
}

// not finished yet
func GetReportHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	resp := []getReportResponse{}
	sess := r.Context().Value("User").(*auth.User)

	payload := r.FormValue("schedule_id")
	if !helper.IsEmpty(payload) {
		scheduleID, err := strconv.ParseInt(payload, 10, 64)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("Invalid Request"))
			return
		}
		gradeResp, statusCode, err := handleGradeBySchedule(scheduleID, sess.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(statusCode).
				AddError("Invalid Request"))
			return
		}
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(gradeResp))
		return
	}

	schedulesID, err := cs.SelectScheduleIDByUserID(sess.ID, cs.PStatusStudent)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if len(schedulesID) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
		return
	}

	courses, err := cs.SelectByScheduleID(schedulesID, cs.StatusScheduleActive)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if len(courses) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
		return
	}

	schedulesID = []int64{}
	for _, val := range courses {
		schedulesID = append(schedulesID, val.Schedule.ID)
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
	scheduleGP := map[int64][]cs.GradeParameter{}
	for _, gp := range gps {
		gpsID = append(gpsID, gp.ID)
		scheduleGP[gp.ScheduleID] = append(scheduleGP[gp.ScheduleID], gp)
	}

	assignments, err := asg.SelectByGP(gpsID, false)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var asgID []int64
	gpAsg := map[int64][]asg.Assignment{}
	for _, val := range assignments {
		asgID = append(asgID, val.ID)
		gpAsg[val.GradeParameterID] = append(gpAsg[val.GradeParameterID], val)
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

	asgSubmit := map[int64]asg.UserAssignment{}
	for _, val := range submitted {
		asgSubmit[val.AssignmentID] = val
	}

	attReport, err := att.CountByUserSchedule(sess.ID, schedulesID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	for _, c := range courses {
		rep := getReportResponse{
			CourseName: c.Course.Name,
			ScheduleID: c.Schedule.ID,
			Assignment: "-",
			Attendance: "-",
			Quiz:       "-",
			Mid:        "-",
			Final:      "-",
			Total:      "-",
		}
		total := float64(0)
		for _, gp := range scheduleGP[c.Schedule.ID] {
			scoreFloat64 := float64(0)
			if gp.Type == "ATTENDANCE" {
				attendance := attReport[gp.ScheduleID]
				if attendance.MeetingTotal > 0 {
					scoreFloat64 = (float64(attendance.AttendanceTotal) / float64(attendance.MeetingTotal)) * float64(gp.Percentage) / 100
				}
				rep.Attendance = fmt.Sprintf("%.3g", scoreFloat64)
			} else {
				count := len(gpAsg[gp.ID])
				for _, assignment := range gpAsg[gp.ID] {
					submit, exist := asgSubmit[assignment.ID]
					if !exist {
						continue
					}
					if submit.Score.Valid {
						scoreFloat64 += submit.Score.Float64
					} else {
						count--
					}
				}
				if count > 0 {
					scoreFloat64 = scoreFloat64 / float64(count)
				}
				total += (scoreFloat64 * float64(gp.Percentage) / 100)
				switch gp.Type {
				case "ASSIGNMENT":
					rep.Assignment = fmt.Sprintf("%.3g", scoreFloat64)
				case "QUIZ":
					rep.Quiz = fmt.Sprintf("%.3g", scoreFloat64)
				case "MID":
					rep.Mid = fmt.Sprintf("%.3g", scoreFloat64)
				case "FINAL":
					rep.Final = fmt.Sprintf("%.3g", scoreFloat64)
				}
			}
		}
		rep.Total = fmt.Sprintf("%.3g", total)
		resp = append(resp, rep)
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// GetGradeByAdmin ..
func GetGradeByAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := detailScoreParams{
		AssignmentID: r.FormValue("assignment_id"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	scheduleID, err := cs.GetScheduleIDByGP(gradeParameterID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	if !asg.IsAssignmentExist(args.AssignmentID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}
	if !cs.IsAssistant(sess.ID, scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("You dont have privillage"))
		return
	}

	assignments := asg.GetAssignmentByID(args.AssignmentID)
	res := detailAssignmentResponse{}
	us := []userAssignment{}
	res = detailAssignmentResponse{
		ID:      assignments.ID,
		Name:    assignments.Name,
		DueDate: assignments.DueDate.Format("Monday, 2 January 2006 15:04:05"),
	}
	if assignments.Status == 1 || (asg.IsExistSubmitted(assignments.ID) && assignments.Status == 0) {
		ids := []int64{}
		users, err := asg.SelectUserScoreByID(assignments.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		for _, val := range users {
			ids = append(ids, val.UserID)
		}
		usrDetail, err := usr.SelectConciseUserByID(ids)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		if users != nil {
			for i := 0; i < len(usrDetail); i++ {
				var score float32
				if users[i].Score.Valid {
					score = float32(users[i].Score.Float64)
				}
				us = append(us, userAssignment{
					ID:        usrDetail[i].ID,
					NPM:       usrDetail[i].IdentityCode,
					Name:      usrDetail[i].Name,
					Submitted: users[i].UpdatedAt.Format("Monday, 2 January 2006 15:04:05"),
					Grade:     score,
				})
			}
		}
		res.User = us
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(res))
		return
	}
	if !asg.IsExistSubmitted(assignments.ID) {
		ids, err := cs.SelectIDBySchedule(scheduleID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		usrDetail, err := usr.SelectConciseUserByID(ids)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		for i := 0; i < len(usrDetail); i++ {
			us = append(us, userAssignment{
				ID:   usrDetail[i].ID,
				NPM:  usrDetail[i].IdentityCode,
				Name: usrDetail[i].Name,
			})
		}
		res.User = us
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(res))
		return
	}
}

// UpdateScoreHandler ...
func UpdateScoreHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAssignment, rg.RoleXUpdate, rg.RoleUpdate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}
	params := updateScoreParams{
		AssignmentID: r.FormValue("assignment_id"),
		UserID:       r.FormValue("users_id"),
		Score:        r.FormValue("scores"),
	}
	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	gradeParameterID := cs.GetGradeParametersID(args.AssignmentID)
	scheduleID, err := cs.GetScheduleIDByGP(gradeParameterID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}
	if !asg.IsAssignmentExist(args.AssignmentID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}
	if !cs.IsAssistant(sess.ID, scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("You dont have privillage"))
		return
	}
	if !cs.IsAllUsersEnrolled(scheduleID, args.UserID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Wrong user list"))
		return
	}
	tx, err := conn.DB.Beginx()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	err = asg.CreateScore(args.AssignmentID, args.UserID, args.Score, tx)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	tx.Commit()
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Score Updated"))
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
