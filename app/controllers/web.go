package controllers

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
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

// UserGuideHandler ...
func UserGuideHandler(w http.ResponseWriter, r *http.Request) {
	filters.RenderHTML(w, r, "app/views/userguide.html", nil)
}

// ContactForm ...
type ContactForm struct {
	Name    string
	Email   string
	Comment string
}

// FieldMap ...
func (s *ContactForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
		&s.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&s.Comment: binding.Field{
			Form:     "comment",
			Required: true,
		},
	}
}

// ContactHandler ...
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	form := new(ContactForm)
	errs := binding.Bind(r, form)
	fmt.Println(errs)
	if errs.Handle(w) {
		return
	}
	body := "Name: " + form.Name + "\nEmail: " + form.Email + "\nComment: " + form.Comment
	err := helpers.SendEmail("contact@mitty.co", "admin@mitty.co", "Contact from Mitty Web", body)
	fmt.Println(err)

	jsonStr := `{"text":"` + body + `"}`
	url := "https://hooks.slack.com/services/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"
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"result": "ok",
	})
}
