package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/download_droplet_plugin/droplet"
)

//DownloadDropletCmd is the plugin objectx
type DownloadDropletCmd struct {
	Drop droplet.Droplet
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
				HelpText: "Download a droplet to the local machine.",
				UsageDetails: plugin.Usage{
					Usage: "cf download-droplet",
				},
			},
		},
	}
}

func (cmd *DownloadDropletCmd) usageAndExit() {
	fmt.Println("Usage: cf download-droplet APP_NAME PATH")
	os.Exit(1)
}

//Run runs the plugin
func (cmd *DownloadDropletCmd) Run(cli plugin.CliConnection, args []string) {
	log.Println(args)
	if len(args) != 3 {
		cmd.usageAndExit()
	}
	command := args[0]
	appName := args[1]
	path := args[2]
	if command == "download-droplet" {
		fmt.Printf("Saving %s's droplet to %s", appName, path)
		cmd.Drop.SaveDroplet(appName, path)
	} else {
		fmt.Printf("%s is an unknown command.", args[0])
		cmd.usageAndExit()
	}
}

func main() {
	d := new(droplet.CFDroplet)
	cmd := new(DownloadDropletCmd)
	cmd.Drop = d
	plugin.Start(cmd)

}
