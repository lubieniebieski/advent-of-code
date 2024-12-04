#!/bin/bash

echo "Running all Advent of Code solutions..."
echo "-------------------------------------"

# Find all day directories
for dir in day[0-9][0-9]; do
    if [ -d "$dir" ] && [ -f "$dir/main.go" ]; then
        echo "$dir:"
        cd "$dir"
        go run main.go
        cd ..
        echo "-------------------------------------"
    fi
done
