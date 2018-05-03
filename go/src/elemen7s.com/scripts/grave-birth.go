package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"ztaylor.me/js"
)

const GraveBirthID = "grave-birth"

func init() {
	games.Scripts[GraveBirthID] = GraveBirth
}

type GraveBirthMode struct {
	Card  *vii.GameCard
	Stack *games.Event
}

func (mode *GraveBirthMode) Name() string {
	return "choice"
}

func (mode *GraveBirthMode) Json(e *games.Event, g *games.Game, s *games.Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"choice":   "Grave Birth",
		"username": s.Username,
		"timer":    int(e.Duration.Seconds()),
	}
}

func (mode *GraveBirthMode) OnActivate(e *games.Event, g *games.Game) {
	games.AnimateGraveBirthChoice(g.GetSeat(e.Username), g)
}

func (mode *GraveBirthMode) OnSendCatchup(e *games.Event, g *games.Game, s *games.Seat) {
	if e.Username == s.Username {
		games.AnimateGraveBirthChoice(s, g)
	}
}

func (mode *GraveBirthMode) OnResolve(e *games.Event, g *games.Game) {
	g.RegisterToken(e.Username, mode.Card)
	games.BroadcastAnimateSeatUpdate(g, g.GetSeat(e.Username))
	mode.Stack.Activate(g)
}

func (mode *GraveBirthMode) OnReceive(e *games.Event, g *games.Game, s *games.Seat, json js.Object) {
	log := g.Log().Add("Username", s.Username).Add("Choice", json.Val("choice"))

	if s.Username != e.Username {
		log.Add("HotSeat", e.Username).Warn(GraveBirthID + ": not your choice")
		return
	}

	if gcid := json.Sval("choice"); gcid == "" {
		log.Warn(GraveBirthID + ": choice not found")
	} else if card := g.Cards[gcid]; card == nil {
		log.Warn(GraveBirthID + ": gcid not found")
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Warn(GraveBirthID + ": not type body")
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Add("CardOwner", card.Username).Warn(GraveBirthID + ": card owner not found")
	} else if !ownerSeat.HasPastCard(gcid) {
		log.Add("CardOwner", card.Username).Add("Past", ownerSeat.Graveyard.String()).Warn(GraveBirthID + ": card not in past")
	} else {
		mode.Card = vii.NewGameCard(card.Card, card.CardText)
		log.Add("CardId", mode.Card.Card.Id).Info(GraveBirthID + ": confirmed card")
		g.TimelineJoin(nil)
	}
}

func GraveBirth(g *games.Game, s *games.Seat, target interface{}) {
	event := games.NewEvent(s.Username)
	event.EMode = &GraveBirthMode{
		Stack: g.Active,
	}
	g.TimelineJoin(event)
}
