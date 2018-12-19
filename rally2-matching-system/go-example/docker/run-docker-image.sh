#!/bin/bash

HOST_PORT=8080

#run the docker image in detached mode - forward a host port
docker run -d -p $HOST_PORT:8080 unknown-rally2-matching-system
