package db

type Scanner interface {
	Scan(...interface{}) error
}
