package cards

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/power"
	"taylz.io/db"
)

func GetAll(conn *db.DB) card.Prototypes {
	cards := card.Prototypes{}

	// select all cards
	rows, err := conn.Query("SELECT id, type, name, text FROM cards")
	if err != nil {
		return nil
	}

	for rows.Next() {
		c, err := scanCard(rows)
		if err != nil {
			rows.Close()
			return nil
		}
		cards[c.ID] = c
	}
	rows.Close()

	if err = loadCardBodies(conn, cards); err != nil {
		return nil
	} else if err = loadCardCosts(conn, cards); err != nil {
		return nil
	} else if err = loadCardsPowers(conn, cards); err != nil {
		return nil
	} else if err = loadCardsPowersCosts(conn, cards); err != nil {
		return nil
	}

	return cards
}

func scanCard(scanner db.Scanner) (*card.Prototype, error) {
	c := card.NewPrototype()
	var kind int

	if err := scanner.Scan(&c.ID, &kind, &c.Name, &c.Text); err != nil {
		return nil, err
	} else if k := card.Kind(kind); k.String() == "error" {
		return nil, errors.New("cards: cardtype not recognized #" + strconv.FormatInt(int64(kind), 10))
	} else {
		c.Kind = k
	}

	return c, nil
}

func loadCardBodies(conn *db.DB, cards card.Prototypes) error {
	rows, err := conn.Query("SELECT cardid, attack, health FROM cards_bodies")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid int
		body := &card.Body{}
		err = rows.Scan(&cardid, &body.Attack, &body.Life)

		if err != nil {
			return err
		} else if cards[cardid] == nil {
			return errors.New(fmt.Sprintf("cards: body matching missing card#%v", cardid))
		}

		cards[cardid].Body = body
	}

	return nil
}

func loadCardCosts(conn *db.DB, cards card.Prototypes) error {
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
		} else if elementid > len(element.Index) || elementid < 0 {
			return errors.New(fmt.Sprintf("cards: invalid element#%v", elementid))
		} else if card := cards[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: cost matching missed card#%v", cardid))
		} else {
			card.Costs[element.T(elementid)] += count
		}
	}

	return nil
}

func loadCardsPowers(conn *db.DB, cards card.Prototypes) error {
	rows, err := conn.Query("SELECT cardid, id, usesturn, useskill, xtrigger, target, script, text FROM cards_powers")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid, usesturn, useskill int
		p := power.New()
		err = rows.Scan(&cardid, &p.ID, &usesturn, &useskill, &p.Trigger, &p.Target, &p.Script, &p.Text)
		p.UsesTurn = usesturn > 0
		p.UsesLife = useskill > 0

		if err != nil {
			return err
		} else if card := cards[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power card#%v id#%v", cardid, p.ID))
		} else if card.Powers[p.ID] != nil {
			return errors.New(fmt.Sprintf("cards: duplicate power card#%v id#%v", cardid, p.ID))
		} else {
			card.Powers[p.ID] = p
		}
	}

	return nil
}

func loadCardsPowersCosts(conn *db.DB, cards card.Prototypes) error {
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
			power.Costs[element.Index[elementid]] += count
		}
	}

	return nil
}
