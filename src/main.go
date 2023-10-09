package main

// TODO: add better logs
// TODO: review Content-Type an what can suffix application/ | e.g. application/json

import (
	"app_router"
	"fmt"
	"log"
	"net/http"
	"time"
	utils "utilities"

	types "example.com/declarations"
)

//----------------  INITIAL TASKS ----------------//

var tasks []types.Task

var tasksMap = make(types.TaskMap)

// defaultTasks assigns initial tasks.
func defaultTasks(tasks *types.TaskMap, defaultTasks []types.Task) {

	for _, task := range defaultTasks {
		(*tasks)[task.ID] = task
	}

}

func init() {
	var task0 = types.Task{
		ID:           "1001",
		TaskName:     "Create Your First Task",
		TaskDetails:  "One task at a time!",
		CreatedDate:  utils.GetDate(),
		ModifiedDate: utils.GetDate(),
	}

	// assign inital tasks.
	defaultTasks(&tasksMap, []types.Task{task0})

}

//----------------  MAIN FUNCTION ----------------//

func main() {

	fmt.Printf("\nServer has started successfully!\n")

	router := NewLogger(app_router.SetUpRouter(&tasksMap))

	log.Fatal(http.ListenAndServe(app_router.Address, router))

}

// Move Logger into its own package

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
