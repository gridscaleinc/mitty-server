package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	goutils "github.com/dongri/goutils"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	"mitty.co/mitty-server/config"
)

// UpdateUserIconHandler ...
func UpdateUserIconHandler(w http.ResponseWriter, r *http.Request) {
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
	contentID, _ := strconv.Atoi(r.URL.Query().Get("contentId"))

	if err := models.SetUserIcon(*tx, currentUserID, contentID); err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"ok": true,
	})

}

// GetUserInfo ...
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
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

	id := r.URL.Query().Get("id")

	userInfo, err := models.GetUserInfo(id)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"userInfo": userInfo,
	})
}

// EmailForm ...
type EmailForm struct {
	Email string
}

// FieldMap ...
func (s *EmailForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
	}
}

// ResetPasswordSendHandler ...
func ResetPasswordSendHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	form := new(EmailForm)
	errs := binding.Bind(r, form)
	if errs.Handle(w) {
		return
	}
	user, err := models.GetUserByEmail(dbmap, form.Email)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	if err == sql.ErrNoRows {
		filters.RenderError(w, r, errors.New("user not found"))
		return
	}
	resetPassword := new(models.ResetPassword)
	resetPassword.Email = user.MailAddress
	resetPassword.Token = helpers.GenerateTokenWithSalt(user.MailAddress)
	if err = resetPassword.Save(dbmap); err != nil {
		filters.RenderError(w, r, err)
		return
	}
	err = helpers.SendEmail("noreply@mitty.co", user.MailAddress, "Reset Password", `
		Reset your password?
	  If you requested a password reset, click the button below. If you didn't make this request, ignore this email.
	  `+"http://dev.mitty.co/api/reset_password/verify?token="+resetPassword.Token)
	fmt.Println(err)

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"ok": true,
	})

}

// ResetPasswordVerifyHandler ...
func ResetPasswordVerifyHandler(w http.ResponseWriter, r *http.Request) {
	dbmap := helpers.GetPostgres()
	token := r.URL.Query().Get("token")
	resetPassword, err := models.GetEmailByToken(dbmap, token)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	output := map[string]interface{}{
		"token": resetPassword.Token,
	}
	filters.RenderHTML(w, r, "app/views/reset_password/input.html", output)
}

// ResetPasswordForm ...
type ResetPasswordForm struct {
	Token    string
	Password string
}

// FieldMap ...
func (f *ResetPasswordForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Token: binding.Field{
			Form:     "token",
			Required: true,
		},
		&f.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

// ResetPasswordResetHandler ...
func ResetPasswordResetHandler(w http.ResponseWriter, r *http.Request) {
	dbmap := helpers.GetPostgres()
	form := new(ResetPasswordForm)
	errs := binding.Bind(r, form)
	if errs.Handle(w) {
		return
	}
	resetPassword, err := models.GetEmailByToken(dbmap, form.Token)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	user, err := models.GetUserByEmail(dbmap, resetPassword.Email)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	hashedPassword := goutils.Sha256Sum256(form.Password + config.CurrentSet.PasswordSalt())
	user.Password = hashedPassword

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

	if err := user.Update(*tx); err != nil {
		filters.RenderError(w, r, err)
		return
	}
	output := map[string]interface{}{
		"title": "reset password done",
	}
	filters.RenderHTML(w, r, "app/views/reset_password/done.html", output)
}
