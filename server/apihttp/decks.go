package apihttp

// import (
// 	"fmt"

// 	"github.com/zachtaylor/7elements"
// 	"ztaylor.me/http"
// 	"ztaylor.me/js"
// )

// func DecksHandler(r *http.Quest) error {
// 	decks, err := vii.DeckService.GetAll()
// 	if err != nil {
// 		return err
// 	}

// 	j := js.Object{}
// 	for deckid, deck := range decks {
// 		j[fmt.Sprintf("%d", deckid)] = deck.Json()
// 	}
// 	r.WriteJson(j)
// 	return nil
// }
