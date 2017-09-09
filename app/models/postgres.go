package models

import (
	"log"

	"gopkg.in/gorp.v1"
)

type MyGorpTracer struct{}

func (t *MyGorpTracer) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// AddTableWithName ...
func AddTableWithName(dbmap *gorp.DbMap) {
	// A
	dbmap.AddTableWithName(ActivityItem{}, "activity_item").SetKeys(true, "ID")
	dbmap.AddTableWithName(Activity{}, "activity").SetKeys(true, "ID")

	// C
	dbmap.AddTableWithName(Contact{}, "contact").SetKeys(true, "ID")
	dbmap.AddTableWithName(Contents{}, "contents").SetKeys(true, "ID")
	dbmap.AddTableWithName(Conversation{}, "conversation").SetKeys(true, "ID")

	// E
	dbmap.AddTableWithName(Event{}, "events").SetKeys(true, "ID")

	// F
	dbmap.AddTableWithName(Footmark{}, "footmark").SetKeys(true, "ID")

	// G
	dbmap.AddTableWithName(Gallery{}, "gallery").SetKeys(true, "ID")

	// I
	dbmap.AddTableWithName(Invitation{}, "invitation").SetKeys(true, "ID")
	dbmap.AddTableWithName(Invitees{}, "invitees").SetKeys(true, "ID")
	dbmap.AddTableWithName(Island{}, "island").SetKeys(true, "ID")

	// L
	dbmap.AddTableWithName(Likes{}, "likes").SetKeys(true, "ID")

	// M
	dbmap.AddTableWithName(Meeting{}, "meeting").SetKeys(true, "ID")

	// N
	dbmap.AddTableWithName(Namecard{}, "namecard").SetKeys(true, "ID")

	// O
	dbmap.AddTableWithName(Offers{}, "offers").SetKeys(true, "ID")

	// P
	dbmap.AddTableWithName(Presence{}, "presence").SetKeys(true, "ID")
	dbmap.AddTableWithName(Profile{}, "profile").SetKeys(true, "ID")
	dbmap.AddTableWithName(Proposal{}, "proposal").SetKeys(true, "ID")

	// R
	dbmap.AddTableWithName(Request{}, "request").SetKeys(true, "ID")
	dbmap.AddTableWithName(ResetPassword{}, "reset_passwords").SetKeys(true, "ID")

	// S
	dbmap.AddTableWithName(SocialID{}, "socialid").SetKeys(true, "ID")
	dbmap.AddTableWithName(SocialLink{}, "sociallink").SetKeys(true, "ID")

	// U
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "ID")

	// Loggerを生成
	tracer := &MyGorpTracer{}

	// Logging有効化
	dbmap.TraceOn("[gorp SQL trace]", tracer)

}
