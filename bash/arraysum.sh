#!/usr/bin/env bash

read -a input_array
sum=0
for i in "${input_array[@]}"; do
  (( sum+=i ))
done
echo "$sum"