package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// Conversation Parameters ...
type ConversationParam struct {
	MeetingID      int64  `json:"meetingId"`
	ReplyToID      int64  `json:"replyToId"`
	Speaking        string `json:"speaking"`
}

// FieldMap defines parameter requirements
func (p *ConversationParam) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.MeetingID: binding.Field{
			Form:     "meetingId",
			Required: true,
		},
		&p.ReplyToID: binding.Field{
			Form:     "startId",
			Required: false,
		},
		&p.Speaking: binding.Field{
			Form:     "speaking",
			Required: true,
		},
	}
}

// Conversation Parameters ...
type ConvFetchParams struct {
	MeetingID      int64  `json:"meetingId"`
	StartID           int64  `json:"startId"`
	Direction        int `json:"direction"`
	NumberLimit  int `json:"numberLimit"`
}

// FieldMap defines parameter requirements
func (p *ConvFetchParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.MeetingID: binding.Field{
			Form:     "meetingId",
			Required: true,
		},
		&p.StartID: binding.Field{
			Form:     "startId",
			Required: false,
		},
		&p.Direction: binding.Field{
			Form:     "direction",
			Required: false,
		},
	}
}

// FetchingConversation ...
func FetchConversationHandler(w http.ResponseWriter, r *http.Request) {
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
	
	p := new(ConvFetchParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputError(w, r, errs)
		return
	}

	conversations, err := models.GetLatestConversation(tx, 90)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"conversations": conversations,
	})

}
