package main

import (
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"taylz.io/log"
)

const (
	buildVersion = 1

	expectedPatch = 4
)

func main() {
	stdout := log.Default()

	env := DefaultENV().ShouldParseDefault()

	runtime, err := server.NewRuntime(env, expectedPatch)
	if runtime == nil {
		stdout.Add("Error", err).Error("vii: failed to parse env")
		return
	}

	stdout.With(map[string]any{
		"web_dir":       env["WWW_PATH"],
		"build_version": buildVersion,
	}).Info("starting")

	if !runtime.IsProd() {
		server.Start(runtime, ":"+env["PORT"])
	} else {
		// production has a different binary so this is never used
		server.StartTLS(runtime, "7tcg.cert", "7tcg.key")
	}
}
