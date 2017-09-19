package controllers

import (
	"net/http"

	"mitty.co/mitty-server/config"

	"mitty.co/mitty-server/app/filters"
)

// StatusHandler PATH: GET /status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	render.JSON(w, http.StatusOK, map[string]string{
		"name":    config.AppName,
		"version": config.ServerVersion,
		"env":     config.CurrentEnv,
	})
}

// NotFoundHandler return 404 NotFound
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/userguide", http.StatusFound)
	//http.Error(w, "404 NotFound", http.StatusNotFound)
}
