#!/bin/bash

mkdir -p builds

# List of architectures to build for
architectures=("amd64" "arm64")

for arch in "${architectures[@]}"; do
    output_name="builds/simuaudio-linux-$arch"
    echo "Building for linux/$arch..."

    env GOOS=linux GOARCH="$arch" go build -o "$output_name"

    if [ $? -ne 0 ]; then
      echo "Failed to build for linux/$arch"
      exit 1
    fi

    echo "Build successful: $output_name"
done

echo "All builds completed and stored in the 'builds' directory."
