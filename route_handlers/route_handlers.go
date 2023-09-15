package routehandlers

import (
	"constants"
	"encoding/json"
	"fmt"
	"net/http"
	utils "utilities"

	types "example.com/declarations"
	"github.com/gorilla/mux"
)

func taskError(w http.ResponseWriter, id string) {

	fmt.Printf("\n\nClient requested task id: %v which doesn't exist in the list of tasks", id)

	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("there was an issue retrieving your data, TaskId: %v may not exist", id)})

}

// metaData represents meta data constants in a struct.
var metaData = constants.HeaderData()

//---------------- HOME PAGE ROUTE HANDLER ----------------//

func HomePage(tasks *[]types.Task) types.RouteHandlerFunc {

	fmt.Println("I am the home page!")

	return func(w http.ResponseWriter, r *http.Request) {}

}

// TODO: refactor operations from O(n) --> O(1)

//---------------- GET ALL TASK ROUTE HANDLER ----------------//

// GetTasks returns all tasks as a json encoded response to the requesting application.
func GetTasks(tasks *[]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// HTTP Header setup
		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		// write a response of json data back to the client
		json.NewEncoder(w).Encode(*tasks)

		fmt.Printf("\n\nTasks sent to client:\n\n%v", *tasks)
	}

}

//---------------- GET SINGLE TASK ROUTE HANDLER ----------------//

// GetTasks returns the specified task as a json encoded response to the requesting application.
func GetTask(tasks *[]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		// parameters recieved with client request
		params := mux.Vars(r)

		isPresent := false

		fmt.Printf("\nroute variables:\n\n%v", params)

		for i := 0; i < len(*tasks); i++ {

			if params["id"] == (*tasks)[i].ID {
				isPresent = true

				json.NewEncoder(w).Encode((*tasks)[i])

				fmt.Printf("\n\nTask id: %v sent to client:\n\n%v", params["id"], (*tasks)[i])

				break
			}
		}

		if isPresent == false {

			taskError(w, params["id"])

		}
	}

}

//---------------- CREATED TASK ROUTE HANDLER ----------------//

func CreateTask(tasks *[]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		var newTask types.Task

		// decode the request recieved from the client
		json.NewDecoder(r.Body).Decode(&newTask)

		// assign new task an id
		newTask.ID = utils.GetId()

		currentTime := utils.GetDate()

		newTask.CreatedDate = currentTime

		// append recieved decoded task to tasks Slice
		*tasks = append(*tasks, newTask)

		// send the new task Slice as a response
		json.NewEncoder(w).Encode(*tasks)

		fmt.Printf("\n\nNew task created: \n\n%v", newTask)
	}

}

//---------------- UPDATED TASK ROUTE HANDLER ----------------//

func UpdateTask(tasks *[]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		params := mux.Vars(r)

		isPresent := false

		for index, task := range *tasks {

			if task.ID == params["id"] {
				isPresent = true

				*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)

				var updatedTask types.Task

				json.NewDecoder(r.Body).Decode(&updatedTask)

				currentTime := utils.GetDate()

				updatedTask.CreatedDate = currentTime

				*tasks = append(*tasks, updatedTask)

				fmt.Printf("\n\nUpdated task: %v", updatedTask)

				fmt.Printf("\n\ntasks with appended updated task: %v", *tasks)

				json.NewEncoder(w).Encode(*tasks)

				break

			}

		}

		if isPresent == false {
			taskError(w, params["id"])
		}
	}

}

//---------------- DELETE TASK ROUTE HANDLER ----------------//

func DeleteTask(tasks *[]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		params := mux.Vars(r)

		isPresent := false

		for index, task := range *tasks {

			if params["id"] == task.ID {

				isPresent = true

				*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)

				json.NewEncoder(w).Encode(*tasks)

				fmt.Printf("\n\ntask deleted: \n\n%v", task)

				break
			}
		}

		if isPresent == false {

			taskError(w, params["id"])

		}
	}
}
