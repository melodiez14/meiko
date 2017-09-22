package webserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/handler/assignment"
	"github.com/melodiez14/meiko/src/webserver/handler/attendance"
	"github.com/melodiez14/meiko/src/webserver/handler/course"
	"github.com/melodiez14/meiko/src/webserver/handler/notification"
	"github.com/melodiez14/meiko/src/webserver/handler/user"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {
	r.POST("/api/v1/user/register", auth.OptionalAuthorize(user.SignUpHandler))
	r.POST("/api/v1/user/verify", auth.OptionalAuthorize(user.EmailVerificationHandler))
	r.GET("/api/v1/user/verified", auth.MustAuthorize(user.ReadUserHandler))
	r.POST("/api/v1/user/signin", auth.OptionalAuthorize(user.SignInHandler))
	r.POST("/api/v1/user/forgot", auth.OptionalAuthorize(user.ForgotHandler))
	r.DELETE("/api/v1/user/signout", auth.MustAuthorize(user.SignOutHandler))
	r.POST("/api/v1/user/profile", auth.MustAuthorize(user.UpdateProfileHandler))
	r.GET("/api/v1/user/profile", auth.MustAuthorize(user.GetProfileHandler))
	r.POST("/api/v1/user/activate", auth.MustAuthorize(user.ActivationHandler))
	r.POST("/api/v1/user/changepassword", auth.MustAuthorize(user.ChangePasswordHandler))

	r.GET("/api/v1/course/summary", auth.MustAuthorize(course.GetSummaryHandler))
	r.GET("/api/v1/assignment/incomplete", auth.MustAuthorize(assignment.GetIncompleteHandler))
	r.GET("/api/v1/assignment/summary", auth.MustAuthorize(assignment.GetSummaryHandler))
	r.GET("/api/v1/attendance/summary", auth.MustAuthorize(attendance.GetSummaryHandler))
	r.GET("/api/v1/notification", auth.MustAuthorize(notification.GetHandler))
}
