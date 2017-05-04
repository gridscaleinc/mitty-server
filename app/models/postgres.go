package models

import "gopkg.in/gorp.v1"

// AddTableWithName ...
func AddTableWithName(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "ID")
	dbmap.AddTableWithName(Event{}, "events").SetKeys(true, "ID")
	dbmap.AddTableWithName(Gallery{}, "gallery").SetUniqueTogether("ID","SEQ")
	dbmap.AddTableWithName(Island{}, "island").SetKeys(true, "ID")
	dbmap.AddTableWithName(Contents{}, "contents").SetKeys(true, "ID")
	dbmap.AddTableWithName(Activity{}, "activity").SetKeys(true, "ID")
	dbmap.AddTableWithName(ActivityItem{}, "activity_item").SetUniqueTogether("ACTIVITY_ID","EVENT_ID")
}
