package main

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/download_droplet_plugin/droplet"
)

type Droplet interface {
	SaveDroplet(name string, path string) error
}

//Droplet utility for saving and whatnot.
type CFDroplet struct {
	Cli        plugin.CliConnection
	Downloader droplet.Downloader
}

//SaveDroplet to the local filesystem.
func (droplet *CFDroplet) SaveDroplet(name string, path string) error {
	guid, err := droplet.getGUID(name)
	if nil != err {
		return err
	}
	data, err := droplet.Downloader.GetDroplet(guid)
	if nil != err {
		return err
	}
	err = droplet.Downloader.SaveDropletToFile(path, data)
	if nil != err {
		return err
	}
	return nil
}

func (droplet *CFDroplet) getGUID(appName string) (string, error) {
	app, err := droplet.Cli.GetApp(appName)
	return app.Guid, err
}
