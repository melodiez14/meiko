package attendance

import (
	"fmt"
	"net/http"

	"github.com/melodiez14/meiko/src/util/helper"

	"github.com/julienschmidt/httprouter"
	atd "github.com/melodiez14/meiko/src/module/attendance"
	cs "github.com/melodiez14/meiko/src/module/course"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	usr "github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// u := r.Context().Value("User").(*auth.User)

	// courses, err := course.GetByUserID(u.ID)
	// if err != nil {
	// 	template.RenderJSONResponse(w, new(template.Response).
	// 		SetCode(http.StatusInternalServerError).
	// 		AddError(err.Error()))
	// 	return
	// }

	// var res []summaryResponse
	// var percentage float32
	// for _, c := range courses {
	// 	a, err := attendance.GetByUserCourseID(u.ID, c.ID)
	// 	if err != nil {
	// 		template.RenderJSONResponse(w, new(template.Response).
	// 			SetCode(http.StatusInternalServerError).
	// 			AddError(err.Error()))
	// 		return
	// 	}

	// 	if len(a) > 0 {
	// 		percentage = (float32(len(a)) * 100) / float32(len(a))
	// 	} else {
	// 		percentage = 0
	// 	}

	// 	res = append(res, summaryResponse{
	// 		Course:     c.Name,
	// 		Percentage: fmt.Sprintf("%.4g%%", percentage),
	// 	})
	// }

	// template.RenderJSONResponse(w, new(template.Response).
	// 	SetCode(http.StatusOK).
	// 	SetData(res))
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK))
	return
}

