package jsbuilder

import (
	"ztaylor.me/log"
)

func SetPath(path string) {
	if watcher == nil {
	} else if err := watcher.Add(path); err != nil {
		log.Add("Path", path).Add("Error", err).Error("jsbuilder: SetPath")
	} else {
		Options.path = path
	}
}
