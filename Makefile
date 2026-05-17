VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X main.Version=$(VERSION)

# Build the kuberik CLI binary.
.PHONY: build
build:
	go build -ldflags '$(LDFLAGS)' -o bin/kuberik ./cmd/kuberik

# Install the kuberik CLI to $GOPATH/bin.
.PHONY: install
install:
	go install -ldflags '$(LDFLAGS)' ./cmd/kuberik

# Run unit tests.
.PHONY: test
test:
	go test -race -coverprofile=coverage.txt ./...

# Tidy go modules.
.PHONY: tidy
tidy:
	go mod tidy

# Format Go code.
.PHONY: fmt
fmt:
	gofmt -s -w cmd

# Combined pre-PR sanity check: tidy, format, vet, test.
.PHONY: presubmit
presubmit: tidy fmt vet test

# Run go vet.
.PHONY: vet
vet:
	go vet ./...

# Lint with golangci-lint (must be installed).
.PHONY: lint
lint:
	golangci-lint run ./...

# Remove build artifacts.
.PHONY: clean
clean:
	rm -rf bin/ dist/ coverage.txt