// ListStudentHandler not functional yet
func ListStudentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAttendance, rg.RoleXRead, rg.RoleRead) && !sess.IsHasRoles(rg.ModuleUser, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := listStudentParams{
		meetingID: r.FormValue("meeting_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	// get attendance
	meeting, err := atd.GetMeetingByID(args.meetingID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	// get list of enrolled users_id
	enrolledID, err := cs.SelectEnrolledStudentID(meeting.ScheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	resp := []listStudentResponse{}
	if len(enrolledID) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
		return
	}

	users, err := usr.SelectByID(enrolledID, true, usr.ColID, usr.ColName, usr.ColIdentityCode)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	atdUser, err := atd.SelectUserIDByMeetingID(meeting.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var status string
	for _, val := range users {
		status = "absent"
		if helper.Int64InSlice(val.ID, atdUser) {
			status = "present"
		}
		resp = append(resp, listStudentResponse{
			IdentityCode: val.IdentityCode,
			StudentName:  val.Name,
			Status:       status,
		})
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return

}

// CreateMeetingHandler ...
func CreateMeetingHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAttendance, rg.RoleXCreate, rg.RoleCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := createMeetingParams{
		subject:       r.FormValue("subject"),
		meetingNumber: r.FormValue("meeting_number"),
		description:   r.FormValue("description"),
		scheduleID:    r.FormValue("schedule_id"),
		date:          r.FormValue("date"),
		users:         r.FormValue("student"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if !cs.IsAssistant(sess.ID, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	if atd.IsExistMeeting(args.meetingNumber, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			AddError("Meeting number already exists"))
		return
	}

	// insert without without identity code
	if len(args.userIdentityCodes) < 1 {
		_, err := atd.InsertMeeting(args.subject, args.meetingNumber, args.description, args.date, args.scheduleID, nil)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}

		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetMessage("Meeting sucessfully inserted"))
		return
	}

	usersID, err := usr.SelectIDByIdentityCode(args.userIdentityCodes)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	countEnrolled, err := cs.CountEnrolled(usersID, args.scheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if countEnrolled != len(args.userIdentityCodes) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid request"))
		return
	}

	tx := conn.DB.MustBegin()
	meetingID, err := atd.InsertMeeting(args.subject, args.meetingNumber, args.description, args.date, args.scheduleID, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	err = atd.Insert(usersID, meetingID, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tx.Commit()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Meeting sucessfully inserted"))
	return
}

// UpdateMeetingHandler ...
func UpdateMeetingHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAttendance, rg.RoleXUpdate, rg.RoleUpdate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := updateMeetingParams{
		id:            ps.ByName("meeting_id"),
		subject:       r.FormValue("subject"),
		meetingNumber: r.FormValue("meeting_number"),
		scheduleID:    r.FormValue("schedule_id"),
		description:   r.FormValue("description"),
		isForceUpdate: r.FormValue("is_force"),
		date:          r.FormValue("date"),
		users:         r.FormValue("student"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	if !cs.IsAssistant(sess.ID, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	meeting, err := atd.GetMeetingByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid request"))
		return
	}

	if meeting.ScheduleID != args.scheduleID {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid request"))
		return
	}

	isExist := atd.IsExistByMeetingID(meeting.ID)
	if isExist && !args.isForceUpdate {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			SetMessage("This meeting has an attendant. Are you sure want to update this meeting?"))
		return
	}

	if meeting.Number != args.meetingNumber {
		if atd.IsExistMeeting(args.meetingNumber, args.scheduleID) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusConflict).
				AddError("Meeting number already exist"))
			return
		}
	}

	var usersID []int64
	if len(args.userIdentityCodes) > 0 {
		// check if all inputted attendance registered in p_users_schedule
		usersID, err = usr.SelectIDByIdentityCode(args.userIdentityCodes)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}

		countEnrolled, err := cs.CountEnrolled(usersID, args.scheduleID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}

		if countEnrolled != len(args.userIdentityCodes) {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusBadRequest).
				AddError("Invalid request"))
			return
		}
	}

	attendUser, err := atd.SelectUserIDByMeetingID(meeting.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var delete []int64
	for _, val := range attendUser {
		if !helper.Int64InSlice(val, usersID) {
			delete = append(delete, val)
		}
	}

	var insert []int64
	for _, val := range usersID {
		if !helper.Int64InSlice(val, attendUser) {
			insert = append(insert, val)
		}
	}

	tx := conn.DB.MustBegin()
	err = atd.UpdateMeeting(args.id, args.subject, args.meetingNumber, args.description, args.date, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	if len(delete) > 0 {
		err = atd.Delete(delete, meeting.ID, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	if len(insert) > 0 {
		err = atd.Insert(insert, meeting.ID, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	tx.Commit()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Meeting sucessfully updated"))
	return
}

// DeleteMeetingHandler ...
func DeleteMeetingHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAttendance, rg.RoleXDelete, rg.RoleDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := deleteMeetingParams{
		id:            ps.ByName("meeting_id"),
		isForceDelete: r.FormValue("is_force"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	meeting, err := atd.GetMeetingByID(args.id)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid request"))
		return
	}

	if !cs.IsAssistant(sess.ID, meeting.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	isExist := atd.IsExistByMeetingID(meeting.ID)
	if isExist && !args.isForceDelete {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			SetMessage("This meeting has an attendant. Are you sure want to delete this meeting?"))
		return
	}

	tx := conn.DB.MustBegin()
	if isExist {
		err = atd.DeleteByMeetingID(args.id, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}
	err = atd.DeleteMeeting(args.id, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tx.Commit()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Meeting sucessfully deleted"))
	return
}

// ReadMeetingHandler ...
func ReadMeetingHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAttendance, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readMeetingParams{
		scheduleID: r.FormValue("schedule_id"),
		page:       r.FormValue("pg"),
		total:      r.FormValue("ttl"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid request"))
		return
	}

	if !cs.IsAssistant(sess.ID, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	if args.total > 100 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Max total should be less than or equal to 100"))
		return
	}

	offset := (args.page - 1) * args.total
	meetings, count, err := atd.SelectMeetingByPage(args.scheduleID, args.total, offset, true)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	enrolled, err := cs.SelectEnrolledStudentID(args.scheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	respMeetings := []readMeetings{}
	for _, val := range meetings {
		respMeetings = append(respMeetings, readMeetings{
			ID:             val.ID,
			Subject:        val.Subject,
			MeetingNumber:  val.Number,
			Date:           val.Date.Unix(),
			TotalAttendant: val.TotalAttendant,
		})
	}

	totalStudent := len(enrolled)
	totalPage := count / args.total
	if count%args.total > 0 {
		totalPage++
	}

	if totalPage < 1 {
		totalPage = 1
	}

	resp := readMeetingResponse{
		TotalPage:    totalPage,
		Page:         args.page,
		Meetings:     respMeetings,
		TotalStudent: totalStudent,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

func ReadMeetingDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleAttendance, rg.RoleXRead, rg.RoleRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readMeetingDetailParams{
		meetingID: ps.ByName("meeting_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	meeting, err := atd.GetMeetingByID(args.meetingID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound).
			AddError("Meeting ID not found"))
		return
	}

	if !cs.IsAssistant(sess.ID, meeting.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	usersID, err := cs.SelectEnrolledStudentID(meeting.ScheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	users, err := usr.SelectByID(usersID, true, usr.ColID, usr.ColIdentityCode, usr.ColName)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	presentUserID, err := atd.SelectUserIDByMeetingID(meeting.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	std := []student{}
	for _, val := range users {
		status := "absent"
		if helper.Int64InSlice(val.ID, presentUserID) {
			status = "present"
		}

		std = append(std, student{
			IdentityCode: val.IdentityCode,
			Name:         val.Name,
			Status:       status,
		})
	}

	resp := readMeetingDetailResponse{
		ID:            meeting.ID,
		Subject:       meeting.Subject,
		Description:   meeting.Description.String,
		MeetingNumber: meeting.Number,
		Date:          meeting.Date.Unix(),
		Student:       std,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

func GetAttendanceHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := getAttendanceParams{
		scheduleID: r.FormValue("schedule_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	if !cs.IsEnrolled(sess.ID, args.scheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You are not authorized"))
		return
	}

	var meetingsID []int64
	var present, absent, totalMeeting int
	var percentage float32

	meetingsID, err = atd.SelectMeetingIDByScheduleID(sess.ID, args.scheduleID)
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
		absent = totalMeeting - present
		percentage = float32(present) * 100 / float32(totalMeeting)
	}

	resp := getAttendanceResponse{
		Absent:       absent,
		Present:      present,
		TotalMeeting: totalMeeting,
		Percentage:   fmt.Sprintf("%.3g%%", percentage),
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}
