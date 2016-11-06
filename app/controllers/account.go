package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

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
	fmt.Print("===========================")
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
	fmt.Println(err)
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
	user.Password = p.Password
	err = user.Insert(*tx)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": err,
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
			"err": err,
		})
		return
	}

	if user.Password != p.Password {
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
