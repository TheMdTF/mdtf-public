#!/bin/bash

COMPANY_NAME=$1

if [ -z $COMPANY_NAME ]; then
  COMPANY_NAME=unknown
fi

cd ../
docker build -t ${COMPANY_NAME}-rally2-matching-system -f docker/Dockerfile .
cd docker
