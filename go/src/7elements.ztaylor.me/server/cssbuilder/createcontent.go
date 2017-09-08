package cssbuilder

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"io/ioutil"
	"ztaylor.me/log"
)

var content string
var minifier = minify.New()

func init() {
	minifier.AddFunc("text/css", css.Minify)
}

func CreateContent() {
	log := log.Add("Path", Options.path).Add("Options.Minify", Options.Minify)

	content = ""
	fileIncludes := make([]string, 0)

	files, _ := ioutil.ReadDir(Options.path)
	for _, f := range files {
		var path = Options.path + f.Name()

		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Add("Error", err).Error("cssbuilder: file read")
			return
		}

		content += string(file)
		fileIncludes = append(fileIncludes, f.Name())
	}

	log.Add("Files", fileIncludes)

	if Options.Minify {
		if newContent, err := minifier.String("text/css", content); err != nil {
			log.Add("Error", err).Error("cssbuilder: minification error")
		} else {
			content = newContent
			log.Debug("cssbuilder: compile")
		}
	}
}
