package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/geo"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// EventParams ...
type EventParams struct {
	ID                int64     `json:"id"`
	Type              string    `json:"type"`
	Tag               string    `json:"tag"`
	Title             string    `json:"title"`
	LogoID            int       `json:"logoId"`
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
	Participation     string    `json:"participation"`
	AccessControl     string    `json:"accessControl"`
	Language          string    `json:"language"`
	RelatedActivityID int64     `json:"relatedActivityId"`
	AsMainEvent       bool      `json:"asMainEvent"`
}

// FieldMap defines parameter requirements
func (p *EventParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.ID: binding.Field{
			Form:     "id",
			Required: false,
		},
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
		&p.LogoID: binding.Field{
			Form:     "logoId",
			Required: false,
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
		&p.Participation: binding.Field{
			Form:     "participation",
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

// Validate ...
func (p *EventParams) Validate(req *http.Request) error {
	if len(p.Type) > 20 {
		return errors.New("type is too long")
	}
	if len(p.Tag) > 50 {
		return errors.New("tag is too long")
	}
	if len(p.Title) > 100 {
		return errors.New("title is too long")
	}
	if len(p.PriceName1) > 100 {
		return errors.New("price_name1 is too long")
	}
	if len(p.PriceName2) > 100 {
		return errors.New("price_name2 is too long")
	}
	if len(p.Currency) > 3 {
		return errors.New("currency is too long")
	}
	if len(p.PriceInfo) > 200 {
		return errors.New("price info is too long")
	}
	if len(p.ContactTel) > 20 {
		return errors.New("contact tel is too long")
	}
	if len(p.ContactFax) > 20 {
		return errors.New("contact fax is too long")
	}
	if len(p.ContactMail) > 50 {
		return errors.New("contact mail is too long")
	}
	if len(p.OfficialURL) > 200 {
		return errors.New("official url is too long")
	}
	if len(p.Organizer) > 100 {
		return errors.New("organizer url is too long")
	}
	if len(p.SourceName) > 100 {
		return errors.New("organizer url is too long")
	}
	if len(p.SourceURL) > 200 {
		return errors.New("organizer url is too long")
	}
	if len(p.Participation) > 20 {
		return errors.New("organizer url is too long")
	}
	if len(p.AccessControl) > 20 {
		return errors.New("organizer url is too long")
	}
	if len(p.Language) > 10 {
		return errors.New("language url is too long")
	}
	return nil
}

// PostEventHandler ...
func PostEventHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	currentUserID := filters.GetCurrentUserID(r)
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
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	m := new(models.Meeting)
	m.Name = p.Title
	m.Type = models.EventType
	if err := m.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	e := new(models.Event)
	e.Type = p.Type
	e.Tag = p.Tag
	e.Title = p.Title
	e.LogoID = p.LogoID
	e.Action = p.Action
	e.StartDatetime = p.StartDatetime
	e.EndDatetime = p.EndDatetime
	e.AlldayFlag = p.AlldayFlag
	e.IslandID = p.IslandID
	e.MeetingID = m.ID
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
	e.Participation = p.Participation
	e.AccessControl = p.AccessControl
	e.Language = p.Language
	e.PublisherID = currentUserID
	if err := e.Save(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	if p.RelatedActivityID != 0 {
		activityItem := new(models.ActivityItem)
		activityItem.ActivityID = p.RelatedActivityID
		activityItem.EventID = e.ID
		activityItem.Title = e.Title
		activityItem.Notification = false
		if err := activityItem.Insert(*tx); err != nil {
			filters.RenderError(w, r, err)
			return
		}

		if p.AsMainEvent == true {
			activity, err := models.GetActivityByID(tx, p.RelatedActivityID)
			if err != nil && err != sql.ErrNoRows {
				filters.RenderError(w, r, err)
				return
			}
			activity.MainEventID = activityItem.EventID
			if err := activity.Update(*tx); err != nil {
				filters.RenderError(w, r, err)
				return
			}
		}
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"eventId": e.ID,
	})
}

//UpdateEventHandler ... イベント編集した後に更新する。
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	currentUserID := filters.GetCurrentUserID(r)
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
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	if p.ID == 0 {
		filters.RenderError(w, r, errors.New("Event ID needed"))
		return
	}

	e, err := models.GetEventByID(tx, p.ID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	if e.PublisherID != currentUserID {
		filters.RenderError(w, r, errors.New("Not owner"))
		return
	}

	e.Type = p.Type
	e.Tag = p.Tag
	e.Title = p.Title
	e.LogoID = p.LogoID
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
	e.Participation = p.Participation
	e.AccessControl = p.AccessControl
	e.Language = p.Language

	if err := e.Update(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"OK": true,
	})
}

// FindEventByGeoHashHandler ...
func FindEventByGeoHashHandler(w http.ResponseWriter, r *http.Request) {
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

	rangeFlage := r.URL.Query().Get("range")
	latitude, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	longitude, err := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	var events []models.EventDetail
	if strings.Compare(rangeFlage, "near") == 0 {
		geohash := geo.GenerateHashID(latitude, longitude, 12)
		events, err = models.SearchByGeohashL12(tx, geohash)
	} else if strings.Compare(rangeFlage, "middle") == 0 {
		geohash := geo.GenerateHashID(latitude, longitude, 10)
		events, err = models.SearchByGeohashL10(tx, geohash)
	} else {
		geohash := geo.GenerateHashID(latitude, longitude, 8)
		events, err = models.SearchByGeohashL8(tx, geohash)
	}

	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"events": events,
	})
}

