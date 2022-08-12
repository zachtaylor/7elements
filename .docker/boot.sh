#!/bin/sh
echo -e "installing..."
go mod download
echo -e "building..."
go build -o /build/api.7tcg /app/cmd/api.7tcg
echo -e "booting..."
/build/api.7tcg
