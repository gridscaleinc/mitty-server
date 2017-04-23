package controllers

import (
	"net/http"
	"time"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// EventParams ...
type EventParams struct {
	Type          string
	Category      string
	Title         string
	Action        string
	StartDatetime time.Time
	EndDatetime   time.Time
	AlldayFlag    bool
	IslandID      int
	PriceName1    string
	Price1        int
	PriceName2    string
	Price2        int
	Currency      string
	PriceInfo     string
	Description   string
	ContactTel    string
	ContactFax    string
	ContactMail   string
	OfficialURL   string
	Organizer     string
	SourceName    string
	SourceURL     string
	Anticipation  string
	AccessControl string
	Language      string
}

// FieldMap defines parameter requirements
func (p *EventParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Type: binding.Field{
			Form:     "type",
			Required: true,
		},
		&p.Category: binding.Field{
			Form:     "category",
			Required: true,
		},
		&p.Title: binding.Field{
			Form:     "title",
			Required: true,
		},
		&p.Action: binding.Field{
			Form:     "action",
			Required: true,
		},
		&p.StartDatetime: binding.Field{
			Form:     "startDate",
			Required: true,
		},
		&p.EndDatetime: binding.Field{
			Form:     "endDate",
			Required: true,
		},
		&p.AlldayFlag: binding.Field{
			Form:     "allDayFlag",
			Required: true,
		},
		&p.IslandID: binding.Field{
			Form:     "islandId",
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
		&p.Currency: binding.Field{
			Form:     "currency",
			Required: true,
		},
		&p.PriceInfo: binding.Field{
			Form:     "priceInfo",
			Required: true,
		},
		&p.Description: binding.Field{
			Form:     "description",
			Required: true,
		},
		&p.ContactTel: binding.Field{
			Form:     "contactTel",
			Required: true,
		},
		&p.ContactFax: binding.Field{
			Form:     "contactFax",
			Required: true,
		},
		&p.ContactMail: binding.Field{
			Form:     "contactMail",
			Required: true,
		},
		&p.OfficialURL: binding.Field{
			Form:     "officialUrl",
			Required: true,
		},
		&p.Organizer: binding.Field{
			Form:     "organizer",
			Required: true,
		},
		&p.SourceName: binding.Field{
			Form:     "sourceName",
			Required: true,
		},
		&p.SourceURL: binding.Field{
			Form:     "sourceUrl",
			Required: true,
		},
		&p.Anticipation: binding.Field{
			Form:     "anticipation",
			Required: true,
		},
		&p.AccessControl: binding.Field{
			Form:     "accessControl",
			Required: true,
		},
		&p.Language: binding.Field{
			Form:     "language",
			Required: true,
		},
	}
}

// PostEventHandler ...
func PostEventHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(EventParams)
	if errors := binding.Bind(r, p); errors != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}

	e := new(models.Event)
	e.Type = p.Type
	e.Category = p.Category
	e.Title = p.Title
	e.Action = p.Action
	e.StartDatetime = p.StartDatetime
	e.EndDatetime = p.EndDatetime
	e.AlldayFlag = p.AlldayFlag
	e.IslandID = p.IslandID
	e.PriceName1 = p.PriceName1
	e.Price1 = p.Price1
	e.PriceName2 = p.PriceName2
	e.Price2 = p.Price2
	e.Currency = p.Currency
	e.PriceInfo = p.PriceInfo
	e.Description = p.Description
	e.ContactTel = p.ContactTel
	e.ContactFax = p.ContactFax
	e.ContactMail = p.ContactMail
	e.OfficialURL = p.OfficialURL
	e.Organizer = p.Organizer
	e.SourceName = p.SourceName
	e.SourceURL = p.SourceURL
	e.Anticipation = p.Anticipation
	e.AccessControl = p.AccessControl
	e.Language = p.Language
	if err := e.Save(*tx); err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"eventId": e.ID,
	})
}
