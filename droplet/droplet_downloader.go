package droplet

import (
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

//Downloader utility class to download droplets.
type Downloader struct {
	Cli    plugin.CliConnection
	Writer FileWriter
}

//FileWriter test shim for writing to a file.
type FileWriter interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

//GetGUID the GUID of an app
func (downloader *Downloader) GetGUID(appName string) (string, error) {
	app, err := downloader.Cli.GetApp(appName)
	return app.Guid, err
}

//GetDroplet from CF
func (downloader *Downloader) GetDroplet(guid string) ([]byte, error) {
	downloadURL := "/v2/apps/" + guid + "/droplet/download"
	droplet, err := downloader.Cli.CliCommandWithoutTerminalOutput("curl", downloadURL)
	return []byte(droplet[0]), err
}

//SaveDropletToFile writes a downloaded droplet to file
func (downloader *Downloader) SaveDropletToFile(filePath string, data []byte) error {
	downloader.Writer.WriteFile(filePath, data, 0644)
	return nil
}
