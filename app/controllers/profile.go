package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// ProfileParams ...
type ProfileParams struct {
	ID             int64  `json:"id"`
	MittyID        int    `json:"mitty_id"`
	Gender         string `json:"gender"`
	OneWordSpeech  string `json:"one_word_speech"`
	Constellation  string `json:"constellation"`
	HomeIslandID   int64  `json:"home_island_id"`
	BirthIslandID  int64  `json:"birth_island_id"`
	AgeGroup       string `json:"age_group"`
	AppearanceTag  string `json:"appearance_tag"`
	OccupationTag1 string `json:"occupation_tag1"`
	OccupationTag2 string `json:"occupation_tag2"`
	OccupationTag3 string `json:"occupation_tag3"`
	HobbyTag1      string `json:"hobby_tag1"`
	HobbyTag2      string `json:"hobby_tag2"`
	HobbyTag3      string `json:"hobby_tag3"`
	HobbyTag4      string `json:"hobby_tag4"`
	HobbyTag5      string `json:"hobby_tag5"`
}

// FieldMap defines parameter requirements
func (p *ProfileParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.ID: "id",
		&p.Gender: binding.Field{
			Form:     "gender",
			Required: true,
		},
		&p.OneWordSpeech: binding.Field{
			Form:     "one_word_speech",
			Required: false,
		},
		&p.Constellation: binding.Field{
			Form:     "constellation",
			Required: false,
		},
		&p.HomeIslandID: binding.Field{
			Form:     "home_island_id",
			Required: false,
		},
		&p.BirthIslandID: binding.Field{
			Form:     "birth_island_id",
			Required: false,
		},
		&p.AgeGroup: binding.Field{
			Form:     "age_group",
			Required: false,
		},
		&p.AppearanceTag: binding.Field{
			Form:     "appearance_tag",
			Required: false,
		},
		&p.OccupationTag1: binding.Field{
			Form:     "occupation_tag1",
			Required: false,
		},
		&p.OccupationTag2: binding.Field{
			Form:     "occupation_tag2",
			Required: false,
		},
		&p.OccupationTag3: binding.Field{
			Form:     "occupation_tag3",
			Required: false,
		},
		&p.HobbyTag1: binding.Field{
			Form:     "hobby_tag1",
			Required: false,
		},
		&p.HobbyTag2: binding.Field{
			Form:     "hobby_tag2",
			Required: false,
		},
		&p.HobbyTag3: binding.Field{
			Form:     "hobby_tag3",
			Required: false,
		},
		&p.HobbyTag4: binding.Field{
			Form:     "hobby_tag4",
			Required: false,
		},
		&p.HobbyTag5: binding.Field{
			Form:     "hobby_tag5",
			Required: false,
		},
	}
}

// Validate ...
func (p *ProfileParams) Validate(req *http.Request) error {
	if len(p.Gender) > 10 {
		return errors.New("Gender is too long")
	}
	if len(p.OneWordSpeech) > 200 {
		return errors.New("OneWordSpeech is too long")
	}
	if len(p.Constellation) > 20 {
		return errors.New("Constellation is too long")
	}
	if len(p.AgeGroup) > 20 {
		return errors.New("AgeGroup is too long")
	}
	if len(p.AppearanceTag) > 20 {
		return errors.New("AppearanceTag is too long")
	}
	if len(p.OccupationTag1) > 20 {
		return errors.New("OccupationTag1 is too long")
	}
	if len(p.OccupationTag2) > 20 {
		return errors.New("OccupationTag2 is too long")
	}
	if len(p.OccupationTag3) > 20 {
		return errors.New("OccupationTag3 is too long")
	}
	if len(p.HobbyTag1) > 20 {
		return errors.New("HobbyTag1 is too long")
	}
	if len(p.HobbyTag2) > 20 {
		return errors.New("HobbyTag2 is too long")
	}
	if len(p.HobbyTag3) > 20 {
		return errors.New("HobbyTag3 is too long")
	}
	if len(p.HobbyTag4) > 20 {
		return errors.New("HobbyTag4 is too long")
	}
	if len(p.HobbyTag5) > 20 {
		return errors.New("HobbyTag5 is too long")
	}
	return nil
}

// PostProfileHandler ...
func PostProfileHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(ProfileParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	m := new(models.Profile)

	m.ID = p.ID
	m.MittyID = currentUserID
	m.Gender = p.Gender
	m.OneWordSpeech = p.OneWordSpeech
	m.Constellation = p.Constellation
	m.HomeIslandID = p.HomeIslandID
	m.BirthIslandID = p.BirthIslandID
	m.AgeGroup = p.AgeGroup
	m.AppearanceTag = p.AppearanceTag
	m.OccupationTag1 = p.OccupationTag1
	m.OccupationTag2 = p.OccupationTag2
	m.OccupationTag3 = p.OccupationTag3
	m.HobbyTag1 = p.HobbyTag1
	m.HobbyTag2 = p.HobbyTag2
	m.HobbyTag3 = p.HobbyTag3
	m.HobbyTag4 = p.HobbyTag4
	m.HobbyTag5 = p.HobbyTag5

	if err := m.Save(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"id": m.ID,
	})
}

// GetMyProfileHandler ...
func GetMyProfileHandler(w http.ResponseWriter, r *http.Request) {
	currentUserID := filters.GetCurrentUserID(r)
	fetchProfile(w, r, currentUserID)
}

// GetUserProfileHandler ...
func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	forUserID, err := strconv.Atoi(r.URL.Query().Get("mitty_id"))
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	fetchProfile(w, r, forUserID)
}

// fetchProfile ...
func fetchProfile(w http.ResponseWriter, r *http.Request, userID int) {
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

	profile, err := models.GetProfileByUserID(tx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			profile = new(models.Profile)
			profile.MittyID = userID
		} else {
			filters.RenderError(w, r, err)
			return
		}
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"profile": profile,
	})
}
