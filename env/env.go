package env

import "ztaylor.me/db/env"

// Service exports `env.Service`
type Service = env.Service

func NewService() env.Service {
	return env.Service{
		"ENV":       "dev",
		"PORT":      "80",
		"LOG_LEVEL": "info",
		"LOG_PATH":  "./",
		"DB_PWSALT": "",
	}.Merge("DB_", env.NewService())
}
