package SE

import (
	"time"
)

type AccountPack struct {
	Username string
	ArtId    string
	Register time.Time
}

// persistence headers
var AccountsPacks = struct {
	Cache  map[string][]*AccountPack
	Get    func(string) ([]*AccountPack, error)
	Insert func(string) error
	Delete func(string) error
}{make(map[string][]*AccountPack), nil, nil, nil}
