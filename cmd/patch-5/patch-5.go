package main

import (
	pkg_env "ztaylor.me/db/env"
	_ "github.com/zachtaylor/7elements/scripts"
	"ztaylor.me/db/mysql"
	"ztaylor.me/log"
)

const SelectAccountDecks = `SELECT username, id, name, cover FROM accounts_decks`

const InsertDeck = `INSERT INTO decks(name, user, cover) VALUES (?, ?, ?)`

type AccountDeck struct {
	Username string
	Name     string
	ID       int
	Cover    int
}

func main() {
	stdout := log.StdOutService(log.LevelDebug)
	stdout.Formatter().CutSourcePath(2)
	env := pkg_env.NewService().ParseDefault()
	conn, err := mysql.Open(pkg_env.BuildDSN(env.Match("DB_")))
	if err != nil {
		stdout.Error("db connection failed", err)
		return
	}
	stdout.Trace(SelectAccountDecks)
	res, err := conn.Query(SelectAccountDecks)
	if err != nil {
		stdout.Error("select account decks", err)
		return
	}
	accounts_decks := make([]AccountDeck, 0)
	for res.Next() {
		ad := AccountDeck{}
		err = res.Scan(&ad.Username, &ad.Name, &ad.ID, &ad.Cover)
		if err != nil {
			stdout.Error("parse account deck", err)
			return
		}
		stdout.Trace("parse account deck", ad)
		accounts_decks = append(account_decks, ad)
	}
	res.Close()
	stdout.Trace("closed accounts_decks")
	stdout.Info("select accounts_decks", len(accounts_decks))
	id_map := make(map[string]map[int]int)
	for _, ad := range accounts_decks {
		res, err = conn.Exec(InsertDeck, ad.Name, ad.User, ad.Cover)
		id_map[res.LastInsertId()
	}
}
