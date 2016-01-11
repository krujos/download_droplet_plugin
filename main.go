package main

import (
	"fmt"
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
	command := args[0]
	switch command {
	case "download-droplet":
		if len(args) != 3 {
			cmd.usageAndExit()
		}
		appName := args[1]
		path := args[2]
		fmt.Printf("Saving %s's droplet to %s\n", appName, path)
		err := cmd.Drop.SaveDroplet(appName, path)
		if nil != err {
			fmt.Println(err)
		}
	case "CLI-MESSAGE-UNINSTALL":
		fmt.Println("Thanks for using droplet downloader!")
	default:
		cmd.usageAndExit()
	}
}

func main() {
	d := new(droplet.CFDroplet)
	cmd := new(DownloadDropletCmd)
	cmd.Drop = d
	plugin.Start(cmd)

}
