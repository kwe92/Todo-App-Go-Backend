package main

// TODO: separate into diffrent packages | structures etc when you know how to have a multi module workspace

// TODO: convert to a random number function

// TODO: add better logs

// TODO: Retrive from a database

import (
	"encoding/json" // needed to encode json data and send it back to the requesting application
	"fmt"           // used to format text
	"log"
	"math/rand" // the Go math package from the standard library
	"net/http"
	"strconv"
	"time"

	types "example.com/declarations"

	"github.com/gorilla/mux" // used to create a router
)

//! ----------------  CONSTANTS ----------------//

// Pseudo Enum

const (
	ContentTypeHeader = "Content-Type"
	MediaTypeJson     = "application/json"
)

const (
	GetTask    = "/gettask/{id}"
	GetTasks   = "/gettasks"
	CreateTask = "/create"
	UpdateTask = "/update/{id}"
	DeleteTask = "/delete/{id}"
)

//! ----------------  INITIAL TASKS ----------------//

var tasks []types.Task

func allTasks() {
	task0 := types.Task{
		ID:          "1001",
		TaskName:    "Complete What You Start",
		TaskDetails: "If you start it, finish it by the continuous beginning of the task set before you.",
		CreatedDate: "2023-08-15",
	}

	task1 := types.Task{
		ID:          "1002",
		TaskName:    "Work On Pay Off!",
		TaskDetails: "It pays to payoff, so payoff as you pay",
		CreatedDate: "2023-08-15",
	}

	task2 := types.Task{
		ID:          "1003",
		TaskName:    "In God We Find Strength!",
		TaskDetails: "But they that wait upon the Lord shall renew their strength;",
		CreatedDate: "2023-09-04",
	}

	tasks = append(tasks, task0, task1, task2)
}

//! ---------------- HOME PAGE ROUTE HANDLER ----------------//

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am the home page!")
}

//! ---------------- GET ALL TASKS ROUTE HANDLER ----------------//

// a callback that returns a json encoded response to the requesting application.

func gettasks(w http.ResponseWriter, r *http.Request) {

	// HTTP Header setup
	w.Header().Set(ContentTypeHeader, MediaTypeJson)

	// returns json data back to the client
	json.NewEncoder(w).Encode(tasks)

	fmt.Printf("\n\nTasks sent to client:\n\n%v", tasks)

}

//! ---------------- GET SINGLE TASK ROUTE HANDLER ----------------//

// a callback that returns a json encoded response to the requesting application.

func gettask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(ContentTypeHeader, MediaTypeJson)

	// parameters recieved with client request
	params := mux.Vars(r)

	isPresent := false

	fmt.Printf("\nmux.Vars(r) value:\n\n%v", params)

	for i := 0; i < len(tasks); i++ {

		if params["id"] == tasks[i].ID {
			isPresent = true

			json.NewEncoder(w).Encode(tasks[i])

			fmt.Printf("\n\nTask id: %v sent to client:\n\n%v", params["id"], tasks[i])

			break
		}
	}

	if isPresent == false {

		fmt.Printf("\n\nClient requested task id: %v which doesn't exist in the list of tasks", params["id"])

		// TODO: replace with a utility function | DON'T REPEAT YOURSELF

		// if there is no matching id send a json error hashmap to the client as a response

		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("there was an issue retrieving your data, TaskId: %v may not exist", params["id"])})

	}

}

//! ---------------- CREATED TASK ROUTE HANDLER ----------------//

func createTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(ContentTypeHeader, MediaTypeJson)

	var newTask types.Task

	// decode the request recieved from the client

	json.NewDecoder(r.Body).Decode(&newTask)

	// assign new task an id

	newTask.ID = strconv.Itoa(rand.Intn(1000))

	// TODO: replace with a utility function | DON'T REPEAT YOURSELF
	currentTime := time.Now().Format("01-02-2006")

	newTask.CreatedDate = currentTime

	// append recieved decoded task to tasks Slice

	tasks = append(tasks, newTask)

	// send the new task Slice as a response

	json.NewEncoder(w).Encode(tasks)

	fmt.Printf("\n\nNew task created: \n\n%v", newTask)

}

//! ---------------- UPDATED TASK ROUTE HANDLER ----------------//

