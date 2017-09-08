package main

import (
	"7elements.ztaylor.me/db"
	"7elements.ztaylor.me/options"
	"7elements.ztaylor.me/server/cssbuilder"
	"7elements.ztaylor.me/server/jsbuilder"
	"ztaylor.me/log"
)

func init() {
	if logpath := options.String("log-path"); logpath != "" {
		log.SetFile(logpath)
	}
	log.SetLevel(options.String("log-level"))

	jsbuilder.SetPath(options.String("js-path"))
	jsbuilder.Options.Minify = options.Bool("js-minify")
	go jsbuilder.StartWatch()

	cssbuilder.SetPath(options.String("css-path"))
	cssbuilder.Options.Minify = options.Bool("css-minify")
	go cssbuilder.StartWatch()

	db.Open(options.String("db-path"))
}
