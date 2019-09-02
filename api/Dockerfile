FROM golang:1.12.9-alpine3.10

RUN apk add git \
    && go get -u github.com/golang/dep/cmd/dep \
    && mkdir /go/src/api

# dep init 