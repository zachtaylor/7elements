package SE

type CardText struct {
	CardId      uint
	Language    string
	Name        string
	Description string
	Flavor      string
}

// persistence headers
var CardTexts = struct {
	Cache     map[string]map[uint]*CardText
	LoadCache func(string) error
}{make(map[string]map[uint]*CardText), nil}
