package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Meeting struct
type Meeting struct {
	ID      int64     `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	Type    string    `db:"type" json:"type"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (s *Meeting) Insert(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Meeting) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}

type MeetingInfo struct {
	ID      int64     `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	Type    string    `db:"type" json:"type"`
	Title string `db:"title" json:"title"`
	LogoUrl string `db:"logo_url" json:"logoUrl"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

// Get Latest Conversation ...
func GetEventMeetingList(tx *gorp.Transaction,  userID int64) ([]MeetingInfo, error) {
	eventMeeting := []MeetingInfo{}
	_, err := tx.Select(&eventMeeting, `
	select 
	    meeting.*,
	    events.title as title,
	    COALESCE(contents.link_url, '') as logo_url
	from 
	    meeting inner join events on events.meeting_id=meeting.id 
	    inner join activity_item on activity_item.event_id=events.id 
	    inner join activity on activity.id=activity_item.activity_id 
	    left outer join contents on contents.id=events.logo_id 
	where 
	   activity.owner_id=$1 
	order by 
	   events.start_datetime, events.end_datetime;`, userID)
	return eventMeeting, err
}
