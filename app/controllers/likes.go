package controllers

import (
	// 	"database/sql"
	// 	"errors"
	// 	"fmt"
	"net/http"

	// 	goutils "github.com/dongri/goutils"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	// 	"mitty.co/mitty-server/config"
)

// LikesForm ...
type LikesForm struct {
	Type string
	ID   int64
}

// FieldMap ...
func (s *LikesForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.Type: binding.Field{
			Form:     "type",
			Required: true,
		},
		&s.ID: binding.Field{
			Form:     "id",
			Required: true,
		},
	}
}

// SendLikeHandler ...
func SendLikeHandler(w http.ResponseWriter, r *http.Request) {
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

	p := new(LikesForm)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputError(w, r, errs)
		return
	}

	likes := new(models.Likes)
	likes.MittyID = currentUserID
	likes.EntityType = p.Type
	likes.EntityID = p.ID

	if err := likes.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"ok": true,
	})
}

// RemoveLikeHandler ...
func RemoveLikeHandler(w http.ResponseWriter, r *http.Request) {
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

	p := new(LikesForm)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputError(w, r, errs)
		return
	}

	entityType := p.Type
	entityID := p.ID

	err = models.RemoveLikesByID(tx, currentUserID, entityType, entityID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}
