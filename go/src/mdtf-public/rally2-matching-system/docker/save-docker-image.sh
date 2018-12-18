#!/bin/bash

company=$1

if [ -z $company ]; then
  company=unknown
fi

docker save $company-rally2-matching-system:latest | gzip > $company-rally2-matching-system.tgz
