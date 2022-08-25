#!/usr/bin/env bash
#read -r n
i=0
declare -a arr=()
until [[ $i -eq $n ]] ;do
  read -a arr
  i+=$(( i + 1 ))
done
echo "${!arr[@]}"
