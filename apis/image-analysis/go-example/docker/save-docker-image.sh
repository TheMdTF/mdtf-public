#!/bin/bash

COMPANY_NAME=$1

if [ -z $COMPANY_NAME ]; then
  COMPANY_NAME=unknown
fi

docker save ${COMPANY_NAME}-image-analysis:latest | gzip > ${COMPANY_NAME}-image-analysis.tgz
