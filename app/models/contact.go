package models

import "time"

// Contact ...
type Contact struct {
	ID             int64     `db:"id" json:"id"`
	RelatedEventID int64     `db:"related_event_id" json:"related_event_id"`
	NameCardID     int64     `db:"name_card_id" json:"name_card_id"`
	ContctedDate   time.Time `contacted_date:"id" json:"contacted_date"`
}
