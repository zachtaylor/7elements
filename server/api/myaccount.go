package api

import (
	"time"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

func MyAccountHandler(r *http.Request) error {
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := vii.AccountService.Get(r.Username); account == nil {
		return err
	} else if accountcards, err := vii.AccountCardService.Get(r.Username); err != nil {
		return err
	} else if accountdecks, err := vii.AccountDeckService.Get(r.Username); err != nil {
		return err
	} else {
		games := make([]js.Object, 0)
		for _, gameid := range vii.GameService.GetPlayerGames(r.Username) {
			if game := vii.GameService.Get(gameid); game != nil {
				games = append(games, game.Json(r.Username))
			}
		}
		r.WriteJson(js.Object{
			"username":    r.Username,
			"email":       account.Email,
			"sessionlife": r.Expire.Sub(time.Now()).String(),
			"coins":       account.Coins,
			"cards":       accountcards.Json(),
			"decks":       accountdecks.Json(),
			"games":       games,
		})
		return nil
	}
}
