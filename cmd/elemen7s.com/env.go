package main

import "taylz.io/env"

func DefaultENV() env.Builder {
	return env.Builder{
		"ENV":         "dev",
		"PORT":        "80",
		"LOG_LEVEL":   "info",
		"LOG_PATH":    "",
		"LOG_DIR":     "./",
		"DB_PWSALT":   "",
		"DB_HOST":     "",
		"DB_TABLE":    "",
		"DB_NAME":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_PORT":     "",
		"WWW_PATH":    "/srv/www",
	}
}
