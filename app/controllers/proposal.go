package controllers

import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
	"net/http"
// 	"reflect"
	"time"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// RequestParams ...
type ProposalParams struct {
    ReplyToRequestID		int64     `json:"reply_to_request_id"`
    ContactTel		string     `json:"contact_tel"`
    ContactEmail		string     `json:"contact_email"`
    ProposedIslandID		int64     `json:"proposed_island_id"`
    ProposedIslandID2		int64     `json:"proposed_island_id2"`
    GalleryID		int64     `json:"gallery_id"`
    PriceName1		string     `json:"priceName1"`
    Price1		int     `json:"price1"`
    PriceName2		string     `json:"priceName2"`
    Price2		int     `json:"price2"`
    PriceCurrency		string     `json:"price_currency"`
    PriceInfo		string     `json:"price_info"`
    ProposedDatetime1		time.Time `json:"proposed_datetime1"`
    ProposedDatetime2		time.Time `json:"proposed_datetime2"`
    AdditionalInfo		string     `json:"additional_info"`
    ProposerInfo		string     `json:"proposer_info"`
    ConfirmTel		string     `json:"confirm_tel"`
    ConfirmEmail		string     `json:"confirm_email"`
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
			Required: true,
		},
		&p.ProposedIslandID: binding.Field{
			Form:     "proposed_island_id",
			Required: true,
		},
       &p.ProposedIslandID2: binding.Field{
			Form:     "proposed_island_id2",
			Required: true,
		},
		&p.GalleryID: binding.Field{
			Form:     "gallery_id",
			Required: true,
		},
       &p.PriceName1: binding.Field{
			Form:     "priceName1",
			Required: true,
		},
    	&p.Price1: binding.Field{
			Form:     "price1",
			Required: true,
		},
       &p.PriceName2: binding.Field{
			Form:     "priceName2",
			Required: true,
		},
		&p.Price2: binding.Field{
			Form:     "price2",
			Required: true,
		},
      &p.PriceCurrency: binding.Field{
			Form:     "price_currency",
			Required: true,
		},
      &p.PriceInfo: binding.Field{
			Form:     "price_info",
			Required: true,
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
			Required: true,
		},
      &p.ConfirmTel: binding.Field{
			Form:     "confirm_tel",
			Required: true,
		},
     &p.ConfirmEmail: binding.Field{
			Form:     "confirm_email",
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
		filters.RenderInputError(w, r, errs)
		return
	}

	m := new(models.Proposal)
	
	m.ReplyToRequestID  =  p.ReplyToRequestID
    m.ContactTel = p.ContactTel
    m.ContactEmail	= p.ContactEmail	
    m.ProposedIslandID = p.ProposedIslandID
    m.ProposedIslandID2 = p.ProposedIslandID2
    m.GalleryID	= p.GalleryID	
    m.PriceName1 = p.PriceName1
    m.Price1 = p.Price1
    m.PriceName2	 = p.PriceName2	
    m.Price2	= p.Price2		
    m.PriceCurrency	= p.PriceCurrency		
    m.PriceInfo	= p.PriceInfo		
    m.ProposedDatetime1	 = p.ProposedDatetime1		
    m.ProposedDatetime2	 = p.ProposedDatetime2		
    m.AdditionalInfo  = p.AdditionalInfo	
    m.ProposerID = currentUserID
    m.ProposerInfo = p.ProposerInfo	
    m.ProposedDatetime = time.Now().UTC()
    m.	AcceptStatus = "NONE"
    m.ConfirmTel = p.ConfirmTel
    m.ConfirmEmail = p.ConfirmEmail
    m.ApprovalStatus = "NONE"
    
	if err := m.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": m.ID,
	})
}
