package webserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/handler/assignment"
	"github.com/melodiez14/meiko/src/webserver/handler/attendance"
	"github.com/melodiez14/meiko/src/webserver/handler/bot"
	"github.com/melodiez14/meiko/src/webserver/handler/course"
	"github.com/melodiez14/meiko/src/webserver/handler/file"
	"github.com/melodiez14/meiko/src/webserver/handler/information"
	"github.com/melodiez14/meiko/src/webserver/handler/place"
	"github.com/melodiez14/meiko/src/webserver/handler/rolegroup"
	"github.com/melodiez14/meiko/src/webserver/handler/tutorial"
	"github.com/melodiez14/meiko/src/webserver/handler/user"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {

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
	r.PATCH("/api/admin/v1/user/:id", auth.MustAuthorize(user.UpdateHandler))
	r.PATCH("/api/admin/v1/user/:id/:status", auth.MustAuthorize(user.ActivationHandler))
	r.DELETE("/api/admin/v1/user/:id", auth.MustAuthorize(user.DeleteHandler))

	// Public
	r.GET("/api/v1/util/time", user.GetTimeHandler)
	// ======================== End User Handler ========================

	// ======================== Rolegroup Handler =======================
	// User Section
	r.GET("/api/v1/role", auth.OptionalAuthorize(rolegroup.GetPrivilege))
	// Admin section
	r.GET("/api/admin/v1/role", auth.MustAuthorize(rolegroup.ReadHandler))
	r.POST("/api/admin/v1/role", auth.MustAuthorize(rolegroup.CreateHandler))
	r.GET("/api/admin/v1/role/:rolegroup_id", auth.MustAuthorize(rolegroup.ReadDetailHandler))
	r.PATCH("/api/admin/v1/role/:rolegroup_id", auth.MustAuthorize(rolegroup.UpdateHandler))
	r.DELETE("/api/admin/v1/role/:rolegroup_id", auth.MustAuthorize(rolegroup.DeleteHandler))
	// ====================== End Rolegroup Handler =====================

	// ========================== File Handler ==========================
	r.GET("/api/v1/filerouter", auth.OptionalAuthorize(file.RouterFileHandler))
	r.GET("/api/v1/file/:payload/:filename", file.GetFileHandler)
	r.GET("/api/v1/image/:payload", auth.MustAuthorize(file.GetProfileHandler))
	r.POST("/api/v1/image/profile", auth.MustAuthorize(file.UploadProfileImageHandler))
	r.POST("/api/admin/v1/image/information", auth.MustAuthorize(file.UploadInformationImageHandler))
	// r.POST("/api/v1/file/assignment", auth.MustAuthorize(file.UploadAssignmentHandler))
	r.POST("/api/v1/file", auth.MustAuthorize(file.UploadFileHandler))
	r.GET("/static/*filepath", file.StaticHandler)
	// ======================== End File Handler ========================

	// ========================= Course Handler =========================
	// User section
	r.GET("/api/v1/course", auth.MustAuthorize(course.GetHandler))
	r.GET("/api/v1/course/:schedule_id", auth.MustAuthorize(course.GetDetailHandler))
	r.GET("/api/v1/course/:schedule_id/assistant", auth.MustAuthorize(course.GetAssistantHandler))
	r.POST("/api/v1/course/:schedule_id/enrollment", auth.MustAuthorize(course.EnrollRequestHandler))

	// Admin section
	r.GET("/api/admin/v1/course", auth.MustAuthorize(course.ReadHandler))
	r.POST("/api/admin/v1/course", auth.MustAuthorize(course.CreateHandler))
	r.GET("/api/admin/v1/course/:schedule_id", auth.MustAuthorize(course.ReadDetailHandler))                      //read
	r.GET("/api/admin/v1/course/:schedule_id/parameter", auth.MustAuthorize(course.ReadScheduleParameterHandler)) //read
	r.PATCH("/api/admin/v1/course/:schedule_id", auth.MustAuthorize(course.UpdateHandler))
	r.DELETE("/api/admin/v1/course/:schedule_id", auth.MustAuthorize(course.DeleteScheduleHandler))
	r.POST("/api/admin/v1/course/:schedule_id/assistant", auth.MustAuthorize(course.AddAssistantHandler))
	r.GET("/api/admin/v1/list/course/parameter", auth.MustAuthorize(course.ListParameterHandler))
	r.GET("/api/admin/v1/list/course/search", auth.MustAuthorize(course.SearchHandler))
	// ======================== End Course Handler ======================

	// ======================== Tutorial Handler ========================
	r.GET("/api/v1/tutorial", auth.MustAuthorize(tutorial.ReadHandler)) // for admin and user
	r.POST("/api/admin/v1/tutorial", auth.MustAuthorize(tutorial.CreateHandler))
	r.GET("/api/admin/v1/tutorial/:tutorial_id", auth.MustAuthorize(tutorial.ReadDetailHandler))
	r.PATCH("/api/admin/v1/tutorial/:tutorial_id", auth.MustAuthorize(tutorial.UpdateHandler))
	r.DELETE("/api/admin/v1/tutorial/:tutorial_id", auth.MustAuthorize(tutorial.DeleteHandler))
	// ====================== End Tutorial Handler ======================

	// ======================= Attendance Handler =======================
	// Admin section
	r.GET("/api/v1/attendance/list", auth.MustAuthorize(attendance.ListStudentHandler))
	r.GET("/api/v1/attendance/summary", auth.MustAuthorize(attendance.GetAttendanceHandler))
	r.GET("/api/admin/v1/attendance", auth.MustAuthorize(attendance.ReadMeetingHandler))
	r.POST("/api/admin/v1/attendance", auth.MustAuthorize(attendance.CreateMeetingHandler))
	r.GET("/api/admin/v1/attendance/:meeting_id", auth.MustAuthorize(attendance.ReadMeetingDetailHandler))
	r.DELETE("/api/admin/v1/attendance/:meeting_id", auth.MustAuthorize(attendance.DeleteMeetingHandler))
	r.PATCH("/api/admin/v1/attendance/:meeting_id", auth.MustAuthorize(attendance.UpdateMeetingHandler))
	// ===================== End Attendance Handler =====================
	// =========================== Bot Handler ==========================
	// User section
	r.GET("/api/v1/bot", auth.MustAuthorize(bot.LoadHistoryHandler))
	r.POST("/api/v1/bot", auth.MustAuthorize(bot.BotHandler))
	// ========================= End Bot Handler ========================

	// ========================= Assignment Handler ========================
	// r.POST("/api/admin/v1/assignment/create", auth.MustAuthorize(assignment.CreateHandler))
	// r.GET("/api/admin/v1/assignment/:id", auth.MustAuthorize(assignment.DetailHandler))
	// r.POST("/api/admin/v1/assignment/update/:id", auth.MustAuthorize(assignment.UpdateHandler))
	// r.POST("/api/admin/v1/assignment/delete/:assignment_id", auth.MustAuthorize(assignment.DeleteAssignmentHandler))
	// r.GET("/api/admin/v1/assignment", auth.MustAuthorize(assignment.GetAllAssignmentHandler))
	// r.GET("/api/admin/v1/assignment/:id/:assignment_id", auth.MustAuthorize(assignment.GetUploadedAssignmentByAdminHandler))
	// r.GET("/api/admin/v1/score/:schedule_id/:assignment_id", auth.MustAuthorize(assignment.GetDetailAssignmentByAdmin))
	// r.POST("/api/admin/v1/score/:schedule_id/:assignment_id/create", auth.MustAuthorize(assignment.CreateScoreHandler)) // update score

	r.GET("/api/v1/assignment/:id", auth.MustAuthorize(assignment.GetDetailHandler)) // read detail assignments definitions
	r.PUT("/api/v1/assignment/:id", auth.MustAuthorize(assignment.UploadHandler))    // update upload by users
	// r.POST("/api/v1/assignment", auth.MustAuthorize(assignment.CreateHandlerByUser))                                     // create upload by user
	// r.GET("/api/v1/assignment/:id/:schedule_id/:assignment_id", auth.MustAuthorize(assignment.GetUploadedDetailHandler)) // detail user assignments
	// r.GET("/api/v1/assignment", auth.MustAuthorize(assignment.GetAssignmentHandler))                                     // List assignments
	// r.GET("/api/v1/assignment-schedule", auth.MustAuthorize(assignment.GetAssignmentByScheduleHandler))                  // List assignments
	// r.GET("/api/v1/grade", auth.MustAuthorize(assignment.GradeSummary))
	// r.GET("/api/v1/grade/:id", auth.MustAuthorize(assignment.GradeBySchedule))
	// ===================== End Assignment Handler =====================

	// ======================= Information Handler ======================
	// User section
	r.POST("/api/admin/v1/information", auth.MustAuthorize(information.CreateHandler))              // create infomation by admin
	r.PATCH("/api/admin/v1/information/:id", auth.MustAuthorize(information.UpdateHandler))         // update infomation by admin
	r.DELETE("/api/admin/v1/information/:id", auth.MustAuthorize(information.DeleteHandler))        // delete information by admin
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
	r.NotFound = http.HandlerFunc(file.IndexHandler)
	// r.MethodNotAllowed = http.RedirectHandler("/", http.StatusPermanentRedirect)
}
