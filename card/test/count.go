package main

import (
	"fmt"

	"github.com/zachtaylor/7elements/card"
)

func main() {
	log := fmt.Println
	c := card.Count{
		1: 1,
		4: 4,
		7: 7,
	}
	var cfmt string
	if fmt, err := c.Format(); err != nil {
		log("format error", err)
		return
	} else {
		cfmt = fmt
	}
	cans := `1a4d7g`
	if cfmt != cans {
		log("expected", cans)
		log("actual", cfmt)
		return
	}
	if copy, err := card.ParseCount(cfmt); err != nil {
		log("parse error", err)
		return
	} else {
		for k, v := range copy {
			if c[k] != v {
				log("expected", c[k])
				log("actual", v)
			}
		}
		for k, v := range c {
			if copy[k] != v {
				log("expected", v)
				log("actual", copy[k])
			}
		}

		if fmt, err := copy.Format(); err != nil {
			log("format error", err)
			return
		} else {
			// in n out
			if cfmt != fmt {
				log("expected", cans)
				log("actual", cfmt)
				return
			}
		}
	}
	log("pass")
}
