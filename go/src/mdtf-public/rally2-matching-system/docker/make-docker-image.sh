#!/bin/bash

company=$1

if [ -z $company ]; then
  company=unknown
fi

cd ../
docker build -t $company-rally2-matching-system .
cd docker