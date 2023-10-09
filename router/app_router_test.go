package app_router

import (
	"encoding/json"
	"fmt"
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

	tasksMap := make(types.TaskMap)
	tasksMap1 := make(types.TaskMap)

	var task0 = types.Task{
		ID:           "1001",
		TaskName:     "Create Your First Task",
		TaskDetails:  "One task at a time!",
		CreatedDate:  utils.GetDate(),
		ModifiedDate: utils.GetDate(),
	}

	tasksMap[task0.ID] = task0

	router := SetUpRouter(&tasksMap)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	err = json.NewDecoder(response.Body).Decode(&tasksMap1)

	checkError(err)

	if response.Result().StatusCode != 200 {
		t.Fatalf("the status code should be [%d] but received [%d]",
			200,
			response.Result().StatusCode,
		)
	}

	if fmt.Sprint(tasksMap1) != fmt.Sprint(tasksMap) {
		t.Fatalf("the response body should be [%s] but received [%s]",
			fmt.Sprint(tasksMap),
			fmt.Sprint(tasksMap1),
		)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("failed to create request: %s", err.Error())
	}
}
