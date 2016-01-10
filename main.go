package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

//DownloadDropletCmd is the plugin objectx
type DownloadDropletCmd struct {
	Cli plugin.CliConnection
}

//GetMetadata returns metatada to the CF cli
func (cmd *DownloadDropletCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "download-droplet",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "download-droplet",
				HelpText: "Download a droplet",
				UsageDetails: plugin.Usage{
					Usage: "cf download-droplet",
				},
			},
		},
	}
}

//Run runs the plugin
func (cmd *DownloadDropletCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "download-droplet" {
		fmt.Println("Downloading Droplet!")
	}
}

func main() {
	plugin.Start(new(DownloadDropletCmd))
}

//GetGUID the GUID of an app
func (cmd *DownloadDropletCmd) GetGUID(appName string) (string, error) {
	app, err := cmd.Cli.GetApp(appName)
	return app.Guid, err
}

//GetDroplet from CF
func (cmd *DownloadDropletCmd) GetDroplet(guid string) ([]byte, error) {
	downloadURL := "/v2/apps/" + guid + "/droplet/download"
	droplet, err := cmd.Cli.CliCommandWithoutTerminalOutput("curl", downloadURL)
	return []byte(droplet[0]), err
}
