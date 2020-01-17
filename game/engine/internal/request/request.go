package request // import "github.com/zachtaylor/7elements/game/engine/internal/request"

import "ztaylor.me/log"

// CutLogPath cuts "engine/internal/"
func CutLogPath(fmt log.Formatter) {
	fmt.CutSourcePath(1)
}
