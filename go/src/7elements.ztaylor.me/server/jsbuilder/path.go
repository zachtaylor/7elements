package jsbuilder

import (
	"7elements.ztaylor.me/log"
)

func SetPath(newPath string) {
	log.Add("Path", newPath)
	if watcher == nil {
		log.Warn("jsbuilder: watcher unavailable")
	} else if err := watcher.Add(newPath); err != nil {
		log.Add("Error", err).Error("jsbuilder: AddPath")
	} else {
		Options.path = newPath
		log.Add("Path", newPath).Debug("jsbuilder: setpath")
	}
}
