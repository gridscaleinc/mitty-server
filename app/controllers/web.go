package controllers

import (
	"fmt"
	"net/http"
)

// WelcomeHandler ...
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h2>welcome mitty!</h2>")
}
