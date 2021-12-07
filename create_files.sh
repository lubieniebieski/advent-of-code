#!/bin/bash
DAY=$1
DIR="."

for y in `seq -w 2021 2021`; do
  mkdir -p $DIR/$y/solutions
  mkdir -p $DIR/$y/spec
  mkdir -p $DIR/$y/inputs

  for i in `seq $DAY`; do
    if [ $i -lt 10 ]; then 
      i="0$i"
    fi
    PREFIX=D$i
    touch $DIR/$y/solutions/$PREFIX-Readme.txt
    touch $DIR/$y/inputs/$PREFIX.txt
    touch $DIR/$y/spec
    SPEC_FILE=$DIR/$y/spec/${PREFIX}_spec.rb
    SOLUTION_FILE=$DIR/$y/solutions/$PREFIX.rb
    cp -n templates/solution.rb $SOLUTION_FILE
    cp -n templates/spec.rb $SPEC_FILE
    sed -i '' -e "s/D00/$PREFIX/g" $SOLUTION_FILE
    sed -i '' -e "s/D00/$PREFIX/g" $SPEC_FILE
  done
done
