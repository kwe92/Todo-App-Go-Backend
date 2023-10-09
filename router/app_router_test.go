package app_router

// TODO: add comments on Table Driven Tests and Unit Testing

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	utils "utilities"

	types "example.com/declarations"
)

func TestRouter(t *testing.T) {

	// test response content
	testTaskData := make(types.TaskMap)

	var task0 = types.Task{
		ID:           "1001",
		TaskName:     "Create Your First Task",
		TaskDetails:  "One task at a time!",
		CreatedDate:  utils.GetDate(),
		ModifiedDate: utils.GetDate(),
	}

	testTaskData[task0.ID] = task0

	// setup the router / http multiplexer
	router := SetUpRouter(&testTaskData)

	// args: holds the created *http.Request as test arguments
	type args struct {
		req *http.Request
	}

	// Table Driven Test collection
	tests := []struct {
		name     string                  // name of the test
		args     func(t *testing.T) args // creates and returns the request
		wantCode int                     // expected HTTP status code
		wantBody string                  // expected response content
	}{
		// {
		// 	name: "must return all tasks",
		// 	args: func(t *testing.T) args {

		// 		req, err := http.NewRequest(Get, endpoints.GetTasks, nil)

		// 		checkError(err)

		// 		return args{

		// 			req: req,
		// 		}
		// 	}, wantCode: 200,
		// 	wantBody: fmt.Sprint(testTaskData),
		// },

		// TODO: trouble shoot why response is not right | maybe separate into multiple tests
		{
			name: "get single task",
			args: func(t *testing.T) args {

				req, err := http.NewRequest(Get, "/gettask/1001", nil)

				// params := req.URL.Query()

				// params.Add("id", "1001")

				// req.URL.RawQuery = params.Encode()

				checkError(err)

				return args{
					req: req,
				}
			},
			wantCode: 200,
			wantBody: fmt.Sprint(testTaskData),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tArgs := tt.args(t)

			receivedContent := make(types.TaskMap)

			response := httptest.NewRecorder()

			router.ServeHTTP(response, tArgs.req)

			// respBody, _ := io.ReadAll(response.Body)

			// fmt.Println("\nRESPONSE BODY:", string(respBody))

			utils.JsonDecode(io.NopCloser(response.Body), &receivedContent)

			// checkError(err)

			if response.Result().StatusCode != tt.wantCode {
				t.Fatalf("the status code should be [%d] but received [%d]",
					tt.wantCode,
					response.Result().StatusCode,
				)
			}

			if fmt.Sprint(receivedContent) != tt.wantBody {
				t.Fatalf("the response body should be [%s] but received [%s]",
					fmt.Sprint(tt.wantBody),
					fmt.Sprint(receivedContent),
				)
			}

		})
	}

}

// func TestGetTasksHandler(t *testing.T) {

// 	req, err := http.NewRequest(Get, endpoints.GetTasks, nil)

// 	checkError(err)

// 	expectedContent := make(types.TaskMap)
// 	receivedContent := make(types.TaskMap)

// 	var task0 = types.Task{
// 		ID:           "1001",
// 		TaskName:     "Create Your First Task",
// 		TaskDetails:  "One task at a time!",
// 		CreatedDate:  utils.GetDate(),
// 		ModifiedDate: utils.GetDate(),
// 	}

// 	expectedContent[task0.ID] = task0

// 	router := SetUpRouter(&expectedContent)

// 	response := httptest.NewRecorder()

// 	router.ServeHTTP(response, req)

// 	utils.JsonDecode(io.NopCloser(response.Body), &receivedContent)

// 	if response.Result().StatusCode != 200 {
// 		t.Fatalf("the status code should be [%d] but received [%d]",
// 			200,
// 			response.Result().StatusCode,
// 		)
// 	}

// 	if fmt.Sprint(receivedContent) != fmt.Sprint(expectedContent) {
// 		t.Fatalf("the response body should be [%s] but received [%s]",
// 			fmt.Sprint(expectedContent),
// 			fmt.Sprint(receivedContent),
// 		)
// 	}

// }

func checkError(err error) {
	if err != nil {
		log.Fatalf("\nfailed to create request: %s", err.Error())
	}
}
