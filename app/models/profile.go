package models

import (
	gorp "gopkg.in/gorp.v1"
)

// Profile ...
type Profile struct {
	ID             int64  `db:"id" json:"id"`
	MittyID        int    `db:"mitty_id" json:"mitty_id"`
	Gender         string `db:"gender" json:"gender"`
	OneWordSpeech  string `db:"one_word_speech" json:"one_word_speech"`
	Constellation  string `db:"constellation" json:"constellation"`
	HomeIslandID   int64  `db:"home_island_id" json:"home_island_id"`
	BirthIslandID  int64  `db:"birth_island_id" json:"birth_island_id"`
	AgeGroup       string `db:"age_group" json:"age_group"`
	AppearanceTag  string `db:"appearance_tag" json:"appearance_tag"`
	OccupationTag1 string `db:"occupation_tag1" json:"occupation_tag1"`
	OccupationTag2 string `db:"occupation_tag2" json:"occupation_tag2"`
	OccupationTag3 string `db:"occupation_tag3" json:"occupation_tag3"`
	HobbyTag1      string `db:"hobby_tag1" json:"hobby_tag1"`
	HobbyTag2      string `db:"hobby_tag2" json:"hobby_tag2"`
	HobbyTag3      string `db:"hobby_tag3" json:"hobby_tag3"`
	HobbyTag4      string `db:"hobby_tag4" json:"hobby_tag4"`
	HobbyTag5      string `db:"hobby_tag5" json:"hobby_tag5"`
}

// Contactee ...
type Contactee struct {
	ContacteeName string `db:"contactee_name" json:"contactee_name"`
	Profile
}

// Save ...
func (s *Profile) Save(tx gorp.Transaction) error {
	if s.ID == 0 {
		err := s.Insert(tx)
		return err
	}
	err := s.Update(tx)
	return err
}

// Insert ...
func (s *Profile) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Profile) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Profile) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// GetProfileByUserID ...
func GetProfileByUserID(tx *gorp.Transaction, ID int) (*Profile, error) {
	result := new(Profile)
	if err := tx.SelectOne(&result, `select *
		from Profile
		where mitty_id = $1;
		`, ID); err != nil {
		return nil, err
	}
	return result, nil
}

// GetContacteeListByUserID ...
func GetContacteeListByUserID(tx *gorp.Transaction, userID int) ([]Contactee, error) {
	contacteeList := []Contactee{}
	_, err := tx.Select(&contacteeList, `
		select
		   users.user_name as contactee_name,
			 profile.*
		from
			 users
			 inner join profile on users.id=profile.mitty_id
		where
			 users.id in (
				 select namecard.mitty_id
				 from contact
				 inner join namecard on namecard.id=contact.name_card_id
				 where contact.mitty_id=$1);
		`, userID)
	return contacteeList, err
}
