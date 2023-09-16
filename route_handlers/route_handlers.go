package routehandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	utils "utilities"

	types "example.com/declarations"
	"github.com/gorilla/mux"
)

// taskError returns an error messange if a task does not exist.
func taskError(w http.ResponseWriter, id string) {

	fmt.Printf("\n\nClient requested task id: %v which doesn't exist in the list of tasks", id)

	errorMap := map[string]string{"error": fmt.Sprintf("there was an issue retrieving your data, TaskId: %v may not exist", id)}

	utils.JsonEncode(w, errorMap)

}

//---------------- HOME PAGE ROUTE HANDLER ----------------//

func HomePage(tasks *map[string]types.Task) types.RouteHandlerFunc {

	fmt.Println("I am the home page!")

	return func(w http.ResponseWriter, r *http.Request) {}

}

//---------------- GET ALL TASK ROUTE HANDLER ----------------//

// GetTasks returns all tasks as a json encoded response to the requesting application.
func GetTasks(tasks *map[string]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.SetHeader(w)

		// write a response of json data back to the client
		utils.JsonEncode(w, *tasks)

		fmt.Printf("\n\nTasks sent to client:\n\n%v", *tasks)
	}

}

//---------------- GET SINGLE TASK ROUTE HANDLER ----------------//

// GetTasks returns the specified task as a json encoded response to the requesting application.
func GetTask(tasks *map[string]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.SetHeader(w)

		isPresent := false

		// parameters recieved with client request
		params := mux.Vars(r)

		taskId := params["id"]

		task, keyExsists := (*tasks)[taskId]

		fmt.Printf("\nroute variables:\n\n%v", params)

		if keyExsists {
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

//---------------- CREATED TASK ROUTE HANDLER ----------------//

func CreateTask(tasks *map[string]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.SetHeader(w)

		var newTask types.Task

		// decode the request recieved from the client
		json.NewDecoder(r.Body).Decode(&newTask)

		taskId := utils.GetId()

		// assign new task an id
		newTask.ID = taskId

		currentTime := utils.GetDate()

		newTask.CreatedDate = currentTime

		// append recieved decoded task to tasks Slice
		(*tasks)[taskId] = newTask

		// send the new task Slice as a response
		utils.JsonEncode(w, *tasks)

		fmt.Printf("\n\nNew task created: \n\n%v", newTask)
	}

}

//---------------- UPDATED TASK ROUTE HANDLER ----------------//

func UpdateTask(tasks *map[string]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.SetHeader(w)

		isPresent := false

		params := mux.Vars(r)

		taskId := params["id"]

		_, keyExists := ((*tasks)[taskId])

		if keyExists {

			isPresent = true

			var updatedTask types.Task

			json.NewDecoder(r.Body).Decode(&updatedTask)

			currentTime := utils.GetDate()

			updatedTask.CreatedDate = currentTime

			(*tasks)[taskId] = updatedTask

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

func DeleteTask(tasks *map[string]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		utils.SetHeader(w)

		isPresent := false

		params := mux.Vars(r)

		taskId := params["id"]

		deleteTask, keyExists := (*tasks)[taskId]

		if keyExists {

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
