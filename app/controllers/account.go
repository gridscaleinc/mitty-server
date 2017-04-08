package controllers

import (
	"database/sql"
	"net/http"

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
			Required: true,
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
	if errors := binding.Bind(r, p); errors != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}

	u, err := models.GetUserByUserName(*tx, p.UserName)
	if err != nil && err != sql.ErrNoRows {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": err,
		})
		return
	}
	if u != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": "Username has already been taken",
		})
		return
	}

	user := new(models.User)
	user.UserName = p.UserName
	hashedPassword := goutils.Sha256Sum256(p.Password + config.CurrentSet.PasswordSalt())
	user.Password = hashedPassword
	user.MailAddress = p.MailAddress
	err = user.Insert(*tx)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
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
	if errors := binding.Bind(r, p); errors != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}

	user, err := models.GetUserByUserName(*tx, p.UserName)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	hashedPassword := goutils.Sha256Sum256(p.Password + config.CurrentSet.PasswordSalt())

	if user.Password != hashedPassword {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": "Password Error",
		})
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"user_id":      user.ID,
		"access_token": user.AccessToken,
	})
}
