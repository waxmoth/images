# Go images
The image service writen by golang

## Features
1. Download image from url
2. Crop the image;
3. Compress the image size;
4. Upload to the cloud service

## Dependencies
1. [Docker](https://www.docker.com/)
2. [Serverless](https://www.serverless.com/)


## Development
* Build image in your local
```shell script
docker-compose build
```

* Run the binary file build and watch
```shell script
docker-compose up -d
```

* Start the serverless offline, there no runner for runtime go1.x in the serverless offline, so use the docker to run it
```shell script
./node_modules/.bin/serverless offline start --useDocker --stage=dev
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