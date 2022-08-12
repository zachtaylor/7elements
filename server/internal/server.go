package internal

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/content"
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/internal/user"
	"github.com/zachtaylor/7elements/match"
	"taylz.io/db"
	"taylz.io/http"
	"taylz.io/http/session"
	"taylz.io/http/websocket"
	"taylz.io/keygen"
	"taylz.io/log"
)

// Server is the internally exposed server api
type Server interface {
	IsProd() bool
	Accounts() *account.Cache
	Content() content.T
	DB() *db.DB
	Games() game.Manager
	HashPassword(string) string
	Log() log.Writer
	MatchMaker() *match.Maker
	Ping()
	Sessions() session.Manager
	user.Server
	Websockets() *websocket.Cache
	http.Handler
	websocket.MessageHandler
}

type runtime struct {
	isprod   bool
	hostpath string
	accounts *account.Cache
	fork     *http.Fork
	wsFork   *websocket.MessageFork
	log      log.Writer
	db       *db.DB
	content  content.T
	sessions session.Manager
	sockets  *websocket.Cache
	users    user.Manager
	games    game.Manager
	cupid    *match.Maker
	hashPass func(string) string
}

func (rt *runtime) IsProd() bool { return rt.isprod }

// Content implements Server
func (rt *runtime) Accounts() *account.Cache { return rt.accounts }

// Content implements Server
func (rt *runtime) Content() content.T { return rt.content }

// Games implements Server
func (rt *runtime) Games() game.Manager { return rt.games }

// GetDB implements Server
func (rt *runtime) DB() *db.DB { return rt.db }

// Sessions implements Server
func (rt *runtime) Sessions() session.Manager { return rt.sessions }

// HashPassword implements Server
func (rt *runtime) HashPassword(val string) string { return rt.hashPass(val) }

// Log implements Server
func (rt *runtime) Log() log.Writer { return rt.log }

// MatchMaker implements Server
func (rt *runtime) MatchMaker() *match.Maker { return rt.cupid }

// Users implements Server
func (rt *runtime) Users() user.Manager { return rt.users }

// Websockets implements Server
func (rt *runtime) Websockets() *websocket.Cache { return rt.sockets }

func (rt *runtime) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.fork.ServeHTTP(w, r)
}

func (rt *runtime) ServeWS(ws *websocket.T, msg *websocket.Message) {
	rt.wsFork.ServeWS(ws, msg)
}

// Ping sends updated user counts
func (rt runtime) Ping() {
	users, _ := accounts.Count(rt.db)
	bytes := websocket.NewMessage("/ping", map[string]any{
		"ping":   rt.Websockets().Count(),
		"online": rt.Sessions().Count(),
		"users":  users,
	}).ShouldMarshal()
	rt.Websockets().Each(func(id string, ws *websocket.T) {
		ws.Write(websocket.MessageText, bytes)
	})
}

func NewRuntime(
	isprod bool,
	db *db.DB,
	hashPass func(string) string,
	log log.Writer,
	sessionSettings session.Settings,
	userSettings user.Settings,
	gameSettings game.Settings,
	origins []string,
	routes func(wsUpgrader http.Handler) []RouteBuilder,
	wsRoutes []WSRouteBuilder,
) (Server, error) {
	accounts := account.NewCache()

	content, err := content.Build(db)
	if err != nil {
		return nil, err
	}

	sessions := session.NewService(sessionSettings, keygen.DefaultFunc())

	rt := &runtime{
		isprod:   isprod,
		accounts: accounts,
		content:  content,
		log:      log,
		db:       db,
		hashPass: hashPass,
		sessions: sessions,
	}

	var wsUpgrader http.Handler
	rt.users, rt.sockets, wsUpgrader = user.NewServiceHandler(
		userSettings,
		keygen.DefaultFunc(),
		sessions,
		rt,
	)

	rt.games = game.NewService(
		gameSettings,
		keygen.DefaultFunc(),
		rt,
	)

	rt.cupid = match.NewMaker(rt)

	rt.wsFork = websocket.NewMessageFork()
	for _, route := range wsRoutes {
		rt.wsFork.Path(route.Router, route.Provider(rt))
	}

	cors := CORS(origins)
	rt.fork = http.NewFork()
	for _, route := range routes(wsUpgrader) {
		rt.fork.Path(route.Router, cors(route.Provider(rt)))
	}

	rt.sessions.Observe(OnSession(rt))
	rt.sockets.Observe(OnSocket(rt))
	rt.users.Observe(OnUser(rt))

	return rt, nil
}
