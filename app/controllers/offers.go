package controllers

import (
	// 	"database/sql"
	// 	"errors"
	// 	"fmt"
	"errors"
	"net/http"
	"strings"
	"time"

	// 	goutils "github.com/dongri/goutils"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	// 	"mitty.co/mitty-server/config"
)

// OffersForm ...
type OffersForm struct {
	ID          int64  `json:"id"`
	ToMittyID   int    `json:"to_mitty_id"`
	Type        string `json:"type"`
	Message     string `json:"message"`
	ReplyStatus string `json:"reply_status"`
	OfferredID  int64  `json:"offerred_id"`
	RepliedID   int64  `json:"replied_id"`
}

// FieldMap ...
func (s *OffersForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.ID: "id",
		&s.ToMittyID: binding.Field{
			Form:     "to_mitty_id",
			Required: true,
		},
		&s.Type: binding.Field{
			Form:     "type",
			Required: true,
		},
		&s.Message: binding.Field{
			Form:     "message",
			Required: true,
		},
		&s.ReplyStatus: binding.Field{
			Form:     "reply_status",
			Required: true,
		},
		&s.OfferredID: binding.Field{
			Form:     "offerred_id",
			Required: true,
		},
		&s.RepliedID: binding.Field{
			Form:     "replied_id",
			Required: false,
		},
	}
}

// Validate ...
func (s *OffersForm) Validate(req *http.Request) error {
	if len(s.Type) > 50 {
		return errors.New("type is too long")
	}
	if len(s.ReplyStatus) > 50 {
		return errors.New("reply status is too long")
	}
	return nil
}

// PostOfferHandler ...
func PostOfferHandler(w http.ResponseWriter, r *http.Request) {
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

	p := new(OffersForm)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	offers := new(models.Offers)
	offers.FromMittyID = currentUserID
	offers.ToMittyID = p.ToMittyID
	offers.Type = p.Type
	offers.Message = p.Message
	offers.OfferredID = p.OfferredID
	offers.RepliedID = p.RepliedID
	offers.ReplyStatus = "NONE"

	if err := offers.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": offers.ID,
	})
}

// AcceptOffersHandler ...
func AcceptOffersHandler(w http.ResponseWriter, r *http.Request) {
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

	// todo : Maybe should check the id, only the targeted mitty id can accept
	// currentUserID := filters.GetCurrentUserID(r)

	p := new(OffersForm)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	offer := new(models.Offers)
	offer.ID = p.ID

	err = offer.Load(*tx)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	offer.ReplyStatus = p.ReplyStatus
	err = offer.Update(*tx)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	if strings.EqualFold(offer.Type, "NAMECARD") && strings.EqualFold(offer.ReplyStatus, "ACCEPTED") {
		meetingID := int64(0)
		existA, err := models.ExistContactFromIDs(*tx, offer.FromMittyID, offer.RepliedID)
		if err != nil {
			filters.RenderError(w, r, err)
			return
		}

		existB, err := models.ExistContactFromIDs(*tx, offer.ToMittyID, offer.OfferredID)
		if err != nil {
			filters.RenderError(w, r, err)
			return
		}

		if existA != nil {
			meetingID = existA.MeetingID
		} else if existB != nil {
			meetingID = existB.MeetingID
		}

		if meetingID == 0 {
			meeting := new(models.Meeting)
			meeting.Name = "Contact"
			meeting.Type = "CONTACT"
			err = meeting.Insert(*tx)
			if err != nil {
				filters.RenderError(w, r, err)
				return
			}
			meetingID = meeting.ID
		}

		if existA == nil {
			contactA := new(models.Contact)
			contactA.MittyID = offer.FromMittyID
			contactA.NameCardID = offer.RepliedID
			contactA.ContctedDate = time.Now().UTC()
			contactA.MeetingID = meetingID
			err = contactA.Insert(*tx)
			if err != nil {
				filters.RenderError(w, r, err)
				return
			}
		}

		if existB == nil {
			contactB := new(models.Contact)
			contactB.MittyID = offer.ToMittyID
			contactB.NameCardID = offer.OfferredID
			contactB.ContctedDate = time.Now().UTC()
			contactB.MeetingID = meetingID
			err = contactB.Insert(*tx)
			if err != nil {
				filters.RenderError(w, r, err)
				return
			}
		}
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}

// GetOfferListHandler ...
func GetOfferListHandler(w http.ResponseWriter, r *http.Request) {
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

	offerList, err := models.GetOfferListByUserID(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"offerList": offerList,
	})
}
