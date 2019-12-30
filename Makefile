DIST := dist
IMPORT := github.com/countstarlight/homo
 
GO ?= env GO111MODULE=on go
SED_INPLACE := sed -i

ifeq ($(OS), Windows_NT)
	EXECUTABLE_MASTER := homo-master.exe
else
	EXECUTABLE_MASTER := homo-master
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		SED_INPLACE := sed -i ''
	endif
endif

GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")
GOBINS := ${GOPATH}/bin
GOFMT ?= gofmt -s

GOFLAGS := -mod=vendor -v

PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: all
all: build

.PHONY: deps
deps:
	echo "Installing sphinxbase..."; \
	git clone https://github.com/countstarlight/sphinxbase.git; \
	cd sphinxbase && ./autogen.sh && ./configure && make -j 4 && sudo make install; \
	cd .. && rm -rf sphinxbase; \
	echo "Installing PocketSphinx..."; \
	git clone https://github.com/countstarlight/pocketsphinx.git; \
	cp -r pocketsphinx/model/en-us sphinx/; \
	cd pocketsphinx && ./autogen.sh && ./configure && make -j 4 && sudo make install; \
	cd .. && rm -rf pocketsphinx

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
	cd ./master; \
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@; \
	mv $@ ../
