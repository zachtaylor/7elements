package main

import (
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/log"
)

const Patch = 4

func main() {
	stdout := log.Default()

	stdout.Info("vii: start")

	env := DefaultENV().ParseDefault()
	runtime, err := runtime.Parse(env, Patch)
	if runtime == nil {
		stdout.Add("Error", err).Error("vii: failed to parse env")
		return
	}

	if runtime.IsDevEnv {
		server.Start(runtime, ":"+env["PORT"])
	} else {
		// production has a different binary so this is never used
		server.StartTLS(runtime, "7tcg.cert", "7tcg.key")
	}
}
