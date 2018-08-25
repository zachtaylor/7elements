package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/zachtaylor/7elements/cards"
	"github.com/zachtaylor/7elements/cards/texts"
	"github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/decks"
	"github.com/zachtaylor/7elements/games"
	"github.com/zachtaylor/7elements/options"
	"ztaylor.me/log"
)

func main() {
	fmt.Print("starting... press enter to kill\n")
	reader := bufio.NewReader(os.Stdin)
	defer reader.ReadString('\n')

	g := games.New()
	log.SetLevel("debug")
	g.Logger.SetLevel("debug")

	db.Open(options.String("db-path"))
	cards.LoadCache()
	texts.LoadCache("en-US")

	decka := decks.New()
	decka.Username = "YIN"
	decka.Cards[1] = 3
	decka.Cards[2] = 3
	decka.Cards[3] = 3
	decka.Cards[4] = 3
	decka.Cards[5] = 3
	decka.Cards[6] = 3
	decka.Cards[7] = 3

	deckb := decks.New()
	deckb.Username = "YANG"
	deckb.Cards[1] = 3
	deckb.Cards[2] = 3
	deckb.Cards[3] = 3
	deckb.Cards[4] = 3
	deckb.Cards[5] = 3
	deckb.Cards[6] = 3
	deckb.Cards[7] = 3

	sa := g.Register(decka, "en-US")
	sb := g.Register(deckb, "en-US")

	games.ConnectAI(g, sa)
	games.ConnectAI(g, sb)

	games.Start(g)
}

func delay(d time.Duration, f func()) {
	go (func() {
		<-time.After(d)
		f()
	})()
}
