package constants

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

func AppConstants() (HeaderMetaData, EndPoints) {
	headerMetaData := HeaderMetaData{
		ContentTypeHeader: "Content-Type",
		MediaTypeJson:     "application/json",
	}

	endPoints := EndPoints{
		GetTask:    "/gettask/{id}",
		GetTasks:   "/gettasks",
		CreateTask: "/create",
		UpdateTask: "/update/{id}",
		DeleteTask: "/delete/{id}",
	}

	return headerMetaData, endPoints
}
