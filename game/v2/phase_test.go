package game_test

import "github.com/zachtaylor/7elements/game/v2"

func TriggerPhaseIsOnFinisher(r *game.TriggerPhase) game.OnFinishPhaser { return r }

func TriggerTargetPhaseIsOnRequester(r *game.TriggerTargetPhase) game.OnRequestPhaser { return r }
func TriggerTargetPhaseIsOnFinisher(r *game.TriggerTargetPhase) game.OnFinishPhaser   { return r }
