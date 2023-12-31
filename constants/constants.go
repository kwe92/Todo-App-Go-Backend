package constants

import (
	types "example.com/declarations"
)

// Pseudo Enums

const (
	ContentTypeHeader = "Content-Type"
	MediaTypeJson     = "application/json"
)

const (
	Home       = "/"
	GetTask    = "/gettask/{id}"
	GetTasks   = "/gettasks"
	CreateTask = "/create"
	UpdateTask = "/update/{id}"
	DeleteTask = "/delete/{id}"
)

// HeaderData returns header meta data constants in a struct.
func HeaderData() types.HeaderMetaData {
	return types.HeaderMetaData{
		ContentTypeHeader: ContentTypeHeader,
		MediaTypeJson:     MediaTypeJson,
	}
}

// Endpoints returns endpoint constants in a struct.
func Endpoints() types.EndPoints {
	return types.EndPoints{
		Home:       Home,
		GetTask:    GetTask,
		GetTasks:   GetTasks,
		CreateTask: CreateTask,
		UpdateTask: UpdateTask,
		DeleteTask: DeleteTask,
	}
}

// Enumerated Types Go

//   - there are no default enums in Go
//   - a set of constant values can be defined and used throughout your Go application
