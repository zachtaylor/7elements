package vii

type CardText struct {
	Name        string
	Powers      map[int]string
	Description string
	Flavor      string
}

func NewCardText() *CardText {
	return &CardText{
		Powers: make(map[int]string),
	}
}

var CardTextService interface {
	Start() error
	GetCardText(lang string, cardid int) (*CardText, error)
}
