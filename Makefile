.PHONY:
.SILENT:
.DEFAULT_GOAL := all

build:
	go build -o ./.bin/app ./cmd/mw_task/main.go

run: build
	./.bin/app