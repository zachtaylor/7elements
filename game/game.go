package game

// type T struct {
// 	id    string
// 	in    chan *Request
// 	close chan bool
// 	obj   map[string]any
// 	rules Rules
// 	log   log.Writer
// 	Seats *seat.List
// 	State *State
// }

// func New(id string, rules Rules, log log.Writer) *T {
// 	return &T{
// 		id: id,
// 		// in:    make(chan *Request),
// 		close: make(chan bool),
// 		obj:   make(map[string]any),
// 		rules: rules,
// 		log:   log,
// 		Seats: seat.NewList(),
// 	}
// }

// func (game *T) ID() string { return game.id }

// func (game *T) Log() log.Writer { return game.log }

// func (game *T) Rules() Rules { return game.rules }

// func (game *T) String() string { return "Game#" + game.id }

// func (game *T) Close() {
// 	if game.in != nil {
// 		close(game.in)
// 		game.in = nil
// 		close(game.close)
// 		game.close = nil
// 	}
// 	game.log.Close()
// 	// game.chat.Destroy()
// }

// // RequestChan returns the raw ordered game input
// func (game *T) RequestChan() <-chan *Request { return game.in }

// // Request starts a go routine to call RequestSync
// func (game *T) Request(username, uri string, json map[string]any) {
// 	go game.RequestSync(NewReq(username, uri, json))
// }

// // RequestSync waits to request the game engine
// func (game *T) RequestSync(r *Request) {
// 	if game.in != nil {
// 		game.in <- r
// 	}
// }

// // func (game *T) Chat(source, message string) { game.chat.Add(source, message) }

// func (game *T) Phase() string { return game.State.Phase.Name() }

// func (game *T) JSON(seat *seat.T) map[string]any {
// 	return map[string]any{
// 		"id":       game.ID(),
// 		"stateid":  game.State.ID(),
// 		"username": seat.Username,
// 		"seats":    game.Seats.Keys(),
// 	}
// }

// func (t *T) Register(deck *deck.T, writer seat.Writer) *seat.T {
// 	log := t.Log().Add("Username", deck.User)

// 	if t.Seats.Get(deck.User) != nil {
// 		log.Warn("username already registered")
// 		return nil
// 	}

// 	seat := seat.New(t.rules.StartingLife, deck, writer)

// 	t.Seats.Add(seat)

// 	for _, card := range deck.Cards {
// 		t.RegisterCard(card)
// 	}

// 	log.Add("Name", deck.Proto.Name).Info("register")

// 	return seat
// }
