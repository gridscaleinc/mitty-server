package models

import (
	"strconv"
	"time"

	"mitty.co/mitty-server/app/helpers"

	gorp "gopkg.in/gorp.v1"
)

// Request ...
type Request struct {
	ID                 int64     `db:"id" json:"id"`
	Title              string    `db:"title" json:"title"`
	Tag                string    `db:"tag" json:"tag"`
	Description        string    `db:"description" json:"description"`
	ForActivityID      int       `db:"for_activity_id" json:"for_activity_id"`
	PreferredDatetime1 time.Time `db:"preferred_datetime1" json:"preferred_datetime1"`
	PreferredDatetime2 time.Time `db:"preferred_datetime2" json:"preferred_datetime2"`
	PreferredPrice1    int       `db:"preferred_price1" json:"preferred_price1"`
	PreferredPrice2    int       `db:"preferred_price2" json:"preferred_price2"`
	StartPlace         string    `db:"start_place" json:"start_place"`
	TerminatePlace     string    `db:"terminate_place" json:"terminate_place"`
	Oneway             bool      `db:"oneway" json:"oneway"`
	Status             string    `db:"status" json:"status"`
	ExpiryDate         time.Time `db:"expiry_date" json:"expiry_date"`
	GalleryID          int64     `db:"gallery_id" json:"gallery_id"`
	NumOfPerson        int       `db:"num_of_person" json:"num_of_person"`
	NumOfChildren      int       `db:"num_of_children" json:"num_of_children"`
	AcceptedProposalID int       `db:"accepted_proposal_id" json:"accepted_proposal_id"`
	AcceptedDate       time.Time `db:"accepted_date" json:"accepted_date"`
	MeetingID          int64     `db:"meeting_id" json:"meeting_id"`
	OwnerID            int       `db:"owner_id" json:"owner_id"`
	Created            time.Time `db:"created" json:"created"`
	Updated            time.Time `db:"updated" json:"updated"`
	Currency           *string   `db:"currency" json:"currency"`
}

//RequestInfo ...
type RequestInfo struct {
	Request
	NumOfLikes    int     `db:"num_of_likes" json:"num_of_likes"`
	NumOfProposal int     `db:"num_of_proposal" json:"num_of_proposal"`
	OwnerName     *string `db:"owner_name" json:"owner_name"`
	OwnerIconURL  *string `db:"owner_icon_url" json:"owner_icon_url"`
}

// Save ...
func (s *Request) Save(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	if err == nil {
		go func() {
			helpers.ESIndex("mitty", "request", s, strconv.FormatInt(s.ID, 10))
		}()
	}
	return err
}

// Update ...
func (s *Request) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	if err == nil {
		go func() {
			helpers.ESIndex("mitty", "request", s, strconv.FormatInt(s.ID, 10))
		}()
	}
	return err
}

// Delete ...
func (s *Request) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	if err == nil {
		go func() {
			helpers.ESDelete("mitty", "request", strconv.FormatInt(s.ID, 10))
		}()
	}
	return err
}

// GetRequestDetailByID ...
func GetRequestDetailByID(tx *gorp.Transaction, ID int64) (*RequestInfo, error) {

	requestDetail := new(RequestInfo)
	if err := tx.SelectOne(&requestDetail, `select request.*,
		(select count(id) from likes where entity_type='REQUEST' and entity_id=$1) as num_of_likes,
		(select count(id) from Proposal where reply_to_request_id=$1) as num_of_proposal,
		users.name as owner_name,
		users.icon as owner_icon_url
		from request
		join users on users.id = request.owner_id
		where request.id = $1;
		`, ID); err != nil {
		return nil, err
	}
	return requestDetail, nil
}

// GetRequestByUserID ...
func GetRequestByUserID(tx *gorp.Transaction, userID int, key string) ([]Request, error) {
	requests := []Request{}
	_, err := tx.Select(&requests, `
		select
      *
      from request
      where owner_id=$1 and title like '%$2' or action like '%$2'
		`, userID, key)
	return requests, err
}
