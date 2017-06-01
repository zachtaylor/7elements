#!/bin/bash
export GOPATH=$PWD/go/
go get -v golang.org/x/net/websocket
go get -v github.com/mattn/go-sqlite3
go get -v github.com/Sirupsen/logrus
go get -v github.com/cznic/mathutil
go get -v github.com/fsnotify/fsnotify
go get -v github.com/tdewolff/minify
go install -v 7elements.ztaylor.me/7elements
