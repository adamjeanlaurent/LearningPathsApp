#!/bin/bash

# Directory where your Go project is located
project_dir="./src/cmd/main"

# Directory where you want to run go build
build_dir="../../../bin"

# Change to the project directory
cd "$project_dir" || exit

echo "Building...."
# Run go build with the -o flag to specify the output path
go build -o "$build_dir/learningPathsServer"

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful. Binary saved to: ./bin/learningPathsServer"
else
  echo "Build failed."
  exit 1
fi