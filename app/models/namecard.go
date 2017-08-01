package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Likes struct
type Namecard struct {
	ID              int64     `db:"id" json:"id"`
	MittyID         int       `db:"mitty_id" json:"mitty_id"`
	BusinessNae     string    `db:"business_name" json:"business_name"`
	BusinessSubName string    `db:"business_sub_name" json:"business_sub_name"`
	BusinessTitle   string    `db:"business_title" json:"business_title"`
	AddressLine1    string    `db:"address_line1" json:"address_line1"`
	AddressLine2    string    `db:"address_line2" json:"address_line2"`
	Phone           string    `db:"phone" json:"phone"`
	Fax             string    `db:"fax" json:"fax"`
	MobilePhone     string    `db:"mobile_phone" json:"mobile_phone"`
	Webpage         string    `db:"webpage" json:"webpage"`
	Email           string    `db:"email" json:"email"`
	Created         time.Time `db:"created" json:"created"`
	Updated         time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (s *Namecard) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Namecard) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Namecard) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}
