package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Invitees ...
type Invitees struct {
	ID           int64     `db:"id" json:"id"`
	InvitationID int64     `db:"invitation_id" json:"invitation_id"`
	InviteeID    int       `db:"invitee_id" json:"invitee_id"`
	ReplyStatus  string    `db:"reply_status" json:"reply_status"`
	ReplyTime    time.Time `db:"reply_time" json:"reply_time"`
}

// Insert ...
func (s *Invitees) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Invitees) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Invitees) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}
