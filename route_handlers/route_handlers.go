package routehandlers

import (
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

		fmt.Printf("\n\nrequest URL: %v", utils.GetUrl(r))

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

		fmt.Printf("\n\nrequest URL: %v", utils.GetUrl(r))

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

		fmt.Printf("\n\nrequest URL: %v", utils.GetUrl(r))

		utils.SetHeader(w)

		var newTask types.Task

		// decode the request recieved from the client
		utils.JsonDecode(r, &newTask)

		taskId := utils.GetId()

		// assign new task an id
		newTask.ID = taskId

		currentDate := utils.GetDate()

		newTask.CreatedDate = currentDate

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

		fmt.Printf("\n\nrequest URL: %v", utils.GetUrl(r))

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

			updatedTask.CreatedDate = currentTime

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

		fmt.Printf("\n\nrequest URL: %v", utils.GetUrl(r))

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

// mux.Vars(r)

//   - retrieves the parameters specified from a URI
//       - e.g.localhost/gettask/{id} | 127.0.0.1/gettask/1001

// HTTP Header setup

//   - w.Header().Set("Content-Type", "application/json")
