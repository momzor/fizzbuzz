COVER_OUT=${PWD}/$(IGNORED_FOLDER)/coverage.out

.PHONY:  deps up down doc mock unit-test cover

##
## install project dependencies
##
deps:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/golang/mock/mockgen@latest

install: ## Download and install go mod
	@go mod download

## up with docker-compose
up:
	@D_PATH=Dockerfile docker-compose up --remove-orphans --build -d
	@docker-compose logs -f lbc_test_fizzbuzz

## down with docker-compose
down:
	@docker-compose down

## generate swagger documentation for web api
doc:
	@swag init -g ./pkg/webapi/server.go

## generate mock for unit tests
mock:
	@go generate ./...

## run mock defined above, then launch unit tests
unit-test: mock
	@mkdir -p .ignore
	@go test -gcflags=-l -count=1 -race -coverprofile=${COVER_OUT} -covermode=atomic ./...

## Cover
cover:
	@if [ ! -e ${COVER_OUT} ]; then \
		echo "Error: ${COVER_OUT} doesn't exists. Please run \`make test\` then retry."; \
		exit 1; \
	fi
	@go tool cover -func=${COVER_OUT}

## TODO: ADD make dev for hot reload with Air