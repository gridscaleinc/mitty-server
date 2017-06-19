package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// EventParams ...
type MeetingParams struct {
	MeetingID int64 `json:"meetingId"`
}

func (p *MeetingParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.MeetingID: binding.Field{
			Form:     "meetingId",
			Required: true,
		},
	}
}

// FetchingConversation ...
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
		filters.RenderInputError(w, r, errs)
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
