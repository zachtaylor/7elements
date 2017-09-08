package games

import (
	"7elements.ztaylor.me/decks"
	"7elements.ztaylor.me/games/cards"
	"7elements.ztaylor.me/games/decks"
	"7elements.ztaylor.me/games/seats"
	"7elements.ztaylor.me/games/turns"
	"time"
	"ztaylor.me/ctxpert"
)

var Cache = make(map[int]*Game)

type Game struct {
	Id         int
	GamePhase  int
	TurnPhase  int
	nextTurn   int
	nextCardId int
	results    map[string]bool
	Seats      []*gameseats.GameSeat
	History    []*gameturns.GameTurn
	Patience   time.Duration
	*ctxpert.Context
}

func New() *Game {
	return &Game{
		results:  make(map[string]bool),
		Seats:    make([]*gameseats.GameSeat, 0),
		History:  make([]*gameturns.GameTurn, 0),
		Patience: 5 * time.Minute,
		Context:  ctxpert.New(),
	}
}

func (g *Game) MarkWinner(username string) {
	g.results[username] = true
}

func (g *Game) GetWinners() []string {
	winners := make([]string, 0)
	for username, isWinner := range g.results {
		if isWinner {
			winners = append(winners, username)
		}
	}
	return winners
}

func (g *Game) MarkLoser(username string) {
	g.results[username] = false
}

func (g *Game) GetLosers() []string {
	losers := make([]string, 0)
	for username, isWinner := range g.results {
		if !isWinner {
			losers = append(losers, username)
		}
	}
	return losers
}

func (g *Game) GetSeat(username string) *gameseats.GameSeat {
	for _, seat := range g.Seats {
		if username == seat.Username {
			return seat
		}
	}
	return nil
}

func (g *Game) NextTurn() int {
	i := g.nextTurn
	g.nextTurn++
	if g.nextTurn >= len(g.Seats) {
		g.nextTurn = 0
	}
	return i
}

func (g *Game) NextCardId() int {
	i := g.nextCardId
	g.nextCardId++
	return i
}

func (g *Game) CurrentTurn() *gameturns.GameTurn {
	if len(g.History) == 0 {
		return nil
	}
	return g.History[len(g.History)-1]
}

func BuildGameSeat(deck *decks.Deck, game *Game) *gameseats.GameSeat {
	seat := gameseats.New()
	seat.Username = deck.Username
	seat.DeckId = deck.Id
	seat.Deckname = deck.Name
	seat.Deck = gamedecks.New()

	for cardid, count := range deck.Cards {
		for i := 0; i < count; i++ {
			seat.Deck.Append(gamecards.Build(cardid, deck.Username, game.NextCardId()))
		}
	}

	return seat
}

func (g *Game) BuildNextTurn() *gameturns.GameTurn {
	seat := g.Seats[g.NextTurn()]
	turn := &gameturns.GameTurn{
		GameId:   g.Id,
		Id:       len(g.History),
		Username: seat.Username,
	}
	g.History = append(g.History, turn)

	return turn
}
