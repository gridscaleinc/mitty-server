package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mholt/binding"
	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	goutils "github.com/dongri/goutils"

// 	"mitty.co/mitty-server/config"

// NameCardParams ...
type NameCardParams struct {
	ID              int64  `json:"id"`
	MittyID         int    `json:"mitty_id"`
	Name            string `json:"name"`
	BusinessName    string `json:"business_name"`
	BusinessLogoID  int64  `json:"business_logo_id"`
	BusinessSubName string `json:"business_sub_name"`
	BusinessTitle   string `json:"business_title"`
	AddressLine1    string `json:"address_line1"`
	AddressLine2    string `json:"address_line2"`
	Phone           string `json:"phone"`
	Fax             string `json:"fax"`
	MobilePhone     string `json:"mobile_phone"`
	Webpage         string `json:"webpage"`
	Email           string `json:"email"`
}

// FieldMap defines parameter requirements
func (p *NameCardParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.ID: "id",
		&p.MittyID: binding.Field{
			Form:     "mitty_id",
			Required: true,
		},
		&p.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
		&p.BusinessName: binding.Field{
			Form:     "business_name",
			Required: true,
		},
		&p.BusinessLogoID: binding.Field{
			Form:     "business_logo_id",
			Required: false,
		},
		&p.BusinessSubName: binding.Field{
			Form:     "business_sub_name",
			Required: false,
		},
		&p.BusinessTitle: binding.Field{
			Form:     "business_title",
			Required: false,
		},
		&p.AddressLine1: binding.Field{
			Form:     "address_line1",
			Required: false,
		},
		&p.AddressLine2: binding.Field{
			Form:     "address_line2",
			Required: false,
		},
		&p.Phone: binding.Field{
			Form:     "phone",
			Required: false,
		},
		&p.Fax: binding.Field{
			Form:     "fax",
			Required: false,
		},
		&p.MobilePhone: binding.Field{
			Form:     "mobile_phone",
			Required: false,
		},
		&p.Webpage: binding.Field{
			Form:     "webpage",
			Required: false,
		},
		&p.Email: binding.Field{
			Form:     "email",
			Required: false,
		},
	}
}

// Validate ...
func (p *NameCardParams) Validate(req *http.Request) error {
	if len(p.BusinessName) > 200 {
		return errors.New("business name is too long")
	}
	if len(p.BusinessSubName) > 200 {
		return errors.New("business sub name is too long")
	}
	if len(p.BusinessTitle) > 200 {
		return errors.New("business title is too long")
	}
	if len(p.AddressLine1) > 100 {
		return errors.New("address line1 is too long")
	}
	if len(p.AddressLine2) > 100 {
		return errors.New("address line2 is too long")
	}
	if len(p.Phone) > 20 {
		return errors.New("phone is too long")
	}
	if len(p.Fax) > 20 {
		return errors.New("fax is too long")
	}
	if len(p.MobilePhone) > 20 {
		return errors.New("mobile phone is too long")
	}
	if len(p.Webpage) > 100 {
		return errors.New("webpage is too long")
	}
	if len(p.Email) > 100 {
		return errors.New("email is too long")
	}
	if len(p.Name) > 200 {
		return errors.New("name is too long")
	}
	return nil
}

// PostNameCardHandler ...
func PostNameCardHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(NameCardParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	m := new(models.Namecard)

	m.ID = p.ID
	m.MittyID = currentUserID
	m.Name = p.Name
	m.BusinessName = p.BusinessName
	m.BusinessLogoID = p.BusinessLogoID
	m.BusinessSubName = p.BusinessSubName
	m.BusinessTitle = p.BusinessTitle
	m.AddressLine1 = p.AddressLine1
	m.AddressLine2 = p.AddressLine2
	m.Webpage = p.Webpage
	m.Email = p.Email
	m.Phone = p.Phone
	m.MobilePhone = p.MobilePhone
	m.Fax = p.Fax

	if err := m.Save(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": m.ID,
	})
}

// GetMyNamecardsHandler ...
func GetMyNamecardsHandler(w http.ResponseWriter, r *http.Request) {
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

	namecards, err := models.GetNamecardsByUserID(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"namecards": namecards,
	})
}

// GetNamecardHandler ...
func GetNamecardHandler(w http.ResponseWriter, r *http.Request) {
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

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	namecard, err := models.GetNamecardByID(tx, id)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"namecard": namecard,
	})
}
