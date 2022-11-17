#!/bin/sh

# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chan-linux main.go
# 苹果
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o chan-darwin main.go
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o chan-windows.exe main.go
