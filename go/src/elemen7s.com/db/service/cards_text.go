package dbservice

import (
	"elemen7s.com"
	"elemen7s.com/db"
	"errors"
	"fmt"
)

func init() {
	vii.CardTextService = CardTextService{}
}

type CardTextService map[string]map[int]*vii.CardText

func (texts CardTextService) GetCardText(lang string, cardid int) (*vii.CardText, error) {
	return texts.GetAll(lang)[cardid], nil
}

func (texts CardTextService) GetAll(lang string) map[int]*vii.CardText {
	if texts[lang] == nil {
		texts[lang] = make(map[int]*vii.CardText)
	}
	return texts[lang]
}

func (texts CardTextService) Start() error {
	return texts.LoadLanguage("en-US")
}

func (texts CardTextService) LoadLanguage(lang string) error {
	rows, err := db.Connection.Query("SELECT cardid, language, name, description, flavor FROM cards_text WHERE language=?",
		lang,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	langText := texts.GetAll(lang)

	for rows.Next() {
		text := vii.NewCardText()
		var cardid int
		var language string
		err := rows.Scan(&cardid, &language, &text.Name, &text.Description, &text.Flavor)
		if err != nil {
			return err
		}
		langText[cardid] = text
	}

	if err = texts.loadPowersCache(lang); err != nil {
		return err
	}

	return nil
}

func (texts CardTextService) loadPowersCache(lang string) error {
	rows, err := db.Connection.Query("SELECT cardid, powerid, description FROM cards_powers_texts WHERE language=?",
		lang,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	langText := texts.GetAll(lang)

	for rows.Next() {
		var cardid, powerid int
		var text string
		err = rows.Scan(&cardid, &powerid, &text)

		if err != nil {
			return err
		} else if text := langText[cardid]; text == nil {
			return errors.New(fmt.Sprintf("power-texts: missing card#%v powerid#%v", cardid, powerid))
		} else if text.Powers[powerid] != "" {
			return errors.New(fmt.Sprintf("power-texts: duplicate card#%v powerid#%v", cardid, powerid))
		}

		langText[cardid].Powers[powerid] = text
	}

	return nil
}
