package game

import "github.com/zachtaylor/7elements/power"

type PriorityContext Priority

func (phase PriorityContext) Priority() Priority { return Priority(phase) }

type LoserPhase struct {
	PriorityContext
	player string
}

func NewLoserPhase(priority Priority, player string) *LoserPhase {
	return &LoserPhase{
		PriorityContext: PriorityContext(priority),
		player:          player,
	}
}

func (*LoserPhase) Type() string { return "loser" }
func (r *LoserPhase) JSON() map[string]any {
	return map[string]any{
		"player": r.player,
	}
}

func NewTriggerPhase(g *G, token *Token, power *power.T) Phaser {
	if power.Target == "self" {
		return &TriggerPhase{
			PriorityContext: PriorityContext(g.NewPriority(token.Player())),
			token:           token.ID(),
			power:           power,
			targets:         []string{token.ID()},
		}
	} else {
		return &TriggerTargetPhase{
			PriorityContext: PriorityContext(Priority{token.Player()}),
			token:           token.ID(),
			power:           power,
		}
	}
}

type TriggerPhase struct {
	PriorityContext
	token   string
	power   *power.T
	targets []string
}

func (*TriggerPhase) Type() string { return "trigger" }
func (r *TriggerPhase) JSON() map[string]any {
	return map[string]any{
		"tokenid": r.token,
		"power":   r.power.JSON(),
		"targets": r.targets,
	}
}
func (r *TriggerPhase) OnFinish(g *G, state *State) []Phaser {
	g.Log().Trace("finish trigger power", r.Priority(), r.token, r)
	return RunScript(g, NewScriptContext(r.power.Script, r.token, r.Priority()[0], nil, r.targets))
}

// TriggerTargetPhase occurs when a target is unresolved
type TriggerTargetPhase struct {
	PriorityContext
	token   string
	power   *power.T
	targets []string
}

func (*TriggerTargetPhase) Type() string { return "target" }
func (r *TriggerTargetPhase) JSON() map[string]any {
	return map[string]any{
		"type": r.power.Target,
		"text": r.power.Text,
	}
}
func (r *TriggerTargetPhase) OnRequest(g *G, state *State, player *Player, json map[string]any) {
	if player.ID() != r.Priority()[0] {
		g.log.Warn("request trigger target not allowed", player.ID(), r.Priority())
		return
	}

	if target, _ := json["target"].(string); target != "" {
		r.targets = []string{target}
		state.T.React.Add(player.ID())
	}
}
func (r *TriggerTargetPhase) OnFinish(g *G, state *State) []Phaser {
	return []Phaser{&TriggerPhase{
		PriorityContext: PriorityContext(g.NewPriority(r.Priority()[0])),
		token:           r.token,
		power:           r.power,
		targets:         r.targets,
	}}
}
