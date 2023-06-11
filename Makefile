.PHONY: clean build package

build:
	@echo "Building lambda package ..."
	go build -ldflags="-s -w" -o bin/main ./src/lambda/main.go

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
	sls deploy --verbose
