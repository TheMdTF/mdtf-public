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
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
  }' \
  http://${HOST}:${PORT}/v1/create-template
