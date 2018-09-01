BUILD_TIME=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GO=${shell which go}
GOOS=$(shell ${GO} env GOOS)
GOARCH=$(shell ${GO} env GOARCH)
GOCACHE ?= /tmp/go-build
LD_FLAGS=-X ${GO_PKG}/api/build.GitVersion=${GIT_VERSION} -X ${GO_PKG}/api/build.GitRefs=${GIT_REFS} -X ${GO_PKG}/api/build.BuildTime=${BUILD_TIME}
OUTPUT_DIR=output
OUTPUT_NAME=play
GO_PKG=github.com/play

build:
	GOOS=${GOOS} GOARCH=${GOARCH} GOCACHE=${GOCACHE} ${GO} build -ldflags "${LD_FLAGS}" -gcflags "${GC_FLAGS}" -o ${OUTPUT_DIR}/${OUTPUT_NAME} ${GO_PKG}

clean:
	$(GO) clean
	rm -fr ${OUTPUT}/${OUTPUT_NAME}

test:
	${GO} test ./...

linters:
	gometalinter --config gometalinter-config.json ./...
