#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <number_of_times>"
    exit 1
fi

num_times=$1

if [ "$num_times" -gt 100 ]; then
    num_times=100
fi

for ((i=1; i<=num_times; i++)); do
    echo "$i times I've printed juanjosecollgarcia"
done