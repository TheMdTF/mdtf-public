#!/bin/bash

HOST=$1
PORT=$2

if [ -z $HOST ]; then
  HOST=localhost
fi

if [ -z $PORT ]; then
  PORT=8080
fi

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
  "ImageData": "$(base64 -w 0 ./test-images/S223-01-t10_01.png)"
  }' \
  http://${HOST}:${PORT}/v1/analyze-image
