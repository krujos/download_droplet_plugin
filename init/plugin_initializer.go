package init

import "github.com/cloudfoundry/cli/plugin"

//PluginInitializer provides IOC for plugin initialization
type PluginInitializer interface {
	InitializePlugin(cmd *DownloadDropletCmd, cli plugin.CliConnection) error
}
