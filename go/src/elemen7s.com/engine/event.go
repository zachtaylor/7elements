package engine

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

type Event interface {
	// Name is the refferential event type
	Name() string
	// Priority indicates that the event cannot resolve without further input
	Priority(*vii.Game, *Timeline) bool
	// OnStart is called by the host Timeline when its' lifetime begins counting
	OnStart(*vii.Game, *Timeline)
	// OnReconnect is called by the host Timeline whenever a GameSeat (re)joins
	OnReconnect(*vii.Game, *Timeline, *vii.GameSeat)
	// OnStop is called by the host Timeline when its' lifetime reaches 0
	// Returns the next Event
	OnStop(*vii.Game, *Timeline) *Timeline
	// Json() create a representation of this Event for client consumption
	Json(*vii.Game, *Timeline) js.Object
	// Receive is called when data is sent to this event
	Receive(*vii.Game, *Timeline, *vii.GameSeat, js.Object)
}
