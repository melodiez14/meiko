package webserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/handler"
	"github.com/melodiez14/meiko/src/webserver/handler/assignment"
	"github.com/melodiez14/meiko/src/webserver/handler/attendance"
	"github.com/melodiez14/meiko/src/webserver/handler/course"
	"github.com/melodiez14/meiko/src/webserver/handler/file"
	"github.com/melodiez14/meiko/src/webserver/handler/information"
	"github.com/melodiez14/meiko/src/webserver/handler/notification"
	"github.com/melodiez14/meiko/src/webserver/handler/place"
	"github.com/melodiez14/meiko/src/webserver/handler/rolegroup"
	"github.com/melodiez14/meiko/src/webserver/handler/user"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {

	// Home Handler
	r.GET("/", handler.HelloHandler)

	// ==================================================================
	// ========================== User Handler ==========================
	// ==================================================================

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

	// ==================================================================
	// ======================== End User Handler ========================
	// ==================================================================

	// ==================================================================
	// ======================== Rolegroup Handler =======================
	// ==================================================================

	r.GET("/api/v1/role", auth.OptionalAuthorize(auth.OptionalAuthorize(rolegroup.GetPrivilege)))

	// ==================================================================
	// ====================== End Rolegroup Handler =====================
	// ==================================================================

	// File Handler
	r.POST("/api/v1/image", auth.MustAuthorize(file.UploadImageHandler))
	r.GET("/api/v1/image/:payload", auth.MustAuthorize(file.GetProfile))

	// Course Handler
	r.POST("/api/admin/v1/course", auth.MustAuthorize(course.CreateHandler))
	r.GET("/api/admin/v1/course", auth.MustAuthorize(course.ReadHandler))
	r.POST("/api/admin/v1/course/:id", auth.MustAuthorize(course.UpdateHandler)) // patch
	r.GET("/api/v1/course", auth.MustAuthorize(course.GetHandler))
	r.GET("/api/v1/course/assistant", auth.MustAuthorize(course.GetAssistantHandler))
	r.GET("/api/v1/course/summary", auth.MustAuthorize(course.GetSummaryHandler))

	// Assignment Handler
	r.GET("/api/v1/assignment/incomplete", auth.MustAuthorize(assignment.GetIncompleteHandler))
	r.GET("/api/v1/assignment/summary", auth.MustAuthorize(assignment.GetSummaryHandler))

	// Attendance Handler
	r.GET("/api/v1/attendance/summary", auth.MustAuthorize(attendance.GetSummaryHandler))
	r.GET("/api/v1/notification", auth.MustAuthorize(notification.GetHandler))

	// Information Handler
	r.GET("/api/v1/information", auth.MustAuthorize(information.GetSummaryHandler))

	// Place Handler
	r.GET("/api/v1/place/search", place.SearchHandler)

	// Catch
	r.NotFound = http.RedirectHandler("/", http.StatusPermanentRedirect)
	r.MethodNotAllowed = http.RedirectHandler("/", http.StatusPermanentRedirect)
}
