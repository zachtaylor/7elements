package vii

// Deck is a system premade deck
type Deck struct {
	ID    int
	Name  string
	Level int
	Color string
	Cards map[int]int
}

func NewDeck() *Deck {
	return &Deck{
		Cards: make(map[int]int),
	}
}

func (deck *Deck) Count() int {
	total := 0
	for _, count := range deck.Cards {
		total += count
	}
	return total
}

func (deck *Deck) Json() Json {
	return Json{
		"id":    deck.ID,
		"name":  deck.Name,
		"level": deck.Level,
		"color": deck.Color,
		"cards": deck.Cards,
	}
}

type Decks map[int]*Deck

func (decks Decks) Json() []Json {
	data := make([]Json, 0)
	for _, deck := range decks {
		data = append(data, deck.Json())
	}
	return data
}

// DeckService provides access to Decks
var DeckService interface {
	GetAll() (Decks, error)
	Get(int) (*Deck, error)
}
