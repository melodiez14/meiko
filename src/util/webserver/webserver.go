package webserver

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"fmt"

	"github.com/julienschmidt/httprouter"
	router "github.com/melodiez14/lastcake/src"
	"github.com/tokopedia/grace"
)

// Config is used for the setting of web server
type Config struct {
	Port string
}

type requestLogger struct {
	Handle http.Handler
	Logger *log.Logger
}

// Start is used by app.go for starting the webserver
func Start(cfg Config) {
	l := log.New(os.Stdout, "[lastcake] ", 0)
	port := fmt.Sprintf(":%s", cfg.Port)
	r := httprouter.New()
	router.Load(r)
	grace.Serve(port, requestLogger{Handle: r, Logger: l})
}

func (rl requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rl.Handle.ServeHTTP(w, r)
	log.Printf("[%s] %s %s in %v", strings.Replace(os.Args[0], "./", "", -1), r.Method, r.URL.Path, time.Since(start))
}
