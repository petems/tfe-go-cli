CLI_NAME := tfe-go-cli
GIT_VERSION := $(shell git log -1 --pretty=format:"%h" .)
INT_VERSION := $(shell echo "INTEGRATION_TESTS")

test:
	go test ./...

int_test:
	go test integration/*.go

int_update_golden:
	go test integration/*.go -update

build:
	go build -ldflags "-X \"cmd.gitCommit=$(GIT_VERSION)\"" \
		-o "$(GOPATH)/bin/$(CLI_NAME)" \
		./main.go

int_build:
	go build -ldflags "-X \"cmd.gitCommit=$(INT_VERSION)\"" \
		-o "./$(CLI_NAME)-int-testing" \
		./main.go

.DEFAULT_GOAL := build
