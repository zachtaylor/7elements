package game

import "taylz.io/log"

type Runner interface {
	Run(*G, log.Writer)
}
