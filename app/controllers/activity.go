package controllers

import (
	"net/http"
	"strconv"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// GetActivityListHandler ...
func GetActivityListHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	currentUserID := filters.GetCurrentUserID(r)
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

	key := r.URL.Query().Get("key")

	activities, err := models.GetActivityListByKey(tx, currentUserID, key)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	count := len(activities)
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"count":      count,
		"activities": activities,
	})
}

// ActivityParams ...
type ActivityParams struct {
	Title       string `json:"title"`
	MainEventID int64  `json:"mainEventId"`
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
	currentUserID := filters.GetCurrentUserID(r)
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
		filters.RenderInputError(w, r, errs)
		return
	}

	activity := new(models.Activity)
	activity.Title = p.Title
	activity.MainEventID = p.MainEventID
	activity.Memo = p.Memo
	activity.OwnerID = currentUserID
	if err := activity.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}
	
	// Insert Activity Item also if mainEventId was set
	if (p.MainEventID != 0) {
	    activityItem := new(models.ActivityItem)
	    activityItem.ActivityID = activity.ID
	    activityItem.EventID = p.MainEventID
	    activityItem.Participation = "PARTICIPATING"
	    activityItem.Notification = false
	    if err := activityItem.Insert(*tx); err != nil {
		    filters.RenderError(w, r, err)
		    return
	    }
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"activityId": activity.ID,
	})
}

// GetActivityDetailHandler ...
func GetActivityDetailHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	currentUserID := filters.GetCurrentUserID(r)
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

	id := r.URL.Query().Get("id")

	details, err := models.GetActivityDetailsByID(tx, currentUserID, id)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	activity, err := models.GetActivityByID(tx, intID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"activity": activity,
		"details":  details,
	})

}

// GetDestinationListHandler ...
func GetDestinationListHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	currentUserID := filters.GetCurrentUserID(r)
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

	destinations, err := models.GetDestinationList(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"destinations": destinations,
	})

}
