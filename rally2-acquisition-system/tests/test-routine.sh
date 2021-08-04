#!/bin/bash

echo -e "\nRally Acquisition End Point Test Routine Started..."

# Set the mode, required

mode=$1

if [ -z $mode ] || ([ $mode != "face" ] && [ $mode != "iris" ] && [ $mode != "finger" ]); then
  echo "No mode provided, should be face//finger//iris.. exiting"
  exit
fi

# Set the service address and port, if passed

host=$2
port=$3
image=$4

if [ -z $host ]; then
  host=localhost
fi

if [ -z $port ]; then
  port=8080
fi
echo "Mode: " $mode
echo "Host: " $host
echo -e "Port: " $port "\n"
echo "Image: " $image

# test image acquisition end points
# face

if [ $mode == "face" ]; then
	echo "Test acquiring face image " $image
	curl --header "Content-Type: application/json" \
	  --request POST \
	  --data @- \
		http://$host:$port/v1/face-captures << JSON
	  {
		"FaceImageData": "$(base64 -w 0 $image)",
		"StationID": "TestStation"
	  }
JSON
	  exit
fi

#iris

if [ $mode == "iris" ]; then
	echo "Test acquiring right and left iris images " $image
	curl --header "Content-Type: application/json" \
	  --request POST \
	  --data @- \
		http://$host:$port/v1/iris-captures << JSON
	  {
		"LeftIrisImageData": "$(base64 -w 0 $image)",
		"RightIrisImageData": "$(base64 -w 0 $image)",
		"StationID": "TestStation"
	  }
JSON
	  exit
fi

#fingerprint

if [ $mode == "finger" ]; then
	echo "Test acquiring four fingerprint images " $image
	curl --header "Content-Type: application/json" \
	  --request POST \
	  --data @- \
		http://$host:$port/v1/finger-captures << JSON
	  {
	  	"FingerImages": [
			{
				"FingerImageData": "$(base64 -w 0 $image)",
				"FingerType": "RightIndexFinger"
			},
			{
				"FingerImageData": "$(base64 -w 0 $image)",
				"FingerType": "RightMiddleFinger"
			},
			{
				"FingerImageData": "$(base64 -w 0 $image)",
				"FingerType": "RightRingFinger"
			},
			{
				"FingerImageData": "$(base64 -w 0 $image)",
				"FingerType": "RightLittleFinger"
			}
		],
		"StationID": "TestStation"
	  }
JSON
	  exit
fi

echo -e "\n\n... Rally Matching System Test Routine Complete\n"
