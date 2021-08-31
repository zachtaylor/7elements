package phase

import (
	"reflect"
	"strconv"

	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func NewSunrise(seat string) game.Phaser {
	return &Sunrise{
		R:   R(seat),
		Ans: make(map[string]element.T),
	}
}

type Sunrise struct {
	R
	Ans map[string]element.T
}

func (r *Sunrise) Name() string { return "sunrise" }

func (r *Sunrise) String() string {
	return "sunrise (" + r.Seat() + ")"
}

// OnActivate implements game.OnActivatePhaser
func (r *Sunrise) OnActivate(game *game.T) (rs []game.Phaser) {
	game.Log().Add("Seat", r.Seat()).Trace("activate")
	seat := game.Seats.Get(r.Seat())
	seat.Karma.Reactivate()
	game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
	for _, seatName := range game.Seats.Keys() {
		seat := game.Seats.Get(seatName)
		for _, token := range seat.Present {
			rs = append(rs, game.Engine().WakeToken(game, token)...)
		}
	}
	return
}
func (r *Sunrise) onActivatePhaser() game.OnActivatePhaser { return r }

// OnConnect implements game.OnConnectPhaser
func (r *Sunrise) OnConnect(game *game.T, seat *seat.T) {
	game.Log().Add("Seat", seat).Trace("connect")
	if seat == nil {
		go game.Chat("sunrise", r.Seat())
	}
}
func (r *Sunrise) onConnectPhaser() game.OnConnectPhaser { return r }

// Finish implements game.OnFinishPhaser
func (r *Sunrise) OnFinish(game *game.T) (rs []game.Phaser) {
	for _, seatName := range game.Seats.Keys() {
		seat := game.Seats.Get(seatName)
		log := game.Log().Add("Username", seat.Username)

		if el := r.Ans[seat.Username]; el < element.White || el > element.Black {
			log.Warn("el is out of bounds")
			continue
		} else {
			seat.Karma.Append(el, false)
		}
		log.Add("Karma", seat.Karma).Trace("finish")

		if seatName != game.State.Phase.Seat() {
			game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
		}
	}
	if seat := game.Seats.Get(game.State.Phase.Seat()); seat.Deck.Count() > 0 {
		rs = append(rs, game.Engine().DrawCard(game, seat, 1)...)
	} else {
		game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
	}
	return
}

// type check
func (r *Sunrise) onFinishPhaser() game.OnFinishPhaser { return r }

func (r *Sunrise) GetNext(game *game.T) game.Phaser {
	return NewMain(r.Seat())
}

func (r *Sunrise) Data() websocket.MsgData { return nil }

// OnRequest implements game.OnRequestPhaser
func (r *Sunrise) OnRequest(game *game.T, seat *seat.T, json websocket.MsgData) {
	var choice int
	if choicebuff, _ := json["choice"].(string); len(choicebuff) < 1 {
		game.Log().Add("choice", json["choice"]).Add("type", reflect.TypeOf(json["choice"])).Warn("choice missing")
		seat.Writer.Write(wsout.ErrorJSON("vii", "choice missing"))
		return
	} else if choiceparse, err := strconv.ParseInt(choicebuff, 10, 0); err != nil {
		game.Log().Add("buff", choicebuff).Add("error", err.Error()).Error("choice parse")
		seat.Writer.Write(wsout.ErrorJSON("vii", "choice missing"))
	} else {
		choice = int(choiceparse)
	}
	log := game.Log().Add("Seat", seat).Add("Choice", choice)

	el := element.T(choice)
	log.Info("confirm")
	r.Ans[seat.Username] = el
	game.State.Reacts[seat.Username] = "ok"
	game.Seats.Write(wsout.GameReact(game.State.ID(), seat.Username, "ok", game.State.Timer).EncodeToJSON())
}
func (r *Sunrise) onRequestPhaser() game.OnRequestPhaser { return r }
