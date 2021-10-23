package main

import (
	"net/http"

	"github.com/zachtaylor/7elements/env"
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/log"
)

const Patch = 4

func main() {
	stdout := log.Default()

	stdout.Info("vii: start")

	env := env.New().ParseDefault()
	runtime, err := runtime.Parse(env, Patch)
	if runtime == nil {
		stdout.Add("Error", err).Error("vii: failed to parse env")
		return
	}

	fs := http.FileSystem(http.Dir(env["WWW_PATH"]))

	if runtime.IsDevEnv {
		server.Start(runtime, fs, ":"+env["PORT"])
	} else {
		// production has a different binary so this is never used
		server.StartTLS(runtime, fs, "7elements.cert", "7elements.key")
	}
}
