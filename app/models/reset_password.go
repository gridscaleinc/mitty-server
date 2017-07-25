package models

import (
	"database/sql"
	"time"

	gorp "gopkg.in/gorp.v1"
)

// OTP Const
const (
	CodeExpireSecond = 60 * 60 * 6
)

// ResetPassword ...
type ResetPassword struct {
	ID      int       `db:"id" json:"id"`
	Email   string    `db:"email" json:"email"`
	Token   string    `db:"token" json:"token"`
	Expire  time.Time `db:"expire" json:"expire"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

// Save ...
func (rp *ResetPassword) Save(dbmap *gorp.DbMap) error {
	tx, err := dbmap.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		return
	}()
	rp.Expire = time.Now().UTC().Add(time.Duration(CodeExpireSecond) * time.Second)
	rp.Created = time.Now().UTC()
	rp.Updated = time.Now().UTC()
	if err := tx.Insert(rp); err != nil {
		return err
	}
	return nil
}

// GetEmailByToken ...
func GetEmailByToken(dbmap *gorp.DbMap, token string) (*ResetPassword, error) {
	resetPasswords := []ResetPassword{}
	if _, err := dbmap.Select(&resetPasswords, `SELECT *
		FROM reset_passwords WHERE token = $1 AND expire >= $2 ORDER By created DESC`, token, time.Now().UTC()); err != nil {
		return nil, err
	}
	if len(resetPasswords) > 0 {
		return &resetPasswords[0], nil
	}
	return nil, sql.ErrNoRows
}
