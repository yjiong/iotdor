.PHONY: build clean test update-vendor FORCE
BUILD_TIME := $(shell date "+%F %T")
PROJECT := iotdor
MKFILEPATH := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
THISDIR := $(patsubst %/, %, $(MKFILEPATH))
PKGS := $(shell go list ./... | grep -v /vendor)
VERSION := $(shell git describe --always |sed -e "s/^v//" |cut -d- -f 1)
XENV := -X 'main.version=$(VERSION)'
DFLAGS = $(GO_EXTRA_BUILD_ARGS) -v $(BUILDTAGS) -gcflags '-N -l' -ldflags "$(XENV)"
#ENV := CGO_CFLAGS='-g -O2 -Wno-return-local-addr' CGO_ENABLED=1
#ENV := CGO_CFLAGS='-Wno-return-local-addr' CGO_ENABLED=0
ifeq ("$(origin D)", "command line")
	FLAGS = $(DFLAGS)
else
	FLAGS = $(GO_EXTRA_BUILD_ARGS) -v $(BUILDTAGS) -ldflags "-s -w $(XENV)"
endif

ifeq ("$(GOOS)", "windows")
	ENV = CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 
endif

ifeq ("$(GOARCH)", "arm64")
	ENV = CC=aarch64-linux-gnu-gcc CXX=aarch64-linux-gnu-g++ CGO_ENABLED=1 GOOS=linux
endif

build:
	@echo "Compiling source for ENV=$(ENV) GOOS=$(GOOS) GOARCH=$(GOARCH) BUILDTAGS=$(BUILDTAGS)"
	@echo $(FLAGS)
	@mkdir -p build
	@env $(ENV) go build $(FLAGS) -o ./build/$(PROJECT) ./cmd/${PROJECT}/main.go

swagger:
	@swag init -g api/http_server.go -o api/docs

clean:
	@echo "Cleaning up workspace"
	@rm -rf build 

test:
	@echo "Running tests"
	@for pkg in $(PKGS) ; do \
		golint $$pkg ; \
	done
	@go vet $(PKGS)
	@go test -p 1 -v $(PKGS)

run: build 
	@cd build && $(THISDIR)/build/./$(PROJECT) -L 5

debug:swagger FORCE
	@mkdir -p build
	@env $(ENV) go build $(DFLAGS) -o ./build/$(PROJECT) ./cmd/$(PROJECT)/main.go
	@cd build && dlv exec ./$(PROJECT) -- -L 5

gupdate-vendor:
	@echo "Updating vendored packages"
	@go mod vendor

