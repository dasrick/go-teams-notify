LIST_ALL := $(shell go list ./... | grep -v /vendor/)

# Force using Go Modules and always read the dependencies from
# the `vendor` folder.
export GO111MODULE = on
#export GOFLAGS = -mod=vendor

all: lint test

.PHONY: install
install: ## Install the dependencies
	@go mod vendor

.PHONY: update
update: ## Update the dependencies
	@go mod tidy

.PHONY: clean
clean: ## Remove binaries and ZIP files based on directory `./cmd/`
	@rm -rf "$(go env GOCACHE)"
	@rm -f coverage.out

.PHONY: lint
lint: ## Lint all files (via golint)
	@go fmt ${LIST_ALL}
	@golint -set_exit_status ${LIST_ALL}

.PHONY: test
test: clean ## Run unit tests (incl. race and coverprofile)
	@go test -race -cover -short -timeout=90s -coverprofile=coverage.out ${LIST_ALL}

.PHONY: coverage
coverage: test ## Generate coverage report
	@go tool cover -func coverage.out

.PHONY: report
report: coverage ## Open the coverage report in browser
	@go tool cover -html=coverage.out

# ----------------------------------------------------------------------------------------------------------------------

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
