package app_router

import (
	"constants"
	"log"
	"net/http"
	"routehandlers"

	types "example.com/declarations"
	"github.com/gorilla/mux"
)

// endpoints represent endpoint constants in a struct.
var endpoints = constants.Endpoints()

// handleRoutes handles all the routes for this API.
func HandleRoutes(tasks *[]types.Task) {

	// router instance
	router := mux.NewRouter()

	// all API endpoints
	router.HandleFunc(endpoints.Home, routehandlers.HomePage(tasks)).Methods("GET")

	router.HandleFunc(endpoints.GetTasks, routehandlers.GetTasks(tasks)).Methods("GET")

	router.HandleFunc(endpoints.GetTask, routehandlers.GetTask(tasks)).Methods("GET")

	router.HandleFunc(endpoints.CreateTask, routehandlers.CreateTask(tasks)).Methods("POST")

	router.HandleFunc(endpoints.UpdateTask, routehandlers.UpdateTask(tasks)).Methods("PUT")

	router.HandleFunc(endpoints.DeleteTask, routehandlers.DeleteTask(tasks)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", router))
}
