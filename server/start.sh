#!/bin/bash

# Execute the build script
./build.sh

# Check if the build was successful
if [ $? -eq 0 ]; then
  # Execute the run script
  ./run.sh
else
  echo "Build failed. Aborting run."
fi