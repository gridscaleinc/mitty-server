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

type NamecardInfo struct {
	CardInfo        Namecard `json:"cardInfo"`
	BusinessLogoUrl string   `db:"business_logo_url" json:"businessLogoUrl"`
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

// Delete ...
func (s *Namecard) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// GetNamecardsByUserID ...
func GetNamecardsByUserID(tx *gorp.Transaction, ID int) (*[]NamecardInfo, error) {
	results := []NamecardInfo{}

	if err := tx.SelectOne(&results, `select *,
		COALESCE(contents.link_url, '') as business_logo_url
		from Namecard
		 left join Contents on Namecard.business_logo_id=Contents.id
		where mitty_id = $1;
		`, ID); err != nil {
		return nil, err
	}

	return &results, nil
}
