package db

import (
	"github.com/zachtaylor/7elements/card/pack"
	"ztaylor.me/db"
)

func NewPackService(db *db.DB) pack.Service {
	return &PackService{
		conn: db,
	}
}

type PackService struct {
	conn  *db.DB
	cache pack.Prototypes
}

func (ps *PackService) Start() error {
	packs, err := ps.reloadPacks()
	if err != nil {
		return err
	}
	packscards, err := ps.reloadPacksCards()
	if err != nil {
		return err
	}

	for _, chance := range packscards {
		if pack := packs[chance.PackID]; pack != nil {
			pack.Cards = append(pack.Cards, chance)
		}
	}

	ps.cache = packs
	return nil
}

func (ps *PackService) reloadPacks() (pack.Prototypes, error) {
	packs := make(pack.Prototypes)
	rows, err := ps.conn.Query(`SELECT id, name, size, cost, image FROM packs`)
	if err != nil {
		return packs, err
	}
	defer rows.Close()
	for rows.Next() {
		pack := pack.NewPrototype()
		if err = rows.Scan(&pack.ID, &pack.Name, &pack.Size, &pack.Cost, &pack.Image); err == nil {
			packs[pack.ID] = pack
		} else {
			break
		}
	}
	return packs, err
}

func (ps *PackService) reloadPacksCards() ([]*pack.Chance, error) {
	packscards := make([]*pack.Chance, 0)
	rows, err := ps.conn.Query(`SELECT packid, cardid, weight FROM packs_cards`)
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

func (ps *PackService) Get(id int) (*pack.Prototype, error) {
	packs, err := ps.GetAll()
	if packs == nil {
		return nil, err
	}
	return packs[id], err
}

func (ps *PackService) GetAll() (pack.Prototypes, error) {
	if ps.cache == nil {
		ps.Start()
	}
	return ps.cache, nil
}
