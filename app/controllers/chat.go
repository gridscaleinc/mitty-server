package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
)

// TalkIndexHandler ...
func TalkIndexHandler(w http.ResponseWriter, r *http.Request) {
	filters.RenderHTML(w, r, "talk/talk.html", nil)
}
