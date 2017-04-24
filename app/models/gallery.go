package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Gallery struct
type Gallery struct {
	ID        int64     `db:"id" json:"id"`
	Seq       int       `db:"seq" json:"seq"`
	Caption   string    `db:"caption" json:"caption"`
	BriefInfo string    `db:"brief_info" json:"brief_info"`
	ContentID int64     `db:"content_id" json:"content_id"`
	FreeText  string    `db:"free_text" json:"free_text"`
	Created   time.Time `db:"created" json:"created"`
	Updated   time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (s *Gallery) Insert(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Gallery) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}
