package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// MeetingParams ...
type MeetingParams struct {
	MeetingID int64 `json:"meetingId"`
}

// FieldMap ... Mapping input value to MeetingParams
func (p *MeetingParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.MeetingID: binding.Field{
			Form:     "meetingId",
			Required: true,
		},
	}
}

// GetEventMeeting ...
func GetEventMeeting(w http.ResponseWriter, r *http.Request) {
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
	meetingList, err := models.GetEventMeetingList(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"eventMeetingList": meetingList,
	})
}

// GetRequestMeeting ...
func GetRequestMeeting(w http.ResponseWriter, r *http.Request) {
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
	meetingList, err := models.GetRequestMeetingList(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"requestMeetingList": meetingList,
	})
}

// GetContactMeeting ...
func GetContactMeeting(w http.ResponseWriter, r *http.Request) {
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
	meetingList, err := models.GetContactMeetingList(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"contactMeetingList": meetingList,
	})
}

// GetLatestConversation ...
func GetLatestConversation(w http.ResponseWriter, r *http.Request) {
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

	p := new(MeetingParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	talks, err := models.GetLatestConversation(tx, p.MeetingID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	count := len(talks)
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"count":         count,
		"conversations": talks,
	})
}
