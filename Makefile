PROJECT_NAME := "go-etl-pipeline"
PKG := "gitlab.com/bosamatheus/$(PROJECT_NAME)"
MAIN_APP := "$(PKG)/cmd/pipeline"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

.PHONY: all dep build clean test coverage coverhtml lint

all: build

run: ## Run the server application
	@echo "Running server application"
	@go run ${MAIN_APP}

format: ## Format the files
	@echo "Formatting files"
	@go fmt ${PKG_LIST}

lint: ## Lint the files
	@echo "Linting files using revive"
	@revive -config ./revive.toml -formatter friendly ${PKG_LIST}

race: ## Run data race detector
	@echo "Running data race detector"
	@go test -race -short ${PKG_LIST}
	
test: ## Run unit tests
	@echo "Running unit tests"
	@go test -v ${PKG_LIST} -coverprofile coverage.out

coverage: ## Generate global code coverage report
	@echo "Generating coverage report"
	@go tool cover -func coverage.out

coverhtml: ## Generate global code coverage report in HTML
	@echo "Generating coverage report in HTML"
	@go tool cover -html=coverage.out -o coverage.html

dep: ## Get the dependencies
	@echo "Getting dependencies"
	@go get -v -d ./...
	@go get -u github.com/mgechev/revive

build: ## Build the binary file
	@echo "Building the project"
	@go build -v $(PKG)/cmd/${PROJECT_NAME}

clean: ## Remove previous build
	@echo "Cleaning the project"
	@rm -f $(PROJECT_NAME) coverage.out coverage.html

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
