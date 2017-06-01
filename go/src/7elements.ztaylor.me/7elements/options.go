package main

import (
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/options"
	"7elements.ztaylor.me/persistence"
	"7elements.ztaylor.me/server/cssbuilder"
	"7elements.ztaylor.me/server/jsbuilder"
)

func init() {
	log.SetLevel(options.String("log-level"))
	if logpath := options.String("log-path"); logpath != "" {
		log.SetFile(logpath)
	}

	jsbuilder.SetPath(options.String("js-path"))
	jsbuilder.Options.Minify = options.Bool("js-minify")
	go jsbuilder.StartWatch()

	cssbuilder.SetPath(options.String("css-path"))
	cssbuilder.Options.Minify = options.Bool("css-minify")
	go cssbuilder.StartWatch()

	persistence.SetConnection(options.String("db-path"))
}
