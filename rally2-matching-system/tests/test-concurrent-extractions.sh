#!/bin/bash

#
# This script can be used to test concurrent requests to the `/v1/create-template` endpoint.
# Adjust the `workers` and `iterations` parameters to set how many concurrent workers will make requests, and the total
# number of requests.
#

# define concurrency settings
workers=4
iterations=5

# Set the service address and port, if passed
host=$1
port=$2


if [ -z $host ]; then
  host=localhost
fi

if [ -z $port ]; then
  port=8080
fi
echo "Host: " $host
echo -e "Port: " $port "\n"

# Get a list of images

cd test-routine-images
sh ./list-test-routine-images.sh test-routine-images.dat
cd ..

#make output arrays
images=()
templates=()

# define a work function
doExtractions()
{
  for i in `seq 1 $iterations`; do
    while read p; do
      images+=($p)
      echo "Creating Template for" $(basename -- $p)
      template=$(curl --header "Content-Type: application/json" \
        --request POST \
        --silent \
        --data @- \
        http://$host:$port/v1/create-template << JSON
{
  "ImageData": "$(base64 -w 0 $p)"
}
JSON
      )
      echo "worker $1 extracted template: $(echo $template | jq '.Template')"
    done < ./test-routine-images/test-routine-images.dat
  done
}

#for num workers
pids=()
for i in `seq 1 $workers`; do
  echo "starting worker $i"
  doExtractions $i &
  pids[${i}]=$!
done

sleep 1

echo $pids

# wait for workers
worker_count=1
for pid in ${pids[*]}; do
  wait $pid
  echo "worker $worker_count returned"
  worker_count=$(( $worker_count + 1 ))
done
