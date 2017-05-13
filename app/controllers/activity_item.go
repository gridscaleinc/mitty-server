package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// ActivityItemParams ...
type ActivityItemParams struct {
	ActivityID           int       `json:"activityId"`
	EventID              int64     `json:"eventId"`
	Title                string    `json:"title"`
	Memo                 string    `json:"memo"`
	Notification         bool      `json:"notification"`
	NotificationDateTime time.Time `json:"notificationDateTime"`
	AsMainEvent          bool      `json:"asMainEvent"`
}

// FieldMap defines parameter requirements
func (p *ActivityItemParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
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
		helpers.RenderInputError(w, r, errs)
		return
	}

	activityItem := new(models.ActivityItem)
	activityItem.ActivityID = p.ActivityID
	activityItem.EventID = p.EventID
	activityItem.Title = p.Title
	activityItem.Memo = p.Memo
	activityItem.Notification = p.Notification
	activityItem.NotificationDateTime = p.NotificationDateTime
	if err := activityItem.Insert(*tx); err != nil {
		helpers.RenderDBError(w, r, err)
		return
	}

	if p.AsMainEvent == true {
		activity, err := models.GetActivityByID(tx, p.ActivityID)
		if err != nil && err != sql.ErrNoRows {
			helpers.RenderDBError(w, r, err)
			return
		}
		activity.MainEventID = sql.NullInt64{Int64:activityItem.EventID, Valid:true}
		if err := activity.Update(*tx); err != nil {
			helpers.RenderDBError(w, r, err)
			return
		}
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{})
}
