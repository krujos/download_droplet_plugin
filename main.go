package main

import (
	"github.com/krujos/download_droplet_plugin/cmd"
	"github.com/krujos/download_droplet_plugin/init"
)

func main() {
	cmd := cmd.NewDownloadDropletCmd(&init.DownloadDropletCmdInitiliazer{})
	cmd.Start()
}
