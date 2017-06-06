package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"mitty.co/mitty-server/app/controllers"
	"mitty.co/mitty-server/app/filters"
    //"mitty.co/mitty-server/app/talk"
)

// BuildRouter creates and returns a router which hold whole handler functions.
func BuildRouter() http.Handler {
	cssHandler := http.FileServer(http.Dir("./public/css/"))
	fontsHandler := http.FileServer(http.Dir("./public/fonts/"))
	imagesHandler := http.FileServer(http.Dir("./public/img/"))
	javascriptHandler := http.FileServer(http.Dir("./public/js/"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", fontsHandler))
	http.Handle("/img/", http.StripPrefix("/images/", imagesHandler))
	http.Handle("/js/", http.StripPrefix("/js/", javascriptHandler))

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
	r.Handle("/talk", basicAuth(controllers.TalkIndexHandler)).Methods("GET")
}

func publicRoutes(r *mux.Router) {
		// Configure websocket route
	// r.Handle("/ws/", apiAuth(talk.WebsocketHandler))
	r.HandleFunc("/status", controllers.StatusHandler).Methods("GET")
	r.HandleFunc("/signup", controllers.SignUpHandler).Methods("POST")
	r.HandleFunc("/signin", controllers.SignInHandler).Methods("POST")
	r.HandleFunc("/new/event", controllers.PostEventHandler).Methods("POST")
	r.Handle("/gallery/content", apiAuth(controllers.PostGalleryContentHandler)).Methods("POST")
	r.HandleFunc("/search/event", controllers.SearchEventHandler).Methods("GET")
	r.HandleFunc("/new/activity", controllers.PostActivityHandler).Methods("POST")
	r.HandleFunc("/new/activity/item", controllers.PostActivityItemHandler).Methods("POST")
	r.HandleFunc("/new/island", controllers.PostIslandHandler).Methods("POST")
	r.HandleFunc("/event/of", controllers.EventFetchingHandler).Methods("GET")
	r.HandleFunc("/activity/list", controllers.GetActivityListHandler).Methods("GET")
	r.HandleFunc("/activity/details", controllers.GetActivityDetailHandler).Methods("GET")
	r.HandleFunc("/island/info", controllers.GetIslandHandler).Methods("GET")
	r.HandleFunc("/mycontents/list", controllers.GetMyContentsHandler).Methods("GET")
	r.Handle("/upload/content", apiAuth(controllers.GetMyContentsHandler)).Methods("POST")
	r.HandleFunc("/event/meeting", controllers.GetEventMeeting).Methods("GET")
}

func basicAuth(handler http.HandlerFunc) http.Handler {
	return filters.BasicAuthHandler(handler)
}

func apiAuth(handler http.HandlerFunc) http.Handler {
	return filters.APIAuthHandler(handler)
}
