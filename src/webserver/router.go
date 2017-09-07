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
	r.POST("/api/v1/user/updateuser", auth.MustAuthorize(user.UpdateUserAccountHandler))
	r.POST("/api/v1/user/changepassword", auth.MustAuthorize(user.ChangePasswordHandler))
	r.GET("/api/v1/user/edit", auth.MustAuthorize(user.GetUserAccountHandler))
	r.POST("/api/v1/user/verified", auth.OptionalAuthorize(user.RequestVerifiedUserHandler))
	r.DELETE("/api/v1/user/logout", auth.MustAuthorize(user.LogoutHandler))
	r.POST("/api/v1/user/verify", auth.OptionalAuthorize(user.RequestVerifiedUserHandler))
	r.POST("/api/v1/user/register", auth.OptionalAuthorize(user.SignUpHandler))
	r.POST("/api/v1/user/forgot/request", auth.OptionalAuthorize(user.ForgotRequestHandler))
	r.POST("/api/v1/user/forgot/confirmation", auth.OptionalAuthorize(user.ForgotConfirmation))
	r.POST("/api/v1/user/login", auth.OptionalAuthorize(user.LoginHandler))
	r.GET("/api/v1/user/activate", auth.MustAuthorize(user.ActivationHandler))
	r.GET("/api/v1/user/validated", auth.MustAuthorize(user.GetValidatedUser))
	r.GET("/api/v1/course/summary", auth.MustAuthorize(course.GetSummaryHandler))
	r.GET("/api/v1/assignment/incomplete", auth.MustAuthorize(assignment.GetIncompleteHandler))
	r.GET("/api/v1/assignment/summary", auth.MustAuthorize(assignment.GetSummaryHandler))
	r.GET("/api/v1/attendance/summary", auth.MustAuthorize(attendance.GetSummaryHandler))
	r.GET("/api/v1/notification", auth.MustAuthorize(notification.GetHandler))
}
