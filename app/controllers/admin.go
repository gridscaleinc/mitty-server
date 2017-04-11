package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

// AdminUsersHandler ...
func AdminUsersHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	users, err := models.GetAdminUsers(dbmap)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	output := map[string]interface{}{
		"users": users,
	}
	filters.RenderHTML(w, r, "app/views/admin/users.html", output)
}
