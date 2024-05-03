export PATH := $(PATH):`go env GOPATH`/bin
export GO111MODULE=on
LDFLAGS := -s -w

all: env fmt build

build: apifs

env:
	@go version

fmt:
	go fmt ./...

apifs:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -tags frps -o bin/apifs ./cmd/apifs

test: gotest

gotest:
	go test -v --cover ./cmd/...
	go test -v --cover ./pkg/...

clean:
	rm -f ./bin/apifs
