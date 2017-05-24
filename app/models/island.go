package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Island struct
type Island struct {
	ID            int64     `db:"id" json:"id"`
	Nickname      string    `db:"nickname" json:"nickname"`
	Name          string    `db:"name" json:"name"`
	LogoID        int       `db:"logo_id" json:"logo_id"`
	Description   string    `db:"description" json:"description"`
	Category      string    `db:"category" json:"category"`
	MobilityType  string    `db:"mobility_type" json:"mobility_type"`
	RealityType   string    `db:"reality_type" json:"reality_type"`
	OwnershipType string    `db:"ownership_type" json:"ownership_type"`
	OwnerName     string    `db:"owner_name" json:"owner_name"`
	OwnerID       int       `db:"owner_id" json:"owner_id"`
	CreatorID     int       `db:"creator_id" json:"creator_id"`
	MeetingID     int64     `db:"meeting_id" json:"meeting_id"`
	GalleryID     int64     `db:"gallery_id" json:"gallery_id"`
	Tel           string    `db:"tel" json:"tel"`
	Fax           string    `db:"fax" json:"fax"`
	MailAddress   string    `db:"mailaddress" json:"mailaddress"`
	Webpage       string    `db:"webpage" json:"webpage"`
	Likes         int       `db:"likes" json:"likes"`
	CountryCode   string    `db:"country_code" json:"country_code"`
	CountryName   string    `db:"country_name" json:"country_name"`
	State         string    `db:"state" json:"state"`
	City          string    `db:"city" json:"city"`
	Postcode      string    `db:"postcode" json:"postcode"`
	Thoroghfare   string    `db:"thoroghfare" json:"thoroghfare"`
	Subthroghfare string    `db:"subthroghfare" json:"subthroghfare"`
	BuildingName  string    `db:"building_name" json:"building_name"`
	FloorNumber   string    `db:"floor_number" json:"floor_number"`
	RoomNumber    string    `db:"room_number" json:"room_number"`
	Address1      string    `db:"address1" json:"address1"`
	Address2      string    `db:"address2" json:"address2"`
	Address3      string    `db:"address3" json:"address3"`
	Latitude      float64   `db:"latitude" json:"latitude"`
	Longitude     float64   `db:"longitude" json:"longitude"`
	Created       time.Time `db:"created" json:"created"`
	Updated       time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (s *Island) Insert(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Island) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}

// GetIslandByID ...
func GetIslandByID(tx *gorp.Transaction, ID int) (*Island, error) {
	island := new(Island)
	if err := tx.SelectOne(&island, "select * from island where id = $1", ID); err != nil {
		return nil, err
	}
	return island, nil
}

// SearchIslandByName ...
func SearchIslandByName(tx *gorp.Transaction, name string) (interface{}, error) {
	// Island struct
	type island struct {
		Nickname  string `db:"nickname" json:"nickname"`
		Name      string `db:"name" json:"name"`
		LogoURL   string `db:"logo_url" json:"logoUrl"`
		Address1  string `db:"address1" json:"address1"`
		Address2  string `db:"address2" json:"address2"`
		Address3  string `db:"address3" json:"address3"`
		Latitude  string `db:"latitude" json:"latitude"`
		Longitude string `db:"longitude" json:"longitude"`
	}
	islands := []island{}

	if _, err := tx.Select(&islands, `
		select
		  i.nickname, i.name, i.address1, i.address2, i.address3,
			i.latitude, i.longitude,
			COALESCE(c.link_url, '') as logo_url
		from island as i
		left join contents as c on c.id = i.logo_id
		where i.name like '%' || $1 || '%'`, name); err != nil {
		return nil, err
	}
	return islands, nil
}
