package game

import (
	"strings"

	"taylz.io/types"
)

func SourcePath() string {
	sourcePath := types.NewSource(0).File()
	sourcePath = sourcePath[:strings.LastIndex(sourcePath, "/")]
	return sourcePath
}
