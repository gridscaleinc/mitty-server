package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Conversation struct

type Conversation struct {
	ID                       int64     `db:"id" json:"id"`
	MeetingID           int64     `db:"meeting_id" json:"meetingId"`
	ReplyToID           int64     `db:"reply_to_id" json:"replyToId"`
    Speaking             string    `db:"speaking" json:"speaking"`
    SpeakerID           int64      `db:"speaker_id" json:"speakerId"`
    SpeakTime          time.Time  `db:"speak_time" json:"speakTime"`
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

// Delete
func (s *Conversation) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// Get Latest Conversation ...
func GetLatestConversation(tx *gorp.Transaction,  meetingID int64) ([]Conversation, error) {
	conversations := []Conversation{}
	_, err := tx.Select(&conversations, `select * from conversation where meeting_id=$1 order by speak_time desc;`, meetingID)
	return conversations, err
}
