package vii

import (
	"ztaylor.me/js"
)

type Card struct {
	Id    int
	Image string
	CardType
	Costs ElementMap
	*CardBody
	Powers
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

func (c *Card) JsonWithText(text *CardText) js.Object {
	json := js.Object{
		"id":          c.Id,
		"image":       c.Image,
		"name":        text.Name,
		"type":        c.CardType.String(),
		"description": text.Description,
		"powers":      c.Powers.JsonWithText(text),
		"flavor":      text.Flavor,
		"costs":       c.Costs.Copy(),
		"body":        c.CardBody.Json(),
	}

	if power := c.GetPlayPower(); power != nil {
		json["target"] = power.Target
	}

	return json
}

var CardService interface {
	Start() error
	GetCard(cardid int) (*Card, error)
	GetAllCards() map[int]*Card
}
