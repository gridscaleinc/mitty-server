package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Event struct
type Event struct {
	ID                  int64     `db:"id" json:"id"`
	Type                string    `db:"type" json:"type"`
	Category            string    `db:"category" json:"category"`
	Theme               string    `db:"theme" json:"theme"`
	Title               string    `db:"title" json:"title"`
	Action              string    `db:"action" json:"action"`
	StartDatetime       time.Time `db:"start_datetime" json:"start_datetime"`
	EndDatetime         time.Time `db:"end_datetime" json:"end_datetime"`
	AlldayFlag          bool      `db:"allday_flag" json:"allday_flag"`
	IslandID            int       `db:"islandId" json:"islandId"`
	LogoID              int       `db:"logo_id" json:"logo_id"`
	GalleryID           int64     `db:"gallery_id" json:"gallery_id"`
	MeetingID           int       `db:"meeting_id" json:"meeting_id"`
	PriceName1          string    `db:"price_name1" json:"price_name1"`
	Price1              int       `db:"price1" json:"price1"`
	PriceName2          string    `db:"price_name2" json:"price_name2"`
	Price2              int       `db:"price2" json:"price2"`
	Currency            string    `db:"currency" json:"currency"`
	PriceInfo           string    `db:"price_info" json:"price_info"`
	Description         string    `db:"description" json:"description"`
	ContactTel          string    `db:"contact_tel" json:"contact_tel"`
	ContactFax          string    `db:"contact_fax" json:"contact_fax"`
	ContactMail         string    `db:"contact_mail" json:"contact_mail"`
	OfficialURL         string    `db:"official_url" json:"official_url"`
	Organizer           string    `db:"organizer" json:"organizer"`
	SourceName          string    `db:"source_name" json:"source_name"`
	SourceURL           string    `db:"source_url" json:"source_url"`
	NumberOfAnticipants int       `db:"number_of_anticipants" json:"number_of_anticipants"`
	Anticipation        string    `db:"anticipation" json:"anticipation"`
	AccessControl       string    `db:"access_control" json:"access_control"`
	Likes               int       `db:"likes" json:"likes"`
	Status              string    `db:"status" json:"status"`
	Language            string    `db:"language" json:"language"`
	Created             time.Time `db:"created" json:"created"`
	PublisherID         int       `db:"publisher_id" json:"publisher_id"`
	OrgnizationID       int       `db:"orgnization_id" json:"orgnization_id"`
	Lastupdated         time.Time `db:"lastupdated" json:"lastupdated"`
	AmenderID           int       `db:"amender_id" json:"amender_id"`
}

// Save ...
func (s *Event) Save(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Lastupdated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Event) Update(tx gorp.Transaction) error {
	s.Lastupdated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}

// GetAdminEvents ...
func GetAdminEvents(dbmap *gorp.DbMap) ([]Event, error) {
	events := []Event{}
	_, err := dbmap.Select(&events, "select * from events")
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetEventByID ...
func GetEventByID(tx *gorp.Transaction, ID int) (*Event, error) {
	event := new(Event)
	if err := tx.SelectOne(&event, "select * from events where id = $1", ID); err != nil {
		return nil, err
	}
	return event, nil
}
