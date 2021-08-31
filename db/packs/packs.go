package packs

import (
	"github.com/zachtaylor/7elements/card/pack"
	"taylz.io/db"
)

func GetAll(conn *db.DB) (pack.Prototypes, error) {
	packs, err := get(conn)
	if err != nil {
		return nil, err
	}
	packscards, err := getCards(conn)
	if err != nil {
		return nil, err
	}

	for _, chance := range packscards {
		if pack := packs[chance.PackID]; pack != nil {
			pack.Cards = append(pack.Cards, chance)
		}
	}

	return packs, nil
}

func get(conn *db.DB) (pack.Prototypes, error) {
	packs := make(pack.Prototypes)
	rows, err := conn.Query(`SELECT id, name, size, cost FROM packs`)
	if err != nil {
		return packs, err
	}
	defer rows.Close()
	for rows.Next() {
		pack := pack.NewPrototype()
		if err = rows.Scan(&pack.ID, &pack.Name, &pack.Size, &pack.Cost); err == nil {
			packs[pack.ID] = pack
		} else {
			break
		}
	}
	return packs, err
}

func getCards(conn *db.DB) ([]*pack.Chance, error) {
	packscards := make([]*pack.Chance, 0)
	rows, err := conn.Query(`SELECT packid, cardid, weight FROM packs_cards`)
	if err != nil {
		return packscards, err
	}
	defer rows.Close()
	for rows.Next() {
		chance := &pack.Chance{}
		if err = rows.Scan(&chance.PackID, &chance.CardID, &chance.Weight); err == nil {
			packscards = append(packscards, chance)
		} else {
			break
		}
	}
	return packscards, err
}
