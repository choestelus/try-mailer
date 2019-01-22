#!/bin/bash

docker run --rm -it -v "$(pwd)":/go/src/github.com/choestelus/try-mailer \
    -w /go/src/github.com/choestelus/try-mailer golang:1.11-alpine sh -c 'echo "build..." && apk update && \
    go build -ldflags "-X main.authKey=seup_key_here" -o ./build/mailer ./cmd'
