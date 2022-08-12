package game

type Manager interface {
	Get(string) *G
	New(rules Rules, runner Runner, e1, e2 Entry) (game *G)
}
