.PHONY: clean build

APP_BINARY_NAME=main

build:
	@echo "Building lambda package ..."
	go build -ldflags="-s -w" -o ../bin/main ../src/lambda/main.go

clean:
	@echo "Cleaning package ..."
	@rm -rf ../bin/main*

test:
	@echo "Testing the project ..."
	go test -v ../tests

deploy: test clean build
	@echo "Deploying application ..."
	sls deploy --verbose

## Makes for development
build-dev: stop-dev clean
	@echo "Building development package ..."
	go build -mod=readonly -ldflags="-s -w" -o ../bin/main-dev ../src/main.go

run-dev: build-dev
	@echo "Starting the application ..."
	../bin/main-dev & echo $$! > /tmp/image-go.pid

stop-dev:
	@echo "Stopping the application ..."
	@kill `cat /tmp/image-go.pid` || true
