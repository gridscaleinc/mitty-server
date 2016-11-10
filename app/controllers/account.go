package controllers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	"mitty.co/mitty-server/config"

	"github.com/mholt/binding"
)

// SignUpParams ...
type SignUpParams struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// FieldMap defines parameter requirements
func (p *SignUpParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.UserName: binding.Field{
			Form:     "username",
			Required: true,
		},
		&p.Password: binding.Field{
			Form:     "password",
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
	bytes := sha256.Sum256([]byte(p.Password + config.CurrentSet.PasswordSalt()))
	user.Password = hex.EncodeToString(bytes[:])
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

	bytes := sha256.Sum256([]byte(p.Password + config.CurrentSet.PasswordSalt()))
	inputPassword := hex.EncodeToString(bytes[:])

	if user.Password != inputPassword {
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
