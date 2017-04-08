package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
)

// WelcomeHandler ...
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{
		"title": "welcome mitty",
	}
	filters.RenderHTML(w, r, "app/views/index.html", output)
}
