# Environment variable default values.
BINARY_NAME=alberto

# Default target.
.PHONY: all
all: build

# Unless VERBOSE env var is defined, do not output executed commands.
# Warning: do _not_ put this at top above default target.
ifndef VERBOSE
.SILENT:
endif

.PHONY: build
build: vendor
	echo Compiling binary
	go build -o $(BINARY_NAME)

.PHONY: test
test: vendor
	echo Running tests
	ENV=test go test ./...

.PHONY: clean
clean:
	echo Removing compiled binaries
	go clean
	echo Cleaning cached tests
	go clean -testcache
	echo Uninstalling dependencies
	rm -rf vendor

.PHONY: run
run: build
	./$(BINARY_NAME)

vendor: Gopkg.toml Gopkg.lock
	echo Installing dependencies
	dep ensure
