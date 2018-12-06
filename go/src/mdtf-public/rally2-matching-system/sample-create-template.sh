#!/bin/bash

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="}' \
  http://localhost:8080/v1/create-template