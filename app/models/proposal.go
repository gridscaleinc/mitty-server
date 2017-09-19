package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Proposal struct
type Proposal struct {
	ID                int64     `db:"id" json:"id"`
	ReplyToRequestID  int64     `db:"reply_to_request_id" json:"reply_to_request_id"`
	ContactTel        string    `db:"contact_tel" json:"contact_tel"`
	ContactEmail      string    `db:"contact_email" json:"contact_email"`
	ProposedIslandID  int64     `db:"proposed_island_id" json:"proposed_island_id"`
	ProposedIslandID2 int64     `db:"proposed_island_id2" json:"proposed_island_id2"`
	GalleryID         int64     `db:"gallery_id" json:"gallery_id"`
	PriceName1        string    `db:"priceName1" json:"priceName1"`
	Price1            int       `db:"price1" json:"price1"`
	PriceName2        string    `db:"priceName2" json:"priceName2"`
	Price2            int       `db:"price2" json:"price2"`
	PriceCurrency     string    `db:"price_currency" json:"price_currency"`
	PriceInfo         string    `db:"price_info" json:"price_info"`
	ProposedDatetime1 time.Time `db:"proposed_datetime1" json:"proposed_datetime1"`
	ProposedDatetime2 time.Time `db:"proposed_datetime2" json:"proposed_datetime2"`
	AdditionalInfo    string    `db:"additional_info" json:"additional_info"`
	ProposerID        int       `db:"proposer_id" json:"proposer_id"`
	ProposerInfo      string    `db:"proposer_info" json:"proposer_info"`
	ProposedDatetime  time.Time `db:"proposed_datetime" json:"proposed_datetime"`
	AcceptStatus      string    `db:"accept_status" json:"accept_status"`
	AcceptDatetime    time.Time `db:"accept_datetime" json:"accept_datetime"`
	ConfirmTel        string    `db:"confirm_tel" json:"confirm_tel"`
	ConfirmEmail      string    `db:"confirm_email" json:"confirm_email"`
	ApprovalStatus    string    `db:"approval_status" json:"approval_status"`
	ApprovalDate      time.Time `db:"approval_date" json:"approval_date"`
}

// ProposalInfo ...
type ProposalInfo struct {
	Proposal
	NumberOfLikes   int    `db:"num_of_likes" json:"num_of_likes"`
	IslandName      string `db:"island_name" json:"island_name"`
	ProposerName    string `db:"proposer_name" json:"proposer_name"`
	ProposerIconURL string `db:"proposer_icon_url" json:"proposer_icon_url"`
}

// Insert ...
func (s *Proposal) Insert(tx gorp.Transaction) error {
	s.ProposedDatetime = time.Now().UTC()
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

// GetProposalByID ...
func GetProposalByID(tx gorp.Transaction, ID int64) (*Proposal, error) {
	p := new(Proposal)
	if err := tx.SelectOne(&p, "SELECT * FROM proposal WHERE id = $1", ID); err != nil {
		return nil, err
	}
	return p, nil
}

// GetProposalsOf ...
func GetProposalsOf(tx *gorp.Transaction, requestID int64) ([]ProposalInfo, error) {
	proposals := []ProposalInfo{}
	_, err := tx.Select(&proposals, `
    select
      proposal.*,
      island.Name as island_name,
      (select count(id) from likes where entity_type='PROPOSAL' and entity_id=proposal.id) as num_of_likes,
  		users.user_name as proposer_name,
  		users.icon as proposer_icon_url
      from proposal
      inner join island on island.id=proposal.proposed_island_id
      join users on users.id = proposal.proposer_id
      where reply_to_request_id=$1
    `, requestID)
	return proposals, err
}

// CountOfProposalByUserID ...
func CountOfProposalByUserID(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := tx.SelectInt(`select count(*) from proposal
	    inner join request on request.id=proposal.reply_to_request_id
	    where proposer_id=$1 and request.expiry_date>current_date;`, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}
