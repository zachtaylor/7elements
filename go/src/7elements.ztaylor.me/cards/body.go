package cards

type Body struct {
	CardId       int
	GameCardId   int
	Attack       int
	Health       int
	IsUnkillable bool
	IsHidden     bool
}

func (body *Body) Copy() *Body {
	return &Body{
		CardId:       body.CardId,
		GameCardId:   body.GameCardId,
		Attack:       body.Attack,
		Health:       body.Health,
		IsUnkillable: body.IsUnkillable,
		IsHidden:     body.IsHidden,
	}
}
