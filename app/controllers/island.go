package controllers

import (
	"errors"
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/geo"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// IslandParams ...
type IslandParams struct {
	Nickname      string  `json:"nickname"`
	Name          string  `json:"name"`
	LogoID        int     `json:"logo_id"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
	MobilityType  string  `json:"mobilityType"`
	RealityType   string  `json:"realityType"`
	OwnershipType string  `json:"ownershipType"`
	OwnerName     string  `json:"ownerName"`
	OwnerID       int     `json:"ownerId"`
	CreatorID     int     `json:"creatorId"`
	MeetingID     int     `json:"meetingId"`
	GalleryID     int64   `json:"galleryId"`
	Tel           string  `json:"tel"`
	Fax           string  `json:"fax"`
	MailAddress   string  `json:"mailaddress"`
	Webpage       string  `json:"webpage"`
	Likes         int     `json:"likes"`
	CountryCode   string  `json:"countryCode"`
	CountryName   string  `json:"countryName"`
	State         string  `json:"state"`
	City          string  `json:"city"`
	Postcode      string  `json:"postcode"`
	Thoroghfare   string  `json:"thoroghfare"`
	Subthroghfare string  `json:"subthroghfare"`
	BuildingName  string  `json:"buildingName"`
	FloorNumber   string  `json:"floorNumber"`
	RoomNumber    string  `json:"roomNumber"`
	Address1      string  `json:"address1"`
	Address2      string  `json:"address2"`
	Address3      string  `json:"address3"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

// FieldMap defines parameter requirements
func (p *IslandParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Nickname: binding.Field{
			Form:     "nickname",
			Required: false,
		},
		&p.Name: binding.Field{
			Form:     "name",
			Required: false,
		},
		&p.LogoID: binding.Field{
			Form:     "logoId",
			Required: false,
		},
		&p.Description: binding.Field{
			Form:     "description",
			Required: false,
		},
		&p.Category: binding.Field{
			Form:     "category",
			Required: true,
		},
		&p.MobilityType: binding.Field{
			Form:     "mobilityType",
			Required: true,
		},
		&p.RealityType: binding.Field{
			Form:     "realityType",
			Required: true,
		},
		&p.OwnershipType: binding.Field{
			Form:     "ownershipType",
			Required: true,
		},
		&p.OwnerName: binding.Field{
			Form:     "ownerName",
			Required: false,
		},
		&p.OwnerID: binding.Field{
			Form:     "ownerId",
			Required: false,
		},
		&p.CreatorID: binding.Field{
			Form:     "creatorId",
			Required: false,
		},
		&p.MeetingID: binding.Field{
			Form:     "meetingId",
			Required: false,
		},
		&p.GalleryID: binding.Field{
			Form:     "galleryId",
			Required: false,
		},
		&p.Tel: binding.Field{
			Form:     "tel",
			Required: false,
		},
		&p.Fax: binding.Field{
			Form:     "fax",
			Required: false,
		},
		&p.MailAddress: binding.Field{
			Form:     "mailaddress",
			Required: false,
		},
		&p.Webpage: binding.Field{
			Form:     "webpage",
			Required: false,
		},
		&p.Likes: binding.Field{
			Form:     "likes",
			Required: false,
		},
		&p.CountryCode: binding.Field{
			Form:     "countryCode",
			Required: false,
		},
		&p.CountryName: binding.Field{
			Form:     "countryName",
			Required: false,
		},
		&p.State: binding.Field{
			Form:     "state",
			Required: false,
		},
		&p.City: binding.Field{
			Form:     "city",
			Required: false,
		},
		&p.Postcode: binding.Field{
			Form:     "postcode",
			Required: false,
		},
		&p.Thoroghfare: binding.Field{
			Form:     "thoroghfare",
			Required: false,
		},
		&p.Subthroghfare: binding.Field{
			Form:     "subthroghfare",
			Required: false,
		},
		&p.BuildingName: binding.Field{
			Form:     "buildingName",
			Required: false,
		},
		&p.FloorNumber: binding.Field{
			Form:     "floorNumber",
			Required: false,
		},
		&p.RoomNumber: binding.Field{
			Form:     "roomNumber",
			Required: false,
		},
		&p.Address1: binding.Field{
			Form:     "address1",
			Required: false,
		},
		&p.Address2: binding.Field{
			Form:     "address2",
			Required: false,
		},
		&p.Address3: binding.Field{
			Form:     "address3",
			Required: false,
		},
		&p.Latitude: binding.Field{
			Form:     "latitude",
			Required: false,
		},
		&p.Longitude: binding.Field{
			Form:     "longitude",
			Required: false,
		},
	}
}

// Validate ...
func (p *IslandParams) Validate(req *http.Request) error {
	if len(p.Nickname) > 89 {
		return errors.New("nickname is too long")
	}
	if len(p.Name) > 80 {
		return errors.New("name is too long")
	}
	if len(p.Category) > 20 {
		return errors.New("category is too long")
	}
	if len(p.MobilityType) > 20 {
		return errors.New("mobility type is too long")
	}
	if len(p.RealityType) > 20 {
		return errors.New("reality type is too long")
	}
	if len(p.OwnershipType) > 20 {
		return errors.New("ownership type is too long")
	}
	if len(p.OwnerName) > 80 {
		return errors.New("owner name is too long")
	}
	if len(p.Tel) > 20 {
		return errors.New("tel is too long")
	}
	if len(p.Fax) > 20 {
		return errors.New("fax is too long")
	}
	if len(p.MailAddress) > 50 {
		return errors.New("maill address is too long")
	}
	if len(p.Webpage) > 50 {
		return errors.New("webpage is too long")
	}
	if len(p.CountryCode) > 2 {
		return errors.New("country code is too long")
	}
	if len(p.CountryName) > 30 {
		return errors.New("country name is too long")
	}
	if len(p.State) > 30 {
		return errors.New("state is too long")
	}
	if len(p.City) > 30 {
		return errors.New("city is too long")
	}
	if len(p.Postcode) > 20 {
		return errors.New("postcode is too long")
	}
	if len(p.Thoroghfare) > 30 {
		return errors.New("thoroghfare is too long")
	}
	if len(p.Subthroghfare) > 30 {
		return errors.New("subthroghfare is too long")
	}
	if len(p.BuildingName) > 50 {
		return errors.New("building name is too long")
	}
	if len(p.FloorNumber) > 3 {
		return errors.New("floor number is too long")
	}
	if len(p.RoomNumber) > 10 {
		return errors.New("room number is too long")
	}
	if len(p.Address1) > 100 {
		return errors.New("address1 is too long")
	}
	if len(p.Address2) > 100 {
		return errors.New("address2 is too long")
	}
	if len(p.Address3) > 100 {
		return errors.New("address3 is too long")
	}
	return nil
}

// PostIslandHandler ...
func PostIslandHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(IslandParams)
	if errs := binding.Bind(r, p); errs != nil {
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	m := new(models.Meeting)
	m.Name = p.Name
	m.Type = models.IslandType
	if err := m.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	island := new(models.Island)
	island.Nickname = p.Nickname
	island.Name = p.Name
	island.LogoID = p.LogoID
	island.Description = p.Description
	island.Category = p.Category
	island.MobilityType = p.MobilityType
	island.RealityType = p.RealityType
	island.OwnershipType = p.OwnershipType
	island.OwnerName = p.OwnerName
	island.OwnerID = p.OwnerID
	island.CreatorID = p.CreatorID
	island.MeetingID = m.ID
	island.GalleryID = p.GalleryID
	island.Tel = p.Tel
	island.Fax = p.Fax
	island.MailAddress = p.MailAddress
	island.Webpage = p.Webpage
	island.Likes = p.Likes
	island.CountryCode = p.CountryCode
	island.CountryName = p.CountryName
	island.State = p.State
	island.City = p.City
	island.Postcode = p.Postcode
	island.Thoroghfare = p.Thoroghfare
	island.Subthroghfare = p.Subthroghfare
	island.BuildingName = p.BuildingName
	island.FloorNumber = p.FloorNumber
	island.RoomNumber = p.RoomNumber
	island.Address1 = p.Address1
	island.Address2 = p.Address2
	island.Address3 = p.Address3
	island.Latitude = p.Latitude
	island.Longitude = p.Longitude

	// if geo info set, calculate the geohashes.
	if island.Latitude != 999 && island.Longitude != 999 {
		island.GeohashLevel8 = geo.GenerateHashID(island.Latitude, island.Longitude, 8)
		island.GeohashLevel10 = geo.GenerateHashID(island.Latitude, island.Longitude, 10)
		island.GeohashLevel12 = geo.GenerateHashID(island.Latitude, island.Longitude, 12)
	}

	if err := island.Insert(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"islandId": island.ID,
	})
}

// GetIslandHandler ...
func GetIslandHandler(w http.ResponseWriter, r *http.Request) {
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
	name := r.URL.Query().Get("name")
	islands, err := models.SearchIslandByName(tx, name)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"islands": islands,
	})

}
