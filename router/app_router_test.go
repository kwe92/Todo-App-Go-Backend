package app_router

// TODO: add comments on Table Driven Tests and Unit Testing

// TODO: implement all tests with t.Run to name them | Add comments to tests

import (
	"bytes"
	"fmt"
	"io"
	"log"
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
		fmt.Sprintf("GetTasks returns all tasks when endpoint %s has been requested", endpoints.GetTasks),
		func(t *testing.T) {
			// create request to desired endpoint
			req, err := http.NewRequest(Get, endpoints.GetTasks, nil)

			// check if there was an error generating the request
			checkError(err)

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
			matchStatusCode(t, response.Result().StatusCode)

			matchContent(t, expected, received)
		},
	)

}

func TestGetTask(t *testing.T) {

	tasksMap := make(types.TaskMap)

	tasksMap[testTask.ID] = testTask

	req, err := http.NewRequest(Get, fmt.Sprintf("/gettask/%s", testTask.ID), nil)

	router := SetUpRouter(&tasksMap)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	expected := testTask

	var received types.Task

	utils.JsonDecode(io.NopCloser(response.Body), &received)

	matchStatusCode(t, response.Result().StatusCode)

	matchContent(t, expected, received)

	checkError(err)

}

func TestCreateTask(t *testing.T) {

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

	checkError(err)

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

	matchStatusCode(t, response.Result().StatusCode)

	matchContent(t, expected, received)

}

func TestUpdateTask(t *testing.T) {

	var reqBuffer bytes.Buffer

	updatedTask := testTask

	expected := "Man shall not live by bread alone, but by every word that proceedth out of the mouth of God."

	updatedTask.TaskDetails = expected

	utils.JsonEncode(&reqBuffer, updatedTask)

	tasksMap := make(types.TaskMap)

	responseTaskMap := make(types.TaskMap)

	tasksMap[testTask.ID] = testTask

	req, err := http.NewRequest(Put, "/update/1001", &reqBuffer)

	checkError(err)

	response := httptest.NewRecorder()

	router := SetUpRouter(&tasksMap)

	router.ServeHTTP(response, req)

	utils.JsonDecode[types.TaskMap](io.NopCloser(response.Body), &responseTaskMap)

	// fmt.Println("\nResponse TaskMap:", responseTaskMap["1001"].TaskDetails)

	received := responseTaskMap["1001"].TaskDetails

	matchStatusCode(t, response.Result().StatusCode)

	matchContent(t, expected, received)

}

func TestDeleteTask(t *testing.T) {

	req, err := http.NewRequest(Delete, fmt.Sprintf("/delete/%s", testTask.ID), nil)

	checkError(err)

	response := httptest.NewRecorder()

	tasksMap := make(types.TaskMap)

	tasksMap[testTask.ID] = testTask

	router := SetUpRouter(&tasksMap)

	router.ServeHTTP(response, req)

	received := make(types.TaskMap)

	utils.JsonDecode[types.TaskMap](io.NopCloser(response.Body), &received)

	expected := make(types.TaskMap)

	// fmt.Println("\nResponse TaskMap:", received)

	matchStatusCode(t, response.Result().StatusCode)

	matchContent(t, expected, received)

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("\nfailed to create request: %s", err.Error())
	}
}

func matchStatusCode(t *testing.T, statusCode int, expectedStatusCode ...int) {

	var expectedCode int

	if len(expectedStatusCode) > 0 {
		expectedCode = expectedStatusCode[0]
	} else {
		expectedCode = http.StatusOK
	}

	if statusCode != expectedCode {
		t.Fatalf("the status code should be [%d] but received [%d]",
			expectedCode,
			statusCode,
		)
		return
	}
}

func matchContent[T any](t *testing.T, expected T, received T) {
	if fmt.Sprint(received) != fmt.Sprint(expected) {

		t.Fatalf("the response body should be [%s] but received [%s]",
			fmt.Sprint(expected),
			fmt.Sprint(received),
		)

	}

}

// func TestRouter(t *testing.T) {

// 	// test response content
// 	testTaskData := make(types.TaskMap)

// 	var task0 = types.Task{
// 		ID:           "1001",
// 		TaskName:     "Create Your First Task",
// 		TaskDetails:  "One task at a time!",
// 		CreatedDate:  utils.GetDate(),
// 		ModifiedDate: utils.GetDate(),
// 	}

// 	testTaskData[task0.ID] = task0

// 	// setup the router / http multiplexer
// 	router := SetUpRouter(&testTaskData)

// 	// args: holds the created *http.Request as test arguments
// 	type args struct {
// 		req *http.Request
// 	}

// 	// Table Driven Test collection
// 	tests := []struct {
// 		name     string                  // name of the test
// 		args     func(t *testing.T) args // creates and returns the request
// 		wantCode int                     // expected HTTP status code
// 		wantBody string                  // expected response content
// 	}{
// 		// {
// 		// 	name: "must return all tasks",
// 		// 	args: func(t *testing.T) args {

// 		// 		req, err := http.NewRequest(Get, endpoints.GetTasks, nil)

// 		// 		checkError(err)

// 		// 		return args{

// 		// 			req: req,
// 		// 		}
// 		// 	}, wantCode: 200,
// 		// 	wantBody: fmt.Sprint(testTaskData),
// 		// },

// 		// TODO: trouble shoot why response is not right | maybe separate into multiple tests
// 		{
// 			name: "get single task",
// 			args: func(t *testing.T) args {

// 				req, err := http.NewRequest(Get, "/gettask/1001", nil)

// 				// params := req.URL.Query()

// 				// params.Add("id", "1001")

// 				// req.URL.RawQuery = params.Encode()

// 				checkError(err)

// 				return args{
// 					req: req,
// 				}
// 			},
// 			wantCode: 200,
// 			wantBody: fmt.Sprint(testTaskData),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			tArgs := tt.args(t)

// 			// receivedContent := make(types.TaskMap)

// 			response := httptest.NewRecorder()

// 			router.ServeHTTP(response, tArgs.req)

// 			respBody, _ := io.ReadAll(response.Body)

// 			fmt.Println("\nRESPONSE BODY:", string(respBody))

// 			// utils.JsonDecode(io.NopCloser(response.Body), &receivedContent)

// 			// checkError(err)

// 			if response.Result().StatusCode != tt.wantCode {
// 				t.Fatalf("the status code should be [%d] but received [%d]",
// 					tt.wantCode,
// 					response.Result().StatusCode,
// 				)
// 			}

// 			if fmt.Sprint(respBody) != tt.wantBody {
// 				t.Fatalf("the response body should be [%s] but received [%s]",
// 					fmt.Sprint(tt.wantBody),
// 					fmt.Sprint(string(respBody)),
// 				)
// 			}

// 		})
// 	}

// }
