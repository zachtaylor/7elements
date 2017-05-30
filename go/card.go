package SE

type Card struct {
	Id           uint
	Image        string
	CardType     *CardType
	ElementCosts []*ElementCost
}

// persistence headers
var Cards = struct {
	Cache     map[uint]*Card
	LoadCache func() error
	Insert    func(uint) error
	Delete    func(uint) error
}{make(map[uint]*Card), nil, nil, nil}
