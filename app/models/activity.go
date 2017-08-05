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

// ActivityList ...
type ActivityList struct {
	ID            int        `db:"id" json:"id"`                       //  Activity のID
	EventID       int        `db:"eventId" json:"eventId"`             // ActivityのMainEventId
	Title         string     `db:"title" json:"title"`                 // ActivityのTitle
	StartDateTime *time.Time `db:"startDateTime" json:"startDateTime"` // MainEventのstart_datetime
	EventLogoURL  *string    `db:"eventLogoUrl" json:"eventLogoUrl"`   // MainEventのLogoIDから結びつけるContentsのLinkURL
}

// ActivityDetail ...
type ActivityDetail struct {
	ID               int64     `db:"id" json:"id"`
	ActivityID       int64     `db:"activity_id" json:"activity_id"`
	Title            string    `db:"title" json:"title"`
	Memo             string    `db:"memo" json:"memo"`
	EventID          int64     `db:"eventId" json:"eventId"`
	Notification     bool      `db:"notification" json:"notification"`
	NotificationTime time.Time `db:"notificationTime" json:"notificationTime"`
	EventTitle       string    `db:"eventTitle" json:"eventTitle"`
	StartDateTime    time.Time `db:"startDateTime" json:"startDateTime"`
	EndDateTime      time.Time `db:"endDateTime" json:"endDateTime"`
	AllDayFlag       bool      `db:"allDayFlag" json:"allDayFlag"`
	EventLogoURL     string    `db:"eventLogoUrl" json:"eventLogoUrl"`
	IslandName       string    `db:"islandName" json:"islandName"`
	IslandNickname   string    `db:"islandNickname" json:"islandNickname"`
	IslandLogoURL    string    `db:"islandLogoUrl" json:"islandLogoUrl"`
}

// Destination ...
type Destination struct {
	IslandID       int       `db:"island_id" json:"islandId"`
	IslandNickName string    `db:"island_nickname" json:"IslandNickName"`
	IslandName     string    `db:"island_name" json:"islandName"`
	Latitude       float64   `db:"latitude" json:"latitude"`
	Longitude      float64   `db:"longitude" json:"longitude"`
	IslandLogo     string    `db:"island_logo" json:"islandLogo"`
	EventID        int       `db:"event_id" json:"eventId"`
	EventTitle     string    `db:"event_title" json:"eventTitle"`
	EventTime      time.Time `db:"event_time" json:"eventTime"`
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
func GetActivityByID(tx *gorp.Transaction, ID int64) (*Activity, error) {
	activity := new(Activity)
	if err := tx.SelectOne(&activity, "select * from activity where id = $1", ID); err != nil {
		return nil, err
	}
	return activity, nil
}

// GetActivityListByKey ...
func GetActivityListByKey(tx *gorp.Transaction, userID int, key string) ([]ActivityList, error) {
	activities := []ActivityList{}
	_, err := tx.Select(&activities, `
		select
      a.id,
      a.main_event_id as eventId,
      a.title,
      e.start_datetime as startDateTime,
      c.link_url as eventLogoUrl
		from activity as a
	    left outer join events as e on a.main_event_id=e.id
	    left outer join contents as c on e.logo_id=c.id
		where
		  a.owner_id=$1 and
			(a.title like '%' || $2 || '%' or a.memo like '%' || $2 || '%');
		`, userID, key)
	return activities, err
}

// GetActivityDetailsByID ...
func GetActivityDetailsByID(tx *gorp.Transaction, userID int, id string) ([]ActivityDetail, error) {
	details := []ActivityDetail{}
	_, err := tx.Select(&details, `
		select
		   i.id as id,
			 i.activity_id,
		   i.event_Id as eventId,
		   i.title,
		   i.memo,
		   i.notification,
		   notificationdatetime as notificationTime,
		   e.title as eventTitle,
		   e.start_datetime as startDateTime,
		   e.end_datetime as endDateTime,
		   e.allday_flag as allDayFlag,
		   COALESCE(c.link_url, '') as eventLogoUrl,
		   l.name as islandName,
		   l.nickname as islandNickname,
		   COALESCE(c2.link_url, '') as islandLogoUrl
		from
		   activity as a
		   inner join activity_item as i on a.id=i.activity_id
		   inner join events as e on i.event_id=e.id
		   inner join island as l on e.islandid=l.id
		   left outer join contents as c on e.logo_id=c.id
		   left outer join contents as c2 on l.logo_id=c2.id
		where
		   a.id=$1
		   and
		   a.owner_id=$2;
		`, id, userID)
	return details, err

}

// GetDestinationList ...
func GetDestinationList(tx *gorp.Transaction, userID int) ([]Destination, error) {
	destinations := []Destination{}
	_, err := tx.Select(&destinations, `
		select
      island.id as island_id,
      island.nickname as island_nickname,
      island.name as island_name,
      island.latitude,
      island.longitude,
      COALESCE(contents.link_url, '') as island_logo,
      events.id as event_id,
      events.title as event_title,
      events.start_datetime as event_time
  from island
      left join contents on island.logo_id=contents.id
      inner join events on island.id=events.islandid
      inner join activity_item on activity_item.event_id=events.id
      inner join activity on activity.id=activity_item.activity_id
  where activity.owner_id=$1
  order by events.start_datetime;
		`, userID)
	return destinations, err
}

// Save ...
func (s *Activity) Save(tx *gorp.Transaction) error {
	if _, err := tx.Exec("Update Activity set title=$3, memo=$4 WHERE owner_id = $2 and id=$1", s.ID, s.OwnerID, s.Title, s.Memo); err != nil {
		return err
	}
	return nil
}
