FROM golang:1.12.9-alpine3.10

RUN apk add git \
    && go get -u github.com/golang/dep/cmd/dep \
    && mkdir /go/src/api \
    && cd /go/src/api \
    && dep init 

WORKDIR /go/src/api

COPY base.go main.go

RUN pwd && ls -l

RUN dep ensure -add "go.mongodb.org/mongo-driver/mongo@~1.0.0"

COPY main.go .