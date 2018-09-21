package db

import (
	"github.com/zachtaylor/7elements"
)

func init() {
	vii.PackService = &PackService{}
}

type PackService struct {
	Packs map[int]*vii.Pack
}

func (service *PackService) Start() error {
	packs, err := service.reloadPacks()
	if err != nil {
		return err
	}
	packscards, err := service.reloadPacksCards()
	if err != nil {
		return err
	}

	for _, chance := range packscards {
		if pack := packs[chance.PackID]; pack != nil {
			pack.Cards = append(pack.Cards, chance)
		}
	}

	service.Packs = packs
	return nil
}

func (service *PackService) reloadPacks() (vii.Packs, error) {
	packs := make(vii.Packs)
	rows, err := Conn.Query(`SELECT id, name, size, cost, image FROM packs`)
	if err != nil {
		return packs, err
	}
	defer rows.Close()
	for rows.Next() {
		pack := vii.NewPack()
		if err = rows.Scan(&pack.ID, &pack.Name, &pack.Size, &pack.Cost, &pack.Image); err == nil {
			packs[pack.ID] = pack
		} else {
			break
		}
	}
	return packs, err
}

func (service *PackService) reloadPacksCards() ([]*vii.PackChance, error) {
	packscards := make([]*vii.PackChance, 0)
	rows, err := Conn.Query(`SELECT packid, cardid, weight FROM packs_cards`)
	if err != nil {
		return packscards, err
	}
	defer rows.Close()
	for rows.Next() {
		chance := &vii.PackChance{}
		if err = rows.Scan(&chance.PackID, &chance.CardID, &chance.Weight); err == nil {
			packscards = append(packscards, chance)
		} else {
			break
		}
	}
	return packscards, err
}

func (service *PackService) Get(id int) (*vii.Pack, error) {
	packs, err := service.GetAll()
	if packs == nil {
		return nil, err
	}
	return packs[id], err
}

func (service *PackService) GetAll() (vii.Packs, error) {
	if service.Packs == nil {
		service.Start()
	}
	return service.Packs, nil
}
