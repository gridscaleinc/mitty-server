package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Namecard ...
type Namecard struct {
	ID              int64     `db:"id" json:"id"`
	MittyID         int       `db:"mitty_id" json:"mitty_id"`
	Name            string    `db:"name" json:"name"`
	BusinessName    string    `db:"business_name" json:"business_name"`
	BusinessLogoID  int64     `db:"business_logo_id" json:"business_logo_id"`
	BusinessSubName string    `db:"business_sub_name" json:"business_sub_name"`
	BusinessTitle   string    `db:"business_title" json:"business_title"`
	AddressLine1    string    `db:"address_line1" json:"address_line1"`
	AddressLine2    string    `db:"address_line2" json:"address_line2"`
	Phone           string    `db:"phone" json:"phone"`
	Fax             string    `db:"fax" json:"fax"`
	MobilePhone     string    `db:"mobile_phone" json:"mobile_phone"`
	Webpage         string    `db:"webpage" json:"webpage"`
	Email           string    `db:"email" json:"email"`
	Created         time.Time `db:"created" json:"created"`
	Updated         time.Time `db:"updated" json:"updated"`
}

// NamecardInfo ...
type NamecardInfo struct {
	Namecard
	BusinessLogoURL string `db:"business_logo_url" json:"business_logo_url"`
}

// ContacteeNamecard ...
type ContacteeNamecard struct {
	NamecardID      int64      `db:"name_card_id" json:"name_card_id"`
	ContactID       int64      `db:"contact_id" json:"contact_id"`
	BusinessName    string     `db:"business_name" json:"business_name"`
	BusinessLogoURL string     `db:"business_logo_url" json:"business_logo_url"`
	RelatedEventID  int64      `db:"related_event_id" json:"related_event_id"`
	ContctedDate    *time.Time `db:"contacted_date" json:"contacted_date"`
}

// Save ...
func (s *Namecard) Save(tx gorp.Transaction) error {
	if s.ID == 0 {
		err := s.Insert(tx)
		return err
	}
	err := s.Update(tx)
	return err
}

// Insert ...
func (s *Namecard) Insert(tx gorp.Transaction) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Namecard) Update(tx gorp.Transaction) error {
	s.Updated = time.Now().UTC()
	_, err := tx.Update(s)
	return err
}

// GetNamecardByID ...
func GetNamecardByID(tx *gorp.Transaction, ID int64) (*NamecardInfo, error) {

	namecard := new(NamecardInfo)

	if err := tx.SelectOne(&namecard, `select namecard.*,
		COALESCE(contents.link_url, '') as business_logo_url
		from Namecard
		 left join Contents on Namecard.business_logo_id=Contents.id
		where namecard.id = $1;
		`, ID); err != nil {
		return nil, err
	}

	return namecard, nil
}

// Delete ...
func (s *Namecard) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// GetNamecardsByUserID ...
func GetNamecardsByUserID(tx *gorp.Transaction, ID int) ([]NamecardInfo, error) {
	results := []NamecardInfo{}

	if _, err := tx.Select(&results, `select namecard.*,
		COALESCE(contents.link_url, '') as business_logo_url
		from Namecard
		 left join Contents on Namecard.business_logo_id=Contents.id
		where mitty_id = $1;
		`, ID); err != nil {
		return nil, err
	}

	return results, nil
}

// GetContacteeNamecards ...
func GetContacteeNamecards(tx *gorp.Transaction, fromUserID int, contacteeUserID int) ([]ContacteeNamecard, error) {
	results := []ContacteeNamecard{}
	if _, err := tx.Select(&results, `select distinct Namecard.id as name_card_id,
		COALESCE(contact.id, 0) as contact_id,
		Namecard.business_name,
		COALESCE(contents.link_url, '') as business_logo_url,
		COALESCE(contact.related_event_id, 0) as related_event_id,
		contact.contacted_date
		from namecard
		   left join contact on namecard.id=contact.name_card_id and contact.mitty_id=$1
			 left join Contents on Namecard.business_logo_id=Contents.id
		where namecard.mitty_id=$2;
		`, fromUserID, contacteeUserID); err != nil {
		return nil, err
	}
	return results, nil
}
