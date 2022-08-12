package main

import "taylz.io/env"

func DefaultENV() env.Builder {
	return env.Builder{
		"ENV":         "dev",
		"PORT":        "7000",
		"ORIGINS":     "",
		"LOG_LEVEL":   "info",
		"LOG_PATH":    "",
		"LOG_DIR":     "./",
		"DB_HOST":     "",
		"DB_TABLE":    "",
		"DB_NAME":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_PORT":     "",
		"DB_PWSALT":   "",
	}
}
