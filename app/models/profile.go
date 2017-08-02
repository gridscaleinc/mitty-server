package models

// Profile ...
type Profile struct {
	ID             int64  `db:"id" json:"id"`
	MittyID        int    `db:"mitty_id" json:"mitty_id"`
	Gender         string `db:"gender" json:"gender"`
	OneWordSpeech  string `db:"one_word_speech" json:"one_word_speech"`
	Constellation  string `db:"constellation" json:"constellation"`
	HomeIslandID   int8   `db:"home_island_id" json:"home_island_id"`
	BirthIslandID  int8   `db:"birth_island_id" json:"birth_island_id"`
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
