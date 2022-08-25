#!/usr/bin/env bash
DONE=false
arr=()
until $DONE ;do
  read -r k || DONE=true
  t=$(echo "$k" | sed "s/^./\./" )
  arr+=("$t")
done
echo "${arr[@]}"
