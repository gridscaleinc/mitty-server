package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	"mitty.co/mitty-server/config"

	goutils "github.com/dongri/goutils"
	"github.com/mholt/binding"
)

// SignUpParams ...
type SignUpParams struct {
	UserName    string
	Password    string
	MailAddress string
}

// FieldMap defines parameter requirements
func (p *SignUpParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.UserName: binding.Field{
			Form:     "user_name",
			Required: true,
		},
		&p.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
		&p.MailAddress: binding.Field{
			Form:     "mail_address",
			Required: false,
		},
	}
}

// SignUpHandler ...
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(SignUpParams)
	if errs := binding.Bind(r, p); errs != nil {
		helpers.RenderInputError(w, r, errs)
		return
	}

	u, err := models.GetUserByUserName(*tx, p.UserName)
	if err != nil && err != sql.ErrNoRows {
		helpers.RenderDBError(w, r, err)
		return
	}
	if u != nil {
		err = errors.New("Username has already been taken")
		helpers.RenderDBError(w, r, err)
		return
	}

	emailAddress := ""
	if p.MailAddress != "" {
		email, e := mail.ParseAddress(p.MailAddress)
		if e != nil {
			helpers.RenderDBError(w, r, e)
			return
		}
		emailAddress = email.Address
		e = helpers.SendEmail("noreply@mitty.co", emailAddress, "Confirm", "confirm email address")
		fmt.Println(e)
	}

	user := new(models.User)
	user.UserName = p.UserName
	hashedPassword := goutils.Sha256Sum256(p.Password + config.CurrentSet.PasswordSalt())
	user.Password = hashedPassword
	user.MailAddress = emailAddress
	err = user.Insert(*tx)
	if err != nil {
		helpers.RenderDBError(w, r, err)
		return
	}
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"user_id":      user.ID,
		"access_token": user.AccessToken,
	})
}

// SignInHandler ...
func SignInHandler(w http.ResponseWriter, r *http.Request) {
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
	p := new(SignUpParams)
	if errs := binding.Bind(r, p); errs != nil {
		helpers.RenderInputError(w, r, errs)
		return
	}

	user, err := models.GetUserByUserName(*tx, p.UserName)
	if err != nil {
		helpers.RenderDBError(w, r, err)
		return
	}

	hashedPassword := goutils.Sha256Sum256(p.Password + config.CurrentSet.PasswordSalt())

	if user.Password != hashedPassword {
		err = errors.New("Password Error")
		helpers.RenderDBError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"user_id":      user.ID,
		"access_token": user.AccessToken,
	})
}
