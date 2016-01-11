package command

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/download_droplet_plugin/droplet"
)

//DownloadDropletCmd is the plugin objectx
type DownloadDropletCmd struct {
	Drop        droplet.Droplet
	initializer PluginInitializer
}

//PluginInitializer provides IOC for plugin initialization
type PluginInitializer interface {
	InitializePlugin(cmd *DownloadDropletCmd, cli plugin.CliConnection) error
}

//DownloadDropletCmdInitiliazer is the default plugin initilization implementation
type DownloadDropletCmdInitiliazer struct{}

//InitializePlugin default (and real) implementation. We do it this way because
//we don't have all the objects we need (the cli) to initialze it in main, but we
//have to give the CLI back something to call run on, and then we need to be able
//to test it.
func (initalizer *DownloadDropletCmdInitiliazer) InitializePlugin(
	cmd *DownloadDropletCmd, cli plugin.CliConnection) error {
	return nil
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
		cmd.initializer.InitializePlugin(cmd, cli)
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

//Start the plugin
func (cmd *DownloadDropletCmd) Start() {
	plugin.Start(cmd)
}

//NewDownloadDropletCmd constructor / factory
func NewDownloadDropletCmd(initializer PluginInitializer) *DownloadDropletCmd {
	return nil
}
