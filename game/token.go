package game

import (
	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/cast"
)

// Token is an in-play game object
type Token struct {
	ID       string
	Card     *Card
	Username string
	IsAwake  bool
	Body     *vii.Body
	Powers   vii.Powers
}

func NewToken(card *Card, username string) *Token {
	return &Token{
		Card:     card,
		Username: username,
		Body:     card.Card.Body.Copy(),
		Powers:   card.Card.Powers.Copy(),
	}
}

// Tokens is a map of tid->Token
type Tokens map[string]*Token

func (t *Token) String() string {
	return "Token#" + t.ID + ":" + t.Card.Card.Name
}

// JSON returns a representation of a game token
func (t *Token) JSON() cast.JSON {
	if t == nil {
		return nil
	}

	json := t.Card.JSON()
	json["id"] = t.ID
	json["cardid"] = t.Card.ID
	json["owner"] = t.Card.Username
	json["username"] = t.Username
	json["awake"] = t.IsAwake
	json["powers"] = t.Powers.JSON()
	json["body"] = t.Body.JSON()

	return json
}

// JSON returns a representation of a set of game tokens
func (tokens Tokens) JSON() []cast.JSON {
	data := make([]cast.JSON, 0)
	for _, c := range tokens {
		data = append(data, c.JSON())
	}
	return data
}
