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

// Package gorilla/mux

//   - Implements a request router and dispatcher that matches incoming requests to their respective route handlers
//   - the mux.Router is an implementation of the http.Handler interface required by http.ListenAndServe
//   - the mux.Router instance matches a request to a list of registered routes,
//     if a match is found for the URL the associated handler is called

// HandleFunc | mux package

//   - registers a new route, mapping the URL path to the route handler
//   - if an incoming request URL matches a registered path the associated route handler is called
//   - three things you need:
//       - route name (endpoint name / path)
//       - route handler callback
//       - route method (using dot notation)
//   - "/" path represents the home page
//   - similar to larvel or node.js

// Matching Routes

//   - there are serveral matchers that can be added to a registered route
//   - the most common are .Methods for HTTP methods and .Headers for HTTP headers
//   - see documentation for more matchers

// Path Variables

//   - paths can have variables defined in the format {name} or {name: pattern} where pattern = regular expression

// http.ListenAndServe

//   - Starts an HTTP server with a given address and handler

// Pinging Local Host

//   - you can ping local host with Postman desktop client

//   ~ you can use localhost or 127.0.0.1 as a prefix to the port you specified

//       + e.g. localhost:8082/gettasks or 127.0.0.1:8082/gettasks

// What Happens when :router.HandleFunc("/create", createTask).Methods("POST") is called?

//   - the createTask callback is passed a request object from the caller
//     that contains the r.Body of the Post request
