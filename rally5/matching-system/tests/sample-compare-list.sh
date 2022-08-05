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
  "ProbeTemplate": {
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
  },
  "TargetTemplateList": [
    {
      "Template": "aSBzZWUgeW91IGhhdmUgYSBjdXJpb3VzIG1pbmQ="
    },
    {
      "Template": "aWYgeW91IGFyZSByZWFkaW5nIHRoaXM="
    },
    {
      "Template": "eW91IHNob3VsZCBiZSB3b3JraW5nIGF0IHRoZSBNZFRGIQ=="
    },
    {
      "Template": "YXBwbGljYXRpb25zIGNhbiBiZSBzZW50IHRvIHJhbGx5QG1kdGYub3Jn"
    },
    {
      "Template": "Z29vZCBsdWNrIGluIHRoZSByYWxseSE="
    }
  ]
  }' \
  http://${HOST}:${PORT}/v1/compare-list
