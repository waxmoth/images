FROM golang:1.17.2

RUN export DEBIAN_FRONTEND=noninteractive

WORKDIR /go/src/image-functions

# Install common tools
RUN apt update && apt install -y -q curl git ssh wget awscli

# Install node and npm
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get install -y -q nodejs
