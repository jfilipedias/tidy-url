include .env
export

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: prints this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## mocks: generate interfaces mocks
.PHONY: mocks
mocks:
	@echo 'Generating mocks'
	go tool mockery

## run: executes the program
.PHONY: run
run: 
	go tool air

# ==================================================================================== #
# TEST
# ==================================================================================== #

## test: run application tests
.PHONY: tests
tests:
	go test ./...
