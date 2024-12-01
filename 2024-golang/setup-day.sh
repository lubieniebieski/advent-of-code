#!/bin/bash
# Usage: ./setup-day.sh 01

if [ -z "$1" ]; then
    echo "Please provide a day number (e.g., 01, 02, etc.)"
    exit 1
fi

cp -r template "$1"
echo "Created directory for day $1" 