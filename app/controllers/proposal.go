package controllers

import (
	// 	"database/sql"
	// 	"encoding/json"
	// 	"fmt"
	"errors"
	"net/http"
	// 	"reflect"
	"time"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// ProposalParams ...
type ProposalParams struct {
	ReplyToRequestID  int64     `json:"reply_to_request_id"`
	ContactTel        string    `json:"contact_tel"`
	ContactEmail      string    `json:"contact_email"`
	ProposedIslandID  int64     `json:"proposed_island_id"`
	ProposedIslandID2 int64     `json:"proposed_island_id2"`
	GalleryID         int64     `json:"gallery_id"`
	PriceName1        string    `json:"priceName1"`
	Price1            int       `json:"price1"`
	PriceName2        string    `json:"priceName2"`
	Price2            int       `json:"price2"`
	PriceCurrency     string    `json:"price_currency"`
	PriceInfo         string    `json:"price_info"`
	ProposedDatetime1 time.Time `json:"proposed_datetime1"`
	ProposedDatetime2 time.Time `json:"proposed_datetime2"`
	AdditionalInfo    string    `json:"additional_info"`
	ProposerInfo      string    `json:"proposer_info"`
	ConfirmTel        string    `json:"confirm_tel"`
	ConfirmEmail      string    `json:"confirm_email"`
}

// FieldMap defines parameter requirements
func (p *ProposalParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.ReplyToRequestID: binding.Field{
			Form:     "reply_to_request_id",
			Required: true,
		},
		&p.ContactTel: binding.Field{
			Form:     "contact_tel",
			Required: true,
		},
		&p.ContactEmail: binding.Field{
			Form:     "contact_email",
			Required: false,
		},
		&p.ProposedIslandID: binding.Field{
			Form:     "proposed_island_id",
			Required: true,
		},
		&p.ProposedIslandID2: binding.Field{
			Form:     "proposed_island_id2",
			Required: false,
		},
		&p.GalleryID: binding.Field{
			Form:     "gallery_id",
			Required: false,
		},
		&p.PriceName1: binding.Field{
			Form:     "priceName1",
			Required: false,
		},
		&p.Price1: binding.Field{
			Form:     "price1",
			Required: false,
		},
		&p.PriceName2: binding.Field{
			Form:     "priceName2",
			Required: false,
		},
		&p.Price2: binding.Field{
			Form:     "price2",
			Required: false,
		},
		&p.PriceCurrency: binding.Field{
			Form:     "price_currency",
			Required: false,
		},
		&p.PriceInfo: binding.Field{
			Form:     "price_info",
			Required: false,
		},
		&p.ProposedDatetime1: binding.Field{
			Form:     "proposed_datetime1",
			Required: true,
		},
		&p.ProposedDatetime2: binding.Field{
			Form:     "proposed_datetime2",
			Required: true,
		},
		&p.AdditionalInfo: binding.Field{
			Form:     "additional_info",
			Required: true,
		},
		&p.ProposerInfo: binding.Field{
			Form:     "proposer_info",
			Required: false,
		},
		&p.ConfirmTel: binding.Field{
			Form:     "confirm_tel",
			Required: false,
		},
		&p.ConfirmEmail: binding.Field{
			Form:     "confirm_email",
			Required: false,
		},
	}
}

// Validate ...
func (p *ProposalParams) Validate(req *http.Request) error {
	if len(p.ContactTel) > 20 {
		return errors.New("ContactTel is too long")
	}
	if len(p.ContactEmail) > 50 {
		return errors.New("ContactEmail is too long")
	}
	if len(p.PriceName1) > 50 {
		return errors.New("PriceName1 is too long")
	}
	if len(p.PriceName2) > 50 {
		return errors.New("PriceName2 is too long")
	}
	if len(p.PriceCurrency) > 3 {
		return errors.New("PriceCurrency is too long")
	}
	if len(p.PriceInfo) > 1000 {
		return errors.New("PriceInfo is too long")
	}
	if len(p.ProposerInfo) > 1000 {
		return errors.New("ProposerInfo is too long")
	}
	if len(p.ConfirmTel) > 20 {
		return errors.New("ConfirmTel is too long")
	}
	if len(p.ConfirmEmail) > 50 {
		return errors.New("ConfirmEmail is too long")
	}
	return nil
}

// ProposalStatusParams ...
type ProposalStatusParams struct {
	ProposalID   int64  `json:"proposal_id"`
	ConfirmTel   string `json:"confirm_tel"`
	ConfirmEmail string `json:"confirm_email"`
	Status       string `json:"status"`
}

// FieldMap defines parameter requirements
func (p *ProposalStatusParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.ProposalID: binding.Field{
			Form:     "proposal_id",
			Required: true,
		},
		&p.ConfirmTel: binding.Field{
			Form:     "confirm_tel",
			Required: false,
		},
		&p.ConfirmEmail: binding.Field{
			Form:     "confirm_email",
			Required: false,
		},
		&p.Status: binding.Field{
			Form:     "status",
			Required: true,
		},
	}
}

// PostProposalHandler ...
func PostProposalHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	currentUserID := filters.GetCurrentUserID(r)
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
	p := new(ProposalParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	m := new(models.Proposal)

	m.ReplyToRequestID = p.ReplyToRequestID
	m.ContactTel = p.ContactTel
	m.ContactEmail = p.ContactEmail
	m.ProposedIslandID = p.ProposedIslandID
	m.ProposedIslandID2 = p.ProposedIslandID2
	m.GalleryID = p.GalleryID
	m.PriceName1 = p.PriceName1
	m.Price1 = p.Price1
	m.PriceName2 = p.PriceName2
	m.Price2 = p.Price2
	m.PriceCurrency = p.PriceCurrency
	m.PriceInfo = p.PriceInfo
	m.ProposedDatetime1 = p.ProposedDatetime1
	m.ProposedDatetime2 = p.ProposedDatetime2
	m.AdditionalInfo = p.AdditionalInfo
	m.ProposerID = currentUserID
	m.ProposerInfo = p.ProposerInfo
	m.ProposedDatetime = time.Now().UTC()
	m.AcceptStatus = models.None
	m.ConfirmTel = p.ConfirmTel
	m.ConfirmEmail = p.ConfirmEmail
	m.ApprovalStatus = models.None

	if err := m.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": m.ID,
	})
}

// PostAcceptProposalHandler ...
func PostAcceptProposalHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	currentUserID := filters.GetCurrentUserID(r)
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

	p := new(ProposalStatusParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	proposal, err := models.GetProposalByID(*tx, p.ProposalID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	// check status and ownership
	request, err := models.GetRequestDetailByID(tx, proposal.ReplyToRequestID)
	if request.OwnerID != currentUserID {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Only Requester user can accept a proposal!",
		})
		return
	}

	proposal.ConfirmEmail = p.ConfirmEmail
	proposal.ConfirmTel = p.ConfirmTel
	proposal.AcceptStatus = p.Status
	proposal.AcceptDatetime = time.Now().UTC()

	proposal.Update(*tx)

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": proposal.ID,
	})
}

// PostApproveProposalHandler ...
func PostApproveProposalHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	currentUserID := filters.GetCurrentUserID(r)
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

	p := new(ProposalStatusParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	proposal, err := models.GetProposalByID(*tx, p.ProposalID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	// check status and ownership
	if proposal.ProposerID != currentUserID {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Only Proposed user can approve a proposal!",
		})
		return
	}

	proposal.AcceptStatus = p.Status
	proposal.ApprovalDate = time.Now().UTC()

	proposal.Update(*tx)

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": proposal.ID,
	})
}
