package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Presence ...
type Presence struct {
	ID             int       `db:"id" json:"id"`
	MittyID        int       `db:"mitty_id" json:"mitty_id"`
	LastLogedin    time.Time `db:"last_logedin" json:"last_logedin"`
	OnlineStatus   string    `db:"online_status" json:"online_status"`
	Visiability    string    `db:"visiability" json:"visiability"`
	Traceablity    string    `db:"traceablity" json:"traceablity"`
	Latitude       float64   `db:"latitude" json:"latitude"`
	Longitude      float64   `db:"longitude" json:"longitude"`
	LocatedTime    time.Time `db:"located_time" json:"located_time"`
	CheckinStatus  string    `db:"checkin_status" json:"checkin_status"`
	CheckinTime    time.Time `db:"checkin_time" json:"checkin_time"`
	CheckinEventID int       `db:"checkin_event_id" json:"checkin_event_id"`
	CheckinImage   int8      `db:"checkin_image_id" json:"checkin_image_id"`
}

// Insert ...
func (s *Presence) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Presence) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Presence) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}
