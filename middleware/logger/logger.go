package logger

import (
	"log"
	"net/http"
	"time"
	utils "utilities"
)

// TODO: add coments

type Logger struct {
	Handler http.Handler
	Prefix  string
}

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
