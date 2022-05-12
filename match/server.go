package match

import "github.com/zachtaylor/7elements/game"

type Server interface {
	// Log() log.Writer
	GetGameVersion() *game.Version
	GetGameManager() *game.Manager
}
