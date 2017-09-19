package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Offers ...
type Offers struct {
	ID               int64     `db:"id" json:"id"`
	FromMittyID      int       `db:"from_mitty_id" json:"from_mitty_id"`
	ToMittyID        int       `db:"to_mitty_id" json:"to_mitty_id"`
	Type             string    `db:"type" json:"type"`
	Message          string    `db:"message" json:"message"`
	ReplyStatus      string    `db:"reply_status" json:"reply_status"`
	OfferredID       int64     `db:"offerred_id" json:"offerred_id"`
	RepliedID        int64     `db:"replied_id" json:"replied_id"`
	OfferredDatetime time.Time `db:"offerred_datetime" json:"offerred_datetime"`
}

// Insert ...
func (s *Offers) Insert(tx gorp.Transaction) error {
	s.OfferredDatetime = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Offers) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Offers) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// Load ...
func (s *Offers) Load(tx gorp.Transaction) error {
	err := tx.SelectOne(&s, `select * from offers where id=$1;`, s.ID)
	return err
}

// GetOfferListByUserID ...
func GetOfferListByUserID(tx *gorp.Transaction, userID int) ([]Offers, error) {
	myOffers := []Offers{}
	_, err := tx.Select(&myOffers, `
	select
			*
	from
			offers
	where
		 to_mitty_id = $1 and reply_status='NONE';`, userID)
	return myOffers, err
}

// CountOfNamecardOfferByUserID ...
func CountOfNamecardOfferByUserID(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := tx.SelectInt(`select count(*) from offers
	    where to_mitty_id=$1 and reply_status='NONE';`, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}
