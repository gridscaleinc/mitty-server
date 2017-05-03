package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// EventParams ...
type EventParams struct {
	Type              string    `json:"type"`
	Tag               string    `json:"tag"`
	Title             string    `json:"title"`
	Action            string    `json:"action"`
	StartDatetime     time.Time `json:"startDate"`
	EndDatetime       time.Time `json:"endDate"`
	AlldayFlag        bool      `json:"allDayFlag"`
	IslandID          int       `json:"islandId"`
	PriceName1        string    `json:"priceName1"`
	Price1            int       `json:"price1"`
	PriceName2        string    `json:"priceName2"`
	Price2            int       `json:"price2"`
	Currency          string    `json:"currency"`
	PriceInfo         string    `json:"priceInfo"`
	Description       string    `json:"description"`
	ContactTel        string    `json:"contactTel"`
	ContactFax        string    `json:"contactFax"`
	ContactMail       string    `json:"contactMail"`
	OfficialURL       string    `json:"officialUrl"`
	Organizer         string    `json:"organizer"`
	SourceName        string    `json:"sourceName"`
	SourceURL         string    `json:"sourceUrl"`
	Anticipation      string    `json:"anticipation"`
	AccessControl     string    `json:"accessControl"`
	Language          string    `json:"language"`
	RelatedActivityID int       `json:"relatedActivityId"`
	AsMainEvent       bool      `json:"asMainEvent"`
}

// FieldMap defines parameter requirements
func (p *EventParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Type: binding.Field{
			Form:     "type",
			Required: true,
		},
		&p.Tag: binding.Field{
			Form:     "tag",
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
			Required: false,
		},
		&p.IslandID: binding.Field{
			Form:     "islandId",
			Required: true,
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
		&p.Currency: binding.Field{
			Form:     "currency",
			Required: false,
		},
		&p.PriceInfo: binding.Field{
			Form:     "priceInfo",
			Required: false,
		},
		&p.Description: binding.Field{
			Form:     "description",
			Required: true,
		},
		&p.ContactTel: binding.Field{
			Form:     "contactTel",
			Required: false,
		},
		&p.ContactFax: binding.Field{
			Form:     "contactFax",
			Required: false,
		},
		&p.ContactMail: binding.Field{
			Form:     "contactMail",
			Required: false,
		},
		&p.OfficialURL: binding.Field{
			Form:     "officialUrl",
			Required: false,
		},
		&p.Organizer: binding.Field{
			Form:     "organizer",
			Required: false,
		},
		&p.SourceName: binding.Field{
			Form:     "sourceName",
			Required: true,
		},
		&p.SourceURL: binding.Field{
			Form:     "sourceUrl",
			Required: false,
		},
		&p.Anticipation: binding.Field{
			Form:     "anticipation",
			Required: false,
		},
		&p.AccessControl: binding.Field{
			Form:     "accessControl",
			Required: false,
		},
		&p.Language: binding.Field{
			Form:     "language",
			Required: true,
		},
		&p.RelatedActivityID: binding.Field{
			Form:     "relatedActivityId",
			Required: false,
		},
		&p.AsMainEvent: binding.Field{
			Form:     "asMainEvent",
			Required: false,
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
	if errs := binding.Bind(r, p); errs != nil {
		helpers.RenderInputError(w, r, errs)
		return
	}

	e := new(models.Event)
	e.Type = p.Type
	e.Tag = p.Tag
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
		helpers.RenderDBError(w, r, err)
		return
	}

	if p.RelatedActivityID != 0 {
		activityItem := new(models.ActivityItem)
		activityItem.ActivityID = p.RelatedActivityID
		activityItem.EventID = e.ID
		activityItem.Title = e.Title
		activityItem.Notification = false
		if err := activityItem.Insert(*tx); err != nil {
			helpers.RenderDBError(w, r, err)
			return
		}

		if p.AsMainEvent == true {
			activity, err := models.GetActivityByID(tx, p.RelatedActivityID)
			if err != nil && err != sql.ErrNoRows {
				helpers.RenderDBError(w, r, err)
				return
			}
			activity.MainEventID = activityItem.EventID
			if err := activity.Update(*tx); err != nil {
				helpers.RenderDBError(w, r, err)
				return
			}
		}
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"eventId": e.ID,
	})
}

// SearchEventHandler ...
func SearchEventHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)

	queryParams := r.URL.Query().Get("q")

	matchQuery1 := elastic.NewMatchQuery("action", queryParams)
	matchQuery1.Boost(4)
	matchQuery2 := elastic.NewMatchQuery("description", queryParams)
	matchQuery2.Boost(2)
	matchQuery3 := elastic.NewMatchQuery("source_name", queryParams)
	matchQuery3.Boost(1)
	matchQuery4 := elastic.NewMatchQuery("category", queryParams)
	matchQuery4.Boost(1)
	matchQuery5 := elastic.NewMatchQuery("tag", queryParams)
	matchQuery5.Boost(1)

	query := elastic.NewBoolQuery()
	query.Should(matchQuery1, matchQuery2, matchQuery3, matchQuery4, matchQuery5)

	searchResult, err := helpers.ESSearchBoolQuery("mitty", "event", "id", 0, 100, query)
	if err != nil {
		helpers.RenderDBError(w, r, err)
		return
	}

	var events []models.Event
	var event models.Event
	for _, item := range searchResult.Each(reflect.TypeOf(event)) {
		if t, ok := item.(models.Event); ok {
			fmt.Printf("User by %s: %s\n", t.Title, t.Description)
			events = append(events, t)
		}
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"events": events,
	})
}
