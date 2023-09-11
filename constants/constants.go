package constants

import (
	types "example.com/declarations"
)

// Pseudo Enum
const (
	ContentTypeHeader = "Content-Type"
	MediaTypeJson     = "application/json"
)

const (
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
		GetTask:    GetTask,
		GetTasks:   GetTasks,
		CreateTask: CreateTask,
		UpdateTask: UpdateTask,
		DeleteTask: DeleteTask,
	}
}
