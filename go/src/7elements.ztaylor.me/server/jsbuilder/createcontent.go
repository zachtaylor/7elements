package jsbuilder

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"io/ioutil"
	"ztaylor.me/log"
)

var content string
var minifier = minify.New()

func init() {
	minifier.AddFunc("text/javascript", js.Minify)
}

func CreateContent() {
	log := log.Add("Path", Options.path)

	content = ""
	fileIncludes := make([]string, 0)

	files, _ := ioutil.ReadDir(Options.path)
	for _, f := range files {
		path := Options.path + f.Name()

		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Add("Error", err).Add("File", path).Error("jsbuilder: file read")
			return
		}

		content += string(file)
		fileIncludes = append(fileIncludes, f.Name())
	}

	log.Add("Files", fileIncludes)

	if Options.Minify {
		if newContent, err := minifier.String("text/javascript", content); err != nil {
			log.Add("Error", err).Error("jsbuilder: minification error")
		} else {
			content = newContent
			log.Add("Options.Minify", Options.Minify).Debug("jsbuilder: compile")
		}
	}
}
