#!/bin/bash

COMPANY_NAME=$1

if [ -z $COMPANY_NAME ]; then
  COMPANY_NAME=unknown
fi

cd ../
docker build -t ${COMPANY_NAME}-image-analysis -f docker/Dockerfile .
cd docker
