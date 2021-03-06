GO ?= go

.PHONY: all
all: pinkie

%: cmd/%/main.go pkg/*/*.go
	$(GO) build -o $@ $<

.PHONY: test
test:
	$(GO) test -covermode=atomic -coverpkg=./... -coverprofile=coverage.txt -race ./...
