package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"mitty.co/mitty-server/app/controllers"
)

// BuildRouter creates and returns a router which hold whole handler functions.
func BuildRouter() http.Handler {
	appRouter := mux.NewRouter()

	publicRouter := appRouter.PathPrefix("/api/").Subrouter()
	publicRoutes(publicRouter)

	appRouter.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)

	return appRouter
}

func publicRoutes(r *mux.Router) {
	r.HandleFunc("/status", controllers.StatusHandler).Methods("GET")
}
