package routehandlers

import (
	"constants"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	types "example.com/declarations"
	"github.com/gorilla/mux"
)

// metaData represents meta data constants in a struct.
var metaData = constants.HeaderData()

//?---------------- HOME PAGE ROUTE HANDLER ----------------?//

func HomePage(tasks *[]types.Task) types.RouteHandlerFunc {
	fmt.Println("I am the home page!")
	return func(w http.ResponseWriter, r *http.Request) {}
}

//?---------------- GET ALL TASKSROUTE HANDLER ----------------?//

// GetTasks returns a callback that returns a json encoded response to the requesting application.
func GetTasks(tasks *[]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// HTTP Header setup
		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		// returns json data back to the client
		json.NewEncoder(w).Encode(*tasks)

		fmt.Printf("\n\nTasks sent to client:\n\n%v", *tasks)
	}

}

//?---------------- GET SINGLE TASK ROUTE HANDLER ----------------?//

// GetTask returns a callback that returns a json encoded response to the requesting application.
func GetTask(tasks *[]types.Task) types.RouteHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		// parameters recieved with client request
		params := mux.Vars(r)

		isPresent := false

		fmt.Printf("\nmux.Vars(r) value:\n\n%v", params)

		for i := 0; i < len(*tasks); i++ {

			if params["id"] == (*tasks)[i].ID {
				isPresent = true

				json.NewEncoder(w).Encode((*tasks)[i])

				fmt.Printf("\n\nTask id: %v sent to client:\n\n%v", params["id"], (*tasks)[i])

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

}

//?---------------- CREATED TASK ROUTE HANDLER ----------------?//

func CreateTask(tasks *[]types.Task) types.RouteHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

		var newTask types.Task

		// decode the request recieved from the client

		json.NewDecoder(r.Body).Decode(&newTask)

		// assign new task an id

		newTask.ID = strconv.Itoa(rand.Intn(1000))

		// TODO: replace with a utility function | DON'T REPEAT YOURSELF
		currentTime := time.Now().Format("01-02-2006")

		newTask.CreatedDate = currentTime

		// append recieved decoded task to tasks Slice

		*tasks = append(*tasks, newTask)

		// send the new task Slice as a response

		json.NewEncoder(w).Encode(*tasks)

		fmt.Printf("\n\nNew task created: \n\n%v", newTask)
	}

}

//?---------------- UPDATED TASK ROUTE HANDLER ----------------?//

// TODO: use hashmap instead for constant time O(1) opperations instead of a linear O(n) for loop | a database would be even better

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

				// TODO: replace with a utility function | DON'T REPEAT YOURSELF

				currentTime := time.Now().Format("01-02-2006")

				updatedTask.CreatedDate = currentTime

				*tasks = append(*tasks, updatedTask)

				fmt.Printf("\n\ntasks with appended updatedTask: %v", *tasks)

				json.NewEncoder(w).Encode(*tasks)

				fmt.Printf("\n\nUpdated task: %v", updatedTask)

				break

			}

		}

		if isPresent == false {
			// TODO: replace with a utility function | DON'T REPEAT YOURSELF
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("there was an issue retrieving your data, TaskId: %v may not exist.", params["id"])})
		}
	}

}

//?---------------- DELETE TASK ROUTE HANDLER ----------------?//

func DeleteTask(tasks *[]types.Task) types.RouteHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)
		params := mux.Vars(r)
		isPresent := false

		// O(N) --> O(1) ? Arrays to HashMaps for lookup operations

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

			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("there is no task with the id: %v.", params["id"])})

		}
	}
}