// SearchEventHandler ...
func SearchEventHandler(w http.ResponseWriter, r *http.Request) {
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

	queryParams := r.URL.Query().Get("q")
	offsetParams := r.URL.Query().Get("offset")
	limitParams := r.URL.Query().Get("limit")
	offset := 0
	limit := 30
	if offsetParams != "" {
		offset, _ = strconv.Atoi(offsetParams)
	}
	if limitParams != "" {
		limit, _ = strconv.Atoi(limitParams)
	}

	matchQuery1 := elastic.NewMatchQuery("action", queryParams)
	//matchQuery1.Boost(4)
	matchQuery2 := elastic.NewMatchQuery("description", queryParams)
	matchQuery3 := elastic.NewMatchQuery("sourceName", queryParams)
	matchQuery4 := elastic.NewMatchQuery("category", queryParams)
	matchQuery5 := elastic.NewMatchQuery("tag", queryParams)

	query1 := elastic.NewBoolQuery()
	query1.Should(matchQuery1, matchQuery2, matchQuery3, matchQuery4, matchQuery5)

	query := elastic.NewBoolQuery()
	query.Must(query1, elastic.NewTermQuery("accessControl", "public"), elastic.NewRangeQuery("endDatetime").Gt(time.Now()))
	//query.Must()

	// Debug
	src, err := query.Source()
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	s := string(data)
	fmt.Println(s)

	searchResult, err := helpers.ESSearchBoolQuery("mitty", "event", "id", offset, limit, query)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	userID := -1
	var events []interface{}
	var event models.Event
	for _, item := range searchResult.Each(reflect.TypeOf(event)) {
		if t, ok := item.(models.Event); ok {
			eventDetail, err := models.GetEventDetailByID(tx, userID, t.ID)
			if err != nil && err != sql.ErrNoRows {
				filters.RenderError(w, r, err)
				return
			}
			events = append(events, eventDetail)
		}
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"events": events,
	})
}

// EventFetchingHandler ...
func EventFetchingHandler(w http.ResponseWriter, r *http.Request) {
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
	idParams := r.URL.Query().Get("id")
	eventID, err := strconv.ParseInt(idParams, 10, 64)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	currentUserID := filters.GetCurrentUserID(r)

	event, err := models.GetEventDetailByID(tx, currentUserID, eventID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"event": event,
	})
}
