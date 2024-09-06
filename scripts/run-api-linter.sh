#!/usr/bin/env bash
set -e

if ! which api-linter ; then
    go install github.com/aep-dev/api-linter/cmd/api-linter@latest
fi

if [ ! -d /tmp/googleapis ]; then
    git clone https://github.com/googleapis/googleapis.git --depth=1 /tmp/googleapis
fi

api-linter \
    ./example/bookstore/v1/bookstore.proto \
    -I /tmp/googleapis \
    --set-exit-status