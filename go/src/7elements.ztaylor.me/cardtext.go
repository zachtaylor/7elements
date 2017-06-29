package SE

type CardText struct {
	CardId      int
	Language    string
	Name        string
	Description string
	Flavor      string
}

// persistence headers
var CardTexts = struct {
	Cache     map[string]map[int]*CardText
	LoadCache func(string) error
}{make(map[string]map[int]*CardText), nil}
