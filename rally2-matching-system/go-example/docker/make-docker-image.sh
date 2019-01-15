#!/bin/bash

company=$1

if [ -z $company ]; then
  company=unknown
fi

cd ../
docker build -t $company-rally2-matching-system -f docker/Dockerfile .
cd docker
