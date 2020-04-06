PREFIX?=/usr/local
MODE?=native
MODULES?=hub function function-python3
APPS?=nlu
SRC_FILES:=$(shell find main.go cmd master logger sdk protocol utils -type f -name '*.go')
PLATFORM_ALL:=darwin/amd64 linux/amd64 linux/arm64 linux/386 linux/arm/v7 linux/arm/v6 linux/arm/v5 linux/ppc64le linux/s390x

GIT_REV:=git-$(shell git rev-parse --short HEAD)
GIT_TAG:=$(shell git tag --contains HEAD)
VERSION:=$(if $(GIT_TAG),$(GIT_TAG),$(GIT_REV))

GO_OS:=$(shell go env GOOS)
GO_ARCH:=$(shell go env GOARCH)
GO_ARM:=$(shell go env GOARM)
GO_FLAGS?=-mod=vendor -v -ldflags "-X 'github.com/aiicy/aiicy/cmd.Revision=$(GIT_REV)' -X 'github.com/aiicy/aiicy/cmd.Version=$(VERSION)'"
GO_FLAGS_STATIC=-mod=vendor -v -ldflags '-X "github.com/aiicy/aiicy/cmd.Revision=$(GIT_REV)" -X "github.com/aiicy/aiicy/cmd.Version=$(VERSION)"  -linkmode external -w -extldflags "-static"'
GO_TEST_FLAGS?=-race -short -covermode=atomic -coverprofile=coverage.out
GO_TEST_PKGS?=$(shell go list ./...)

ifndef PLATFORMS
	GO_OS:=$(shell go env GOOS)
	GO_ARCH:=$(shell go env GOARCH)
	GO_ARM:=$(shell go env GOARM)
	PLATFORMS:=$(if $(GO_ARM),$(GO_OS)/$(GO_ARCH)/$(GO_ARM),$(GO_OS)/$(GO_ARCH))
	ifeq ($(GO_OS),darwin)
		PLATFORMS+=linux/amd64
	endif
else ifeq ($(PLATFORMS),all)
	override PLATFORMS:=$(PLATFORM_ALL)
endif

OUTPUT:=out
OUTPUT_DIRS:=$(PLATFORMS:%=$(OUTPUT)/%/aiicy)
OUTPUT_BINS:=$(OUTPUT_DIRS:%=%/bin/aiicy)

OUTPUT_MODS:=$(MODULES:%=aiicy-%)
IMAGE_MODS:=$(MODULES:%=image/aiicy-%) # a little tricky to add prefix 'image/' in order to distinguish from OUTPUT_MODS
NATIVE_MODS:=$(MODULES:%=native/aiicy-%) # a little tricky to add prefix 'native/' in order to distinguish from OUTPUT_MODS
OUTPUT_APPS:=$(APPS:%=application/%)

.PHONY: all $(OUTPUT_MODS)
all: aiicy $(OUTPUT_MODS)

aiicy: $(OUTPUT_BINS)

$(OUTPUT_BINS): $(SRC_FILES)
	@echo "BUILD $@"
	@mkdir -p $(dir $@)
	@# aiicy failed to collect cpu related data on darwin if set 'CGO_ENABLED=0' in compilation
	@$(shell echo $(@:$(OUTPUT)/%/aiicy/bin/aiicy=%)  | sed 's:/v:/:g' | awk -F '/' '{print "GO111MODULE=on GOOS="$$1" GOARCH="$$2" GOARM="$$3" go build"}') -o $@ ${GO_FLAGS} .

$(OUTPUT_MODS):
	@${MAKE} -C $@

$(OUTPUT_APPS):
	@${MAKE} -C $@

.PHONY: apps $(OUTPUT_APPS)
apps: $(OUTPUT_APPS)

.PHONY: build
build: $(SRC_FILES)
	@echo "BUILD aiicy"
ifneq ($(GO_OS),darwin)
	@GO111MODULE=on CGO_ENABLED=1 go build -o aiicy $(GO_FLAGS_STATIC) .
else
	@GO111MODULE=on CGO_ENABLED=1 go build -o aiicy $(GO_FLAGS) .
endif

.PHONY: rebuild
rebuild: clean all

.PHONY: test
test:
	@go test ${GO_TEST_FLAGS} ${GO_TEST_PKGS}
	@go tool cover -func=coverage.out | grep total

.PHONY: install $(NATIVE_MODS)
install: all
	@install -d -m 0755 ${PREFIX}/bin
	@install -m 0755 $(OUTPUT)/$(if $(GO_ARM),$(GO_OS)/$(GO_ARCH)/$(GO_ARM),$(GO_OS)/$(GO_ARCH))/aiicy/bin/aiicy ${PREFIX}/bin/
ifeq ($(MODE),native)
	@${MAKE} $(NATIVE_MODS)
endif
	@tar cf - -C conf/$(MODE) etc var | tar xvf - -C ${PREFIX}/

$(NATIVE_MODS):
	@install -d -m 0755 ${PREFIX}/var/db/aiicy/$(notdir $@)/bin
	@install -m 0755 $(OUTPUT)/$(if $(GO_ARM),$(GO_OS)/$(GO_ARCH)/$(GO_ARM),$(GO_OS)/$(GO_ARCH))/$(notdir $@)/bin/* ${PREFIX}/var/db/aiicy/$(notdir $@)/bin/
	@install -m 0755 $(OUTPUT)/$(if $(GO_ARM),$(GO_OS)/$(GO_ARCH)/$(GO_ARM),$(GO_OS)/$(GO_ARCH))/$(notdir $@)/package.yml ${PREFIX}/var/db/aiicy/$(notdir $@)/

.PHONY: deps
deps:
	go get -mod=readonly github.com/golang/mock/mockgen
	go get -mod=readonly github.com/golang/protobuf/proto
	go get -mod=readonly github.com/gogo/protobuf/proto
	go get -mod=readonly github.com/gogo/protobuf/jsonpb
	go get -mod=readonly github.com/gogo/protobuf/protoc-gen-gogo
	go get -mod=readonly github.com/gogo/protobuf/gogoproto

.PHONY: clean
clean:
	@-rm -rf $(OUTPUT)

.PHONY: gen
gen:
	@go generate -mod=vendor ./...

.PHONY: fmt
fmt:
	@go fmt  ./...

.PHONY: fmt-check
fmt-check:
	# get all go files and run go fmt on them
	@diff=$$(go fmt  ./...); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;
