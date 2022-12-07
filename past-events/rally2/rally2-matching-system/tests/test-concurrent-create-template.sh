#!/bin/bash
#
# This script can be used to perform concurrent requests to the `/v1/create-template` endpoint using custom images. Both
# passed images will be extracted sequentially first, and then again in randomized order by the specified number of concurrent
# workers. A sample range of template bytes will be printed to the console (the range is hardcoded and may need to be
# adjusted depending on the matching system) and color coded by image for human readablity. This is intended to be used
# to debug concurrency related extraction errors using a limited set of local files and an easily configurable
# number of workers.
#
# EXPECTS:
#   * $1 - service address (ex. 172.17.0.2:8080)
#   * $2 - filename1 (ex. ~/Pictures/S001_Right4.png)
#   * $3 - filename2 (ex. ~/Pictures/S002_Right4.png)
#   * $4 - number of workers (ex. 3)

# define output colors
COLORS=('\e[1;34m%s\e[0m' '\e[1;32m%s\e[0m')

# check input params
SERVICE=$1
IMG_FILE_1=$2
IMG_FILE_2=$3
NUM_WORKERS=$4


if [ -z $SERVICE ]; then
  echo "service address must be provided"
  exit 1
fi
if [ -z $IMG_FILE_1 ]; then
  echo "missing one or more image files in params"
  exit 2
fi
if [ -z $IMG_FILE_2 ]; then
  echo "missing one or more image files in params"
  exit 2
fi
if [ -z $NUM_WORKERS ]; then
  echo "number of workers must be provided"
  exit 3
fi

# print vars
printf "AlgoServer:\t%s\n" $SERVICE
printf "${COLORS[0]}:\t%s\n" "Image1" $IMG_FILE_1
printf "${COLORS[1]}:\t%s\n\n" "Image2" $IMG_FILE_2

# get the image bytes as a base64 string
IMGS=()
IMGS[0]="$(cat $IMG_FILE_1 | base64 -w 0)"
IMGS[1]="$(cat $IMG_FILE_2 | base64 -w 0)"

# define a func for generating templates
#  EXPECTS:
#    * $1 name (in the format <description>_<#>)
doExtractions()
{
  # randomize the sequence to make sure the workers end up on different images
  if [ $(( $RANDOM % 2 )) -eq 0 ]; then
    SEQ=(0 1)
  else
    SEQ=(1 0)
  fi

  # for each image
  for INDEX in ${SEQ[*]}; do
    # do the extraction
    RESP=$(curl --header "Content-Type: application/json" \
      --request POST \
      --silent \
      --data @- \
      $SERVICE/v1/create-template << JSON
        {
          "ImageData": "${IMGS[${INDEX}]}"
        }
JSON
    )

    # print a sample of template bytes (adjust/remove the indices in the jq filter if needed)
    template=$(echo $RESP | jq .Template[300:400])
    printf "%s_${COLORS[${INDEX}]}:\t${COLORS[${INDEX}]}\n" "$1" "$(printf "IMG%d" $(( $INDEX + 1 )))" $template
  done
}

# get the expected templates by extracting them sequentially
doExtractions "sequential"
echo "~~~~~~~~~~~~~~~~~~~~~"

# launch background "workers"
pids=()
for i in `seq 1 $NUM_WORKERS`; do
  doExtractions $(printf "con_worker%02d" $i) &
  pids[${i}]=$!
done

# wait for all workers
for pid in ${pids[*]}; do
  wait $pid
done

printf "\nAll done!\n"
