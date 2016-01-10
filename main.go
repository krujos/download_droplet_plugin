package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

//DownloadDropletCmd is the plugin objectx
type DownloadDropletCmd struct {
}

type exiter func(code int)

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
				HelpText: "Download a droplet to the local machine.",
				UsageDetails: plugin.Usage{
					Usage: "cf download-droplet",
				},
			},
		},
	}
}

func (cmd *DownloadDropletCmd) usage() {

}

//Run runs the plugin
func (cmd *DownloadDropletCmd) Run(cli plugin.CliConnection, args []string) {
	if len(args) != 3 {
		cmd.usage()
		os.Exit(1)
	}
	command := args[0]
	appName := args[1]
	path := args[2]
	if command == "download-droplet" {
		fmt.Printf("Saving %s's droplet to %s", appName, path)
	} else {
		fmt.Printf("%s is an unknown command.", args[0])
	}
}

func main() {
	plugin.Start(new(DownloadDropletCmd))
}
