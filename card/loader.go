package card

// Loader loads cards
type Loader interface {
	GetAll() Prototypes
}
