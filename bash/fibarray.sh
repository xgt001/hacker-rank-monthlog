#!/usr/bin/env bash
set -x
read a b n m

for (( i=0; i<n; i++ ))
do
    fib=$((a + b))
    a=$b
    b=$fib
done

out=$((b%m))

echo "$out"