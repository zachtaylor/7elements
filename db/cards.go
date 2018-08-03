package db

import (
	"errors"
	"fmt"

	"github.com/zachtaylor/7tcg"
	"ztaylor.me/db"
)

func init() {
	vii.CardService = CardService{}
}

type CardService map[int]*vii.Card

func (cards CardService) GetCard(id int) (*vii.Card, error) {
	return cards[id], nil
}

func (cards CardService) GetAllCards() map[int]*vii.Card {
	return map[int]*vii.Card(cards)
}

func (cards CardService) Start() error {
	// select all cards
	rows, err := conn.Query("SELECT id, type, image FROM cards")
	if err != nil {
		rows.Close()
		return err
	}

	for rows.Next() {
		card, err := cards.scanCard(rows)
		if err != nil {
			rows.Close()
			return err
		}
		cards[card.Id] = card
	}
	rows.Close()

	if err = cards.loadCardBodies(); err != nil {
		return err
	} else if err = cards.loadCardCosts(); err != nil {
		return err
	} else if err = cards.loadCardsPowers(); err != nil {
		return err
	} else if err = cards.loadCardsPowersCosts(); err != nil {
		return err
	}

	return nil
}

func (cards CardService) loadCardBodies() error {
	rows, err := conn.Query("SELECT cardid, attack, health FROM cards_bodies")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid int
		body := vii.NewCardBody()
		err = rows.Scan(&cardid, &body.Attack, &body.Health)

		if err != nil {
			return err
		} else if cards[cardid] == nil {
			return errors.New(fmt.Sprintf("cards: body matching missing card#%v", cardid))
		}

		cards[cardid].CardBody = body
	}

	return nil
}

func (cards CardService) loadCardCosts() error {
	rows, err := conn.Query("SELECT cardid, element, count FROM cards_element_costs")
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
		} else if card := cards[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: cost matching missed card#%v", cardid))
		} else {
			card.Costs[vii.Element(elementid)] += count
		}
	}

	return nil
}

func (cards CardService) loadCardsPowers() error {
	rows, err := conn.Query("SELECT cardid, id, usesturn, xtrigger, target, script FROM cards_powers")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid, usesturn int
		power := vii.NewPower()
		err = rows.Scan(&cardid, &power.Id, &usesturn, &power.Trigger, &power.Target, &power.Script)
		power.UsesTurn = usesturn > 0

		if err != nil {
			return err
		} else if card := cards[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power card#%v id#%v", cardid, power.Id))
		} else if card.Powers[power.Id] != nil {
			return errors.New(fmt.Sprintf("cards: duplicate power card#%v id#%v", cardid, power.Id))
		} else {
			card.Powers[power.Id] = power
		}
	}

	return nil
}

func (cards CardService) loadCardsPowersCosts() error {
	rows, err := conn.Query("SELECT cardid, powerid, element, count FROM cards_powers_costs")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid, powerid, elementid, count int
		err = rows.Scan(&cardid, &powerid, &elementid, &count)

		if err != nil {
			return err
		} else if card := cards[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power cost card#%v id#%v", cardid, powerid))
		} else if power := card.Powers[powerid]; power == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power cost card#%v id#%v", cardid, powerid))
		} else {
			power.Costs[vii.Elements[elementid]] += count
		}
	}

	return nil
}

func (cards CardService) scanCard(scanner db.Scanner) (*vii.Card, error) {
	card := vii.NewCard()
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
