package queue

import (
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/gameserver"
	"taylz.io/http/user"
	"taylz.io/log"
	"taylz.io/types"
)

type Server struct {
	Settings Settings
	Storer
	live map[Options]string
}

type Settings struct {
	Logger *log.T
	Games  *gameserver.T
}

func NewServer(settings Settings, store Storer) *Server {
	return &Server{
		Settings: settings,
		Storer:   store,
		live:     make(map[Options]string),
	}
}

func (s *Server) Request(options Options, user *user.T, deck *deck.T) (*T, error) {
	if s.Get(deck.User) != nil {
		return nil, ErrExists
	}

	q := &T{
		Entry: &gameserver.Entry{
			Deck:   deck,
			Writer: user,
		},
		Start:  types.NewTime(),
		done:   make(chan string),
		cancel: make(chan bool),
	}

	s.Sync(func(get Getter, set Setter) {
		if match := s.live[options]; match == "" {
			set(deck.User, q)
			s.live[options] = deck.User
		} else if q2 := get(match); q2 == nil { // expired ?
			set(deck.User, q)
			s.live[options] = deck.User
		} else {
			set(match, nil)
			delete(s.live, options)
			go s.match(options, q, q2)
		}
	})
	return q, nil
}
func (s *Server) match(options Options, q, q2 *T) {
	board := s.Settings.Games.New(options.Rules, s.Settings.Logger, q.Entry, q2.Entry)
	q.Finish(board.ID())
	q2.Finish(board.ID())
}
