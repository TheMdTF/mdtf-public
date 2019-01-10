#!/bin/bash

output_file=$1

if [ -z $output_file ]; then
  output_file=test-routine-images.dat
fi

ls -d -1 $PWD/*.png > $output_file