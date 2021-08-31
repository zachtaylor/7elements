package main

// import (
// 	_ "github.com/zachtaylor/7elements/scripts"
// 	"taylz.io/db"
// 	"taylz.io/db/mysql"
// 	"taylz.io/log"
// )

// Patch 5 changes the Database Schema and adapts existing account data
//
// Account decks (accounts_decks, accounts_decks_items) encoded to single string (accounts_decks)
//
// TODO

const SelectAccountDecks = `SELECT username, id, name, cover FROM accounts_decks`

const InsertDeck = `INSERT INTO decks(name, user, cover) VALUES (?, ?, ?)`

type AccountDeck struct {
	Username string
	Name     string
	ID       int
	Cover    int
}

func main() {
	// stdout := log.StdOutService(log.LevelDebug)
	// stdout.Formatter().CutSourcePath(2)
	// env := db.ENV().ParseDefault()
	// conn, err := mysql.Open(pkg_env.BuildDSN(env.Match("DB_")))
	// if err != nil {
	// 	stdout.Error("db connection failed", err)
	// 	return
	// }
	// stdout.Trace(SelectAccountDecks)
	// res, err := conn.Query(SelectAccountDecks)
	// if err != nil {
	// 	stdout.Error("select account decks", err)
	// 	return
	// }
	// accounts_decks := make([]AccountDeck, 0)
	// for res.Next() {
	// 	ad := AccountDeck{}
	// 	err = res.Scan(&ad.Username, &ad.Name, &ad.ID, &ad.Cover)
	// 	if err != nil {
	// 		stdout.Error("parse account deck", err)
	// 		return
	// 	}
	// 	stdout.Trace("parse account deck", ad)
	// 	accounts_decks = append(account_decks, ad)
	// }
	// res.Close()
	// stdout.Trace("closed accounts_decks")
	// stdout.Info("select accounts_decks", len(accounts_decks))
	// id_map := make(map[string]map[int]int)
	// for _, ad := range accounts_decks {
	// 	res, err = conn.Exec(InsertDeck, ad.Name, ad.User, ad.Cover)
	// 	id_map[res.LastInsertId()]
	// }
}
