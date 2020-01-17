package request // import "github.com/zachtaylor/7elements/game/engine/request"

import "ztaylor.me/log"

// CutLogPath cuts ".../game/engine/"
func CutLogPath(fmt log.Formatter) {
	fmt.CutSourcePath(1)
}
