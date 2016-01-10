package main

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/download_droplet_plugin/droplet"
)

//Droplet utility for saving and whatnot.
type Droplet struct {
	Cli        plugin.CliConnection
	Downloader droplet.Downloader
}

//SaveDroplet to the local filesystem.
func (droplet *Droplet) SaveDroplet(name string, path string) error {
	_, err := droplet.getGUID(name)
	if nil != err {
		return err
	}
	return nil
}

//GetGUID the GUID of an app
func (droplet *Droplet) getGUID(appName string) (string, error) {
	app, err := droplet.Cli.GetApp(appName)
	return app.Guid, err
}
