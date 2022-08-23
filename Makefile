.DEFAULT_GOAL := help
.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

hello: ## init hello
		@echo "hello!"

build: ## build binary
		@CGO_ENABLED=0 go build

build-container: ## build container
		@docker build -t restapi .

run-container: ## run container
		@docker run --rm -p 3000:3000 restapi:latest