project_name = leaderboard-api
image_name = leaderboard:latest
port = 1919

help: ## This help dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run main.go

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

up: ## Run the project in a local container
	make up-silent
	make shell

build: ## Generate docker image
	docker build -t $(image_name) .

build-no-cache: ## Generate docker image with no cache
	docker build --no-cache -t $(image_name) .

up-silent: ## Run local container in background
	make delete-container-if-exist
	docker run -d -p $(port):$(port) --env-file=.env --name $(project_name) $(image_name) ./main

up-silent-prefork: ## Run local container in background with prefork
	make delete-container-if-exist
	docker run -d -p $(port):$(port) --env-file=.env  --name $(project_name) $(image_name) ./main -prod

delete-container-if-exist: ## Delete container if it exists
	docker stop $(project_name) || true && docker rm $(project_name) || true

shell: ## Run interactive shell in the container
	docker exec -it $(project_name) /bin/sh

stop: ## Stop the container
	docker stop $(project_name)

start: ## Start the container
	docker start $(project_name)

test: ## Run the tests
	go test
