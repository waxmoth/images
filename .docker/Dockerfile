FROM golang:1.20-alpine

# Install common tools
RUN apk add --update --no-cache \
    curl bash git wget make aws-cli
RUN aws --version

# Install NodeJS for serverless
RUN apk add --update --no-cache \
    nodejs npm
RUN node --version

# Clean the APK packages cache
RUN rm -rf /var/cache/apk/*

WORKDIR /go/src/image-functions
