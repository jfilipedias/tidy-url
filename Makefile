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
	@echo 'Generating template pages'
	go tool templ generate
	@echo 'Generating css files'
	npx @tailwindcss/cli -i ./ui/static/css/style.css -o ./ui/static/css/tailwind.css
	go run ./cmd/app

# ==================================================================================== #
# TEST
# ==================================================================================== #

## test: run application tests
.PHONY: tests
tests:
	go test ./...
