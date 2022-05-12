package game

type CardLocation uint

const (
	CardLocationUnknown CardLocation = iota
	CardLocationHand
	CardLocationPlay
	CardLocationPresent
	CardLocationPast
	CardLocationRemoved
)
