package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Activity struct
type Activity struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	MainEventID int64     `db:"main_event_id" json:"main_event_id"`
	Memo        string    `db:"memo" json:"memo"`
	OwnerID     int       `db:"owner_id" json:"owner_id"`
	Created     time.Time `db:"created" json:"created"`
	Updated     time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (s *Activity) Insert(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Activity) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}

// GetActivityByMainEventID ...
func GetActivityByMainEventID(tx *gorp.Transaction, ID int) (*Activity, error) {
	activity := new(Activity)
	if err := tx.SelectOne(&activity, "select * from activity where main_event_id = $1", ID); err != nil {
		return nil, err
	}
	return activity, nil
}

// GetActivityByID ...
func GetActivityByID(tx *gorp.Transaction, ID int) (*Activity, error) {
	activity := new(Activity)
	if err := tx.SelectOne(&activity, "select * from activity where id = $1", ID); err != nil {
		return nil, err
	}
	return activity, nil
}
