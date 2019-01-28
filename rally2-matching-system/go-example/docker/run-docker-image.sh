#!/bin/bash

COMPANY_NAME=$1
HOST_PORT=$2

if [ -z $COMPANY_NAME ]; then
  COMPANY_NAME=unknown
fi
if [ -z $HOST_PORT ]; then
  HOST_PORT=8080
fi


#run the docker image in detached mode - forward a host port
docker run -d -p $HOST_PORT:8080 ${COMPANY_NAME}-rally2-matching-system
