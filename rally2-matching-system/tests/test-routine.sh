#!/bin/bash

host=$1
port=$2

if [ -z $host ]; then
  host=localhost
fi

if [ -z $port ]; then
  port=8080
fi

cd test-routine-images
sh ./list-test-routine-images.sh test-routine-images.dat
cd ..

images=() 
templates=()
while read p; do
  images+=($p)
  echo "Creating Template for"$p
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
  echo $template | jq '.Template' >> templates.dat
  templates+=( $(echo $template | jq '.Template') )
done < ./test-routine-images/test-routine-images.dat

echo "Created" ${#templates[@]} "templates from" ${#images[@]} "images"

templateList="\"TemplateList\": ["
c=0
for i in "${templates[@]}"; do
  if [ "$c" -gt 0 ]; then
    templateList="$templateList,"
  fi
  templateList="$templateList { \"Template\": $i }"
  c=$(($c+1))
done
templateList="$templateList ] }"

if [ -f "scores.dat" ] ; then
    rm "scores.dat"
fi

echo -e "image1\timage2\tnormalized_score\tscore" >> scores.dat

c1=0
for i in "${templates[@]}"; do
  tmp=$(curl --header "Content-Type: application/json" \
  --request POST \
  --silent \
  --data @- \
    http://$host:$port/v1/compare-list << JSON
  {
    "SingleTemplate": {
      "Template": $i
    }, 
    $templateList
  }
JSON
  )

  scores=( $(echo $tmp | jq -r '.[].Score') )
  nscores=( $(echo $tmp | jq -r '.[].NormalizedScore') )
  echo -e "\t Compare list generated" ${#scores[@]} "scores"
  c2=0
  for i in ${scores[@]}; do
    echo -e ${images[c1]}"\t"${images[c2]}"\t"${nscores[c2]}"\t"$i >> scores.dat
    c2=$((c2+1))
  done
  c1=$((c1+1))
done

echo "Rally Matching System Test Routine Complete"
