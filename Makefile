.DEFAULT_GOAL := build

fmt:
	go fmt ./
.PHONY:fmt

lint: fmt
	golint ./
.PHONY:lint

vet: fmt
	go vet ./
.PHONY:vet

build: vet
	GOOS=windows GOARCH=386 go build -o mynewtool.exe .
.PHONY:build