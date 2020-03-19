package db

import (
	"errors"
	"fmt"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/power"
	"ztaylor.me/db"
)

func NewCardService(db *db.DB) card.PrototypeService {
	return &CardService{
		conn:  db,
		cache: make(card.Prototypes),
	}
}

type CardService struct {
	conn  *db.DB
	cache card.Prototypes
}

func (cs *CardService) Get(id int) (*card.Prototype, error) {
	return cs.cache[id], nil
}

func (cs *CardService) GetAll() card.Prototypes {
	if len(cs.cache) == 0 {
		cs.Start()
	}
	return cs.cache
}

func (cs *CardService) Start() error {
	// select all cards
	rows, err := cs.conn.Query("SELECT id, type, name, text, image FROM cards")
	if err != nil {
		rows.Close()
		return err
	}

	for rows.Next() {
		c, err := cs.scanCard(rows)
		if err != nil {
			rows.Close()
			return err
		}
		cs.cache[c.ID] = c
	}
	rows.Close()

	if err = cs.loadCardBodies(); err != nil {
		return err
	} else if err = cs.loadCardCosts(); err != nil {
		return err
	} else if err = cs.loadCardsPowers(); err != nil {
		return err
	} else if err = cs.loadCardsPowersCosts(); err != nil {
		return err
	}

	return nil
}

func (cs *CardService) loadCardBodies() error {
	rows, err := cs.conn.Query("SELECT cardid, attack, health FROM cards_bodies")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid int
		body := &card.Body{}
		err = rows.Scan(&cardid, &body.Attack, &body.Health)

		if err != nil {
			return err
		} else if cs.cache[cardid] == nil {
			return errors.New(fmt.Sprintf("cards: body matching missing card#%v", cardid))
		}

		cs.cache[cardid].Body = body
	}

	return nil
}

func (cs *CardService) loadCardCosts() error {
	rows, err := cs.conn.Query("SELECT cardid, element, count FROM cards_element_costs")
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
		} else if card := cs.cache[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: cost matching missed card#%v", cardid))
		} else {
			card.Costs[element.T(elementid)] += count
		}
	}

	return nil
}

func (cs *CardService) loadCardsPowers() error {
	rows, err := cs.conn.Query("SELECT cardid, id, usesturn, useskill, xtrigger, target, script, text FROM cards_powers")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid, usesturn, useskill int
		p := power.New()
		err = rows.Scan(&cardid, &p.ID, &usesturn, &useskill, &p.Trigger, &p.Target, &p.Script, &p.Text)
		p.UsesTurn = usesturn > 0
		p.UsesKill = useskill > 0

		if err != nil {
			return err
		} else if card := cs.cache[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power card#%v id#%v", cardid, p.ID))
		} else if card.Powers[p.ID] != nil {
			return errors.New(fmt.Sprintf("cards: duplicate power card#%v id#%v", cardid, p.ID))
		} else {
			card.Powers[p.ID] = p
		}
	}

	return nil
}

func (cs *CardService) loadCardsPowersCosts() error {
	rows, err := cs.conn.Query("SELECT cardid, powerid, element, count FROM cards_powers_costs")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cardid, powerid, elementid, count int
		err = rows.Scan(&cardid, &powerid, &elementid, &count)

		if err != nil {
			return err
		} else if card := cs.cache[cardid]; card == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power cost card#%v id#%v", cardid, powerid))
		} else if power := card.Powers[powerid]; power == nil {
			return errors.New(fmt.Sprintf("cards: unrooted power cost card#%v id#%v", cardid, powerid))
		} else {
			power.Costs[element.Index[elementid]] += count
		}
	}

	return nil
}

func (cs *CardService) scanCard(scanner db.Scanner) (*card.Prototype, error) {
	c := card.NewPrototype()
	var typebuff int

	err := scanner.Scan(&c.ID, &typebuff, &c.Name, &c.Text, &c.Image)
	if err != nil {
		return nil, err
	}

	if t := card.Type(typebuff); t.String() == "error" {
		return nil, errors.New(fmt.Sprintf("cards: cardtype not recognized #%v", typebuff))
	} else {
		c.Type = t
	}

	return c, nil
}
