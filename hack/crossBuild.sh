#! /bin/bash

PLATFORMS=(
    linux/amd64
    linux/386
    darwin/amd64
)

for platform in "${PLATFORMS[@]}"
do
    export GOOS=${platform%/*}
    export GOARCH=${platform##*/}

    echo "Building binaries for $GOOS/$GOARCH"
    hack/build.sh
done
