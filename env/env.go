package env

import "taylz.io/env"

func New() env.Service {
	return env.Service{
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
