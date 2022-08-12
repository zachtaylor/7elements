package phase

import "github.com/zachtaylor/7elements/game"

type Choice struct {
	game.PriorityContext
	Text     string
	Data     map[string]any
	Choices  []map[string]any
	Finisher func(answer any)
	answer   any
}

func NewChoice(playerID, text string, data map[string]any, choices []map[string]any, fin func(any)) *Choice {
	return &Choice{
		PriorityContext: game.PriorityContext{playerID},
		Text:            text,
		Data:            data,
		Choices:         choices,
		Finisher:        fin,
	}
}

func (*Choice) Type() string      { return "choice" }
func (*Choice) Next() game.Phaser { return nil }

func (r *Choice) JSON() map[string]any {
	return map[string]any{
		"choice":  r.Text,
		"data":    r.Data,
		"options": r.Choices,
	}
}

func (r *Choice) OnConnect(g *game.G, player *game.Player) {
	// if player == nil {
	// player = g.Player(r.Priority()[0])
	// } else if player.ID() != r.Priority()[0] {
	// return
	// }
}

// Finish implements game.OnFinishPhaser
func (r *Choice) OnFinish(g *game.G, _ *game.State) []game.Phaser {
	if r.Finisher != nil {
		r.Finisher(r.answer)
	}
	return nil
}

// Request implements game.OnFinishPhaser
func (r *Choice) OnRequest(g *game.G, state *game.State, player *game.Player, json map[string]any) {
	if player.ID() != r.Priority()[0] {
		g.Log().With(map[string]any{
			"Player": player,
			"json":   json,
		}).Warn("choice: receive")
		return
	}

	r.answer = json["choice"]
	if r.answer != "" {
		state.T.React.Add(player.ID())
	}
}

var ChoiceElementsData = []map[string]any{
	{"choice": 1, "display": `<img src="/img/icon/element-1.png">`},
	{"choice": 2, "display": `<img src="/img/icon/element-2.png">`},
	{"choice": 3, "display": `<img src="/img/icon/element-3.png">`},
	{"choice": 4, "display": `<img src="/img/icon/element-4.png">`},
	{"choice": 5, "display": `<img src="/img/icon/element-5.png">`},
	{"choice": 6, "display": `<img src="/img/icon/element-6.png">`},
	{"choice": 7, "display": `<img src="/img/icon/element-7.png">`},
}
