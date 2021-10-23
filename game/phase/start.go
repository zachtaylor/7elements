package phase

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

func NewStart(seat string) game.Phaser {
	return &Start{
		R: R(seat),
	}
}

type Start struct{ R }

func (r *Start) Name() string { return "start" }

func (r *Start) String() string {
	return "start (" + r.Seat() + ")"
}

// OnActivate implements game.OnActivatePhaser
func (r *Start) OnActivate(game *game.T) []game.Phaser {
	game.Log().Trace("activate")
	for _, seatName := range game.Seats.Keys() {
		seat := game.Seats.Get(seatName)
		seat.Life = 7
		seat.Deck.Shuffle()
		_ = game.Engine().DrawCard(game, seat, game.Rules().StartingHand)
	}
	return nil
}
func (r *Start) onActivatePhaser() game.OnActivatePhaser { return r }

// OnConnect implements game.OnConnectPhaser
func (r *Start) OnConnect(game *game.T, seat *seat.T) {
	// if seat == nil {
	// game.Log().Trace("announce")
	// go game.Chat("sunrise", r.Seat())
	// }
}
func (r *Start) onConnectPhaser() game.OnConnectPhaser { return r }

func (r *Start) GetNext(game *game.T) game.Phaser { return NewSunrise(r.Seat()) }

func (r *Start) Data() map[string]interface{} { return nil }

// Request implements Requester
func (r *Start) OnRequest(game *game.T, seat *seat.T, json map[string]interface{}) {
	choice, _ := json["choice"].(string)
	log := game.Log().Add("Seat", seat).Add("Choice", choice)

	if react := game.State.Reacts[seat.Username]; react != "" {
		log.Add("React", react).Warn("already recorded")
		return
	} else if choice == "keep" {
		game.State.Reacts[seat.Username] = "keep"
	} else if choice == "mulligan" {
		game.State.Reacts[seat.Username] = "mulligan"
		for _, card := range seat.Hand {
			seat.Past[card.ID] = card
		}
		seat.Hand = card.Set{}
		_ = game.Engine().DrawCard(game, seat, game.Rules().StartingHand)
	} else {
		log.Warn("unrecognized")
		return
	}

	game.State.Reacts[seat.Username] = choice
	game.Seats.Write(wsout.GameReact(game.State.ID(), seat.Username, choice, game.State.Timer).EncodeToJSON())
	log.Info("confirm")
}
func (r *Start) onRequestPhaser() game.OnRequestPhaser { return r }
