#!/bin/bash -e

OUTPUT=$(pwd)/tested

export GOPATH=$(pwd)/go
export PATH=$GOPATH/bin:$PATH

cd ${GOPATH}/src/github.com/krujos/download_droplet_plugin

go test ./...

cp -Rf * ${OUTPUT}

ls ${OUTPUT}
