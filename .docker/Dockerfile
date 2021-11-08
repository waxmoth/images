FROM golang:1.17.2

RUN export DEBIAN_FRONTEND=noninteractive

WORKDIR /go/src/image-functions

# Install common tools
RUN apt update && apt install -y -q curl git ssh wget awscli

# Install node and serverless
RUN apt install -y -q nodejs npm
