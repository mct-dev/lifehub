#!/usr/bin/env bash
# Do not use this script manually. Use Makefile.

source ./scripts/setup-variables.sh

# Override OS/ARCH
if [[ $# -eq 1 ]]; then
  if [[ $1 == "linux/amd64" ]]; then
    GOOS=linux
    GOARCH=amd64
  elif [[ $1 == "linux/ARMv5" ]]; then
    GOOS=linux
    GOARCH=arm
    GOARM=5
  elif [[ $1 == "windows" ]]; then
    GOOS=windows
    GOARCH=amd64
  elif [[ $1 == "macos" ]]; then
    GOOS=darwin
    GOARCH=amd64
  fi
fi

targetBaseName="$LH_BINARIES_PATH/lifehub"
targetVersion="$LH_VERSION"

targetOsArch="-$GOOS-$GOARCH"
if [[ $GOOS == "darwin" ]]; then
  targetOsArch="-macos"
fi

# Define target extention
ext=""
if [[ $GOOS == "windows" ]]; then
  ext=".exe"
fi

# Target
target=$(printf %s%s%s%s "$targetBaseName" "$targetOsArch" "$targetVersion" "$ext")

echo "Building statically linked $target"
go build -o "$target"
