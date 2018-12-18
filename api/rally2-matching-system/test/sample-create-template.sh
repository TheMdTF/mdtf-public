#!/bin/bash

host=$1
port=$2

if [ -z $host ]; then
  host=localhost
fi

if [ -z $port ]; then
  port=8080
fi

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
  }' \
  http://$host:$port/v1/create-template