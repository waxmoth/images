.PHONY: clean build package

build:
	@echo "Building lambda package ..."
	export GIN_MODE=release && go build -ldflags="-s -w" -o bin/main ./src/main.go

build-lambda:
	@echo "Building lambda package ..."
	export GIN_MODE=release && go build -ldflags="-s -w" -o bin/main ./src/lambda/main.go

clean:
	@echo "Cleaning package ..."
	@rm -rf ./bin/main*

install:
	@echo "Installing package ..."
	go mod download && go mod vendor

test:
	@echo "Testing the project ..."
	export GIN_MODE=release && go test -v ./tests/...

benchmark:
	@echo "Benchmarking the project ..."
	export GIN_MODE=test && go test -v ./tests/... -bench=. -count 5 -benchmem -run=^#

deploy: clean build
	@echo "Deploying application ..."
	.node/node_modules/.bin/sls deploy --verbose

upgrade: clean
	@echo "Updating application ..."
	go get -u ./... && go mod tidy && go mod vendor
	make build && make test

install-swag:
ifeq (, $(shell which swag))
	@echo "Installing the Swagger ..."
	go install github.com/swaggo/swag/cmd/swag@latest;
endif

swag-fmt: install-swag
	@echo "Formatting the Swagger doc annotation ..."
	swag fmt

swag-init: install-swag
	@echo "Generating API doc ..."
	swag init -d src -g main.go -o doc/api/
