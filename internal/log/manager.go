package log

import "taylz.io/log"

type Server interface {
	Log() Writer
}

type Writer = log.Writer
