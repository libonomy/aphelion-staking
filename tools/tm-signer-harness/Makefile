.PHONY: build install docker-image

LIBONOMY_VERSION?=latest
BUILD_TAGS?='libonomy'
BUILD_FLAGS = -ldflags "-X github.com/evdatsion/aphelion-dpos-bft/version.GitCommit=`git rev-parse --short=8 HEAD`"

.DEFAULT_GOAL := build

build:
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -tags $(BUILD_TAGS) -o ../../build/tm-signer-harness main.go

install:
	CGO_ENABLED=0 go install $(BUILD_FLAGS) -tags $(BUILD_TAGS) .

docker-image:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -tags $(BUILD_TAGS) -o tm-signer-harness main.go
	docker build \
		--build-arg LIBONOMY_VERSION=$(LIBONOMY_VERSION) \
		-t libonomy/tm-signer-harness:$(LIBONOMY_VERSION) .
	rm -rf tm-signer-harness
