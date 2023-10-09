package app_router

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

func TestGetTasksHandler(t *testing.T) {

	req, err := http.NewRequest(Get, endpoints.GetTasks, nil)

	checkError(err)

	expectedContent := make(types.TaskMap)
	receivedContent := make(types.TaskMap)

	var task0 = types.Task{
		ID:           "1001",
		TaskName:     "Create Your First Task",
		TaskDetails:  "One task at a time!",
		CreatedDate:  utils.GetDate(),
		ModifiedDate: utils.GetDate(),
	}

	expectedContent[task0.ID] = task0

	router := SetUpRouter(&expectedContent)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	utils.JsonDecode(io.NopCloser(response.Body), &receivedContent)

	checkError(err)

	if response.Result().StatusCode != 200 {
		t.Fatalf("the status code should be [%d] but received [%d]",
			200,
			response.Result().StatusCode,
		)
	}

	if fmt.Sprint(receivedContent) != fmt.Sprint(expectedContent) {
		t.Fatalf("the response body should be [%s] but received [%s]",
			fmt.Sprint(expectedContent),
			fmt.Sprint(receivedContent),
		)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("failed to create request: %s", err.Error())
	}
}
