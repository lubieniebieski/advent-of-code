#!/bin/bash
DAY=$1
DIR="."

for y in `seq -w 2021 2021`; do
  mkdir -p $DIR/$y/solutions
  mkdir -p $DIR/$y/spec
  mkdir -p $DIR/$y/inputs

  for i in `seq -w $DAY`; do
    touch $DIR/$y/solutions/D$(printf %02d $i)-Readme.md
    touch $DIR/$y/solutions/D$(printf %02d $i).rb
    touch $DIR/$y/spec/D$(printf %02d $i)\_spec.rb
    touch $DIR/$y/inputs/D$(printf %02d $i).txt
  done
done
