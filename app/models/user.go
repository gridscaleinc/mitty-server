package models

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// User ...
type User struct {
	ID       int       `db:"id" json:"id"`
	UserName string    `db:"user_name" json:"user_name"`
	Password string    `db:"password" json:"password"`
	Name     string    `db:"name" json:"name"`
	Created  time.Time `db:"created" json:"created"`
	Updated  time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (u *User) Insert(tx gorp.Transaction) error {
	u.Created = time.Now().UTC()
	u.Updated = time.Now().UTC()
	err := tx.Insert(u)
	return err
}

// Update ...
func (u *User) Update(tx gorp.Transaction) error {
	u.Updated = time.Now().UTC()
	_, err := tx.Update(u)
	return err
}

// GetUserByUserName ...
func GetUserByUserName(tx gorp.Transaction, userName string) (*User, error) {
	u := new(User)
	if err := tx.SelectOne(&u, "SELECT * FROM users WHERE user_name = $1", userName); err != nil {
		return nil, err
	}
	return u, nil
}
