package main

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/decks"
	"github.com/zachtaylor/7elements/server/security"
	"time"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

func main() {
	env.Bootstrap()

	log.SetLevel(env.Default("LOG_LEVEL", "info"))

	tStart := time.Now()

	db.Open(env.Default("DB_PATH", "elemen7s.db"))

	db.Connection.Exec(
		`INSERT INTO accounts (username, email, password, skill, coins, packs, language, register, lastlogin) VALUES ("alpha", "", ?, 1000, 700, 700, "en-US", ?, ?)`,
		security.HashPassword("alpha"),
		time.Now().Unix(),
		time.Now().Unix(),
	)

	decks.Store("alpha", decks.NewDecks())
	for _, deck := range decks.Test("alpha") {
		deck.Username = "alpha"
	}
	if err := decks.Insert("alpha", 0); err != nil {
		log.Add("Error", err).Error("grant decks")
		return
	}

	log.Add("Time", time.Now().Sub(tStart)).Debug("created account")

	for i := 1; i <= 50; i++ {
		ti := time.Now()
		for j := 0; j < 7; j++ {
			vii.AccountCardService.InsertCard(&vii.AccountCard{
				Username: "alpha",
				CardId:   i,
				Register: time.Now(),
				Notes:    "",
			})
		}
		log.Add("CardId", i).Add("Time", time.Now().Sub(ti)).Debug("created accountcards")
	}

	log.Add("TotalTime", time.Now().Sub(tStart)).Info("SUCCESS")
}
