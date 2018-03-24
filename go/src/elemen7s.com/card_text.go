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

type CTService interface {
	Start() error
	Get(lang string, cardid int) (*CardText, error)
}

var CardTextService CTService
