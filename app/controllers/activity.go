package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// GetActivityHandler ...
func GetActivityHandler(w http.ResponseWriter, r *http.Request) {
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

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"soon": "TODO",
	})
}

// ActivityParams ...
type ActivityParams struct {
	Title       string `json:"title"`
	MainEventID int    `json:"mainEventId"`
	Memo        string `json:"memo"`
}

// FieldMap defines parameter requirements
func (p *ActivityParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Title: binding.Field{
			Form:     "title",
			Required: true,
		},
		&p.MainEventID: binding.Field{
			Form:     "mainEventId",
			Required: false,
		},
		&p.Memo: binding.Field{
			Form:     "memo",
			Required: false,
		},
	}
}

// PostActivityHandler ...
func PostActivityHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(ActivityParams)
	if errs := binding.Bind(r, p); errs != nil {
		helpers.RenderInputError(w, r, errs)
		return
	}

	activity := new(models.Activity)
	activity.Title = p.Title
	activity.MainEventID = p.MainEventID
	activity.Memo = p.Memo
	if err := activity.Insert(*tx); err != nil {
		helpers.RenderDBError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"activityId": activity.ID,
	})
}
