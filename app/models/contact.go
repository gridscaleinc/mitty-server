package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Contact ...
type Contact struct {
	ID             int64     `db:"id" json:"id"`
	RelatedEventID int64     `db:"related_event_id" json:"related_event_id"`
	NameCardID     int64     `db:"name_card_id" json:"name_card_id"`
	ContctedDate   time.Time `contacted_date:"id" json:"contacted_date"`
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
