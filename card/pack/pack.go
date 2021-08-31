package pack

import (
	"sort"

	"github.com/cznic/mathutil"
	"taylz.io/http/websocket"
)

// Prototype is a definition of a Pack of Cards
type Prototype struct {
	ID    int
	Name  string
	Size  int
	Cost  int
	Cards []*Chance
}

// Chance is a measure of distribution of Cards within a Pack
type Chance struct {
	PackID int
	CardID int
	Weight int
}

// NewPrototype returns a new Prototype
func NewPrototype() *Prototype {
	return &Prototype{
		Cards: make([]*Chance, 0),
	}
}

var rand, _ = mathutil.NewFC32(1, 98, true)

// NewPack returns a new []int containing random and non-duplicate Card IDs from this Pack Prototype
func (p *Prototype) NewPack() []int {
	cards := make([]int, p.Size)

	for i, lenp := 0, len(p.Cards); i < p.Size; i++ {
		cardid := 0
		for ok := true; ok; ok = newPackSkipProtoID(cards, cardid) {
			cardid = p.Cards[int(rand.Next())%lenp].CardID
		}
		cards[i] = cardid
	}

	return cards
}

func newPackSkipProtoID(pack []int, id int) bool {
	for _, cardID := range pack {
		if cardID == id {
			return true
		}
	}
	return false
}

// JSON returns a representation of this Prototype as type websocket.MsgData
func (p *Prototype) JSON() websocket.MsgData {
	cards := make([]int, 0)
	for _, card := range p.Cards {
		cards = append(cards, card.CardID)
	}
	return websocket.MsgData{
		"id":    p.ID,
		"name":  p.Name,
		"size":  p.Size,
		"cost":  p.Cost,
		"cards": cards,
	}
}

// Prototypes is a set of Prototypes mapped by id
type Prototypes map[int]*Prototype

// JSON returns a representation of these Prototypes as type fmt.Stringer
func (packs Prototypes) JSON() []websocket.MsgData {
	json := make([]websocket.MsgData, len(packs))
	keys := make([]int, len(packs))
	var i int
	for k := range packs {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	var j int
	for _, k := range keys {
		json[j] = packs[k].JSON()
		j++
	}
	return json
}

// Service is used to acquire Pack Prototypes
type Service interface {
	Get(int) (*Prototype, error)
	GetAll() (Prototypes, error)
}
