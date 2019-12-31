DIST := dist
IMPORT := github.com/countstarlight/homo

GO ?= env GO111MODULE=on go
SED_INPLACE := sed -i
EXTRA_GOENVS ?=

GIT_REV:=git-$(shell git rev-parse --short HEAD)
GIT_TAG:=$(shell git tag --contains HEAD)
VERSION:=$(if $(GIT_TAG),$(GIT_TAG),$(GIT_REV))

ifeq ($(OS), Windows_NT)
	EXECUTABLE_MASTER := homo-master.exe
	EXTRA_GOENVS = GOOS=windows GOARCH=amd64
else
	EXECUTABLE_MASTER := homo-master
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		SED_INPLACE := sed -i ''
		EXTRA_GOENVS = GOOS=darwin GOARCH=amd64
	else
		EXTRA_GOENVS = GOOS=linux GOARCH=amd64
	endif
endif

GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")
GOBINS := ${GOPATH}/bin
GOFMT ?= gofmt -s

GOFLAGS := -mod=vendor -v
LDFLAGS := $(LDFLAGS) -X "$(IMPORT)/cmd.Revision=$(GIT_REV)" -X "$(IMPORT)/cmd.Version=$(VERSION)"

PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: all
all: build

.PHONY: gen
gen:
	@hash go-bindata > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/jteeuwen/go-bindata; \
		$(GO) get -u github.com/jteeuwen/go-bindata/...; \
	fi
	$(GO) generate github.com/countstarlight/homo/module/view

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -f $(EXECUTABLE_MASTER)

.PHONY: docker
docker:
	docker build -t homo:v0.0.1 .

.PHONY: dockercn
dockercn:
	docker-compose -f docker/homo-compose-zh.yml up -d --build

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

.PHONY: watch
watch: gen $(EXECUTABLE_WEBVIEW)
	./$(EXECUTABLE_WEBVIEW) -d

.PHONY: build
build: $(EXECUTABLE_MASTER)

$(EXECUTABLE_MASTER): $(SOURCES)
	$(EXTRA_GOENVS) $(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -ldflags '-s -w $(LDFLAGS)' -o $@;
