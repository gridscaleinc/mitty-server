package models

import (
	"time"
    sql  "database/sql"
    pg "github.com/lib/pq"
	gorp "gopkg.in/gorp.v1"
)

// Activity struct
type Activity struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	MainEventID sql.NullInt64     `db:"main_event_id" json:"main_event_id"`
	Memo        string    `db:"memo" json:"memo"`
	OwnerID     int       `db:"owner_id" json:"owner_id"`
	Created     time.Time `db:"created" json:"created"`
	Updated     time.Time `db:"updated" json:"updated"`
}

// ActivityList ...
type ActivityList struct {
	ID            int        `db:"id" json:"id"`                       //  Activity のID
	EventID   sql.NullInt64        `db:"eventId" json:"eventId"`             // ActivityのMainEventId
	Title         string     `db:"title" json:"title"`                 // ActivityのTitle
	StartDateTime *time.Time `db:"startDateTime" json:"startDateTime"` // MainEventのstart_datetime
	EventLogoURL  *string    `db:"eventLogoUrl" json:"eventLogoUrl"`   // MainEventのLogoIDから結びつけるContentsのLinkURL
}

// ActivityDetail ...
type ActivityDetail struct {
	ID               int64     `db:"id" json:"id"`
	MainEventID      sql.NullInt64     `db:"main_event_id" json:"main_event_id"`
	Title            string    `db:"title" json:"title"`
	Memo             string    `db:"memo" json:"memo"`
	EventID          sql.NullInt64     `db:"eventId" json:"eventId"`
	Notification     sql.NullBool      `db:"notification" json:"notification"`
	NotificationTime pg.NullTime `db:"notificationTime" json:"notificationTime"`
	EventTitle       string    `db:"eventTitle" json:"eventTitle"`
	StartDateTime    pg.NullTime `db:"startDateTime" json:"startDateTime"`
	EndDateTime     pg.NullTime `db:"endDateTime" json:"endDateTime"`
	AllDayFlag       sql.NullBool      `db:"allDayFlag" json:"allDayFlag"`
	EventLogoURL     sql.NullString    `db:"eventLogoUrl" json:"eventLogoUrl"`
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
		   a.id,
		   a.title,
		   a.memo,
		   a.main_event_id,
		   i.event_Id as eventId,
		   i.notification,
		   notificationdatetime as notificationTime,
		   e.start_datetime as startDateTime,
		   e.end_datetime as endDateTime,
		   e.allday_flag as allDayFlag,
		   c.link_url as eventLogoUrl
		from
		   activity as a
		   left join activity_item as i on a.id=i.activity_id
		   left outer join events as e on i.event_id=e.id
		   left outer join contents as c on e.logo_id=c.id
		where
		   a.id=$1
		   and
		   a.owner_id=$2;
		`, id, userID)
	return details, err

}
