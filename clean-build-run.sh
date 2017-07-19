#!/bin/bash
export GOPATH=$PWD/go/
echo rm 7elements
rm 7elements
echo rm 7elements.db
rm 7elements.db
echo rm -Rf go/pkg/*/7elements.ztaylor.me*
rm -Rf go/pkg/*/7elements.ztaylor.me*
echo ./build.sh
./build.sh
echo build complete. starting 7elements server.
./7elements
