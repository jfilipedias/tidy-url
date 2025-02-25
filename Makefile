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

## run: executes the program
.PHONY: run
run: 
	@echo 'Generating template pages'
	templ generate
	@echo 'Generating css files'
	npx @tailwindcss/cli -i ui/static/css/style.css -o ./ui/static/css/tailwind.css
	go run ./cmd/app

## tools: install development tools
.PHONY: tools
tools:
	@echo 'Installing mockery'
	go install github.com/vektra/mockery/v2@v2.52.3
	@echo 'Installing templ'
	go install github.com/a-h/templ/cmd/templ@latest
	@echo 'Installing Tailwind CSS'
	npm install -g @tailwindcss/cli

# ==================================================================================== #
# TEST
# ==================================================================================== #

## test: run application tests
.PHONY: test
test:
	go test ./...
