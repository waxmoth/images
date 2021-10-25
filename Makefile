.PHONY: clean build package

build:
	@echo "Building lambda package ..."
	go build -ldflags="-s -w" -o bin/main ./src/lambda/main.go

clean:
	@echo "Cleaning package ..."
	@rm -rf ./bin/main*

deploy: clean build
	@echo "Deploying application ..."
	sls deploy --verbose
