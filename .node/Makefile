.PHONY: clean build package

PID = /tmp/image-go.pid

build:
	@echo "Building lambda package ..."
	go build -ldflags="-s -w" -o ../bin/main ../src/lambda/main.go

clean:
	@echo "Cleaning package ..."
	@rm -rf ../bin/main*

deploy: clean build
	@echo "Deploying application ..."
	sls deploy --verbose

## Makes for development
build-dev: clean
	@echo "Building development package ..."
	go build -mod=readonly -ldflags="-s -w" -o ../bin/main-dev ../src/main.go

run-dev: build-dev
	@echo "Starting the application ..."
	../bin/main-dev & echo $$! > $(PID)

stop-dev:
	@echo "Stopping the application ..."
	@kill `cat $(PID)` || true
