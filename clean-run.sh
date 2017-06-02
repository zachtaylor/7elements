#!/bin/bash
export GOPATH=$PWD/go/
rm 7elements.db
rm -Rf go/pkg/*/7elements.ztaylor.me*
./build.sh
./7elements
