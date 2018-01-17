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

// not finished yet
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

	assignment, err := asg.GetByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNoContent))
		return
	}

	// validate if there is submitted task
	// asg.SelectSubmittedByUser

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
