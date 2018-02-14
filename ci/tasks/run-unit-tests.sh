#!/bin/bash -e

OUTPUT=$(pwd)/tested

export GOPATH=$(pwd)/go

cd ${GOPATH}/src/github.com/krujos/download_droplet_plugin

go test ./...

cp -Rf * ${OUTPUT}

ls ${OUTPUT}
