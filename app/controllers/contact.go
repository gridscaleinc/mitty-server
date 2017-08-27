package controllers

import (
	// 	"database/sql"
	// 	"errors"
	// 	"fmt"
	"net/http"

	// 	goutils "github.com/dongri/goutils"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	// 	"mitty.co/mitty-server/config"
)

// GetContateeListHandler ...
func GetContateeListHandler(w http.ResponseWriter, r *http.Request) {
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

	currentUserID := filters.GetCurrentUserID(r)

	contacteeList, err := models.GetContacteeListByUserID(tx, currentUserID)

	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"contacteeList": contacteeList,
	})
}
