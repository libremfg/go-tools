#!/bin/bash

for item in $(cat targets.txt);do
    PLATFORM=$(echo $item | cut -d ":" -f1)
    ARCH=$(echo $item | cut -d ":" -f2)
    echo "Building $PLATFORM $ARCH"
    if [ "windows" = $PLATFORM ]; then
      env GOARCH=$ARCH GOOS=$PLATFORM go build -o ./builds/sort-schema-v1.0.0-$PLATFORM-$ARCH.exe ./sortSchema.go
    else
      env GOARCH=$ARCH GOOS=$PLATFORM go build -o ./builds/sort-schema-v1.0.0-$PLATFORM-$ARCH ./sortSchema.go
    fi
done