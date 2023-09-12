module example.com/task_app

go 1.21.0

require (
	constants v0.0.0-00010101000000-000000000000
	example.com/declarations v0.0.0-00010101000000-000000000000
)

require github.com/gorilla/mux v1.8.0

replace example.com/declarations => ../declarations

replace constants => ../constants

replace app_router => ./router
