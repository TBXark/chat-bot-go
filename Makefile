BUILD=$(shell git rev-parse --short HEAD)@$(shell date +%s)
GO_BUILD=CGO_ENABLED=0 go build -ldflags "-X 'main.buildVersion=$(BUILD)'"

.PHONY: init
init:
	go mod download

.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build:
	$(GO_BUILD) -o ./build/ ./...

.PHONY: buildLinuxX86
buildLinuxX86: build
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./build/ ./...

