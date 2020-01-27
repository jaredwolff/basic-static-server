MAIN_EXEC := app.go

.PHONY: build run

default: build

build:
	GO111MODULE=on go build

run:
	GO111MODULE=on go run $(MAIN_EXEC)