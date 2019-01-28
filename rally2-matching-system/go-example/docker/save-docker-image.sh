#!/bin/bash

COMPANY_NAME=$1

if [ -z $COMPANY_NAME ]; then
  COMPANY_NAME=unknown
fi

docker save ${COMPANY_NAME}-rally2-matching-system:latest | gzip > ${COMPANY_NAME}-rally2-matching-system.tgz