// TODO: use hashmap instead for constant time O(1) opperations instead of a linear O(n) for loop | a database would be even better

func updateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(ContentTypeHeader, MediaTypeJson)

	params := mux.Vars(r)

	isPresent := false

	for index, task := range tasks {

		if task.ID == params["id"] {
			isPresent = true

			tasks = append(tasks[:index], tasks[index+1:]...)

			var updatedTask types.Task

			json.NewDecoder(r.Body).Decode(&updatedTask)

			// TODO: replace with a utility function | DON'T REPEAT YOURSELF

			currentTime := time.Now().Format("01-02-2006")

			updatedTask.CreatedDate = currentTime

			tasks = append(tasks, updatedTask)

			fmt.Printf("\n\ntasks with appended updatedTask: %v", tasks)

			json.NewEncoder(w).Encode(tasks)

			fmt.Printf("\n\nUpdated task: %v", updatedTask)

			break

		}

	}

	if isPresent == false {
		// TODO: replace with a utility function | DON'T REPEAT YOURSELF
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("there was an issue retrieving your data, TaskId: %v may not exist.", params["id"])})
	}

}

//! ---------------- DELETE TASK ROUTE HANDLER ----------------//

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentTypeHeader, MediaTypeJson)
	params := mux.Vars(r)
	isPresent := false

	// O(N) --> O(1) ? Arrays to HashMaps for lookup operations

	for index, task := range tasks {

		if params["id"] == task.ID {

			isPresent = true

			tasks = append(tasks[:index], tasks[index+1:]...)

			json.NewEncoder(w).Encode(tasks)

			fmt.Printf("\n\ntask deleted: \n\n%v", task)

			break
		}
	}

	if isPresent == false {

		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("there is no task with the id: %v.", params["id"])})

	}
}

//! ----------------  HANDLE ALL ROUTES ----------------//

// handleRoutes handles all the routes for this API.

func handleRoutes() {

	// router instance

	router := mux.NewRouter()

	// all API endpoints

	// TODO: Review all http method types

	router.HandleFunc("/", homePage).Methods("GET")

	router.HandleFunc(GetTasks, gettasks).Methods("GET")

	router.HandleFunc(GetTask, gettask).Methods("GET")

	router.HandleFunc(CreateTask, createTask).Methods("POST")

	router.HandleFunc(UpdateTask, updateTask).Methods("PUT")

	router.HandleFunc(DeleteTask, deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", router))
}

//! ----------------  MAIN FUNCTION ----------------//

func main() {

	// setup initial tasks

	allTasks()

	fmt.Printf("\nServer has started successfully!\n")

	handleRoutes()

}

// TODO: cleanup comments | transfer to physical notes?

// w http.ResponseWriter, r *http.Request

//   - required parameters for the callback passed to router.HandleFunc
//     to handle http requests and responses

//   - type wr to auto-complete callback parameters

//   - e.g. func homePage(w http.ResponseWriter, r *http.Request)

//   ~ r *http.Request

//       + used to retrieve query parameters and post request data

//   ~  r.Body

//        + the body of the Post request sent by the caller `client`

//        + json.NewDecoder(r.Body).Decode(&task)

//            * the request body sent by the client is decoded and
//              stored in memory using a pointer by reference to a defined variable

//            * underscore (_) is used because we are using a reference in memory not a value

//   ~ w http.ResponseWriter

//       + writes a response to the caller

//       + json.NewEncoder(w).Encode(tasks)

//           * returns json response back to the client
//           * does not require a return statement

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

//   ~ you can use localhost or 127.0.0.1 as a prefix to the port you specified

//       + e.g. localhost:8082/gettasks or 127.0.0.1:8082/gettasks

// What Happens when :router.HandleFunc("/create", createTask).Methods("POST") is called?

//   - the createTask callback is passed a request object from the caller
//     that contains the r.Body of the Post request

// if there is not enough elements in the fixed-length array a new array is created with all previous elements plus the added element
// tasks = append(tasks, task)

// Enumerated Types Go

//   - there are no default enums in Go
//   - you can define a set of constant values and use them throughout your Go application

// TODO: Review Listen and serve
// log.Fatal(http.ListenAndServe(":8082", router))
