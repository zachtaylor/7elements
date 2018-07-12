package vii

type CardBody struct {
	Attack int
	Health int
}

func NewCardBody() *CardBody {
	return &CardBody{}
}

func (b *CardBody) Copy() *CardBody {
	if b == nil {
		return nil
	}
	return &CardBody{
		Attack: b.Attack,
		Health: b.Health,
	}
}

func (b *CardBody) Json() Json {
	if b == nil {
		return nil
	}
	return Json{
		"attack": b.Attack,
		"health": b.Health,
	}
}
