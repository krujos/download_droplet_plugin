Download Droplet Plugin
===

Sometimes you just need to know what Cloud Foundry is working with. This plugin
downloads the current droplet for an app to your local file system. This works
even if the app won't start, the only requirement is the stager successfully builds
a droplet. It's useful for troubleshooting and you can't figure out what else to
do (like [this](https://github.com/krujos/goservice)).

[![wercker status](https://app.wercker.com/status/5a81018d727eb6dbc91b5c352d8c0e1e/m "wercker status")](https://app.wercker.com/project/bykey/5a81018d727eb6dbc91b5c352d8c0e1e)

# Usage
```
$ cf download-droplet
Usage: cf download-droplet APP_NAME PATH

$ cf download-droplet gocf /tmp/gocf.tar
Saving gocf's droplet to /tmp/gocf.tar
$ file gocf.tar   
gocf.tar: gzip compressed data, from Unix, last modified: Fri Jan  8 22:13:58 2016

Saving gocf's droplet to /tmp/gocf.tar
➜  /tmp  tar -xvf gocf.tar     
x ./: Can't update time for .
x ./staging_info.yml
x ./logs/
x ./tmp/
x ./app/
x ./app/app.go
x ./app/Procfile
x ./app/.profile.d/
x ./app/.profile.d/concurrency.sh
x ./app/.profile.d/go.sh
x ./app/widget.go
x ./app/Godeps/
x ./app/Godeps/Readme
x ./app/Godeps/_workspace/
x ./app/Godeps/_workspace/src/
x ./app/Godeps/_workspace/src/github.com/
x ./app/Godeps/_workspace/src/github.com/julienschmidt/
x ./app/Godeps/_workspace/src/github.com/julienschmidt/httprouter/
x ./app/Godeps/_workspace/src/github.com/julienschmidt/httprouter/README.md
x ./app/Godeps/_workspace/src/github.com/julienschmidt/httprouter/path.go
x ./app/Godeps/_workspace/src/github.com/julienschmidt/httprouter/LICENSE
x ./app/Godeps/_workspace/src/github.com/julienschmidt/httprouter/router.go
x ./app/Godeps/_workspace/src/github.com/julienschmidt/httprouter/tree.go
x ./app/Godeps/_workspace/src/github.com/julienschmidt/httprouter/.travis.yml
x ./app/Godeps/_workspace/pkg/
x ./app/Godeps/_workspace/pkg/darwin_amd64/
x ./app/Godeps/_workspace/pkg/darwin_amd64/github.com/
x ./app/Godeps/_workspace/pkg/darwin_amd64/github.com/julienschmidt/
x ./app/Godeps/_workspace/pkg/darwin_amd64/github.com/julienschmidt/httprouter.a
x ./app/Godeps/Godeps.json
x ./app/bin/
x ./app/bin/goservice
➜  /tmp

```
# Installation
## Install from CLI
```
$ cf add-plugin-repo CF-Community http://plugins.cloudfoundry.org/
$ cf install-plugin 'Download Droplet' -r CF-Community
```


## Install from Source (need to have [Go](http://golang.org/dl/) installed)
```
$ go get github.com/cloudfoundry/cli
$ go get github.com/krujos/download_droplet_plugin
$ cd $GOPATH/src/github.com/krujos/download_droplet_plugin
$ go build
$ cf install-plugin download_droplet_plugin
```
