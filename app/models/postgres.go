package models

import "gopkg.in/gorp.v1"

// AddTableWithName ...
func AddTableWithName(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "ID")
	dbmap.AddTableWithName(Event{}, "events").SetKeys(true, "ID")
	dbmap.AddTableWithName(Gallery{}, "gallery").SetKeys(true, "ID")
	dbmap.AddTableWithName(Island{}, "island").SetKeys(true, "ID")
	dbmap.AddTableWithName(Contents{}, "contents").SetKeys(true, "ID")
	dbmap.AddTableWithName(Activity{}, "activity").SetKeys(true, "ID")
	dbmap.AddTableWithName(ActivityItem{}, "activity_item").SetKeys(true, "ID")
	dbmap.AddTableWithName(Meeting{}, "meeting").SetKeys(true, "ID")
	dbmap.AddTableWithName(Conversation{}, "conversation").SetKeys(true, "ID")
	dbmap.AddTableWithName(Request{}, "request").SetKeys(true, "ID")
}
