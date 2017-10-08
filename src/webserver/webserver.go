package webserver

import (
	"log"
	"net/http"
	"os"
	"time"

	"fmt"

	"github.com/julienschmidt/httprouter"
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
	log.Println("Initializing web server")
	l := log.New(os.Stdout, "[meiko] ", 0)
	port := fmt.Sprintf(":%s", cfg.Port)
	r := httprouter.New()
	loadRouter(r)

	http.ListenAndServe(port, requestLogger{Handle: r, Logger: l})
}

func (rl requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rl.Handle.ServeHTTP(w, r)
	log.Printf("[meiko] %s %s in %v", r.Method, r.URL.Path, time.Since(start))
}
