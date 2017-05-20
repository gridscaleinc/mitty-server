package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

// GetMyContentsHandler ...
func GetMyContentsHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	tx, err := dbmap.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	userID := 0
	contents, err := models.GetContentsByUserID(tx, userID)
	if err != nil {
		helpers.RenderDBError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"contents": contents,
	})
}
