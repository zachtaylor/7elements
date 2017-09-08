package cards

import (
	"7elements.ztaylor.me/cards/types"
	"7elements.ztaylor.me/db"
	"7elements.ztaylor.me/elements"
	"errors"
	"strconv"
)

var CardCache = make(map[int]*Card)

func Test(id int) *Card {
	return CardCache[id]
}

func LoadCache() error {
	rows, err := db.Connection.Query("SELECT id, type, image FROM cards")
	if err != nil {
		return err
	}

	for rows.Next() {
		card, err := scanCard(rows)
		if err != nil {
			return err
		}
		CardCache[card.Id] = card
	}
	rows.Close()

	rows, err = db.Connection.Query("SELECT cardid, element, count FROM cards_element_costs")
	if err != nil {
		return err
	}

	for rows.Next() {
		cardid, element, count, err := scanElementCost(rows)
		if err != nil {
			return err
		}

		card, ok := CardCache[cardid]

		if !ok {
			return errors.New("card cost matching missed #" + strconv.FormatInt(int64(cardid), 10))
		}

		card.Costs[element] += count
	}
	rows.Close()
	return nil
}

func scanCard(scanner db.Scanner) (*Card, error) {
	card := NewCard()
	var cardtypebuff int

	err := scanner.Scan(&card.Id, &cardtypebuff, &card.Image)
	if err != nil {
		return nil, err
	}

	if cardtypebuff > len(ctypes.CardTypes) {
		return nil, errors.New("cardtype not recognized #" + strconv.FormatInt(int64(cardtypebuff), 10))
	}

	card.CardType = ctypes.CardTypes[cardtypebuff]
	return card, nil
}

func scanElementCost(scanner db.Scanner) (int, elements.Element, int, error) {
	var cardid, elementid, count int

	err := scanner.Scan(&cardid, &elementid, &count)

	if err != nil {
		return 0, elements.Null, 0, err
	}

	if elementid > len(elements.Elements) || elementid < 0 {
		return 0, elements.Null, 0, errors.New("invalid element")
	}

	return cardid, elements.Elements[elementid], count, nil
}
