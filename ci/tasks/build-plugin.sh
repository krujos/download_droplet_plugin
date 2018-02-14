#!/bin/bash -e

OUTPUT=$(pwd)/built-plugins
version_file=$(pwd)/version/number

export GOPATH=$(pwd)/go
export PATH=$GOPATH/bin:$PATH

cd ${GOPATH}/src/github.com/krujos/download_droplet_plugin

for os in linux windows darwin; do
    suffix=${os}
    if [ "windows" = "${os}" ]; then
        suffix="windows.exe"
    elif [ "darwin" = "${os}" ]; then
        suffix="macosx"
    fi

    GOOS=${os} GOARCH=amd64 CGO_ENABLED=0 \
      go build -ldflags="-X github.com/krujos/download_droplet_plugin/command.version=$(cat ${version_file})" \
      -o ${OUTPUT}/download_droplet_plugin_${suffix}
done

ls ${OUTPUT}
