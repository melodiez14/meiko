package webserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/handler"
	"github.com/melodiez14/meiko/src/webserver/handler/assignment"
	"github.com/melodiez14/meiko/src/webserver/handler/attendance"
	"github.com/melodiez14/meiko/src/webserver/handler/course"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {
	r.GET("/api/v1/hellomeiko", auth.OptionalAuthorize(handler.HelloMeiko))
	r.GET("/api/v1/course/summary", auth.MustAuthorize(course.GetSummaryHandler))
	r.GET("/api/v1/assignment/incomplete", auth.MustAuthorize(assignment.GetIncompleteHandler))
	r.GET("/api/v1/assignment/summary", auth.MustAuthorize(assignment.GetIncompleteHandler))
	r.GET("/api/v1/attendance/summary", auth.MustAuthorize(attendance.GetSummaryHandler))
}
