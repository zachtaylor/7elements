package main

import (
	pkg_env "github.com/zachtaylor/7elements/env"
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"ztaylor.me/cast"
	"ztaylor.me/http/session"
	"ztaylor.me/log"
)

const Patch = 3

func main() {
	env := pkg_env.NewService().ParseDefault()

	//

	stdout := log.StdOutService(log.LevelDebug)
	stdout.Formatter().CutSourcePath(2)

	runtime, err := pkg_env.NewRuntime(env)
	if err != nil {
		stdout.New().Add("Error", err).Error("failed to launch")
		return
	}

	stdout.New().With(cast.JSON{
		"Patch": Patch,
	}).Source().Debug()

	if envName := env["ENV"]; envName == "dev" {
		runtime.Sessions = session.NewCache(7 * cast.Minute)
		// runtime.Root.Logger = stdout
		server.Start(runtime, ":"+env["PORT"])
	} else if envName == "pro" {
		runtime.Sessions = session.NewCache(7 * cast.Hour)
		logLevel, _ := log.GetLevel(env["LOG_LEVEL"])
		runtime.Root.Logger = log.DailyRollingService(logLevel, env["LOG_PATH"])
		server.StartTLS(runtime, "7elements.cert", "7elements.key")
	} else {
		stdout.New().Error("7elements failed to launch, env error")
	}
}
