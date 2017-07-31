package controllers

import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
	"net/http"
    "time"
    
// 	goutils "github.com/dongri/goutils"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
// 	"mitty.co/mitty-server/config"
)

type	InvitationParams struct {
    ForType    string `json: "forType"`
    IDOfType  int64 `json: "idOfType"`
    Message    string  `json: "message"`
    Invitees   []int  `json: "invitees"`
} 
	
// FieldMap ...
func (s *InvitationParams) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.ForType: binding.Field{
			Form:     "forType",
			Required: true,
		},
		&s.IDOfType: binding.Field{
			Form:     "idOfType",
			Required: true,
		},
		&s.Message: binding.Field{
			Form:     "message",
			Required: true,
		},
		&s.Invitees: binding.Field{
			Form:     "invitees",
			Required: true,
		},
	}
}

// SendInvitationsHandler ... to be done....
func SendInvitationsHandler(w http.ResponseWriter, r *http.Request) {
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
	
	p := new(InvitationParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputError(w, r, errs)
		return
	}

    invitation := new(models.Invitation)
    invitation.InvitaterID = currentUserID
    invitation.ForType = p.ForType
    invitation.IDOfType = p.IDOfType
    invitation.Message = p.Message
    invitation.TimeOfInvitation = time.Now().UTC()
   
	if err := invitation.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}
	
	// insert into isvitees for every invitee that invited.
    for index, inviteeID := range invitation.invitees {
         invitee := new(models.Invitee)
         invitee.InvitationID = invitation.ID
         invitee.inviteeId = inviteeID
         invitee.ReplyStatus = "NONE"
         
         if err := invitee.Insert(*tx); err != nil {
		     filters.RenderError(w, r, err)
		     return
		 }
    }
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": invitation.ID,
	})
}

