#!/bin/bash

echo "Running all Advent of Code solutions..."
echo "-------------------------------------"

# Find all directories that are numbers (days)
for dir in [0-9][0-9]; do
    if [ -d "$dir" ] && [ -f "$dir/main.go" ]; then
        echo "Day $dir:"
        cd "$dir"
        go run main.go
        cd ..
        echo "-------------------------------------"
    fi
done
