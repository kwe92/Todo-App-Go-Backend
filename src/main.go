package main

// TODO: Try using Pointers instead of passing the list by value

// TODO: convert to a random number function

// TODO: add better logs

// TODO: Retrive from a database

import (
	"constants" // needed to encode json data and send it back to the requesting application
	"fmt"       // used to format text
	"log"       // the Go math package from the standard library
	"net/http"
	"routehandlers"

	types "example.com/declarations"
	"github.com/gorilla/mux" // used to create a router
)

//?----------------  INITIAL TASKS ----------------?//

var endpoints = constants.Endpoints()

// var tasks = make([]types.Task, 0)

var tasks []types.Task

func defaultTasks() {

	task0 := types.Task{
		ID:          "1001",
		TaskName:    "Create Your First Task",
		TaskDetails: "One task at a time!",
	}

	tasks = append(tasks, task0)

}

//?---------------- HOME PAGE ROUTE HANDLER ----------------?//

func homePage(tasks []types.Task) types.RouteHandlerFunc {
	fmt.Println("I am the home page!")
	return func(w http.ResponseWriter, r *http.Request) {}
}

//?----------------  HANDLE ALL ROUTES ----------------?//

// handleRoutes handles all the routes for this API.
func handleRoutes() {

	// router instance
	router := mux.NewRouter()

	// all API endpoints
	router.HandleFunc("/", homePage(tasks)).Methods("GET")

	router.HandleFunc(endpoints.GetTasks, routehandlers.GetTasks(tasks)).Methods("GET")

	router.HandleFunc(endpoints.GetTask, routehandlers.GetTask(tasks)).Methods("GET")

	router.HandleFunc(endpoints.CreateTask, routehandlers.CreateTask(tasks)).Methods("POST")

	router.HandleFunc(endpoints.UpdateTask, routehandlers.CreateTask(tasks)).Methods("PUT")

	router.HandleFunc(endpoints.DeleteTask, routehandlers.DeleteTask(tasks)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", router))
}

//?----------------  MAIN FUNCTION ----------------?//

func main() {
	defaultTasks()

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
