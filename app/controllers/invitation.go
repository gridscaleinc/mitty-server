package controllers

import (
	// 	"database/sql"
	// 	"errors"
	// 	"fmt"
	"errors"
	"net/http"
	"strconv"
	"time"

	// 	goutils "github.com/dongri/goutils"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	// 	"mitty.co/mitty-server/config"
)

// InvitationParams ...
type InvitationParams struct {
	ForType  string `json:"forType"`
	IDOfType int64  `json:"idOfType"`
	Message  string `json:"message"`
	Invitees []int  `json:"invitees"`
}

// AcceptInvitationParams ...
type AcceptInvitationParams struct {
	InvitationID int64  `json:"invitation_id"`
	InviteesID   int64  `json:"invitees_id"`
	ReplyStatus  string `json:"reply_status"`
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
		"invitees": binding.Field{
			Binder: func(fieldName string, formVals []string, errs binding.Errors) binding.Errors {
				values := make([]int, len(formVals))

				for _, v := range formVals {
					val, err := strconv.Atoi(v)
					if err != nil {
						// errs.Add(["invitees"])
						errs.Add([]string{"invitees"}, "conversion", "not an integer")
						return errs
					}
					values = append(values, val)
				}

				s.Invitees = values
				return nil
			},
			Form:     "invitees",
			Required: true,
		},
	}
}

// Validate ...
func (s *InvitationParams) Validate(req *http.Request) error {
	if len(s.Message) > 1000 {
		return errors.New("Message is too long")
	}
	return nil
}

// FieldMap ...
func (s *AcceptInvitationParams) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.InvitationID: binding.Field{
			Form:     "invitation_id",
			Required: true,
		},
		&s.InviteesID: binding.Field{
			Form:     "invitees_id",
			Required: true,
		},
		&s.ReplyStatus: binding.Field{
			Form:     "reply_status",
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
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
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
	for _, inviteeID := range p.Invitees {
		invitee := new(models.Invitees)
		invitee.InvitationID = invitation.ID
		invitee.InviteeID = inviteeID
		invitee.ReplyStatus = models.None

		if err := invitee.Insert(*tx); err != nil {
			filters.RenderError(w, r, err)
			return
		}
	}
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": invitation.ID,
	})
}

// GetMyInvitationsHandler ...
func GetMyInvitationsHandler(w http.ResponseWriter, r *http.Request) {
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

	invitationStatus, err := models.GetInvitationStatusByUserID(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"invitationStatus": invitationStatus,
	})
}

// AcceptInvitationHandler ...
func AcceptInvitationHandler(w http.ResponseWriter, r *http.Request) {
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

	p := new(AcceptInvitationParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	invitee, err := models.GetInviteeForAccept(tx, p.InvitationID, p.InviteesID, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	invitation, err := models.GetInvitationByID(tx, p.InvitationID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	event, err := models.GetEventDetailByID(tx, currentUserID, invitation.IDOfType)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	// update invitees Status
	invitee.ReplyStatus = p.ReplyStatus
	invitee.ReplyTime = time.Now().UTC()

	err = invitee.Update(*tx)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	// すでに参加した場合、活動登録を行わない。
	// Rejectの場合も登録しない。
	if event.ParticipationStatus != "NOT" || p.ReplyStatus == "REJECTED" {
		render.JSON(w, http.StatusOK, map[string]interface{}{
			"ok": true,
		})
		return
	}

	// Acceptの場合、Eventの参加をする。
	activity := new(models.Activity)
	activity.MainEventID = event.ID
	activity.Title = event.Title
	activity.Memo = event.Action
	activity.OwnerID = currentUserID
	err = activity.Insert(*tx)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	activityItem := new(models.ActivityItem)
	activityItem.ActivityID = activity.ID
	activityItem.Title = event.Title
	activityItem.EventID = event.ID
	activityItem.Memo = event.Action

	err = activityItem.Insert(*tx)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}
