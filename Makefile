# https://gist.github.com/eleniums/06c30c50befcfe36e7840ca46c9ebacd?permalink_comment_id=4354315#gistcomment-4354315

EXECUTABLE=huec
ROOT_DIR=./cmd
VERSION=$(shell git describe --tags --always --long --dirty)
LOCAL=$(EXECUTABLE)
WINDOWS=$(EXECUTABLE)_windows_amd64_$(VERSION).exe
LINUX=$(EXECUTABLE)_linux_amd64_$(VERSION)
DARWIN=$(EXECUTABLE)_darwin_amd64_$(VERSION)

.PHONY: all test clean

all: test build

test:
	go test ./... $(ROOT_DIR)

build: local
	@echo version: $(VERSION)

build-all: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

local: $(LOCAL) ## Build for local machine

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(LOCAL):
	go build -v -o bin/$(LOCAL) -ldflags="-s -w -X main.version=$(VERSION)" $(ROOT_DIR)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o bin/$(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" $(ROOT_DIR)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o bin/$(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" $(ROOT_DIR)

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o bin/$(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" $(ROOT_DIR)

docker-build:
	@docker build . --target bin \
	--output bin/

clean:
	rm -rf ./bin
	mkdir ./bin

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
