#! /bin/bash

count=$(find . -type f -o -type d | wc -l)
result=$((count*5))
printf "\t\vTotal files * 5: $result\v\n"