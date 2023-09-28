// All the defined types used throughout the server.

package declarations

import "net/http"

type Task struct {
	ID           string `json:"id"`
	TaskName     string `json:"task_name"`
	TaskDetails  string `json:"task_details"`
	CreatedDate  string `json:"created_date"`
	ModifiedDate string `json:"modified_date"`
}

type HeaderMetaData struct {
	ContentTypeHeader string
	MediaTypeJson     string
}

type EndPoints struct {
	Home       string
	GetTask    string
	GetTasks   string
	CreateTask string
	UpdateTask string
	DeleteTask string
}

type RouteHandlerFunc func(http.ResponseWriter, *http.Request)

type TaskMap map[string]Task
