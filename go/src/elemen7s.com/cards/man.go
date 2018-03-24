package cards

import (
	"elemen7s.com"
	"elemen7s.com/db"
	"errors"
	"fmt"
)

var CardCache = make(map[int]*Card)

func Test(id int) *Card {
	return CardCache[id]
}

func LoadCache() error {
	// select all cards
	rows, err := db.Connection.Query("SELECT id, type, image FROM cards")
	if err != nil {
		rows.Close()
		return err
	}

	for rows.Next() {
		card, err := scanCard(rows)
		if err != nil {
			rows.Close()
			return err
		}
		CardCache[card.Id] = card
	}
	rows.Close()

	if err = loadCardBodies(); err != nil {
		return err
	} else if err = loadCardCosts(); err != nil {
		return err
	} else if err = loadCardsPowers(); err != nil {
		return err
	} else if err = loadCardsPowersCosts(); err != nil {
		return err
	}

	return nil
}

func loadCardBodies() error {
	rows, err := db.Connection.Query("SELECT cardid, attack, health FROM cards_bodies")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid int
		body := &Body{}
		err = rows.Scan(&cardid, &body.Attack, &body.Health)

		if err != nil {
			return err
		} else if CardCache[cardid] == nil {
			return errors.New(fmt.Sprintf("cards: body matching missing card#%v", cardid))
		}

		CardCache[cardid].Body = body
	}

	return nil
}

func loadCardCosts() error {
	rows, err := db.Connection.Query("SELECT cardid, element, count FROM cards_element_costs")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid, elementid, count int
		err = rows.Scan(&cardid, &elementid, &count)

		if err != nil {
			return err
		} else if elementid > len(vii.Elements) || elementid < 0 {
			return errors.New(fmt.Sprintf("cards: invalid element#%v", elementid))
		} else if CardCache[cardid] == nil {
			return errors.New(fmt.Sprintf("cards: cost matching missed card#%v", cardid))
		}

		CardCache[cardid].Costs[vii.Elements[elementid]] += count
	}

	return nil
}

func loadCardsPowers() error {
	rows, err := db.Connection.Query("SELECT cardid, id, usesturn, script FROM cards_powers")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid int
		power := NewPower()
		err = rows.Scan(&cardid, &power.Id, &power.UsesTurn, &power.Script)

		if err != nil {
			return err
		} else if card := CardCache[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power card#%v id#%v", cardid, power.Id))
		} else if power := card.Powers[power.Id]; power != nil {
			return errors.New(fmt.Sprintf("cards: duplicate power card#%v id#%v", cardid, power.Id))
		}

		CardCache[cardid].Powers[power.Id] = power
	}

	return nil
}

func loadCardsPowersCosts() error {
	rows, err := db.Connection.Query("SELECT cardid, powerid, element, count FROM cards_powers_costs")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid, powerid, elementid, count int
		err = rows.Scan(&cardid, &powerid, &elementid, &count)

		if err != nil {
			return err
		} else if card := CardCache[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power cost card#%v id#%v", cardid, powerid))
		} else if card.Powers[powerid] == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power cost card#%v id#%v", cardid, powerid))
		}

		CardCache[cardid].Powers[powerid].Costs[vii.Elements[elementid]] += count
	}

	return nil
}

func scanCard(scanner db.Scanner) (*Card, error) {
	card := NewCard()
	var cardtypebuff int

	err := scanner.Scan(&card.Id, &cardtypebuff, &card.Image)
	if err != nil {
		return nil, err
	}

	if ctype := vii.CardType(cardtypebuff); ctype.String() == "error" {
		return nil, errors.New(fmt.Sprintf("cards: cardtype not recognized #%v", cardtypebuff))
	} else {
		card.CardType = ctype
	}

	return card, nil
}
