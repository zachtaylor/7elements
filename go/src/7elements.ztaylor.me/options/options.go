package options

import (
	"flag"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"ztaylor.me/log"
)

type option struct {
	Value  string
	Prompt string
}

var options = map[string]*option{
	"port":          &option{"80", "Server port to open"},
	"log-level":     &option{"info", "One of [debug,info,warn,error,fatal,panic]"},
	"log-path":      &option{"", "Log file path"},
	"use-https":     &option{"false", "Open server port 443 (overrides \"port\" option)"},
	"db-path":       &option{"7elements.db", "Databse path to open"},
	"server-path":   &option{"server-path", "Server root path"},
	"patch-path":    &option{"patch/", "Available patches path"},
	"session-life":  &option{"180", "Session lifetime in seconds"},
	"js-path":       &option{"js/", "Path for javascript compilation"},
	"js-minify":     &option{"false", "Minify 7elements.js after every compilation"},
	"css-path":      &option{"css/", "Path for css compilation"},
	"css-minify":    &option{"false", "Minify 7elements.css after every compilation"},
	"image-path":    &option{"img", "Root for image files"},
	"password-salt": &option{"salt", "Salt string for password hashing"},
	"redirect-host": &option{"", "Redirect \"/\" traffic that don't match this hostname to this hostname if set"},
}

func init() {
	flags := setupflags()
	readfile()
	bindflags(flags)

	log.Add("Options", options).Debug("options: finished")
}

func String(name string) string {
	if options[name] == nil {
		log.Add("Name", name).Warn("options: invalid option")
		return ""
	}
	return options[name].Value
}

func Int(name string) int {
	i, _ := strconv.Atoi(String(name))
	return i
}

func Duration(name string) time.Duration {
	d, _ := time.ParseDuration(String(name))
	return d
}

func Bool(name string) bool {
	b, _ := strconv.ParseBool(String(name))
	return b
}

func setupflags() map[string]*string {
	read := make(map[string]*string)
	for key, option := range options {
		read[key] = flag.String(key, "", option.Prompt)
	}
	flag.Parse()
	return read
}

func readfile() error {
	file, e := ioutil.ReadFile("settings.txt")
	if e != nil {
		return e
	}

	for lineNo, line := range strings.Split(string(file), "\n") {
		line = strings.Trim(line, " ")
		if line == "" || line[0] == "#"[0] {
			continue
		}

		setting := strings.Split(line, "=")
		if options[setting[0]] == nil {
			log.Add("#", lineNo).Add("Line", line).Warn("options: settings.txt line not recognized")
			continue
		}
		if len(setting) == 2 {
			options[setting[0]].Value = setting[1]
			log.Add("Setting", setting[0]).Add("Value", setting[1]).Debug("options: set from settings.txt")
		}
	}

	return nil
}

func bindflags(read map[string]*string) {
	for key, value := range read {
		if *value == "" {
			continue
		}

		options[key].Value = *value
		log.Add("Setting", key).Add("Value", *value).Debug("options: set from flag")
	}
}
