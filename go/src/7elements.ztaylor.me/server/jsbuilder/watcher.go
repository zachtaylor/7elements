package jsbuilder

import (
	"7elements.ztaylor.me/log"
	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func init() {
	if w, err := fsnotify.NewWatcher(); err != nil {
		log.Add("Error", err).Error("jsbuilder: fsnotify.NewWatcher()")
	} else {
		watcher = w
	}
}

func StartWatch() {
	for {
		select {
		case event := <-watcher.Events:
			log.Add("Event", event).Debug("jsbuilder: rebuild")
			CreateContent()
		case err := <-watcher.Errors:
			log.Add("Error", err.Error()).Error("jsbuilder: watch error")
			return
		}
	}
}
