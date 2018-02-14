#!/bin/bash -e

OUTPUT=$(pwd)/dependencies

export GOPATH=$(pwd)/go
export PATH=$GOPATH/bin:$PATH

go get -u github.com/golang/dep/cmd/dep
go install github.com/golang/dep

cd ${GOPATH}/src/github.com/krujos/download_droplet_plugin

dep ensure

cp -Rf * ${OUTPUT}

ls ${OUTPUT}
