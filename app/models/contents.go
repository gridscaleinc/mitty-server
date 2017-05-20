package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Contents struct
type Contents struct {
	ID      int64     `db:"id" json:"id"`
	Mime    string    `db:"mime" json:"mime"`
	Name    string    `db:"name" json:"name"`
	LinkURL string    `db:"link_url" json:"link_url"`
	Width   int       `db:"width" json:"width"`
	Height  int       `db:"height" json:"height"`
	Size    int       `db:"size" json:"size"`
	OwnerID int       `db:"owner_id" json:"owner_id"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (s *Contents) Insert(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Contents) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}

// GetContentsByUserID ...
func GetContentsByUserID(tx *gorp.Transaction, userID int) ([]Contents, error) {
	contents := []Contents{}
	if _, err := tx.Select(&contents, "select * from contents where owner_id = $1", userID); err != nil {
		return nil, err
	}
	return contents, nil
}
