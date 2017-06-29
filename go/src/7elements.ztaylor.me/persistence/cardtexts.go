package persistence

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"errors"
)

func CardTextsLoadCache(lang string) error {
	rows, err := connection.Query("SELECT cardid, language, name, description, flavor FROM cards_text WHERE language=?",
		lang,
	)
	if err != nil {
		return err
	}

	texts := SE.CardTexts.Cache[lang]
	if texts == nil {
		texts = make(map[int]*SE.CardText)
		SE.CardTexts.Cache[lang] = texts
	}

	cardnames := make([]string, 0)

	for rows.Next() {
		cardText, ok := scanCardText(rows)
		if !ok {
			log.Error("cards_text: load cache: scan")
			return errors.New("cards_text: load cache")
		}

		cardnames = append(cardnames, cardText.Name)
		texts[cardText.CardId] = cardText
	}
	rows.Close()

	log.Add("Language", lang).Add("CardTexts", cardnames).Debug("cards_text: load cache")
	return nil
}

func scanCardText(scanner Scanner) (*SE.CardText, bool) {
	cardtext := &SE.CardText{}
	err := scanner.Scan(&cardtext.CardId, &cardtext.Language, &cardtext.Name, &cardtext.Description, &cardtext.Flavor)
	if err != nil {
		log.Add("Error", err)
		return nil, false
	}
	return cardtext, true
}
