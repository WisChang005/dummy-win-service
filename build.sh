#! /bin/bash

export GOARCH=amd64
echo "Building Windows x64 binary..."
go build -o ./bin/x64/dummy_service.exe dummy_service.go

export GOARCH=386
echo "Building Windows win32 binary..."
go build -o ./bin/win32/dummy_service.exe dummy_service.go

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
echo "Building Linux binary..."
go build -o ./bin/linux/dummy_service dummy_service.go
