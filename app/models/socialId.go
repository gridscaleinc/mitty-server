package models

import gorp "gopkg.in/gorp.v1"

// SocialID ...
type SocialID struct {
	ID       int64 `db:"id" json:"id"`
	MittyID  int   `db:"mitty_id" json:"mitty_id"`
	SocialID int64 `db:"social_id" json:"social_id"`
}

// Insert ...
func (s *SocialID) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *SocialID) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *SocialID) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}
