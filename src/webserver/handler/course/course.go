package course

import (
	"database/sql"
	"net/http"

	"github.com/melodiez14/meiko/src/util/conn"

	"github.com/melodiez14/meiko/src/util/helper"

	"github.com/julienschmidt/httprouter"
	cs "github.com/melodiez14/meiko/src/module/course"
	pl "github.com/melodiez14/meiko/src/module/place"
	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

// CreateHandler handles the http request for creating the course. Accessing this handler needs CREATE or XCREATE ability
/*
	@params:
		id			= required
		name		= required, alphabet and space only
		description	= optional
		ucu			= required, positive numeric
		semester	= required, positive numeric
		start_time	= required, positive numeric, minutes
		end_time	= required, positive numeric, minutes
		class		= required, character=1
		day			= required, [monday, tuesday, wednesday, thursday, friday, saturday, sunday]
		place		= required
		is_update	= optional, true
	@example:
		id			= D10K-7D02
		name		= Sistem Informasi Multimedia
		description	= Praktikum ini membahas mengenai Sistem Informasi Multimedia
		ucu			= 3
		semester	= 1
		start_time	= 600
		end_time	= 800
		class		= A
		day			= monday
		place		= UDJT-102
		is_update	= true
	@return
*/
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleSchedule, rg.RoleCreate, rg.RoleXCreate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := createParams{
		ID:          r.FormValue("id"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		UCU:         r.FormValue("ucu"),
		Semester:    r.FormValue("semester"),
		Year:        r.FormValue("year"),
		StartTime:   r.FormValue("start_time"),
		EndTime:     r.FormValue("end_time"),
		Class:       r.FormValue("class"),
		Day:         r.FormValue("day"),
		PlaceID:     r.FormValue("place"),
		IsUpdate:    r.FormValue("is_update"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// is exist course and place
	scExist := cs.IsExistSchedule(args.Semester, args.Year, args.ID, args.Class)
	if scExist {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Schedule already exists"))
		return
	}
	csExist := cs.IsExist(args.ID)
	plExist := pl.IsExistID(args.PlaceID)

	tx := conn.DB.MustBegin()

	// insert new course and check create or xcreate roles
	if !csExist && sess.IsHasRoles(rg.ModuleCourse, rg.RoleCreate, rg.RoleXCreate) {
		err = cs.Insert(args.ID, args.Name, args.Description, args.UCU, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		// update course and check update or xupdate roles
	} else if csExist && args.IsUpdate && sess.IsHasRoles(rg.ModuleCourse, rg.RoleUpdate, rg.RoleXUpdate) {
		err = cs.Update(args.ID, args.Name, args.Description, args.UCU, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		// want to update course but dont have privilege
	} else if args.IsUpdate {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You dont have privilege to create or update course"))
		return
	}

	// insert place
	if !plExist {
		err = pl.Insert(args.PlaceID, sql.NullString{}, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	// insert schedule
	err = cs.InsertSchedule(sess.ID,
		args.StartTime,
		args.EndTime,
		args.Year,
		args.Semester,
		args.Day,
		cs.StatusScheduleActive,
		args.Class,
		args.ID,
		args.PlaceID,
		tx)
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
		SetMessage("Success"))
	return
}

// ReadHandler handles the http request for creating the course. Accessing this handler needs CREATE or XCREATE ability
/*
	@params:
		pg	= required, positive numeric
		ttl	= required, positive numeric
	@example:
		pg=1
		ttl=10
	@return
		[]{name, email, status, identity}
*/
func ReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleCourse, rg.RoleRead, rg.RoleXRead) {
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
	courses, err := cs.SelectByPage(args.Total, offset)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var status string
	var res []readResponse
	for _, val := range courses {

		if val.Schedule.Status == cs.StatusScheduleActive {
			status = "active"
		} else {
			status = "inactive"
		}

		res = append(res, readResponse{
			ID:         val.Course.ID,
			Name:       val.Course.Name,
			Class:      val.Schedule.Class,
			StartTime:  helper.MinutesToTimeString(val.Schedule.StartTime),
			EndTime:    helper.MinutesToTimeString(val.Schedule.EndTime),
			Day:        helper.IntDayToString(val.Schedule.Day),
			Status:     status,
			ScheduleID: val.Schedule.ID,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

func SearchHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleCourse, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := searchParams{
		Text: r.FormValue("q"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("Invalid request"))
		return
	}

	var resp []searchResponse
	if len(args.Text) < 3 || len(args.Text) > 20 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(resp))
		return
	}

	courses, err := cs.SelectByName(args.Text)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	for _, val := range courses {
		resp = append(resp, searchResponse{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description.String,
			UCU:         val.UCU,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// ReadDetailHandler handles the http request for read course details. Accessing this handler needs READ or XREAD ability
/*
	@params:
		schedule_id	= required, positive numeric
	@example:
		schedule_id = 123
	@return
		id = D10K-7D01
		name = Mobile Computing
		description = Mata kuliah ini mengajarkan tentang aplikasi client server
		ucu = 3
		status = 1
		semester = 7
		year = 2017
		start_time = 600
		end_time = 800
		class = A
		day = thursday
		place_id = UDJT102
		schedule_id = 123
*/
func ReadDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleCourse, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := readDetailParams{
		ScheduleID: ps.ByName("schedule_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad request"))
		return
	}

	course, err := cs.GetByScheduleID(args.ScheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	status := "active"
	if course.Schedule.Status == cs.StatusScheduleInactive {
		status = "inactive"
	}

	resp := readDetailResponse{
		ID:          course.Course.ID,
		Name:        course.Course.Name,
		Description: course.Course.Description.String,
		UCU:         course.Course.UCU,
		Status:      status,
		Semester:    course.Schedule.Semester,
		Year:        course.Schedule.Year,
		StartTime:   course.Schedule.StartTime,
		EndTime:     course.Schedule.EndTime,
		Class:       course.Schedule.Class,
		Day:         helper.IntDayToString(course.Schedule.Day),
		PlaceID:     course.Schedule.PlaceID,
		ScheduleID:  course.Schedule.ID,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleSchedule, rg.RoleUpdate, rg.RoleXUpdate) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := updateParams{
		ID:          r.FormValue("id"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		UCU:         r.FormValue("ucu"),
		ScheduleID:  ps.ByName("schedule_id"),
		Status:      r.FormValue("status"),
		Semester:    r.FormValue("semester"),
		Year:        r.FormValue("year"),
		StartTime:   r.FormValue("start_time"),
		EndTime:     r.FormValue("end_time"),
		Class:       r.FormValue("class"),
		Day:         r.FormValue("day"),
		PlaceID:     r.FormValue("place"),
		IsUpdate:    r.FormValue("is_update"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError(err.Error()))
		return
	}

	// check if semester, year, id, class already used by another schedule
	if cs.IsExistSchedule(args.Semester, args.Year, args.ID, args.Class, args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusConflict).
			AddError("Schedule already exist"))
		return
	}

	// is exist course and place
	scExist := cs.IsExistScheduleID(args.ScheduleID)
	if !scExist {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound))
		return
	}
	csExist := cs.IsExist(args.ID)
	plExist := pl.IsExistID(args.PlaceID)

	tx := conn.DB.MustBegin()

	// insert new course and check create or xcreate roles
	if !csExist && sess.IsHasRoles(rg.ModuleCourse, rg.RoleCreate, rg.RoleXCreate) {
		err = cs.Insert(args.ID, args.Name, args.Description, args.UCU, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		// update course and check update or xupdate roles
	} else if csExist && args.IsUpdate && sess.IsHasRoles(rg.ModuleCourse, rg.RoleUpdate, rg.RoleXUpdate) {
		err = cs.Update(args.ID, args.Name, args.Description, args.UCU, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
		// want to update course but dont have privilege
	} else if args.IsUpdate {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You dont have privilege to create or update course"))
		return
	}

	// insert place if not exist
	if !plExist {
		err = pl.Insert(args.PlaceID, sql.NullString{}, tx)
		if err != nil {
			tx.Rollback()
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	err = cs.UpdateSchedule(args.ScheduleID, args.StartTime, args.EndTime, args.Year, args.Semester, args.Day, args.Status, args.Class, args.ID, args.PlaceID, tx)
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
		SetMessage("Success"))
	return
}

// GetHandler handles the http request return course list
/*
	@params:
		payload	= required
	@example:
		pg	= last or current or all
		ttl = 10
	@return
		[]{id, name, description}
*/
func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := getParams{
		Payload: r.FormValue("payload"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	var csID []int64
	var csStatus int8
	switch args.Payload {
	case "last":
		csStatus = cs.StatusScheduleInactive
		csID, err = cs.SelectScheduleIDByUserID(sess.ID, cs.PStatusStudent)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	case "current":
		csStatus = cs.StatusScheduleActive
		csID, err = cs.SelectScheduleIDByUserID(sess.ID, cs.PStatusStudent)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	case "all":
		csStatus = cs.StatusScheduleActive
	default:
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	var courses []cs.CourseSchedule
	if args.Payload == "all" {
		courses, err = cs.SelectByStatus(csStatus)
	} else {
		courses, err = cs.SelectByScheduleID(csID, csStatus)
	}
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	res := []getResponse{}
	for _, val := range courses {
		res = append(res, getResponse{
			ID:          val.Schedule.ID,
			Name:        val.Course.Name,
			Description: val.Course.Description.String,
			Class:       val.Schedule.Class,
			Semester:    val.Schedule.Semester,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
}

// GetAssistantHandler handles the http request return course assistant list
/*
	@params:
		payload	= required
	@example:
		pg	= last or current or all
		ttl = 10
	@return
		[]{id, name, description}
*/
func GetAssistantHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)

	params := getAssistantParams{
		ScheduleID: r.FormValue("id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	if !cs.IsEnrolled(sess.ID, args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	uIDs, err := cs.SelectAssistantID(args.ScheduleID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var users []user.User
	if len(uIDs) > 0 {
		users, err = user.SelectByID(uIDs, user.ColEmail, user.ColPhone, user.ColName)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError))
			return
		}
	}

	res := []getAssistantResponse{}
	for _, val := range users {
		phone := "-"
		if val.Phone.Valid {
			phone = val.Phone.String
		}

		res = append(res, getAssistantResponse{
			Name:  val.Name,
			Email: val.Email,
			Phone: phone,
			Roles: "Assistant",
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}

// DeleteScheduleHandler handles the http request for deleting schedule. Accessing this handler require DELETE or XDELETE ability
/*
	@params:
		schedule_id = required, postitive numeric
	@example:
		schedule_id = 149
	@return
*/
func DeleteScheduleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleSchedule, rg.RoleDelete, rg.RoleXDelete) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	params := deleteScheduleParams{
		ScheduleID: ps.ByName("schedule_id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad request"))
		return
	}

	if !cs.IsExistScheduleID(args.ScheduleID) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusNotFound).
			AddError("Not Found"))
		return
	}

	err = cs.DeleteSchedule(args.ScheduleID)
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

// ListParameterHandler handles the http request returns list of grade parameters. Accessing this handler require READ or XREAD ability
/*
	@params:
		schedule_id = required, postitive numeric
	@example:
		schedule_id = 149
	@return
*/
func ListParameterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess := r.Context().Value("User").(*auth.User)
	if !sess.IsHasRoles(rg.ModuleSchedule, rg.RoleRead, rg.RoleXRead) {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusForbidden).
			AddError("You don't have privilege"))
		return
	}

	resp := []listParameterResponse{
		{
			id:    cs.GradeParameterAttendance,
			value: "Attendance",
		},
		{
			id:    cs.GradeParameterQuiz,
			value: "Quiz",
		},
		{
			id:    cs.GradeParameterMid,
			value: "UTS",
		},
		{
			id:    cs.GradeParameterFinal,
			value: "UAS",
		},
		{
			id:    cs.GradeParameterAssignment,
			value: "Assignment",
		},
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}
