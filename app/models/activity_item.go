package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// ActivityItem struct
type ActivityItem struct {
	ID                   int64     `db:"id" json:"id"`
	ActivityID           int64       `db:"activity_id" json:"activity_id"`
	EventID              int64     `db:"event_id" json:"event_id"`
	Title                string    `db:"title" json:"title"`
	Memo                 string    `db:"memo" json:"memo"`
	Notification         bool      `db:"notification" json:"notification"`
	NotificationDateTime time.Time `db:"notificationDateTime" json:"notificationDateTime"`
	Participation		   string    `db:"participation" json:"participation"`
    CalendarUrl		   string    `db:"calendar_url" json:"calendar_url"`
	Created              time.Time `db:"created" json:"created"`
	Updated              time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (s *ActivityItem) Insert(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *ActivityItem) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}
