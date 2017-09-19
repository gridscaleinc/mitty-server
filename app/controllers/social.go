package controllers

import (
	"net/http"

	gorp "gopkg.in/gorp.v1"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

// GetSocialMirrorHandler PATH: GET /social/mirror
func GetSocialMirrorHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	tx, err := dbmap.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	currentUserID := filters.GetCurrentUserID(r)
	eventCount, err := countEvents(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	todaysEventCount, err := countTodaysEvents(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	eventInvitationCount, err := countEventInvitation(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	namecardOfferCount, err := countNamecardOffer(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	requestCount, err := countRequest(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	proposalCount, err := countProposal(tx, currentUserID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"event":             eventCount,
		"todaysEvent":       todaysEventCount,
		"eventInvitation":   eventInvitationCount,
		"businessCardOffer": namecardOfferCount,
		"request":           requestCount,
		"proposal":          proposalCount,
	})
}

// Count up comming events the user paticipated
func countEvents(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := models.CountOfEventByUserID(tx, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Count up Todays events the user paticipated
func countTodaysEvents(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := models.CountOfTodaysEventByUserID(tx, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Count up invitations of event for the user to take part in
func countEventInvitation(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := models.CountOfEventInvitationByUserID(tx, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Count up NamecardOffers that not accepted by the user
func countNamecardOffer(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := models.CountOfNamecardOfferByUserID(tx, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Count up Request the user posted
func countRequest(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := models.CountOfRequestByUserID(tx, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Count up Proposals the user posted
func countProposal(tx *gorp.Transaction, uid int) (int64, error) {
	count, err := models.CountOfProposalByUserID(tx, uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}
