#!/bin/bash

echo -e "\nTesting Face Matching System Compare Template List Consistency.."

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

# Get three images
image1=test-routine-images/face/1.png
image2=test-routine-images/face/2.png
image3=test-routine-images/face/3.png

# Create template 1
template1=$(curl --header "Content-Type: application/json" \
--request POST \
--silent \
--data @- \
  http://$host:$port/v1/create-template << JSON
{
  "ImageData": "$(base64 -w 0 $image1)"
}
JSON
)
template1=( $(echo $template1 | jq '.Template') )
  
# Create template 2
template2=$(curl --header "Content-Type: application/json" \
--request POST \
--silent \
--data @- \
  http://$host:$port/v1/create-template << JSON
{
  "ImageData": "$(base64 -w 0 $image2)"
}
JSON
)
template2=( $(echo $template2 | jq '.Template') )

# Create template 3
template3=$(curl --header "Content-Type: application/json" \
--request POST \
--silent \
--data @- \
  http://$host:$port/v1/create-template << JSON
{
  "ImageData": "$(base64 -w 0 $image3)"
}
JSON
)
template3=( $(echo $template3 | jq '.Template') )

# Compare 1 and 2
echo "First call to compare-list"
templateList="\"TargetTemplateList\": [ { \"Template\": $template2 } ]"

comp1=$(curl --header "Content-Type: application/json" \
  --request POST \
  --silent \
  --data @- \
    http://$host:$port/v1/compare-list << JSON
  {
    "ProbeTemplate": {
      "Template": $template1
    },
    $templateList
  }
JSON
  )

echo "match(1.png, 2.png):" $(echo $comp1 | jq -r '.[].Score')
echo -e "\n"

# Compare 1 to 2 and 3
echo "Second call to compare-list:"
templateList="\"TargetTemplateList\": [ { \"Template\": $template2 }, {\"Template\": $template3} ]"

comp2=$(curl --header "Content-Type: application/json" \
  --request POST \
  --silent \
  --data @- \
    http://$host:$port/v1/compare-list << JSON
  {
    "ProbeTemplate": {
      "Template": $template1
    },
    $templateList
  }
JSON
  )
  
echo "match(1.png, 2.png):" $(echo $comp2 | jq -r '.[0].Score')
echo "match(1.png, 3.png):" $(echo $comp2 | jq -r '.[1].Score')

echo -e "\n... Face Matching System Compare Template List Consistency Complete\n"

