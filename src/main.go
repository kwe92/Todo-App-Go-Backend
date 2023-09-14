package main

// TODO: add better logs

import (
	"app_router"
	"fmt"

	types "example.com/declarations"
)

//----------------  INITIAL TASKS ----------------//

var tasks []types.Task

// defaultTasks assigns intial tasks.
func defaultTasks(tasks *[]types.Task, defaultTasks []types.Task) {

	for _, task := range defaultTasks {
		*tasks = append(*tasks, task)
	}

}

func init() {
	var task0 = types.Task{
		ID:          "1001",
		TaskName:    "Create Your First Task",
		TaskDetails: "One task at a time!",
	}
	// assign inital tasks.
	defaultTasks(&tasks, []types.Task{task0})

}

//----------------  MAIN FUNCTION ----------------//

func main() {

	fmt.Printf("\nServer has started successfully!\n")

	app_router.HandleRoutes(&tasks)

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
