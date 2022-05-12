package main

import "taylz.io/env"

func DefaultENV() env.Values {
	return env.Values{
		"ENV":         "dev",
		"PORT":        "80",
		"LOG_LEVEL":   "info",
		"LOG_PATH":    "./",
		"DB_PWSALT":   "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_HOST":     "",
		"DB_PORT":     "",
	}
}
