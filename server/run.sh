#!/bin/bash

build_path="./bin/learningPathsServer"

# Check if the file exists
if [ -e "$build_path" ]; then
    $build_path
else
    echo "Output binary does not exist: $build_path"
    exit 1  # Exit with an error code
fi