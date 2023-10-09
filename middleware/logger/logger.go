package logger

import (
	"log"
	"net/http"
	"time"
	utils "utilities"
)

// Logger: middleware that logs request metadata.
type Logger struct {
	Handler http.Handler
	Prefix  string
}

// ServeHTTP: handles the request by logging meta data
// and passing the http.ResponseWrite and http.*Request to the Loggers handler.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if len(l.Prefix) == 0 {
		log.SetPrefix("\nLogged Event:")
	} else {
		log.SetPrefix(l.Prefix)
	}

	log.Printf("\n%s %v", r.Method, utils.ParseURL(r))

	start := time.Now()

	l.Handler.ServeHTTP(w, r)

	log.Println("\nElapsed Time:", time.Since(start))

}

// NewLogger: returns a new Logger object wrapped around the handler.
func NewLogger(handler http.Handler, prefix ...string) *Logger {

	var pf string

	if len(prefix) > 0 {
		pf = prefix[0]
	}
	return &Logger{
		Handler: handler,
		Prefix:  pf,
	}
}
