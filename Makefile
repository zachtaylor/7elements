PWD=$(shell pwd)
GOPATH=$(PWD)/go

COLOR_RED='\033[0;31m'
COLOR_GREEN='\033[0;32m'
COLOR_YELLOW='\033[0;33m'
COLOR_BLUE='\033[0;34m'
COLOR_PURPLE='\033[0;35m'
COLOR_BWHITE='\033[1;37m'
COLOR_GRAY='\033[0;30m'
COLOR_OFF='\033[0m'

help:
	@echo -e $(COLOR_PURPLE)Makefile Available Targets$(COLOR_OFF)
	@echo -e $(COLOR_BWHITE)help$(COLOR_OFF): "hello, world!"
	@echo -e $(COLOR_BWHITE)server$(COLOR_OFF): build the server
	@echo -e $(COLOR_BWHITE)aifight$(COLOR_OFF): build the aifight program
	@echo -e $(COLOR_BWHITE)run, run-server$(COLOR_OFF): run the server
	@echo -e $(COLOR_BWHITE)run-aifight$(COLOR_OFF): run the aifight program

gopath:
	@export GOPATH=$(GOPATH)
	@echo -e $(COLOR_GRAY)7elements build path $$GOPATH $(COLOR_OFF)
	@go get -v golang.org/x/net/websocket
	@go get -v github.com/mattn/go-sqlite3
	@go get -v github.com/sirupsen/logrus
	@go get -v github.com/cznic/mathutil
	@go get -v github.com/fsnotify/fsnotify
	@go get -v github.com/tdewolff/minify
	@go get -v gopkg.in/gomail.v2
	@go get -v ztaylor.me/env
	@go get -v ztaylor.me/events
	@go get -v ztaylor.me/http
	@go get -v ztaylor.me/buildir
	@go get -v ztaylor.me/js
	@go get -v ztaylor.me/log

update:
	@export GOPATH=$(GOPATH)
	@echo -e $(COLOR_BLUE)7elements updating deps$(COLOR_OFF)
	@go get -u -v golang.org/x/net/websocket
	@go get -u -v github.com/mattn/go-sqlite3
	@go get -u -v github.com/sirupsen/logrus
	@go get -u -v github.com/cznic/mathutil
	@go get -u -v github.com/fsnotify/fsnotify
	@go get -u -v github.com/tdewolff/minify
	@go get -u -v gopkg.in/gomail.v2
	@go get -u -v ztaylor.me/env
	@go get -u -v ztaylor.me/events
	@go get -u -v ztaylor.me/http
	@go get -u -v ztaylor.me/buildir
	@go get -u -v ztaylor.me/js
	@go get -u -v ztaylor.me/log

aifight: gopath
	@go install -v elemen7s.com/cmd/aifight

run-aifight: go/bin/aifight
	@echo -e $(COLOR_YELLOW)STARTING AI VS AI FIGHT$(COLOR_OFF)
	@go/bin/aifight

server: gopath
	@echo -e $(COLOR_BWHITE)7elements build server $$GOPATH/src/elemen7s.com/cmd/elemen7s.com $(COLOR_OFF)
	@go install -v elemen7s.com/cmd/elemen7s.com

run-server: go/bin/elemen7s.com
	@echo -e $(COLOR_GREEN)7elements server starting $$GOPATH/bin/elemen7s.com $(COLOR_OFF)
	@go/bin/elemen7s.com

clean:
	@rm -R log/*.log

run: run-server
