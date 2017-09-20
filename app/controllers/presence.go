package controllers

import (
	"errors"
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// CheckinParams ...
type CheckinParams struct {
	EventID        int64  `json:"event_id"`
	IslandID       int64  `json:"island_id"`
	NamecardID     int64  `json:"name_card_id"`
	PictureID      int64  `json:"picture_id"`
	SeatOrRoomInfo string `json:"seat_or_room_info"`
}

// FieldMap defines parameter requirements
func (p *CheckinParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.EventID: binding.Field{
			Form:     "event_id",
			Required: false,
		},
		&p.IslandID: binding.Field{
			Form:     "island_id",
			Required: false,
		},
		&p.NamecardID: binding.Field{
			Form:     "name_card_id",
			Required: false,
		},
		&p.PictureID: binding.Field{
			Form:     "picture_id",
			Required: false,
		},
		&p.SeatOrRoomInfo: binding.Field{
			Form:     "seat_or_room_info",
			Required: false,
		},
	}
}

// Validate ...
func (p *CheckinParams) Validate(req *http.Request) error {
	if len(p.SeatOrRoomInfo) > 100 {
		return errors.New("seat_or_room_info is too long")
	}

	// one of the IDs shold be set as non-zero value
	if p.IslandID == 0 && p.EventID == 0 {
		return errors.New("Ether IslandId or EventID should be set")
	}

	return nil
}

// PostCheckinHandler ...
func PostCheckinHandler(w http.ResponseWriter, r *http.Request) {
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

	p := new(CheckinParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if errs := p.Validate(r); errs != nil {
		filters.RenderError(w, r, errs)
		return
	}

	checkin := new(models.Footmark)
	checkin.EventID = p.EventID
	checkin.IslandID = p.IslandID
	checkin.MittyID = filters.GetCurrentUserID(r)
	checkin.SeatOrRoomInfo = p.SeatOrRoomInfo

	// TODO: IF Already checked in delete it first
	err = checkin.Insert(*tx)

	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}
