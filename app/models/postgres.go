package models

import "gopkg.in/gorp.v1"

// AddTableWithName ...
func AddTableWithName(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "ID")
	dbmap.AddTableWithName(Event{}, "events").SetKeys(true, "ID")
}
