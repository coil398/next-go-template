FROM golang:1.19-alpine AS dev
COPY . /backend
WORKDIR /backend
RUN apk update && apk upgrade && apk add --no-cache git mysql-client \
    && go get -u \
    && apk del --purge git \
    && go install github.com/cosmtrek/air@latest

ENV SERVER_PORT=8000
