#!/bin/bash

VERSION="v1.0.2"
BINARY_NAME="simuaudio"

OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

DOWNLOAD_URL="https://github.com/sharon-xa/simuaudio/releases/download/$VERSION/$BINARY_NAME-$OS-$ARCH"

echo "Downloading $BINARY_NAME for $OS/$ARCH from $DOWNLOAD_URL..."
curl -L -o /usr/local/bin/$BINARY_NAME "$DOWNLOAD_URL"

chmod +x /usr/local/bin/$BINARY_NAME

if command -v $BINARY_NAME > /dev/null 2>&1; then
  echo "$BINARY_NAME installed successfully!"
else
  echo "Installation failed."
  exit 1
fi
