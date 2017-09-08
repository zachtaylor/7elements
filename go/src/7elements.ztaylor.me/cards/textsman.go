package cards

import (
	"7elements.ztaylor.me/db"
	"errors"
)

var TextsCache = make(map[string]map[int]*Texts)

func LoadTextsCache(lang string) error {
	rows, err := db.Connection.Query("SELECT cardid, language, name, description, flavor FROM cards_text WHERE language=?",
		lang,
	)
	if err != nil {
		return err
	}

	texts := TextsCache[lang]
	if texts == nil {
		texts = make(map[int]*Texts)
		TextsCache[lang] = texts
	}

	for rows.Next() {
		cardText, ok := scanCardText(rows)
		if !ok {
			return errors.New("cards_text: load cache")
		}
		texts[cardText.CardId] = cardText
	}
	rows.Close()

	return nil
}

func scanCardText(scanner db.Scanner) (*Texts, bool) {
	cardtext := &Texts{}
	err := scanner.Scan(&cardtext.CardId, &cardtext.Language, &cardtext.Name, &cardtext.Description, &cardtext.Flavor)
	if err != nil {
		return nil, false
	}
	return cardtext, true
}
