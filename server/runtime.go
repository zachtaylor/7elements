package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/match"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/db"
	"taylz.io/http/session"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
	"taylz.io/log"
)

// Runtime is our implementation of internal.Server
type Runtime struct {
	isProd   bool
	log      log.Writer
	db       *db.DB
	cipher   func(string) string
	fs       http.FileSystem
	sessions *session.Manager
	sockets  *websocket.Manager
	users    *user.Manager
	accounts *account.Cache
	games    *game.Manager
	cupid    *match.Maker

	*game.Runtime
}

func NewRuntime(isprod bool, log log.Writer, db *db.DB, cipher func(string) string, fs http.FileSystem, sessionSettings session.Settings, socketSettings websocket.Settings, gameSettings game.Settings, matchSettings match.CacheSetter)

func (rt *Runtime) EnvProd() bool                           { return rt.isProd }
func (rt *Runtime) Log() log.Writer                         { return rt.log }
func (rt *Runtime) GetDB() *db.DB                           { return rt.db }
func (rt *Runtime) HashPassword(input string) string        { return rt.cipher(input) }
func (rt *Runtime) GetFileSystem() http.FileSystem          { return rt.fs }
func (rt *Runtime) GetSessionManager() *session.Manager     { return rt.sessions }
func (rt *Runtime) GetWebsocketManager() *websocket.Manager { return rt.sockets }
func (rt *Runtime) GetUserManager() *user.Manager           { return rt.users }
func (rt *Runtime) GetAccounts() *account.Cache             { return rt.accounts }
func (rt *Runtime) GetGameManager() *game.Manager           { return rt.games }
func (rt *Runtime) GetMatchMaker() *match.Maker             { return rt.cupid }

// Ping sends updated user counts
func (rt *Runtime) Ping() {
	users, _ := accounts.Count(rt.db)
	bytes := wsout.Ping(map[string]any{
		"ping":   rt.sockets.Count(),
		"online": rt.sessions.Count(),
		"users":  users,
	})
	rt.sockets.Each(func(id string, ws *websocket.T) {
		ws.Write(bytes)
	})
}
