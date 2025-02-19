package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"time"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	"mitty.co/mitty-server/config"

	goutils "github.com/dongri/goutils"
	"github.com/mholt/binding"
)

// SignUpParams ...
type SignUpParams struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	MailAddress string `json:"mail_address"`
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

var usernameRegexp = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,15}$`)
var emailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

// Validate ...
func (p *SignUpParams) Validate(req *http.Request) error {
	if !usernameRegexp.MatchString(p.UserName) {
		return errors.New("username is invalid string")
	}
	if !emailRegexp.MatchString(p.MailAddress) {
		return errors.New("email address is invalid")
	}
	if len(p.Password) > 255 {
		return errors.New("password is too long")
	}
	return nil
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
		filters.RenderInputErrors(w, r, errs)
		return
	}

	if inputErr := p.Validate(r); inputErr != nil {
		filters.RenderInputError(w, r, inputErr)
		return
	}

	u, err := models.GetUserByUserName(*tx, p.UserName)
	if err != nil && err != sql.ErrNoRows {
		filters.RenderError(w, r, err)
		return
	}
	if u != nil {
		err = errors.New("Username has already been taken")
		filters.RenderError(w, r, err)
		return
	}

	emailAddress := ""
	if p.MailAddress != "" {
		email, e := mail.ParseAddress(p.MailAddress)
		if e != nil {
			filters.RenderError(w, r, e)
			return
		}
		emailAddress = email.Address
	}

	user := new(models.User)
	user.UserName = p.UserName
	hashedPassword := goutils.Sha256Sum256(p.Password + config.CurrentSet.PasswordSalt())
	user.Password = hashedPassword
	user.MailAddress = emailAddress
	user.MailConfirmed = false
	if emailAddress != "" {
		user.MailToken = goutils.Sha256Sum256(time.Now().String() + config.CurrentSet.PasswordSalt())
		err = helpers.SendEmail("noreply@mitty.co", emailAddress, "Confirm", "confirm email address\nhttp://dev.mitty.co/email/confirm?token="+user.MailToken)
		fmt.Println(err)
	}
	err = user.Insert(*tx)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"user_id":      user.ID,
		"access_token": user.AccessToken,
		"created":      user.Created,
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
		filters.RenderInputErrors(w, r, errs)
		return
	}

	user, err := models.GetUserByUserName(*tx, p.UserName)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	hashedPassword := goutils.Sha256Sum256(p.Password + config.CurrentSet.PasswordSalt())

	if user.Password != hashedPassword {
		err = errors.New("Password Error")
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"user_id":      user.ID,
		"access_token": user.AccessToken,
	})
}
