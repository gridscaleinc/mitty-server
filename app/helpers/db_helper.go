package helpers

import (
	"database/sql"

	"github.com/Sirupsen/logrus"
	"gopkg.in/gorp.v1"
	"mitty.co/mitty-server/config"
)

// SetupDatabase ...
func SetupDatabase(currentSet config.EnvConfigSet) *sql.DB {

	// Postgres
	db, err := sql.Open("postgres", currentSet.PostgresURI())
	if err != nil {
		logrus.Infof(err.Error())
		logrus.Fatal("Failed to connect to Postgres")
		return nil
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	setPostgres(dbmap)
	return db
}

// DbMap ...
var dbMap *gorp.DbMap

// GetPostgres ...
func GetPostgres() *gorp.DbMap {
	return dbMap
}

// SetPostgres ....
func setPostgres(dbmap *gorp.DbMap) {
	dbMap = dbmap
}
