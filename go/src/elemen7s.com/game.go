package vii

import (
	"elemen7s.com/chat"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

type Game struct {
	Key      string
	Cards    GameCards
	Settings *GameSettings
	Results  *GameResults
	In       chan *GameRequest
	Chat     chat.Channel
	Seats    map[string]*GameSeat
	*log.Logger
}

type GameResults struct {
	Winner string
	Loser  string
}

func NewGame() *Game {
	return &Game{
		Key:      keygen.NewVal(),
		Cards:    make(GameCards),
		Settings: NewDefaultGameSettings(),
		In:       make(chan *GameRequest),
		Seats:    make(map[string]*GameSeat),
		Logger:   log.NewLogger(),
	}
}

func (game Game) String() string {
	return game.Key
}

func (game *Game) Log() log.Log {
	return game.Logger.New()
}

func (game *Game) GetSeat(name string) *GameSeat {
	for k, seat := range game.Seats {
		if k == name {
			return seat
		}
	}
	return nil
}

func (game *Game) GetOpponentSeat(name string) *GameSeat {
	for k, seat := range game.Seats {
		if k != name {
			return seat
		}
	}
	return nil
}

func (game *Game) Register(deck *AccountDeck, lang string) *GameSeat {
	log := game.Log().Add("Username", deck.Username)

	if game.Seats[deck.Username] != nil {
		log.Warn("register: username already registered")
		return nil
	}

	seat := game.NewSeat()
	seat.Deck = NewGameDeck()
	seat.Username = deck.Username
	seat.Deck.Username = deck.Username
	seat.Deck.AccountDeckId = deck.Id
	seat.Deck.AccountDeckVersion = deck.Version

	deckSize := 0

	for cardid, copies := range deck.Cards {
		card, _ := CardService.GetCard(cardid)
		if card == nil {
			log.Clone().Add("CardId", cardid).Warn("register: card missing")
			return nil
		}
		text, err := CardTextService.GetCardText(lang, cardid)
		if text == nil {
			log.Clone().Add("CardId", cardid).Add("Error", err).Warn("register: text missing")
			return nil
		}

		for i := 0; i < copies; i++ {
			card := NewGameCard(card, text)
			card.Username = deck.Username
			game.RegisterCard(card)
			game.Cards[card.Id] = card
			seat.Deck.Append(card)
		}
		deckSize += copies
		log.Clone().Add("CardId", cardid).Add("Copies", copies).Debug("register card")
	}

	game.Seats[seat.Username] = seat
	log.Add("DeckSize", deckSize).Debug("registered seat")
	return seat
}

func (game *Game) RegisterCard(card *GameCard) {
	card.Id = keygen.NewVal()
	for game.Cards[card.Id] != nil {
		card.Id = keygen.NewVal()
	}
	game.Cards[card.Id] = card
}

func (game *Game) Send(name string, json Json) {
	for _, seat := range game.Seats {
		seat.Send(name, json)
	}
}

func (game *Game) SendCatchup(name string) {
	game.In <- &GameRequest{
		Username: name,
		Data: Json{
			"event": "reconnect",
		},
	}
}

func (game *Game) Json(name string) Json {
	seat := game.GetSeat(name)
	return Json{
		"gameid":   game.Key,
		"life":     seat.Life,
		"hand":     len(seat.Hand),
		"opponent": game.GetOpponentSeat(name).Username,
	}
}
