package models
    
import (
    	"time"
    
    	gorp "gopkg.in/gorp.v1"
)
    
// Proposal struct
type Proposal struct {
    ID		int64     `db:"id" json:"id"`
    ReplyToRequestID		int64     `db:"reply_to_request_id" json:"reply_to_request_id"`
    ContactTel		string     `db:"contact_tel" json:"contact_tel"`
    ContactEmail		string     `db:"contact_email" json:"contact_email"`
    ProposedIslandID		int64     `db:"proposed_island_id" json:"proposed_island_id"`
    ProposedIslandID2		int64     `db:"proposed_island_id2" json:"proposed_island_id2"`
    GalleryID		int64     `db:"gallery_id" json:"gallery_id"`
    PriceName1		string     `db:"priceName1" json:"priceName1"`
    Price1		int     `db:"price1" json:"price1"`
    PriceName2		string     `db:"priceName2" json:"priceName2"`
    Price2		int     `db:"price2" json:"price2"`
    PriceCurrency		string     `db:"price_currency" json:"price_currency"`
    PriceInfo		string     `db:"price_info" json:"price_info"`
    ProposedDatetime1		time.Time `db:"proposed_datetime1" json:"proposed_datetime1"`
    ProposedDatetime2		time.Time `db:"proposed_datetime2" json:"proposed_datetime2"`
    AdditionalInfo		string     `db:"additional_info" json:"additional_info"`
    ProposerID		int     `db:"proposer_id" json:"proposer_id"`
    ProposerInfo		string     `db:"proposer_info" json:"proposer_info"`
    ProposedDatetime		time.Time `db:"proposed_datetime" json:"proposed_datetime"`
    AcceptStatus		string     `db:"accept_status" json:"accept_status"`
    AcceptDatetime		time.Time `db:"accept_datetime" json:"accept_datetime"`
    ConfirmTel		string     `db:"confirm_tel" json:"confirm_tel"`
    ConfirmEmail		string     `db:"confirm_email" json:"confirm_email"`
    ApprovalStatus		string     `db:"approval_status" json:"approval_status"`
    ApprovalDate		time.Time `db:"approval_date" json:"approval_date"`
}	

// Insert ...
func (s *Proposal) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Proposal) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Proposal) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}