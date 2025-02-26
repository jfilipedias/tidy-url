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

## views: generate the templ files
.PHONY: views
views: 
	go tool templ generate -watch

## views: generate the css files
.PHONY: styles
styles: 
	npx @tailwindcss/cli -i ./view/static/css/style.css -o ./view/static/css/tailwind.css --watch 

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
