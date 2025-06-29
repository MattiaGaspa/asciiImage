#!/bin/bash
os=("linux" "windows" "darwin")
arch=("amd64" "arm64")

for o in "${os[@]}"; do
  for a in "${arch[@]}"; do
    echo "Building for $o-$a"
    if [ "$o" == "windows" ]; then
      env GOOS=$o GOARCH=$a go build -o "bin/asciiImage-$o-$a.exe"
    else
      env GOOS=$o GOARCH=$a go build -o "bin/asciiImage-$o-$a"
    fi
  done
done