package droplet

import (
	"log"

	"github.com/cloudfoundry/cli/plugin"
)

//Droplet interface
type Droplet interface {
	SaveDroplet(name string, path string) error
	GetDownloader() *Downloader
}

//CFDroplet utility for saving and whatnot.
type CFDroplet struct {
	Cli        plugin.CliConnection
	Downloader Downloader
}

//NewCFDroplet builds a new CF droplet
func NewCFDroplet(cli plugin.CliConnection, downloader Downloader) *CFDroplet {
	log.Printf("Downloader = %v ", downloader)
	return &CFDroplet{
		Cli:        cli,
		Downloader: downloader,
	}
}

//SaveDroplet to the local filesystem.
func (d *CFDroplet) SaveDroplet(name string, path string) error {
	guid, err := d.getGUID(name)
	if nil != err {
		return err
	}
	data, err := d.Downloader.GetDroplet(guid)
	if nil != err {
		return err
	}
	err = d.Downloader.SaveDropletToFile(path, data)
	if nil != err {
		return err
	}
	return nil
}

func (d *CFDroplet) getGUID(appName string) (string, error) {
	app, err := d.Cli.GetApp(appName)
	return app.Guid, err
}

//GetDownloader attached to this dropplet.
func (d *CFDroplet) GetDownloader() *Downloader {
	return &d.Downloader
}
