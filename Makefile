LIST_ALL := $(shell go list ./... | grep -v /vendor/)


.PHONY: all lint test race coverage report build package clean dep help


all: lint test race build


lint: ## Lint all files
	@go fmt ${LIST_ALL}
	@golint -set_exit_status ${LIST_ALL}

test: dep ## Run unittests
	@go test -short ${LIST_ALL}

race: dep ## Run data race detector
	@go test -race -short ${LIST_ALL}

coverage: dep # Generate coverage report
	@go test ${LIST_ALL}  -coverprofile coverage.out
	@go tool cover -func coverage.out

report: coverage # Open the coverage report in browser
	@go tool cover -html=coverage.out


build: dep ## Build all binaries based on directory `./`
	@go build ./

#package: build ## Generate ZIP files of binaries based on directory `./`
#	zip princess-of-mland-e.zip princess-of-mland-e && rm -f princess-of-mland-e

clean: ## Remove binaries and ZIP files based on directory `./`
	rm -f princess-of-mland-e


dep: ## Get the dependencies
	@dep ensure

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
