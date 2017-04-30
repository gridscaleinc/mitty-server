package controllers

import (
	"database/sql"
	"net/http"

	"mitty.co/mitty-server/app/helpers"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/models"
)

// WelcomeHandler ...
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{
		"title": "welcome mitty",
	}
	filters.RenderHTML(w, r, "app/views/index.html", output)
}

// EmailConfirmHandler ...
func EmailConfirmHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
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

	output := map[string]interface{}{
		"title": "Email Confirm",
	}

	user, err := models.GetUserByEmailToken(*tx, token)
	if err != nil && err != sql.ErrNoRows {
		output["result"] = err.Error()
	}
	if err == sql.ErrNoRows {
		output["result"] = "Not Found"
	}
	if user != nil {
		user.MailConfirmed = true
		if err = user.Update(*tx); err != nil {
			output["result"] = err.Error()
		}
		output["result"] = user.MailAddress + " success"
	}
	filters.RenderHTML(w, r, "app/views/email_confirm.html", output)
}
