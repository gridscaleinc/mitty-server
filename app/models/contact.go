package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Contact ...
type Contact struct {
	ID             int64     `db:"id" json:"id"`
	MittyID        int       `db:"mitty_id" json:"mitty_id"`
	RelatedEventID int64     `db:"related_event_id" json:"related_event_id"`
	NameCardID     int64     `db:"name_card_id" json:"name_card_id"`
	MeetingID      int64     `db:"meeting_id" json:"meeting_id"`
	ContctedDate   time.Time `db:"contacted_date" json:"contacted_date"`
}

// Insert ...
func (s *Contact) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Contact) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// ExistContactFromIDs ...
func ExistContactFromIDs(tx gorp.Transaction, mittyID int, nameCardID int64) (*Contact, error) {
	contact := new(Contact)
	err := tx.SelectOne(&contact, "select * from contact where mitty_id = $1 and name_card_id = $2", mittyID, nameCardID)
	if err != nil {
		return nil, err
	}
	return contact, nil
}
