package SE

type Card struct {
	Id           int
	Image        string
	CardType     *CardType
	ElementCosts []*ElementCost
}

// persistence headers
var Cards = struct {
	Cache     map[int]*Card
	LoadCache func() error
	Insert    func(int) error
	Delete    func(int) error
}{make(map[int]*Card), nil, nil, nil}
