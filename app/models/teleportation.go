package models

import "time"

// Teleportation ...
type Teleportation struct {
	MittyID   int       `json:"mitty_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	PortTime  time.Time `json:"porttime"`
}
