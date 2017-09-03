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
	r.POST("/api/v1/user/forgot/request", auth.OptionalAuthorize(user.ForgotRequestHandler))
	r.POST("/api/v1/user/forgot/confirmation", auth.OptionalAuthorize(user.ForgotConfirmation))
	r.GET("/api/v1/course/summary", auth.MustAuthorize(course.GetSummaryHandler))
	r.GET("/api/v1/assignment/incomplete", auth.MustAuthorize(assignment.GetIncompleteHandler))
	r.GET("/api/v1/assignment/summary", auth.MustAuthorize(assignment.GetSummaryHandler))
	r.GET("/api/v1/attendance/summary", auth.MustAuthorize(attendance.GetSummaryHandler))
	r.GET("/api/v1/notification", auth.MustAuthorize(notification.GetHandler))
}
