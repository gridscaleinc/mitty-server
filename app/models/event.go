package models

import (
	"strconv"
	"time"

	"mitty.co/mitty-server/app/helpers"

	gorp "gopkg.in/gorp.v1"
)

// Event struct
type Event struct {
	ID                  int64     `db:"id" json:"id"`
	Type                string    `db:"type" json:"type"`
	Category            string    `db:"category" json:"category"`
	Theme               string    `db:"theme" json:"theme"`
	Tag                 string    `db:"tag" json:"tag"`
	Title               string    `db:"title" json:"title"`
	Action              string    `db:"action" json:"action"`
	StartDatetime       time.Time `db:"start_datetime" json:"startDatetime"`
	EndDatetime         time.Time `db:"end_datetime" json:"endDatetime"`
	AlldayFlag          bool      `db:"allday_flag" json:"alldayFlag"`
	IslandID            int       `db:"islandId" json:"islandId"`
	IslandID2            int       `db:"islandId2" json:"islandId2"`
	LogoID              int       `db:"logo_id" json:"logoId"`
	GalleryID           int64     `db:"gallery_id" json:"galleryId"`
	MeetingID           int64     `db:"meeting_id" json:"meetingId"`
	PriceName1          string    `db:"price_name1" json:"priceName1"`
	Price1              int       `db:"price1" json:"price1"`
	PriceName2          string    `db:"price_name2" json:"priceName2"`
	Price2              int       `db:"price2" json:"price2"`
	Currency            string    `db:"currency" json:"currency"`
	PriceInfo           string    `db:"price_info" json:"priceInfo"`
	Description         string    `db:"description" json:"description"`
	ContactTel          string    `db:"contact_tel" json:"contactTel"`
	ContactFax          string    `db:"contact_fax" json:"contactFax"`
	ContactMail         string    `db:"contact_mail" json:"contactMail"`
	OfficialURL         string    `db:"official_url" json:"officialUrl"`
	Organizer           string    `db:"organizer" json:"organizer"`
	SourceName          string    `db:"source_name" json:"sourceName"`
	SourceURL           string    `db:"source_url" json:"sourceUrl"`
	NumberOfAnticipants int       `db:"number_of_anticipants" json:"numberOfAnticipants"`
	Participation       string    `db:"participation" json:"participation"`
	AccessControl       string    `db:"access_control" json:"accessControl"`
	Likes               int       `db:"likes" json:"likes"`
	Status              string    `db:"status" json:"status"`
	Language            string    `db:"language" json:"language"`
	Created             time.Time `db:"created" json:"created"`
	PublisherID         int       `db:"publisher_id" json:"publisherId"`
	OrgnizationID       int       `db:"orgnization_id" json:"orgnizationId"`
	Lastupdated         time.Time `db:"lastupdated" json:"lastupdated"`
	AmenderID           int       `db:"amender_id" json:"amenderId"`
}

// Save ...
func (s *Event) Save(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Lastupdated = time.Now().UTC()
	err := tx.Insert(s)
	if err == nil {
		go func() {
			helpers.ESIndex("mitty", "event", s, strconv.FormatInt(s.ID, 10))
		}()
	}
	return err
}

// Update ...
func (s *Event) Update(tx gorp.Transaction) error {
	s.Lastupdated = time.Now().UTC()
	_, err := tx.Update(s)
	if err == nil {
		go func() {
			helpers.ESIndex("mitty", "event", s, strconv.FormatInt(s.ID, 10))
		}()
	}
	return err
}

// Delete ...
func (s *Event) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	if err == nil {
		go func() {
			helpers.ESDelete("mitty", "event", strconv.FormatInt(s.ID, 10))
		}()
	}
	return err
}

// GetAdminEvents ...
func GetAdminEvents(dbmap *gorp.DbMap) ([]Event, error) {
	events := []Event{}
	_, err := dbmap.Select(&events, "select * from events")
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetEventByID ...
func GetEventByID(tx *gorp.Transaction, ID int) (*Event, error) {
	event := new(Event)
	if err := tx.SelectOne(&event, "select * from events where id = $1", ID); err != nil {
		return nil, err
	}
	return event, nil
}

// GetEventDetailByID ...
func GetEventDetailByID(tx *gorp.Transaction, userID int, ID int) (interface{}, error) {
	type result struct {
		Event
		CoverImageUrl    *string `db:"cover_img_url" json:"coverImageUrl"`
		EventLogoUrl      *string `db:"event_logo_url" json:"eventLogoUrl"`
		IsLandName          *string `db:"island_name" json:"isLandName"`
		IsLandLogoURL       *string `db:"island_logo_url" json:"isLandLogoUrl"`
	    Latitude      float64   `db:"latitude" json:"latitude"`
	    Longitude     float64   `db:"longitude" json:"longitude"`
		PublisherName       *string `db:"publisher_name" json:"publisherName"`
		PublisherIconURL    *string `db:"publisher_icon_url" json:"publisherIconUrl"`
		PublishedDays       int     `db:"published_days" json:"publishedDays"`
		ParticipationStatus string    `db:"participation_status" json:"participationStatus"`
	}

	eventDetail := new(result)
	if err := tx.SelectOne(&eventDetail, `select events.*,
	    COALESCE(c1.link_url, '') as cover_img_url,
	    COALESCE(c2.link_url, '') as event_logo_url,
		island.name as island_name,
		COALESCE(c3.link_url, '') as island_logo_url,
		COALESCE(island.latitude, 999) as latitude,
		COALESCE(island.longitude, 999) as longitude,		
		users.name as publisher_name,
		users.icon as publisher_icon_url,
		DATE 'now' - events.created as published_days,
		COALESCE(actitem.participation, 'NOT') as participation_status
	from events
		left join gallery on events.gallery_id=gallery.id
		left join contents as c1 on gallery.content_id=c1.id and gallery.seq=0
		left join contents as c2 on events.logo_id=c2.id
		inner join island on island.id = events.islandid
		left join contents as c3 on c3.id = island.logo_id
		left join users on users.id = events.publisher_id
		left join (select item.event_id, item.participation from activity, activity_item as item
		    where item.activity_id=activity.id and activity.owner_id=$1 ) as actitem
		   on actitem.event_id = events.id
		where events.id = $2;
		`, userID, ID); err != nil {
		return nil, err
	}
	return eventDetail, nil
}
