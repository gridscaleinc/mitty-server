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
	http.Handle("/img/", http.StripPrefix("/img/", imagesHandler))
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
	r.HandleFunc("/contact", controllers.ContactHandler).Methods("POST")
	r.HandleFunc("/userguide", controllers.UserGuideHandler).Methods("GET")
	r.HandleFunc("/email/confirm", controllers.EmailConfirmHandler).Methods("GET")
	r.Handle("/admin", basicAuth(controllers.AdminIndexHandler)).Methods("GET")
	r.Handle("/admin/users", basicAuth(controllers.AdminUsersHandler)).Methods("GET")
	r.Handle("/admin/events", basicAuth(controllers.AdminEventsHandler)).Methods("GET")
	r.Handle("/talk", basicAuth(controllers.TalkIndexHandler)).Methods("GET")
}

func publicRoutes(r *mux.Router) {
	// Configure websocket route
	r.HandleFunc("/status", controllers.StatusHandler).Methods("GET")
	r.HandleFunc("/signup", controllers.SignUpHandler).Methods("POST")
	r.HandleFunc("/signin", controllers.SignInHandler).Methods("POST")

	r.HandleFunc("/reset_password/send", controllers.ResetPasswordSendHandler).Methods("POST")
	r.HandleFunc("/reset_password/verify", controllers.ResetPasswordVerifyHandler).Methods("GET")
	r.HandleFunc("/reset_password/reset", controllers.ResetPasswordResetHandler).Methods("POST")

	//
	// INSERT
	r.Handle("/new/event", apiAuth(controllers.PostEventHandler)).Methods("POST")
	r.Handle("/new/activity", apiAuth(controllers.PostActivityHandler)).Methods("POST")
	r.Handle("/new/activity/item", apiAuth(controllers.PostActivityItemHandler)).Methods("POST")
	r.Handle("/new/island", apiAuth(controllers.PostIslandHandler)).Methods("POST")
	r.Handle("/new/request", apiAuth(controllers.PostRequestHandler)).Methods("POST")
	r.Handle("/save/namecard", apiAuth(controllers.PostNameCardHandler)).Methods("POST")
	r.Handle("/save/profile", apiAuth(controllers.PostProfileHandler)).Methods("POST")
	r.Handle("/new/proposal", apiAuth(controllers.PostProposalHandler)).Methods("POST")
	r.Handle("/gallery/content", apiAuth(controllers.PostGalleryContentHandler)).Methods("POST")
	r.Handle("/upload/content", apiAuth(controllers.UploadContentsHandler)).Methods("POST")
	r.Handle("/send/like", apiAuth(controllers.SendLikeHandler)).Methods("POST")
	r.Handle("/send/offers", apiAuth(controllers.PostOfferHandler)).Methods("POST")
	r.Handle("/send/invitation", apiAuth(controllers.SendInvitationsHandler)).Methods("POST")

	// UPDATE

	r.Handle("/accept/proposal", apiAuth(controllers.PostAcceptProposalHandler)).Methods("POST")
	r.Handle("/approve/proposal", apiAuth(controllers.PostApproveProposalHandler)).Methods("POST")
	r.Handle("/accept/offers", apiAuth(controllers.AcceptOffersHandler)).Methods("POST")

	r.Handle("/checkin", apiAuth(controllers.PostCheckinHandler)).Methods("POST")

	r.Handle("/update/user/icon", apiAuth(controllers.UpdateUserIconHandler)).Methods("POST")
	r.Handle("/update/activity", apiAuth(controllers.UpdateActivityHandler)).Methods("POST")
	r.Handle("/update/activity/item", apiAuth(controllers.UpdateActivityItemHandler)).Methods("POST")

	// SELECT
	// A
	r.Handle("/activity/details", apiAuth(controllers.GetActivityDetailHandler)).Methods("GET")
	r.Handle("/activity/list", apiAuth(controllers.GetActivityListHandler)).Methods("GET")

	// B
	// C
	r.Handle("/contact/list", apiAuth(controllers.GetContateeListHandler)).Methods("GET")
	r.Handle("/contactee/namecards", apiAuth(controllers.GetContacteeNamecardsHandler)).Methods("GET")

	// D
	r.Handle("/destination/list", apiAuth(controllers.GetDestinationListHandler)).Methods("GET")

	// E
	r.Handle("/event/of", apiAuth(controllers.EventFetchingHandler)).Methods("GET")
	r.Handle("/event/meeting", apiAuth(controllers.GetEventMeeting)).Methods("GET")

	// F
	// G
	r.HandleFunc("/gallery/contents", controllers.GetGalleryContentsHandler).Methods("GET")

	// H
	// I
	r.HandleFunc("/island/info", controllers.GetIslandHandler).Methods("GET")

	// J
	// K
	// L
	r.HandleFunc("/latest/conversation", controllers.GetLatestConversation).Methods("GET")

	// M
	r.Handle("/mycontents/list", apiAuth(controllers.GetMyContentsHandler)).Methods("GET")
	r.Handle("/myrequest", apiAuth(controllers.GetMyRequestHandler)).Methods("GET")
	r.Handle("/myprofile", apiAuth(controllers.GetMyProfileHandler)).Methods("GET")
	r.Handle("/mynamecards", apiAuth(controllers.GetMyNamecardsHandler)).Methods("GET")
	r.Handle("/myoffers", apiAuth(controllers.GetOfferListHandler)).Methods("GET")
	r.Handle("/myinvitation/status", apiAuth(controllers.GetMyInvitationsHandler)).Methods("GET")

	// N
	r.Handle("/namecard/of", apiAuth(controllers.GetNamecardHandler)).Methods("GET")

	// O
	// P
	r.HandleFunc("/proposals/of", controllers.GetProposalsHandler).Methods("GET")

	// Q
	// R
	r.HandleFunc("/request/details", controllers.GetRequestDetailsHandler).Methods("GET")

	// S
	r.HandleFunc("/search/event", controllers.SearchEventHandler).Methods("GET")
	r.Handle("/search/request", apiAuth(controllers.GetSearchRequestHandler)).Methods("GET")
	r.Handle("/social/mirror", apiAuth(controllers.GetSocialMirrorHandler)).Methods("GET")

	// T
	// U
	r.HandleFunc("/user/info", controllers.GetUserInfo).Methods("GET")
	r.Handle("/user/profile", apiAuth(controllers.GetUserProfileHandler)).Methods("GET")

	// V
	// W
	// X
	// Y
	// Z

	// DELETE
	r.Handle("/remove/like", apiAuth(controllers.RemoveLikeHandler)).Methods("POST")
}

func basicAuth(handler http.HandlerFunc) http.Handler {
	return filters.BasicAuthHandler(handler)
}

func apiAuth(handler http.HandlerFunc) http.Handler {
	return filters.APIAuthHandler(handler)
}
