package phase_test

import (
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/v2"
)

func StartIsOnActivatePhaser(r *phase.Start) game.OnActivatePhaser { return r }
func StartIsOnConnectPhaser(r *phase.Start) game.OnConnectPhaser   { return r }
func StartIsOnRequestPhaser(r *phase.Start) game.OnRequestPhaser   { return r }

func SunriseIsOnActivatePhaser(r *phase.Sunrise) game.OnActivatePhaser { return r }
func SunriseIsOnConnectPhaser(r *phase.Sunrise) game.OnConnectPhaser   { return r }
func SunriseIsOnFinishPhaser(r *phase.Sunrise) game.OnFinishPhaser     { return r }
func SunriseIsOnRequestPhaser(r *phase.Sunrise) game.OnRequestPhaser   { return r }

func MainIsOnConnectPhaser(r *phase.Main) game.OnConnectPhaser { return r }

func SunsetIsOnConnectPhaser(r *phase.Sunset) game.OnConnectPhaser { return r }

func PlayIsOnActivatePhaser(r *phase.Play) game.OnActivatePhaser { return r }
func PlayIsOnFinishPhaser(r *phase.Play) game.OnFinishPhaser     { return r }

func AttackIsOnActivatePhaser(r *phase.Attack) game.OnActivatePhaser { return r }
func AttackIsOnRequestPhaser(r *phase.Attack) game.OnRequestPhaser   { return r }
func AttackIsnFinishPhaser(r *phase.Attack) game.OnFinishPhaser      { return r }

func CombatIsOnActivatePhaser(r *phase.Combat) game.OnActivatePhaser { return r }
func CombatIsOnFinishPhaser(r *phase.Combat) game.OnFinishPhaser     { return r }

func ChoiceIsPhaser(r *phase.Choice) game.Phaser                   { return r }
func ChoiceIsOnConnectPhaser(r *phase.Choice) game.OnConnectPhaser { return r }
func ChoiceIsOnRequestPhaser(r *phase.Choice) game.OnRequestPhaser { return r }
func ChoiceIsOnFinishPhaser(r *phase.Choice) game.OnFinishPhaser   { return r }
