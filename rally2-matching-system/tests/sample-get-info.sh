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
  --request GET \
  http://$host:$port/v1/info
