package app_router

// TODO: add comments on Table Driven Tests and Unit Testing

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	utils "utilities"

	types "example.com/declarations"
)

var testTask = types.Task{
	ID:           "1001",
	TaskName:     "Create Your First Task",
	TaskDetails:  "One task at a time!",
	CreatedDate:  utils.GetDate(),
	ModifiedDate: utils.GetDate(),
}

func TestGetTasksHandler(t *testing.T) {
	t.Run(
		fmt.Sprintf("GetTasks returns all tasks when endpoint %s receives a request.", endpoints.GetTasks),
		func(t *testing.T) {
			// create request to desired endpoint
			req, err := http.NewRequest(Get, endpoints.GetTasks, nil)

			// check if there was an error generating the request
			utils.CheckError(err)

			// create a map of expected response
			expected := make(types.TaskMap)

			// create a map of received response
			received := make(types.TaskMap)

			expected[testTask.ID] = testTask

			// set up router
			router := SetUpRouter(&expected)

			// create a ResponseRecorder to act as a ResponseWriter
			response := httptest.NewRecorder()

			// match request URL to a pattern of a registered handler and excute the handler, loading the response body
			router.ServeHTTP(response, req)

			// write the response body bytes to a GO data structure
			utils.JsonDecode(io.NopCloser(response.Body), &received)

			// match status code to expected status code
			utils.MatchStatusCode(t, response.Result().StatusCode)

			utils.CompareTestResults(t, expected, received)
		},
	)

}

func TestGetTask(t *testing.T) {
	t.Run(fmt.Sprintf(
		"GetTask returns a single task %s when endpoint /gettask/%s receives a request.",
		testTask.ID,
		testTask.ID,
	), func(t *testing.T) {

		tasksMap := make(types.TaskMap)

		tasksMap[testTask.ID] = testTask

		req, err := http.NewRequest(Get, fmt.Sprintf("/gettask/%s", testTask.ID), nil)

		router := SetUpRouter(&tasksMap)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, req)

		expected := testTask

		var received types.Task

		utils.JsonDecode(io.NopCloser(response.Body), &received)

		utils.MatchStatusCode(t, response.Result().StatusCode)

		utils.CompareTestResults(t, expected, received)

		utils.CheckError(err)
	})

}

func TestCreateTask(t *testing.T) {

	t.Run("CreateTask returns all current tasks with created tasks appended when endpoint /create receives a request.",
		func(t *testing.T) {
			var reqBuffer bytes.Buffer

			var createdTask = types.Task{
				ID:          "1002",
				TaskName:    "Light of God",
				TaskDetails: "And the light shineth in the darkness, and the darkness comprehended it not.",
			}

			utils.JsonEncode(&reqBuffer, createdTask)

			tasksMap := make(types.TaskMap)

			responseTaskMap := make(types.TaskMap)

			tasksMap[testTask.ID] = testTask

			req, err := http.NewRequest(Post, "/create", &reqBuffer)

			utils.CheckError(err)

			router := SetUpRouter(&tasksMap)

			response := httptest.NewRecorder()

			router.ServeHTTP(response, req)

			utils.JsonDecode(io.NopCloser(response.Body), &responseTaskMap)

			var keys []string

			for key := range responseTaskMap {
				keys = append(keys, key)
			}

			// fmt.Println("\nKEYS: ", keys)
			// fmt.Println("\nkeys[len(keys)-1]: ", keys[len(keys)-1])

			received := responseTaskMap[keys[len(keys)-1]].TaskDetails

			expected := createdTask.TaskDetails

			utils.MatchStatusCode(t, response.Result().StatusCode)

			utils.CompareTestResults(t, expected, received)
		})

}

func TestUpdateTask(t *testing.T) {
	t.Run("UpdateTask modifies the requested task and returns all tasks when endpoint /update/1001 receives a request.",
		func(t *testing.T) {

			var reqBuffer bytes.Buffer

			updatedTask := testTask

			expected := "Man shall not live by bread alone, but by every word that proceedth out of the mouth of God."

			updatedTask.TaskDetails = expected

			utils.JsonEncode(&reqBuffer, updatedTask)

			tasksMap := make(types.TaskMap)

			responseTaskMap := make(types.TaskMap)

			tasksMap[testTask.ID] = testTask

			req, err := http.NewRequest(Put, "/update/1001", &reqBuffer)

			utils.CheckError(err)

			response := httptest.NewRecorder()

			router := SetUpRouter(&tasksMap)

			router.ServeHTTP(response, req)

			utils.JsonDecode[types.TaskMap](io.NopCloser(response.Body), &responseTaskMap)

			// fmt.Println("\nResponse TaskMap:", responseTaskMap["1001"].TaskDetails)

			received := responseTaskMap["1001"].TaskDetails

			utils.MatchStatusCode(t, response.Result().StatusCode)

			utils.CompareTestResults(t, expected, received)
		})

}

func TestDeleteTask(t *testing.T) {
	t.Run("DeleteTask deletes task deletes task 1001 and returns all other tasks when endpoint /delete/1001 receives a request.",
		func(t *testing.T) {

			req, err := http.NewRequest(Delete, fmt.Sprintf("/delete/%s", testTask.ID), nil)

			utils.CheckError(err)

			response := httptest.NewRecorder()

			tasksMap := make(types.TaskMap)

			tasksMap[testTask.ID] = testTask

			router := SetUpRouter(&tasksMap)

			router.ServeHTTP(response, req)

			received := make(types.TaskMap)

			utils.JsonDecode[types.TaskMap](io.NopCloser(response.Body), &received)

			expected := make(types.TaskMap)

			// fmt.Println("\nResponse TaskMap:", received)

			utils.MatchStatusCode(t, response.Result().StatusCode)

			utils.CompareTestResults(t, expected, received)
		})

}
