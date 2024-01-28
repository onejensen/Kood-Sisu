#! /bin/bash

touch a '!' '\' '"'
mkdir '`' 
cp '!' '`'
if [[ "$MOVE_A" == "yes" ]]; then
    mv a '`'
elif [[ "$MOVE_A" == "no" ]]; then
    rm a
elif [[ "$MOVE_A" == "" ]]; then
    :
fi