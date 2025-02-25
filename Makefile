.PHONY: build
build:
	go build -o=./tmp/app ./cmd/app

.PHONY: run
run: build
	air 

.PHONY: test
test:
	go test ./...

.PHONY: tools-install
tools-install: 
	go install github.com/vektra/mockery/v2@v2.52.3
	go install github.com/air-verse/air@latest

.PHONY: test-mock
test-mock:
	mockery
