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

// MeetingInfo ...
type MeetingInfo struct {
	ID      int64     `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	Type    string    `db:"type" json:"type"`
	Title   string    `db:"title" json:"title"`
	LogoURL string    `db:"logo_url" json:"logoUrl"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

// GetEventMeetingList ...
func GetEventMeetingList(tx *gorp.Transaction, userID int) ([]MeetingInfo, error) {
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

// GetRequestMeetingList ...
func GetRequestMeetingList(tx *gorp.Transaction, userID int) ([]MeetingInfo, error) {
	requestMeeting := []MeetingInfo{}
	_, err := tx.Select(&requestMeeting, `
	select
	    meeting.*,
	    request.title as title,
	    '' as logo_url
	from
	    meeting inner join request on request.meeting_id=meeting.id
	where
	   request.owner_id=$1
		 or request.id in (select reply_to_request_id from proposal where proposer_id=$1)
	order by
	   request.preferred_datetime1;`, userID)
	return requestMeeting, err
}

// GetContactMeetingList ...
func GetContactMeetingList(tx *gorp.Transaction, userID int) ([]MeetingInfo, error) {
	contactMeeting := []MeetingInfo{}
	_, err := tx.Select(&contactMeeting, `
	select
	    meeting.*,
	    namecard.name as title,
	    COALESCE(users.icon, '') as logo_url
	from
	    meeting inner join contact on meeting.id=contact.meeting_id
			inner join namecard on namecard.id=contact.name_card_id
			inner join users on users.id=namecard.mitty_id
	where
	   contact.mitty_id=$1
	order by
	   namecard.name;`, userID)
	return contactMeeting, err
}
