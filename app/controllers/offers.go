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
		filters.RenderInputError(w, r, errs)
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
		filters.RenderInputError(w, r, errs)
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

	if offer.Type == "NAMECARD" && offer.ReplyStatus == "ACCEPTED" {
		contact := new(models.Contact)
		contact.MittyID = offer.FromMittyID
		contact.NameCardID = offer.RepliedID
		contact.ContctedDate = time.Now().UTC()
		err = contact.Insert(*tx)
		if err != nil {
			filters.RenderError(w, r, err)
			return
		}

		contact = new(models.Contact)
		contact.MittyID = offer.ToMittyID
		contact.NameCardID = offer.OfferredID
		contact.ContctedDate = time.Now().UTC()
		err = contact.Insert(*tx)
		if err != nil {
			filters.RenderError(w, r, err)
			return
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
