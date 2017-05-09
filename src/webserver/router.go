package webserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/handler"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {
	r.GET("/api/v1/hellomeiko", auth.OptionalAuthorize(handler.HelloMeiko))
}
