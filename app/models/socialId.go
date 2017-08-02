package models

// SocialID ...
type SocialID struct {
	ID       int64 `db:"id" json:"id"`
	MittyID  int   `db:"mitty_id" json:"mitty_id"`
	SocialID int64 `db:"social_id" json:"social_id"`
}
