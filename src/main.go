package main

// TODO: add better logs
// TODO: review Content-Type an what can suffix application/ | e.g. application/json

import (
	"app_router"
	"fmt"
	utils "utilities"

	types "example.com/declarations"
)

//----------------  INITIAL TASKS ----------------//

var tasks []types.Task

var tasksMap = make(types.TaskMap)

// defaultTasks assigns initial tasks.
func defaultTasks(tasks *types.TaskMap, defaultTasks []types.Task) {

	for _, task := range defaultTasks {
		(*tasks)[task.ID] = task
	}

}

func init() {
	var task0 = types.Task{
		ID:           "1001",
		TaskName:     "Create Your First Task",
		TaskDetails:  "One task at a time!",
		CreatedDate:  utils.GetDate(),
		ModifiedDate: utils.GetDate(),
	}

	// assign inital tasks.
	defaultTasks(&tasksMap, []types.Task{task0})

}

//----------------  MAIN FUNCTION ----------------//

func main() {

	fmt.Printf("\nServer has started successfully!\n")

	app_router.HandleRoutes(&tasksMap)

}
