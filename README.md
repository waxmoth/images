# Go images
The image service written by golang.

## Features
1. Download image from url;
2. Crop the image;
3. Compress the image size;
4. Upload to the cloud service;

## Dependencies
1. [Docker](https://www.docker.com/)
2. [Serverless](https://www.serverless.com/)


## Development
* Copy and set environment values
```shell script
cp .env.dist .env
```

* Build image in your local
```shell script
docker-compose build
```

* Install required NodeJS libraries
```shell script
docker-compose run --rm app bash -c "npm --prefix /go/src/image-functions/.node install"
```

* If you want install some NodeJS packages
```shell script
docker-compose run --rm app bash -c "npm --prefix /go/src/image-functions/.node install serverless --save-dev"
```

* Install required Golang libraries
```shell script
docker-compose run --rm app bash -c "go mod download && go mod vendor"
```

* Run the binary file build and watch
```shell script
docker-compose up -d
```

* Build the production file
```shell script
docker-compose run --rm app make build
```

* Start the serverless offline, there no runner for runtime go1.x in the serverless offline, so use the docker to run it
```shell script
.node/node_modules/.bin/serverless offline start --useDocker --stage=dev
```

* Run or Debug it by IDE
We can use the GoLand to build and run the [main.go](./src/main.go)
It will start a http server by using the port 8080

* Use Goland old version (<2020.1.4) to load new Go SDK (1.17+)? Error: unpacked SDK is corrupted
```
vim <GO_SDK_PATH>/src/runtime/internal/sys/zversion.go
# ADD the const TheVersion
const TheVersion = `go1.17.*`
```