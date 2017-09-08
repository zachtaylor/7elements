package cssbuilder

import (
	"ztaylor.me/log"
)

func SetPath(newPath string) {
	if watcher == nil {
	} else if err := watcher.Add(newPath); err != nil {
		log.Add("Path", newPath).Add("Error", err).Error("cssbuilder: AddPath")
	} else {
		Options.path = newPath
	}
}
