#!/bin/bash

docker run --rm -it -v "$(pwd)":/go/src/github.com/choestelus/try-mailer \
    -w /go/src/github.com/choestelus/try-mailer golang:1.11-alpine sh -c 'echo "build..." && apk update && \
    go build -o ./build/mailer ./cmd'
