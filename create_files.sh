#!/bin/bash

DIR="."

for y in `seq -w 2020 2020`; do
  mkdir -p $DIR/$y

  for i in `seq -w 1 25`; do
    touch $DIR/$y/D$i-Readme.md
    touch $DIR/$y/D$i-01.rb
    touch $DIR/$y/D$i-02.rb
    touch $DIR/$y/D$i-input
  done
done
