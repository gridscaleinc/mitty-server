package models

import (
	"time"

	goutils "github.com/dongri/goutils"
	gorp "gopkg.in/gorp.v1"
)

// User ...
type User struct {
	ID            int       `db:"id" json:"id"`
	Name          string    `db:"name" json:"name"`
	UserName      string    `db:"username" json:"username"`
	Password      string    `db:"password" json:"password"`
	AccessToken   string    `db:"access_token" json:"access_token"`
	MailAddress   string    `db:"mail_address" json:"mail_address"`
	MailConfirmed bool      `db:"mail_confirmed" json:"mail_confirmed"`
	Status        string    `db:"status" json:"status"`
	Created       time.Time `db:"created" json:"created"`
	Updated       time.Time `db:"updated" json:"updated"`
}

// Insert ...
func (u *User) Insert(tx gorp.Transaction) error {
	random := new(goutils.Random)
	random.UseNumber()
	random.UseSmallLetter()
	random.UseCapitalLetter()
	r := random.Random(40)
	u.AccessToken = r
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
	if err := tx.SelectOne(&u, "SELECT * FROM users WHERE username = $1", userName); err != nil {
		return nil, err
	}
	return u, nil
}
