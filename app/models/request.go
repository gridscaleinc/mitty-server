package models

import (
	"strconv"
	"time"

	"mitty.co/mitty-server/app/helpers"

	gorp "gopkg.in/gorp.v1"
)

// Request struct
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
	NumOfPerson        int       `db:"num_of_person" json:"num_of_person"`
	NumOfChildren      int       `db:"num_of_children" json:"num_of_children"`
	AcceptedProposalID int       `db:"accepted_proposal_id" json:"accepted_proposal_id"`
	AcceptedDate       time.Time `db:"accepted_date" json:"accepted_date"`
	MeetingID          int64     `db:"meeting_id" json:"meeting_id"`
	OwnerID            int       `db:"owner_id" json:"owner_id"`
	Created            time.Time `db:"created" json:"created"`
	Updated            time.Time `db:"updated" json:"updated"`
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
