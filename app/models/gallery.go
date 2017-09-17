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

// GalleryContent ...
type GalleryContent struct {
	GalleryID int64  `db:"gallery_id" json:"gallery_id"`
	Seq       int    `db:"seq" json:"seq"`
	Caption   string `db:"caption" json:"caption"`
	BriefInfo string `db:"brief_info" json:"brief_info"`
	FreeText  string `db:"free_text" json:"free_text"`
	Contents
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

// GetGalleryByID ...
func GetGalleryByID(tx *gorp.Transaction, ID int64) (*Gallery, error) {
	gallery := new(Gallery)
	if err := tx.SelectOne(&gallery, "select * from gallery where id = $1", ID); err != nil {
		return nil, err
	}
	return gallery, nil
}

// GetGalleryContentsByID ...
func GetGalleryContentsByID(tx *gorp.Transaction, ID int64) (*[]GalleryContent, error) {
	contents := []GalleryContent{}
	_, err := tx.Select(&contents, `
	select
		  gallery.id as gallery_id,
			gallery.seq,
			gallery.caption,
			gallery.brief_info,
			gallery.free_text,
			contents.*
	from
			contents
			inner join gallery on gallery.id=$1 and gallery.content_id=contents.id
	order by
		  gallery.seq;`, ID)
	return &contents, err
}
