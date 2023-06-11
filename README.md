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

* If you want to install some NodeJS packages
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

* Run the unit test
```shell script
docker-compose run --rm app make test
```

* Run the benchmark
```shell script
docker-compose run --rm app make benchmark
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

## Local S3 service based on [LocalStack](https://localstack.cloud/)

* Create and use local S3 bucket
```shell script
aws --endpoint-url http://localhost:4566 s3api create-bucket --bucket test

# List buckets
aws --endpoint-url http://localhost:4566 s3api list-buckets

# Copy one file into S3
aws --endpoint-url http://localhost:4566 s3 cp README.md s3://test/

# List the copied file
aws --endpoint-url http://localhost:4566 s3 ls s3://test/
```
