package declarations

import "net/http"

type Task struct {
	ID          string `json:"id"`
	TaskName    string `json:"taskName"`
	TaskDetails string `json:"taskDetails"`
	CreatedDate string `json:"createdDate"`
}

type HeaderMetaData struct {
	ContentTypeHeader string
	MediaTypeJson     string
}

type EndPoints struct {
	GetTask    string
	GetTasks   string
	CreateTask string
	UpdateTask string
	DeleteTask string
}

type RouteHandlerFunc func(http.ResponseWriter, *http.Request)
