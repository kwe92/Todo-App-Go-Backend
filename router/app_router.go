package app_router

import (
	"constants"
	"log"
	"net/http"
	"routehandlers"

	types "example.com/declarations"
	"github.com/gorilla/mux"
)

const (
	Get     = "GET"
	Post    = "POST"
	Put     = "PUT"
	Delete  = "DELETE"
	Address = ":8082"
)

// endpoints represent endpoint constants in a struct.
var endpoints = constants.Endpoints()

// handleRoutes handles all the routes for this API.
func HandleRoutes(tasks *map[string]types.Task) {

	// router instance
	router := mux.NewRouter()

	// all API endpoints
	router.HandleFunc(endpoints.Home, routehandlers.HomePage(tasks)).Methods(Get)

	router.HandleFunc(endpoints.GetTasks, routehandlers.GetTasks(tasks)).Methods(Get)

	router.HandleFunc(endpoints.GetTask, routehandlers.GetTask(tasks)).Methods(Get)

	router.HandleFunc(endpoints.CreateTask, routehandlers.CreateTask(tasks)).Methods(Post)

	router.HandleFunc(endpoints.UpdateTask, routehandlers.UpdateTask(tasks)).Methods(Put)

	router.HandleFunc(endpoints.DeleteTask, routehandlers.DeleteTask(tasks)).Methods(Delete)

	log.Fatal(http.ListenAndServe(Address, router))
}
