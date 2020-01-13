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
	EXECUTABLE_HUB := homo-hub.exe
	EXTRA_GOENVS = GOOS=windows GOARCH=amd64
else
	EXECUTABLE_MASTER := homo-master
	EXECUTABLE_HUB := homo-hub
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

.PHONY: deps
deps:
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/jsonpb
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/gogoproto

.PHONY: all
all: build

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -f $(EXECUTABLE_MASTER) $(EXECUTABLE_HUB)

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

.PHONY: build
build: $(EXECUTABLE_MASTER) $(EXECUTABLE_HUB)

$(EXECUTABLE_MASTER): $(SOURCES)
	$(EXTRA_GOENVS) $(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -ldflags '-s -w $(LDFLAGS)' -o $@;

$(EXECUTABLE_HUB): $(SOURCES)
	cd ./hub; \
	$(EXTRA_GOENVS) $(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -ldflags '-s -w $(LDFLAGS)' -o $@; \
	mv $@ ../
