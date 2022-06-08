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
  --request GET \
  http://${HOST}:${PORT}/v1/info
