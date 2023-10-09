package routehandlers

import (
	"constants"
	"fmt"
	"net/http"
	utils "utilities"

	types "example.com/declarations"
	"github.com/gorilla/mux"
)

// taskError returns an error message if a task does not exist.
func taskError(w http.ResponseWriter, id string) {

	fmt.Printf("\n\nClient requested task id: %v which doesn't exist in the list of tasks", id)

	errorMap := map[string]string{"error": fmt.Sprintf("there was an issue retrieving your data, TaskId: %v may not exist", id)}

	utils.JsonEncode(w, errorMap)

}

//---------------- HOME PAGE ROUTE HANDLER ----------------//

func HomePage(tasks *types.TaskMap) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {}

}

//---------------- GET ALL TASK ROUTE HANDLER ----------------//

// GetTasks returns all tasks as a json encoded response to the requesting application.
func GetTasks(tasks *types.TaskMap) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.ParseURL(r)

		utils.SetHeader(w)

		// write a response of json data back to the client
		utils.JsonEncode(w, *tasks)

		fmt.Printf("\n\nTasks sent to client:\n\n%v", *tasks)
	}

}

//---------------- GET SINGLE TASK ROUTE HANDLER ----------------//

// GetTask returns the specified task as a json encoded response to the requesting application.
func GetTask(tasks *types.TaskMap) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.ParseURL(r)

		utils.SetHeader(w)

		isPresent := false

		// parameters recieved with client request
		params := mux.Vars(r)

		taskId := params["id"]

		task, keyExists := (*tasks)[taskId]

		fmt.Printf("\nroute variables:\n\n%v", params)

		if keyExists {
			isPresent = true

			utils.JsonEncode(w, task)

			fmt.Printf("\n\nTask id: %v sent to client:\n\n%v", taskId, task)

			return
		}

		if isPresent == false {

			taskError(w, params["id"])

		}
	}

}

//---------------- CREATE TASK ROUTE HANDLER ----------------//

// CreateTask adds the requested task to the tasks map and returns all tasks as a json encoded response to the requesting application.
func CreateTask(tasks *types.TaskMap) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.ParseURL(r)

		utils.SetHeader(w)

		var newTask types.Task

		// decode the request recieved from the client
		utils.JsonDecode(r, &newTask)

		taskId := utils.GetId()

		// assign new task an id
		newTask.ID = taskId

		currentDate := utils.GetDate()

		newTask.CreatedDate = currentDate

		newTask.ModifiedDate = currentDate

		// update task map with the new task value
		(*tasks)[taskId] = newTask

		// send the new tasks map as a response
		utils.JsonEncode(w, *tasks)

		fmt.Printf("\n\nNew task created: \n\n%v", newTask)
	}

}

//---------------- UPDATED TASK ROUTE HANDLER ----------------//

// UpdateTask updates the requested task and returns all tasks as a json encoded response to the requesting application.
func UpdateTask(tasks *types.TaskMap) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.ParseURL(r)

		utils.SetHeader(w)

		isPresent := false

		params := mux.Vars(r)

		taskId := params["id"]

		previousTask, keyExists := ((*tasks)[taskId])

		if keyExists {

			isPresent = true

			var updatedTask types.Task

			utils.JsonDecode(r, &updatedTask)

			currentTime := utils.GetDate()

			updatedTask.CreatedDate = previousTask.CreatedDate

			updatedTask.ModifiedDate = currentTime

			(*tasks)[taskId] = updatedTask

			fmt.Printf("\n\nPrevious task: %v", previousTask)

			fmt.Printf("\n\nUpdated task: %v", updatedTask)

			fmt.Printf("\n\ntasks with appended updated task: %v", *tasks)

			utils.JsonEncode(w, *tasks)

			return

		}

		if isPresent == false {
			taskError(w, params["id"])
		}
	}

}

//---------------- DELETE TASK ROUTE HANDLER ----------------//

// DeleteTask deletes the requested task and returns all tasks as a json encoded response to the requesting application.
func DeleteTask(tasks *types.TaskMap) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.ParseURL(r)

		utils.SetHeader(w)

		isPresent := false

		params := mux.Vars(r)

		taskId := params["id"]

		deleteTask, keyExists := (*tasks)[taskId]

		if keyExists {

			isPresent = true

			delete((*tasks), taskId)

			utils.JsonEncode(w, *tasks)

			fmt.Printf("\n\ntask deleted: \n\n%v", deleteTask)

			return

		}

		if isPresent == false {

			taskError(w, params["id"])

		}
	}
}

const (
	Get     = "GET"
	Post    = "POST"
	Put     = "PUT"
	Delete  = "DELETE"
	Address = ":8082"
)

// endpoints represent endpoint constants in a struct.
var endpoints = constants.Endpoints()

// handleRoutes handles all the routes for this API.
func SetUpRouter(tasks *types.TaskMap) *mux.Router {

	// router instance
	router := mux.NewRouter()

	// all API endpoints
	router.HandleFunc(endpoints.Home, HomePage(tasks)).Methods(Get)

	router.HandleFunc(endpoints.GetTasks, GetTasks(tasks)).Methods(Get)

	router.HandleFunc(endpoints.GetTask, GetTask(tasks)).Methods(Get)

	router.HandleFunc(endpoints.CreateTask, CreateTask(tasks)).Methods(Post)

	router.HandleFunc(endpoints.UpdateTask, UpdateTask(tasks)).Methods(Put)

	router.HandleFunc(endpoints.DeleteTask, DeleteTask(tasks)).Methods(Delete)

	return router
}

// http.Request

//   - a struct defined in the http package containing
//     important fields and method implementations
//     to represent a received HTTP request from a client to the server

// Body Fields | http.Request

//   - contains the data of a request sent by the caller `client`
//   - has a type of io.ReadCloser that needs to be converted to a string
//   - typcially all of the content form a io.ReadCloser gets read into a Slice of bytes
//     using helper packages like ioutil or the content can be buffered using the bytes package

// URL Field | http.Request

//   - Contains parsed URL meta-data that can be accessed such as:
//       - host, path, and query parameters

// Header Method | http.Request

//   - returns an http.Header type that is a map with additional method implementations
//     which represents the key-value pairs in an HTTP header

// http.ResponseWriter

//   - an interface defined in the http package
//   - provides a server the required method signatures to create HTTP responses
//     for received client requests

// Write Method | http.ResponseWriter

//   - an implementation of io.Writer and writes data to the response

// Header Method | http.ResponseWriter

//   - returns an http.Header type, a defined map with additional methods containing http header meta data

// WriteHeader Method | http.ResponseWriter

//   - takes a status code as an argument and sends an HTTP response Header
//     along with the status code

// mux.Vars(r)

//   - Returns the route variables for the current request

// json package

//   - implements JSON encoding and decoding

// json.NewEncoder(http.ResponseWriter)

//   - returns a new encoder that writes to an http.ResponseWriter

// json.NewEncoder(w).Encode(v)

//   - writes the json encoding of v to the stream
//   - returns json response back to the client

// json.NewDecoder(http.Request)

//   - returns a new decoder that reads from r
//   - the decoder buffers and may read beyond the content of the http.Request io.ReadCloser

// json.NewDecoder(http.Request).Decode(&v)

//   - reads the next JSON-encoded value from its input and stores it in the variable pointed to by v
//   = stores io.ReadCloser content in memory using a pointer to a defined variable
