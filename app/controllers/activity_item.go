package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// ActivityItemParams ...
type ActivityItemParams struct {
	ID                   int64     `json:"id"`
	ActivityID           int64     `json:"activityId"`
	EventID              int64     `json:"eventId"`
	Title                string    `json:"title"`
	Memo                 string    `json:"memo"`
	Notification         string    `json:"notification"`
	NotificationDateTime time.Time `json:"notificationDateTime"`
	AsMainEvent          bool      `json:"asMainEvent"`
}

// FieldMap defines parameter requirements
func (p *ActivityItemParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.ID: "id",
		&p.ActivityID: binding.Field{
			Form:     "activityId",
			Required: true,
		},
		&p.EventID: binding.Field{
			Form:     "eventId",
			Required: true,
		},
		&p.Title: binding.Field{
			Form:     "title",
			Required: true,
		},
		&p.Memo: binding.Field{
			Form:     "memo",
			Required: false,
		},
		&p.Notification: binding.Field{
			Form:     "notification",
			Required: true,
		},
		&p.NotificationDateTime: binding.Field{
			Form:     "notificationDateTime",
			Required: false,
		},
		&p.AsMainEvent: binding.Field{
			Form:     "asMainEvent",
			Required: false,
		},
	}
}

// Validate ...
func (p *ActivityItemParams) Validate(req *http.Request) error {
	if len(p.Title) > 200 {
		return errors.New("title is too long")
	}
	return nil
}

// PostActivityItemHandler ...
func PostActivityItemHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(ActivityItemParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	activityItem := new(models.ActivityItem)
	activityItem.ActivityID = p.ActivityID
	activityItem.EventID = p.EventID
	activityItem.Title = p.Title
	activityItem.Memo = p.Memo
	if p.Notification == "true" {
		activityItem.Notification = true
	} else {
		activityItem.Notification = false
	}
	activityItem.NotificationDateTime = p.NotificationDateTime
	if err := activityItem.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	if p.AsMainEvent == true {
		activity, err := models.GetActivityByID(tx, p.ActivityID)
		if err != nil && err != sql.ErrNoRows {
			filters.RenderError(w, r, err)
			return
		}
		activity.MainEventID = activityItem.EventID
		if err := activity.Update(*tx); err != nil {
			filters.RenderError(w, r, err)
			return
		}
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{})
}

// UpdateActivityItemHandler ...
func UpdateActivityItemHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(ActivityItemParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	activityItem := new(models.ActivityItem)
	activityItem.ID = p.ID
	if err := activityItem.Load(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}
	activityItem.Title = p.Title
	activityItem.Memo = p.Memo
	if p.Notification == "true" {
		activityItem.Notification = true
		activityItem.NotificationDateTime = p.NotificationDateTime
	} else {
		activityItem.Notification = false
	}

	if err := activityItem.Update(*tx); err != nil {
		render.JSON(w, http.StatusServiceUnavailable, map[string]interface{}{
			"parameters": p,
			"error":      err,
		})

		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id":         activityItem.ID,
		"activityId": activityItem.ActivityID,
	})
}
