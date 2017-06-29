package persistence

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"errors"
)

func CardsLoadCache() error {
	rows, err := connection.Query("SELECT id, type, image FROM cards")
	if err != nil {
		return err
	}

	for rows.Next() {
		card, ok := scanCard(rows)
		if !ok {
			log.Error("cards: load cache: scan card")
			return errors.New("cards: load cache")
		}
		SE.Cards.Cache[card.Id] = card
	}
	rows.Close()

	rows, err = connection.Query("SELECT cardid, element, count FROM cards_element_costs")
	if err != nil {
		return err
	}

	for rows.Next() {
		cardid, elementcost, ok := scanElementCost(rows)
		if !ok {
			log.Error("cards: load cache: scan element cost")
			return errors.New("cards: load cache")
		}

		card, ok := SE.Cards.Cache[cardid]

		if !ok {
			log.Error("cards: load cache: element cost missing card")
			return errors.New("cards: load cache")
		}

		card.ElementCosts = append(card.ElementCosts, elementcost)
	}
	rows.Close()

	log.Add("Cards", len(SE.Cards.Cache)).Debug("cards: load cache")
	return nil
}

func CardsInsert(id int) error {
	log.Add("CardId", id).Warn("cards: insert")
	return nil
}

func CardsDelete(id int) error {
	log.Add("CardId", id).Warn("cards: delete")
	return nil
}

func scanCard(scanner Scanner) (*SE.Card, bool) {
	card := &SE.Card{ElementCosts: make([]*SE.ElementCost, 0)}
	var cardtypebuff int

	err := scanner.Scan(&card.Id, &cardtypebuff, &card.Image)
	if err != nil {
		log.Add("Error", err)
		return nil, false
	}

	if cardtypebuff > len(SE.CardTypes) {
		log.Add("CardType", cardtypebuff).Add("Error", "scancard: type out of range")
		return nil, false
	}

	card.CardType = SE.CardTypes[cardtypebuff]
	return card, true
}

func scanElementCost(scanner Scanner) (int, *SE.ElementCost, bool) {
	var cardid int
	var elementbuff uint
	elementcost := &SE.ElementCost{}

	err := scanner.Scan(&cardid, &elementbuff, &elementcost.Count)
	if err != nil {
		log.Add("CardId", cardid).Add("Element", elementbuff).Add("Count", elementcost.Count).Add("Error", err)
		return 0, nil, false
	}

	if int(elementbuff) > len(SE.Elements) {
		log.Add("Element", elementbuff).Add("Error", "scanelementcost: type out of range")
		return 0, nil, false
	}

	elementcost.Element = SE.Elements[elementbuff]
	return cardid, elementcost, true
}
