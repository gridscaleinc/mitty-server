package controllers

import (
	"net/http"
	"time"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// RequestParams ...
type RequestParams struct {
	Title              string    `json:"title"`
	Tag                string    `json:"tag"`
	Description        string    `json:"description"`
	ForActivityID      int       `json:"forActivityId"`
	PreferredDatetime1 time.Time `json:"preferredDatetime1"`
	PreferredDatetime2 time.Time `json:"preferredDatetime2"`
	PreferredPrice1    int       `json:"preferredPrice1"`
	PreferredPrice2    int       `json:"preferredPrice2"`
	StartPlace         string    `json:"startPlace"`
	TerminatePlace     string    `json:"terminatePlace"`
	Oneway             bool      `json:"oneway"`
	Status             string    `json:"status"`
	ExpiryDate         time.Time `json:"expiryDate"`
	NumOfPerson        int       `json:"numOfPerson"`
	NumOfChildren      int       `json:"numOfChildren"`
}

// FieldMap defines parameter requirements
func (p *RequestParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Title: binding.Field{
			Form:     "title",
			Required: true,
		},
		&p.Tag: binding.Field{
			Form:     "tag",
			Required: true,
		},
		&p.Description: binding.Field{
			Form:     "description",
			Required: true,
		},
		&p.ForActivityID: binding.Field{
			Form:     "forActivityId",
			Required: false,
		},
		&p.PreferredDatetime1: binding.Field{
			Form:     "preferredDatetime1",
			Required: false,
		},
		&p.PreferredDatetime2: binding.Field{
			Form:     "preferredDatetime2",
			Required: false,
		},
		&p.PreferredPrice1: binding.Field{
			Form:     "preferredPrice1",
			Required: false,
		},
		&p.PreferredPrice2: binding.Field{
			Form:     "preferredPrice2",
			Required: false,
		},
		&p.StartPlace: binding.Field{
			Form:     "startPlace",
			Required: false,
		},
		&p.TerminatePlace: binding.Field{
			Form:     "terminatePlace",
			Required: false,
		},
		&p.Oneway: binding.Field{
			Form:     "oneway",
			Required: false,
		},
		&p.Status: binding.Field{
			Form:     "status",
			Required: false,
		},
		&p.ExpiryDate: binding.Field{
			Form:     "expiryDate",
			Required: false,
		},
		&p.NumOfPerson: binding.Field{
			Form:     "numOfPerson",
			Required: false,
		},
		&p.NumOfChildren: binding.Field{
			Form:     "numOfChildren",
			Required: false,
		},
	}
}

// PostRequestHandler ...
func PostRequestHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(RequestParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputError(w, r, errs)
		return
	}

	meeting := new(models.Meeting)
	meeting.Name = p.Title
	meeting.Type = "REQUEST"
	if err := meeting.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	m := new(models.Request)
	m.Title = p.Title
	m.Tag = p.Tag
	m.Description = p.Description
	m.ForActivityID = p.ForActivityID
	m.PreferredDatetime1 = p.PreferredDatetime1
	m.PreferredDatetime2 = p.PreferredDatetime2
	m.PreferredPrice1 = p.PreferredPrice1
	m.PreferredPrice2 = p.PreferredPrice2
	m.StartPlace = p.StartPlace
	m.TerminatePlace = p.TerminatePlace
	m.Oneway = p.Oneway
	m.Status = p.Status
	m.ExpiryDate = p.ExpiryDate
	m.NumOfPerson = p.NumOfPerson
	m.NumOfChildren = p.NumOfChildren
	m.MeetingID = meeting.ID
	m.OwnerID = currentUserID
	if err := m.Save(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": m.ID,
	})
}
