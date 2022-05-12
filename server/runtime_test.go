package server_test

import (
	"github.com/zachtaylor/7elements/server"
	"github.com/zachtaylor/7elements/server/internal"
)

func RuntimeIsServer(runtime *server.Runtime) internal.Server {
	return runtime
}
