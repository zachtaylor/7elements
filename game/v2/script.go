package game

import "github.com/zachtaylor/7elements/element"

type ScriptFunc = func(*G, ScriptContext) ([]Phaser, error)

var Scripts = map[string]ScriptFunc{}

func RunScript(g *G, ctx ScriptContext) []Phaser {
	if f := Scripts[ctx.Name]; f == nil {
		g.Log().Error("script spell missing", ctx.Name)
	} else if stack, err := f(g, ctx); err != nil {
		g.Log().Error("script error", ctx.Name, err)
	} else {
		return stack
	}
	return nil
}

type ScriptContext struct {
	Name    string
	Source  string
	Player  string
	Karma   element.Count
	Targets []string
}

func NewScriptContext(name, source, player string, karma element.Count, targets []string) ScriptContext {
	if karma == nil {
		karma = element.Count{}
	}
	return ScriptContext{
		Name:    name,
		Source:  source,
		Player:  player,
		Karma:   karma,
		Targets: targets,
	}
}
