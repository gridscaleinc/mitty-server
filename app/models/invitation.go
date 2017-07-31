package models
    
import (
	    "time"
	    
    	gorp "gopkg.in/gorp.v1"
)
    
// Likes struct
type Invitation struct {
    ID    int64     `db:"id" json:"id"`
    InvitaterID    int   `db:"invitater_id" json:"invitater_id"`
    ForType       string `db:"for_type" json:"for_type"`
    IDOfType    int64     `db:"entity_id" json:"id_of_type"`  
    Message    string `db:"message" json:"message"`
    TimeOfInvitation    time.Time `db:"time_of_invitation" json:"time_of_invitation"`
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
