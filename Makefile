ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=input-server
VERSION=0.0.1
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o ${BINARY}

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/$(BINARY)-$(GOOS)-$(GOARCH))))

install:
	go install ${LDFLAGS}

test:
	go test -v ./...

clean:
	rm -rf ${ROOT_DIR}/bin