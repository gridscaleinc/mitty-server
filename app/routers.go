package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"mitty.co/mitty-server/app/controllers"
	"mitty.co/mitty-server/app/filters"
)

// BuildRouter creates and returns a router which hold whole handler functions.
func BuildRouter() http.Handler {
	appRouter := mux.NewRouter()

	webRouter := appRouter.PathPrefix("/").Subrouter()
	webRoutes(webRouter)

	publicRouter := appRouter.PathPrefix("/api/").Subrouter()
	publicRoutes(publicRouter)

	appRouter.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)

	return appRouter
}

func webRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.WelcomeHandler).Methods("GET")
	r.HandleFunc("/email/confirm", controllers.EmailConfirmHandler).Methods("GET")
	r.Handle("/admin", basicAuth(controllers.AdminIndexHandler)).Methods("GET")
	r.Handle("/admin/users", basicAuth(controllers.AdminUsersHandler)).Methods("GET")
	r.Handle("/admin/events", basicAuth(controllers.AdminEventsHandler)).Methods("GET")
}

func publicRoutes(r *mux.Router) {
	r.HandleFunc("/status", controllers.StatusHandler).Methods("GET")
	r.HandleFunc("/signup", controllers.SignUpHandler).Methods("POST")
	r.HandleFunc("/signin", controllers.SignInHandler).Methods("POST")
	r.HandleFunc("/new/event", controllers.PostEventHandler).Methods("POST")
	r.HandleFunc("/gallery/content", controllers.PostGalleryContentHandler).Methods("POST")
	r.HandleFunc("/search/event", controllers.SearchEventHandler).Methods("GET")
}

func basicAuth(handler http.HandlerFunc) http.Handler {
	return filters.BasicAuthHandler(handler)
}
