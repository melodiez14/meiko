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
	"github.com/melodiez14/meiko/src/webserver/handler/user"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {

	// Home Handler
	r.GET("/", handler.HelloHandler)

	// User Handler
	r.POST("/api/v1/user/register", auth.OptionalAuthorize(user.SignUpHandler))
	r.POST("/api/v1/user/verify", auth.OptionalAuthorize(user.EmailVerificationHandler))
	r.GET("/api/v1/user/verified", auth.MustAuthorize(user.ReadHandler))
	r.POST("/api/v1/user/signin", auth.OptionalAuthorize(user.SignInHandler))
	r.POST("/api/v1/user/forgot", auth.OptionalAuthorize(user.ForgotHandler))
	r.DELETE("/api/v1/user/signout", auth.MustAuthorize(user.SignOutHandler))
	r.POST("/api/v1/user/profile", auth.MustAuthorize(user.UpdateProfileHandler))
	r.GET("/api/v1/user/profile", auth.MustAuthorize(user.GetProfileHandler))
	r.POST("/api/v1/user/activate", auth.MustAuthorize(user.ActivationHandler))
	r.POST("/api/v1/user/changepassword", auth.MustAuthorize(user.ChangePasswordHandler))

	// File Handler
	r.POST("/api/v1/image", auth.MustAuthorize(file.UploadImageHandler))
	r.GET("/api/v1/image/:payload", auth.MustAuthorize(file.GetProfile))

	// Course Handler
	r.POST("/api/admin/v1/course", auth.MustAuthorize(course.CreateHandler))
	r.GET("/api/admin/v1/course", auth.MustAuthorize(course.ReadHandler))
	r.PATCH("/api/admin/v1/course/:id", auth.MustAuthorize(course.UpdateHandler))
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
