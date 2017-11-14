package webserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/handler"
	"github.com/melodiez14/meiko/src/webserver/handler/assignment"
	"github.com/melodiez14/meiko/src/webserver/handler/attendance"
	"github.com/melodiez14/meiko/src/webserver/handler/bot"
	"github.com/melodiez14/meiko/src/webserver/handler/course"
	"github.com/melodiez14/meiko/src/webserver/handler/file"
	"github.com/melodiez14/meiko/src/webserver/handler/information"
	"github.com/melodiez14/meiko/src/webserver/handler/place"
	"github.com/melodiez14/meiko/src/webserver/handler/rolegroup"
	"github.com/melodiez14/meiko/src/webserver/handler/user"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {

	// Home Handler
	r.GET("/", handler.HelloHandler)

	// ========================== User Handler ==========================
	// User section
	r.POST("/api/v1/user/register", auth.OptionalAuthorize(user.SignUpHandler))
	r.POST("/api/v1/user/verify", auth.OptionalAuthorize(user.EmailVerificationHandler))
	r.POST("/api/v1/user/signin", auth.OptionalAuthorize(user.SignInHandler))
	r.POST("/api/v1/user/forgot", auth.OptionalAuthorize(user.ForgotHandler))
	r.POST("/api/v1/user/signout", auth.MustAuthorize(user.SignOutHandler)) // delete
	r.POST("/api/v1/user/profile", auth.MustAuthorize(user.UpdateProfileHandler))
	r.GET("/api/v1/user/profile", auth.MustAuthorize(user.GetProfileHandler))
	r.POST("/api/v1/user/changepassword", auth.MustAuthorize(user.ChangePasswordHandler))

	// Admin section
	r.GET("/api/admin/v1/user", auth.MustAuthorize(user.ReadHandler))
	r.POST("/api/admin/v1/user", auth.MustAuthorize(user.CreateHandler))
	r.GET("/api/admin/v1/user/:id", auth.MustAuthorize(user.DetailHandler))
	r.POST("/api/admin/v1/user/:id", auth.MustAuthorize(user.UpdateHandler))              // patch
	r.POST("/api/admin/v1/user/:id/activate", auth.MustAuthorize(user.ActivationHandler)) // patch
	r.POST("/api/admin/v1/user/:id/delete", auth.MustAuthorize(user.DeleteHandler))       // delete
	// ======================== End User Handler ========================

	// ======================== Rolegroup Handler =======================
	// Admin section
	r.GET("/api/v1/role", auth.OptionalAuthorize(auth.OptionalAuthorize(rolegroup.GetPrivilege)))
	// ====================== End Rolegroup Handler =====================

	// ========================== File Handler ==========================
	// User section
	r.GET("/api/v1/files/:payload/:filename", file.GetFileHandler)
	r.GET("/api/v1/image/:payload", auth.MustAuthorize(file.GetProfileHandler))
	r.POST("/api/v1/image/profile", auth.MustAuthorize(file.UploadProfileImageHandler))
	r.POST("/api/v1/file/assignment", auth.MustAuthorize(file.UploadAssignmentHandler))
	// ======================== End File Handler ========================

	// ========================= Course Handler =========================
	// User section
	r.GET("/api/v1/course", auth.MustAuthorize(course.GetHandler))
	r.GET("/api/v1/course/assistant", auth.MustAuthorize(course.GetAssistantHandler))
	// Admin section
	r.POST("/api/admin/v1/course", auth.MustAuthorize(course.CreateHandler))
	r.GET("/api/admin/v1/course", auth.MustAuthorize(course.ReadHandler))
	r.GET("/api/admin/v1/course/:schedule_id", auth.MustAuthorize(course.ReadDetailHandler))                      //read
	r.GET("/api/admin/v1/course/:schedule_id/parameter", auth.MustAuthorize(course.ReadScheduleParameterHandler)) //read
	r.POST("/api/admin/v1/course/:schedule_id", auth.MustAuthorize(course.UpdateHandler))                         //patch
	r.POST("/api/admin/v1/course/:schedule_id/delete", auth.MustAuthorize(course.DeleteScheduleHandler))          //delete
	r.GET("/api/admin/v1/list/course/parameter", auth.MustAuthorize(course.ListParameterHandler))
	r.GET("/api/admin/v1/list/course/search", auth.MustAuthorize(course.SearchHandler))
	r.GET("/api/admin/v1/list/course/student", auth.MustAuthorize(course.ListEnrolledHandler))
	// ======================== End Course Handler ======================

	// ======================= Attendance Handler =======================
	// Admin section
	// r.GET("/api/admin/v1/attendance", auth.MustAuthorize())
	r.GET("/api/v1/attendance/list", auth.MustAuthorize(attendance.ListStudentHandler))
	r.GET("/api/admin/v1/attendance", auth.MustAuthorize(attendance.ReadMeetingHandler))
	r.POST("/api/admin/v1/attendance", auth.MustAuthorize(attendance.CreateMeetingHandler))
	r.GET("/api/admin/v1/attendance/:meeting_id", auth.MustAuthorize(attendance.ReadMeetingDetailHandler))
	r.POST("/api/admin/v1/attendance/:meeting_id/delete", auth.MustAuthorize(attendance.DeleteMeetingHandler))
	r.POST("/api/admin/v1/attendance/:meeting_id/update", auth.MustAuthorize(attendance.UpdateMeetingHandler))
	// ===================== End Attendance Handler =====================

	// =========================== Bot Handler ==========================
	// User section
	r.GET("/api/v1/bot", auth.MustAuthorize(bot.LoadHistoryHandler))
	r.POST("/api/v1/bot", auth.MustAuthorize(bot.BotHandler))
	// ========================= End Bot Handler ========================

	// ======================= Assignment Handler =======================
	r.POST("/api/admin/v1/assignment/create", auth.MustAuthorize(assignment.CreateHandler))
	r.GET("/api/admin/v1/assignment/:id", auth.MustAuthorize(assignment.DetailHandler))
	r.GET("/api/admin/v1/assignment", auth.MustAuthorize(assignment.GetAllAssignmentHandler))
	r.POST("/api/admin/v1/assignment/update/:id", auth.MustAuthorize(assignment.UpdateHandler))
	r.POST("/api/admin/v1/assignment/delete/:assignment_id", auth.MustAuthorize(assignment.DeleteAssignmentHandler))
	r.GET("/api/admin/v1/assignment/:id/:assignment_id", auth.MustAuthorize(assignment.GetUploadedAssignmentByAdminHandler))
	r.POST("/api/admin/v1/assignment/score/:id/:assignment_id", auth.MustAuthorize(assignment.UpdateScoreByAdminHandler)) // update score
	r.POST("/api/v1/assignment/update/:id", auth.MustAuthorize(assignment.UpdateHandlerByUser))                           // update upload by users
	r.POST("/api/v1/assignment/create", auth.MustAuthorize(assignment.CreateHandlerByUser))                               // create upload by user
	r.GET("/api/v1/assignment/:id/:schedule_id/:assignment_id", auth.MustAuthorize(assignment.GetUploadedDetailHandler))  // detail user assignments
	r.GET("/api/v1/assignment/:id", auth.MustAuthorize(assignment.GetAssignmentHandler))                                  // read detail assignments definitions
	r.GET("/api/v1/course/assignment/:schedule_id", auth.MustAuthorize(assignment.GetAssignmentByScheduleHandler))        // List assignments
	// r.GET("/api/v1/assignment/summary", auth.MustAuthorize(assignment.GetSummaryHandler))
	// ===================== End Assignment Handler =====================

	// ======================= Information Handler ======================
	// User section
	r.POST("/api/admin/v1/information/create", auth.MustAuthorize(information.CreateHandler))       // create infomation by admin
	r.POST("/api/admin/v1/information/update/:id", auth.MustAuthorize(information.UpdateHandler))   // update infomation by admin
	r.POST("/api/admin/v1/information/delete/:id", auth.MustAuthorize(information.DeleteHandler))   // delete information by admin
	r.GET("/api/admin/v1/information", auth.MustAuthorize(information.GetListHandler))              // read list information by admin
	r.GET("/api/admin/v1/information/:id", auth.MustAuthorize(information.GetDetailByAdminHandler)) // read detail information by admin
	r.GET("/api/v1/information", auth.MustAuthorize(information.GetSummaryHandler))                 // list informations
	r.GET("/api/v1/information/:id", auth.MustAuthorize(information.GetDetailHandler))              // detail information

	// ===================== End Information Handler ====================

	// ========================== Place Handler =========================
	// Public section
	r.GET("/api/v1/place/search", place.SearchHandler)
	// ======================== End Place Handler =======================

	// Catch
	// r.NotFound = http.RedirectHandler("/", http.StatusPermanentRedirect)
	// r.MethodNotAllowed = http.RedirectHandler("/", http.StatusPermanentRedirect)
}
