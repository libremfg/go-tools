#!/bin/bash

for item in $(cat targets.txt);do
    PLATFORM=$(echo $item | cut -d ":" -f1)
    ARCH=$(echo $item | cut -d ":" -f2)
    echo "Building $PLATFORM $ARCH"
    if [ "windows" = $PLATFORM ]; then
      env GOARCH=$ARCH GOOS=$PLATFORM go build -o ./builds/merge-v0.1.0-$PLATFORM-$ARCH.exe ./merge.go
    else
      env GOARCH=$ARCH GOOS=$PLATFORM go build -o ./builds/merge-v0.1.0-$PLATFORM-$ARCH ./merge.go
    fi
done