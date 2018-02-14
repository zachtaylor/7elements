package texts

import (
	"elemen7s.com/db"
	"errors"
	"fmt"
)

var Cache = make(map[string]map[int]*Text)

func Get(lang string, cardid int) *Text {
	return GetAll(lang)[cardid]
}

func GetAll(lang string) map[int]*Text {
	if Cache[lang] == nil {
		Cache[lang] = make(map[int]*Text)
	}
	return Cache[lang]
}

func LoadCache(lang string) error {
	rows, err := db.Connection.Query("SELECT cardid, language, name, description, flavor FROM cards_text WHERE language=?",
		lang,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	texts := GetAll(lang)

	for rows.Next() {
		text := New()
		err := rows.Scan(&text.CardId, &text.Language, &text.Name, &text.Description, &text.Flavor)
		if err != nil {
			return err
		}
		texts[text.CardId] = text
	}

	if err = loadPowersCache(lang); err != nil {
		return err
	}

	return nil
}

func loadPowersCache(lang string) error {
	rows, err := db.Connection.Query("SELECT cardid, powerid, description FROM cards_powers_texts WHERE language=?",
		lang,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	texts := GetAll(lang)

	for rows.Next() {
		var cardid, powerid int
		var text string
		err = rows.Scan(&cardid, &powerid, &text)

		if err != nil {
			return err
		} else if text := texts[cardid]; text == nil {
			return errors.New(fmt.Sprintf("power-texts: missing card#%v powerid#%v", cardid, powerid))
		} else if text.Powers[powerid] != "" {
			return errors.New(fmt.Sprintf("power-texts: duplicate card#%v powerid#%v", cardid, powerid))
		}

		texts[cardid].Powers[powerid] = text
	}

	return nil
}
