package game

type Phaser interface {
	// Type returns the class name of this phaser
	Type() string

	// Next returns the following Phaser, i.e. after timeout
	Next() Phaser

	// Priority returns player object ids in priority order (turn holder)
	Priority() Priority

	// JSON returns a representation of this State's extra data
	JSON() map[string]any
}

// OnActivater is a Phaser that is triggered by the engine activating a State only the 1st time
type OnActivatePhaser interface {
	// OnActivate is called by the engine exactly once, when the State is mounted the 1st time
	OnActivate(*G) []Phaser
}

func TryOnActivate(g *G, phaser Phaser) []Phaser {
	g.Log().Debug("activate", phaser)
	if activater, _ := phaser.(OnActivatePhaser); activater != nil {
		return activater.OnActivate(g)
	}
	return nil
}

// OnConnectPhaser is a Phaser that is triggered when a player agent (re)connects
type OnConnectPhaser interface {
	// OnConnect is called by the engine whenever a seat.T (re)joins, and when the
	// Phaser re-mounts, as indicated by OnConnect(*T, nil)
	OnConnect(*G, *Player)
}

func TryOnConnect(g *G, phase Phaser, player *Player) {
	g.Log().Debug("connect", phase.Type(), player)
	if connector, ok := phase.(OnConnectPhaser); ok {
		connector.OnConnect(g, player)
	}
}

// OnDisconnecter is a Phaser that is triggered when a player agent disconnects
type OnDisconnectPhaser interface {
	// OnDisconnect is called by the engine whenever a seat.T disconnects
	OnDisconnect(*G, *Player)
}

// OnFinishPhaser is a Phaser that is triggered by the engine finally resolving a State
type OnFinishPhaser interface {
	// OnFinish is called by the engine exactly once, after all responses, or timeout
	OnFinish(*G, *State) []Phaser
}

func TryOnFinish(g *G, state *State) []Phaser {
	g.Log().Debug("finish", state.T.Phase.Type())
	if finisher, ok := state.T.Phase.(OnFinishPhaser); ok {
		return finisher.OnFinish(g, state)
	}
	return nil
}

// OnRequestPhaser is a Phaser that is triggered by the engine when a Request targets the game state ID
type OnRequestPhaser interface {
	// OnRequest is called when a request is sent to this State
	OnRequest(*G, *State, *Player, map[string]any)
}

func TryOnRequest(g *G, state *State, player *Player, json map[string]any) {
	g.Log().Debug("finish", state.T.Phase.Type())
	if requester, ok := state.T.Phase.(OnRequestPhaser); ok {
		requester.OnRequest(g, state, player, json)
	}
}
