package texts

type Text struct {
	CardId      int
	Language    string
	Name        string
	Powers      map[int]string
	Description string
	Flavor      string
}

func New() *Text {
	return &Text{
		Powers: make(map[int]string),
	}
}
