package models
    
import (
    	"time"
    
    	gorp "gopkg.in/gorp.v1"
)
    
// Footmark struct
type Footmark struct {
    ID    int64     `db:"id" json:"id"`
	EventID	int64	 `db:"event_id" json:"event_id"`
	IslandID		int64  `db:"island_id" json:"island_id"`
	MittyID		int	  `db:"mitty_id" json:"mitty_id"`
	NameCardID	 	int64  `db:"name_card_id" json:"name_card_id"`
	PictureID		int64  `db:"picture_id" json:"picture_id"`
	SeatOrRoomInfo		string `db:"seat_or_room_info" json:"seat_or_room_info"`
	CheckinTime		time.Time `db:"checkin_time" json:"checkin_time"`
}	

// Insert ...
func (s *Footmark) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Footmark) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Footmark) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}