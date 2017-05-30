package persistence

type Scanner interface {
	Scan(...interface{}) error
}
