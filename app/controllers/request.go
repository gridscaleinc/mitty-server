package controllers

import (
	"database/sql"
	"net/http"
	"reflect"
	"strconv"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"

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

// GetSearchRequestHandler ...
func GetSearchRequestHandler(w http.ResponseWriter, r *http.Request) {
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

	matchQuery1 := elastic.NewMatchQuery("title", queryParams)
	matchQuery2 := elastic.NewMatchQuery("tag", queryParams)
	matchQuery3 := elastic.NewMatchQuery("description", queryParams)
	matchQuery4 := elastic.NewMatchQuery("start_place", queryParams)
	matchQuery5 := elastic.NewMatchQuery("terminate_place", queryParams)

	query := elastic.NewBoolQuery()
	query.Should(matchQuery1, matchQuery2, matchQuery3, matchQuery4, matchQuery5)

	// Debug
	// src, err := query.Source()
	// if err != nil {
	// 	panic(err)
	// }
	// data, err := json.Marshal(src)
	// if err != nil {
	// 	panic(err)
	// }
	// s := string(data)
	// fmt.Println(s)

	searchResult, err := helpers.ESSearchBoolQuery("mitty", "request", "id", offset, limit, query)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	var requests []interface{}
	var request models.Request
	for _, item := range searchResult.Each(reflect.TypeOf(request)) {
		if t, ok := item.(models.Request); ok {
			requestDetail, err := models.GetRequestDetailByID(tx, t.ID)
			if err != nil && err != sql.ErrNoRows {
				filters.RenderError(w, r, err)
				return
			}
			requests = append(requests, requestDetail)
		}
	}
	render.JSON(w, http.StatusOK, map[string]interface{}{
		"requests": requests,
	})
}

// GetRequestDetailsHandler ...
func GetRequestDetailsHandler(w http.ResponseWriter, r *http.Request) {
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

	reqID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	request, err := models.GetRequestDetailByID(tx, reqID)

	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"request": request,
	})
}

// GetMyRequestHandler ...
func GetMyRequestHandler(w http.ResponseWriter, r *http.Request) {
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
	currentUserID := filters.GetCurrentUserID(r)

	requests, err := models.GetRequestByUserID(tx, currentUserID, queryParams)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"requests": requests,
	})
}

// GetProposalsHandler ...
func GetProposalsHandler(w http.ResponseWriter, r *http.Request) {
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

	requestID, err := strconv.ParseInt(r.URL.Query().Get("requestId"), 10, 64)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	proposals, err := models.GetProposalsOf(tx, requestID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"proposals": proposals,
	})
}
