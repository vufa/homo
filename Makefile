DIST := dist
IMPORT := github.com/countstarlight/homo
 
GO ?= go
SED_INPLACE := sed -i
EXTRA_GOFLAGS ?=

ifeq ($(OS), Windows_NT)
	EXECUTABLE_INTERACT := homo-interact.exe
	EXECUTABLE_SERVER := homo-server.exe
	EXECUTABLE_NATIVE := homo-native.exe
	EXECUTABLE_WEBVIEW := homo-webview.exe
	EXECUTABLE_QT := homo-qt.exe
	#EXTRA_GOFLAGS = -tags netgo -ldflags '-H=windowsgui -extldflags "-static" -s'
else
	EXECUTABLE_INTERACT := homo-interact
	EXECUTABLE_SERVER := homo-server
	EXECUTABLE_NATIVE := homo-native
	EXECUTABLE_WEBVIEW := homo-webview
	EXECUTABLE_QT := homo-qt
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		SED_INPLACE := sed -i ''
		#EXTRA_GOFLAGS = -ldflags '-s -extldflags "-sectcreate __TEXT __info_plist Info.plist"'
	else
		#EXTRA_GOFLAGS = -tags netgo -ldflags '-extldflags "-static" -s'
	endif
endif

GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")
GOBINS := ${GOPATH}/bin
GOFMT ?= gofmt -s

GOFLAGS := -v

PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: all
all: build

.PHONY: gen
gen:
	@hash go-bindata > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/jteeuwen/go-bindata/...; \
	fi
	$(GO) generate github.com/countstarlight/homo/cmd/webview

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -f $(EXECUTABLE_INTERACT)
	rm -f $(EXECUTABLE_SERVER)
	rm -f $(EXECUTABLE_NATIVE)
	rm -f $(EXECUTABLE_WEBVIEW)
	#qt
	#cd cmd/qt && \
	#rm -f moc* && \
	#rm -rf deploy && \
	#rm -f rcc*
	#rm -rf bin
	#mkdir bin

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

.PHONY: native
native: $(EXECUTABLE_NATIVE)

.PHONY: webview
webview: $(EXECUTABLE_WEBVIEW)

.PHONY: qt
qt:
	rm -rf bin/qt; \
	mkdir bin/qt; \
	qtdeploy build desktop github.com/countstarlight/homo/cmd/qt; \
	rm -rf bin/linux; \
	mv cmd/qt/deploy/* bin/qt/

.PHONY: build
build: $(EXECUTABLE_INTERACT) $(EXECUTABLE_SERVER) $(EXECUTABLE_NATIVE) $(EXECUTABLE_WEBVIEW)

$(EXECUTABLE_INTERACT): $(SOURCES)
	cd ./cmd/interact; \
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@; \
	mv $@ ../../

$(EXECUTABLE_SERVER): $(SOURCES)
	cd ./cmd/server; \
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@; \
	mv $@ ../../

$(EXECUTABLE_NATIVE): $(SOURCES)
	cd ./cmd/native; \
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@; \
	mv $@ ../../

$(EXECUTABLE_WEBVIEW): $(SOURCES)
	cd ./cmd/webview; \
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@; \
	mv $@ ../../
