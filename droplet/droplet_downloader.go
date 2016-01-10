package droplet

import "github.com/cloudfoundry/cli/plugin"

//Downloader utility class to download droplets.
type Downloader struct {
	Cli plugin.CliConnection
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
