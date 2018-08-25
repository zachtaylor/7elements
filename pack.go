package vii

type Pack struct {
	ID    int
	Name  string
	Size  int
	Cost  int
	Image string
	Cards []*PackChance
}

type PackChance struct {
	PackID int
	CardID int
	Weight int
}

type Packs map[int]*Pack

func NewPack() *Pack {
	return &Pack{
		Cards: make([]*PackChance, 0),
	}
}

func (p *Pack) Json() Json {
	cards := make([]int, 0)
	for _, card := range p.Cards {
		cards = append(cards, card.CardID)
	}
	return Json{
		"id":    p.ID,
		"name":  p.Name,
		"size":  p.Size,
		"cost":  p.Cost,
		"image": p.Image,
		"cards": cards,
	}
}

func (packs Packs) Json() []Json {
	data := make([]Json, 0)
	for _, pack := range packs {
		data = append(data, pack.Json())
	}
	return data
}

var PackService interface {
	Get(int) (*Pack, error)
	GetAll() (Packs, error)
}
