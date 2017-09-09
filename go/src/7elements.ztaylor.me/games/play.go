package games

import (
	"7elements.ztaylor.me/cards"
	"7elements.ztaylor.me/cards/types"
	"7elements.ztaylor.me/games/cards"
	"ztaylor.me/ctxpert"
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func Play(game *Game, socket *sessions.Socket, data json.Json, log log.Log) {
	seat := game.GetSeat(socket.Username)
	log.Add("GCID", data["gcid"])

	if gcid, ok := data["gcid"].(float64); !ok {
		log.Warn("play: gcid error")
		return
	} else if game.GamePhase != GPHSplay {
		log.Add("GamePhase", game.GamePhase).Warn("play: rejected")
		return
	} else if game.TurnPhase != TPHSplay {
		log.Add("TurnPhase", game.TurnPhase).Warn("play: rejected")
		return
	} else if pos := seat.CardHandPositionGCID(int(gcid)); pos < 0 {
		log.Warn("play: not in hand")
		return
	} else if gcard := seat.Hand[pos]; gcard == nil {
		log.Warn("play: card in hand has no cardid")
		return
	} else if card := cards.Test(gcard.CardId); card == nil {
		log.Add("CardId", gcard.CardId).Warn("play: card missing")
		return
	} else if !seat.Elements.TestStack(card.Costs) {
		log.Warn("play: need more elements")
		socket.Send("animate", json.Json{
			"animate": "not enough elements",
		})
		return
	} else {
		seat.Elements.Deactivate(card.Costs)
		gcard = seat.RemoveHandPosition(pos)

		// start respond phase
		game.GamePhase = GPHSrespond
		playCtx := ctxpert.WithNewTimeout(game.Context, game.Context.Timer()+game.Patience)
		game.Context = ctxpert.WithNewTimeout(ctxpert.New(), game.Patience)
		game.Context.Always(func(ctx *ctxpert.Context) {
			game.GamePhase = GPHSplay
			game.Context = ctxpert.WithNewTimeout(playCtx, game.Patience)
			SendAllSeats(game, "turn", MakeTurnJson(game, game.CurrentTurn()))
		})
		// end

		if card.CardType == ctypes.Body || card.CardType == ctypes.Item {
			seat.Active = append(seat.Active, gcard)
			log.Info("play")

			socket.Send("hand", json.Json{
				"cards": gamecards.Stack(seat.Hand).Json(),
			})

			SendAllSeats(game, "play", json.Json{
				"username": seat.Username,
				"cardid":   gcard.CardId,
				"gcid":     gcard.GameCardId,
				"timer":    int(game.Context.Timer().Seconds() + 1),
			})

			SendAllSeats(game, "elements", json.Json{
				"username": seat.Username,
				"elements": seat.Elements,
			})
		} else if card.CardType == ctypes.Spell {
			log.Warn("play: spell types dont work yet")
		} else {
			log.Add("CardType", card.CardType).Warn("play: card type not recognized")
		}
	}
}
