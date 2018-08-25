package vii

type Card struct {
	Id     int
	Name   string
	Text   string
	Type   CardType
	Image  string
	Costs  ElementMap
	Body   *CardBody
	Powers Powers
}

func NewCard() *Card {
	return &Card{
		Costs:  ElementMap{},
		Powers: NewPowers(),
	}
}

func (c *Card) GetPlayPower() *Power {
	for _, p := range c.Powers {
		if p.Trigger == "play" {
			return p
		}
	}
	return nil
}

func (c *Card) GetDeathPower() *Power {
	for _, p := range c.Powers {
		if p.Trigger == "death" {
			return p
		}
	}
	return nil
}

func (c *Card) Json() Json {
	json := Json{
		"id":     c.Id,
		"image":  c.Image,
		"name":   c.Name,
		"text":   c.Text,
		"type":   c.Type.String(),
		"powers": c.Powers.Json(),
		"costs":  c.Costs.Copy(),
		"body":   c.Body.Json(),
	}

	if power := c.GetPlayPower(); power != nil {
		json["target"] = power.Target
	}

	return json
}

var CardService interface {
	Start() error
	Get(cardid int) (*Card, error)
	GetAll() map[int]*Card
}
