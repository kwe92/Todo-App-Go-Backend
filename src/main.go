package main

import (
	"encoding/json" // needed to encode json data and send it back to the requesting application
	"fmt"           // used to format text
	"log"
	"net/http"

	"github.com/gorilla/mux" // used to create a router
)

type Task struct {
	ID          string `json:"id"`
	TaskName    string `json:"taskName"`
	TaskDetails string `json:"taskDetails"`
	CreatedDate string `json:"createdDate"`
}

var tasks []Task

func allTasks() {
	task0 := Task{
		ID:          "1001",
		TaskName:    "Complete What You Start",
		TaskDetails: "If you start it, finish it by the continuous beginning of the task set before you.",
		CreatedDate: "2023-08-15",
	}

	task1 := Task{
		ID:          "1002",
		TaskName:    "Work On Pay Off!",
		TaskDetails: "It pays to payoff, so payoff as you pay",
		CreatedDate: "2023-08-15",
	}

	tasks = append(tasks, task0, task1)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am the home page!")
}

// a callback that returns a json encoded response to the requesting application.
func gettasks(w http.ResponseWriter, r *http.Request) {

	// HTTP Header setup
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)

}

// a callback that returns a json encoded response to the requesting application.
func gettask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)

	flag := false

	fmt.Printf("\nmux.Vars(r) value:\n\n%v\n", taskId)

	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}

	if flag == false {

		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("there was an issue retrieving your data, TaskId: %v may not exist", taskId["id"])})
	}

}
func createTask(w http.ResponseWriter, r *http.Request) {}
func updateTask(w http.ResponseWriter, r *http.Request) {}
func deleteTask(w http.ResponseWriter, r *http.Request) {}

func handleRoutes() {
	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", gettasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", gettask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	// setup initial tasks
	allTasks()

	fmt.Println("\nServer has started successfully!\n")

	handleRoutes()

	// fmt.Println("\nHello Go Developer!\n")

}

// TODO: cleanup comments

// type wr in the function header and it will auto-complete what you need as parameters for the callback

//   - func homePage(w http.ResponseWriter, r *http.Request)

// handleRoutes function

//   - use the router to create some route
//   - "/" represents the home page
//   - similar to larvel or node.js
//   - three things you need: route name, route function and route method

// HTTP Header setup
// TODO: review Content-Type an what can suffix application/ | e.g. application/json
//   - w.Header().Set("Content-Type", "application/json")

// mux.Vars(r)

//   - retrieves the parameters specified from a URI
//       - e.g.localhost/gettask/{id} | 127.0.0.1/gettask/1001

// pinging local host

//   - you can ping local host with Postman desktop client
//   - you can use localhost or 127.0.0.1 as a prefix to the port you specified
//       - e.g. localhost:8082/gettasks or 127.0.0.1:8082/gettasks
