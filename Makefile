# default version
VERSION ?= v1.0.0

# binaries
BIN_DARWIN_AMD64 = internal-proxy-darwin-amd64-${VERSION}
BIN_DARWIN_ARM64 = internal-proxy-darwin-arm64-${VERSION}
BIN_WINDOWS_AMD64 = internal-proxy-windows-amd64-${VERSION}

.PHONY: all build clean

# targets
all: build

build: ${BIN_DARWIN_AMD64} ${BIN_DARWIN_ARM64} ${BIN_WINDOWS_AMD64}

${BIN_DARWIN_AMD64}:
	GOOS=darwin GOARCH=amd64 go build -o $@ .

${BIN_DARWIN_ARM64}:
	GOOS=darwin GOARCH=arm64 go build -o $@ .

${BIN_WINDOWS_AMD64}:
	GOOS=windows GOARCH=amd64 go build -o $@ .

clean:
	rm -f ${BIN_DARWIN_AMD64} \
		  ${BIN_DARWIN_ARM64} \
		  ${BIN_WINDOWS_AMD64}