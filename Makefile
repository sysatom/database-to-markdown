EXECUTABLE=database-to-markdown
WINDOWS=./dist/$(EXECUTABLE)_windows
LINUX=./dist/$(EXECUTABLE)_linux
DARWIN=./dist/$(EXECUTABLE)_darwin
VERSION=$(shell git describe --tags --always --long --dirty)

all: build

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=386 go build -i -v -o $(WINDOWS)_386.exe -ldflags="-s -w -X main.version=$(VERSION)"  ./main.go
	env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS)_amd64.exe -ldflags="-s -w -X main.version=$(VERSION)"  ./main.go

$(LINUX):
	env GOOS=linux GOARCH=386 go build -i -v -o $(LINUX)_386 -ldflags="-s -w -X main.version=$(VERSION)"  ./main.go
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX)_amd64 -ldflags="-s -w -X main.version=$(VERSION)"  ./main.go

$(DARWIN):
	env GOOS=darwin GOARCH=386 go build -i -v -o $(DARWIN)_386 -ldflags="-s -w -X main.version=$(VERSION)"  ./main.go
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN)_amd64 -ldflags="-s -w -X main.version=$(VERSION)"  ./main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'