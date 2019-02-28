#!/bin/bash

echo -e "\nRally Matching System Test Routine Started..."

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

# Create templates

images=()
templates=()
start=$(date +%s%N | cut -b1-13)
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
  templates+=( $(echo $template | jq '.Template') )
done < ./test-routine-images/test-routine-images.dat
end=$(date +%s%N | cut -b1-13)

n=${#templates[@]}
echo "Created" $n "templates from" ${#images[@]} "images"
avgr=$(bc <<< "scale = 10; ($end-$start) / $n")
echo -e "Average extraction time" $avgr "ms per extraction\n"

# Create a compare list

templateList="\"TemplateList\": ["
c=0
for i in "${templates[@]}"; do
  if [ "$c" -gt 0 ]; then
    templateList="$templateList,"
  fi
  templateList="$templateList { \"Template\": $i }"
  c=$(($c+1))
done
templateList="$templateList ]"

# Output the header of a scores file

if [ -f "scores.dat" ] ; then
    rm "scores.dat"
fi
echo -e "image1\timage2\tnormalized_score\tscore" >> scores.dat

# Execute template comparisons and save the scores

allScores=()
start=$(date +%s%N | cut -b1-13)
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
  echo "Compare list generated" ${#scores[@]} "scores"
  c2=0
  for i in ${scores[@]}; do
    echo -e ${images[c1]}"\t"${images[c2]}"\t"${nscores[c2]}"\t"$i >> scores.dat
    c2=$((c2+1))
  done
  c1=$((c1+1))
  allScores+=( ${scores[@]} )
done
end=$(date +%s%N | cut -b1-13)

echo "Performed" $((c1*c2)) "comparisons"
avgr=$(bc <<< "scale = 10; ($end-$start) / ($c1*$c2)")
echo -e "Average comparison time" $avgr "ms per comparison\n"

# Print the score matrix
echo "Score Matrix:"
index=0
numImgs=${#images[@]}
for s in ${allScores[@]}; do
    if ! (($index % $numImgs)); then
      if ((index !=0)); then
         echo ""
      fi
  fi
  printf "%0.3f " $s
  ((index++))
done

echo -e "\n\n... Rally Matching System Test Routine Complete\n"
