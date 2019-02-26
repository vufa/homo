DIST := dist
IMPORT := github.com/countstarlight/homo
 
GO ?= go
SED_INPLACE := sed -i

ifeq ($(OS), Windows_NT)
	EXECUTABLE_INTERACT := homo-interact.exe
	EXECUTABLE_SERVER := homo-server.exe
else
	EXECUTABLE_INTERACT := homo-interact
    EXECUTABLE_SERVER := homo-server
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		SED_INPLACE := sed -i ''
	endif
endif

GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")
GOBINS := ${GOPATH}/bin
GOFMT ?= gofmt -s

GOFLAGS := -i -v
EXTRA_GOFLAGS ?=

PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: all
all: build

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -f $(EXECUTABLE_INTERACT)
	rm -f $(EXECUTABLE_SERVER)

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	# get all go files and run go fmt on them
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;
.PHONY: build
build: $(EXECUTABLE_INTERACT) $(EXECUTABLE_SERVER)

$(EXECUTABLE_INTERACT): $(SOURCES)
	cd ./cmd/interact; \
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@; \
	mv $@ ../../

$(EXECUTABLE_SERVER): $(SOURCES)
	cd ./cmd/server; \
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@; \
	mv $@ ../../
