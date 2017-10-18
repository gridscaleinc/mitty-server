package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Conversation ...
type Conversation struct {
	ID        int64     `db:"id" json:"id"`
	MeetingID int64     `db:"meeting_id" json:"meetingId"`
	ReplyToID int64     `db:"reply_to_id" json:"replyToId"`
	Type      string    `db:"type" json:"type"`
	Speaking  string    `db:"speaking" json:"speaking"`
	SpeakerID int64     `db:"speaker_id" json:"speakerId"`
	SpeakTime time.Time `db:"speak_time" json:"speakTime"`
	Latitude  float64   `db:"latitude" json:"latitude"`
	Longitude float64   `db:"longitude" json:"longitude"`
}

// Insert ...
func (s *Conversation) Insert(tx gorp.Transaction) error {
	s.SpeakTime = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Conversation) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Conversation) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// GetLatestConversation ...
func GetLatestConversation(tx *gorp.Transaction, meetingID int64) ([]Conversation, error) {
	conversations := []Conversation{}
	_, err := tx.Select(&conversations, `select * from conversation where meeting_id=$1 order by speak_time;`, meetingID)
	return conversations, err
}
