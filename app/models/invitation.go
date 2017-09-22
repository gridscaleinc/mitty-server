package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Invitation ...
type Invitation struct {
	ID               int64     `db:"id" json:"id"`
	InvitaterID      int       `db:"invitater_id" json:"invitater_id"`
	ForType          string    `db:"for_type" json:"for_type"`
	IDOfType         int64     `db:"id_of_type" json:"id_of_type"`
	Message          string    `db:"message" json:"message"`
	TimeOfInvitation time.Time `db:"time_of_invitation" json:"time_of_invitation"`
}

// InvitationStatus ...
type InvitationStatus struct {
	InvitationTitle string `db:"invitation_title" json:"invitation_title"`
	Invitation
	InviteesID  int64  `db:"invitees_id" json:"invitees_id"`
	ReplyStatus string `db:"reply_status" json:"reply_status"`
}

// Insert ...
func (s *Invitation) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Invitation) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Invitation) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// GetInvitationStatusByUserID ...
func GetInvitationStatusByUserID(tx *gorp.Transaction, ID int) ([]InvitationStatus, error) {
	statusList := []InvitationStatus{}
	if _, err := tx.Select(&statusList, `select invitation.*,
		 invitees.id as invitees_id,
		 invitees.reply_status as reply_status,
		 events.title as invitation_title
		 from invitation
		 inner join invitees on invitation.id=invitees.invitation_id
		 inner join events on events.id=NULLIF(invitation.id_of_type, '0')::int
		  where invitees.invitee_id = $1`, ID); err != nil {
		return nil, err
	}
	return statusList, nil
}
