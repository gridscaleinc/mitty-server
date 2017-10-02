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

// CountOfEventInvitationByUserID ...
func CountOfEventInvitationByUserID(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := tx.SelectInt(`select count(*) from invitees
	    where invitee_id=$1 and reply_status='NONE';`, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetInviteeForAccept ...
func GetInviteeForAccept(tx *gorp.Transaction, invitationID int64, inviteesID int64, userID int) (*Invitees, error) {
	invitee := new(Invitees)
	if err := tx.SelectOne(&invitee, "select * from invitees where id = $1 and invitation_id=$2 and invitee_id=$3", inviteesID, invitationID, userID); err != nil {
		return nil, err
	}
	return invitee, nil
}
