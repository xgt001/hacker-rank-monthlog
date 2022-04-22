#!/usr/bin/env bash
read k
printf %.3f $(bc -l <<< "$k")

