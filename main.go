package main

import (
	"github.com/krujos/download_droplet_plugin/cmd"
)

func main() {
	cmd := cmd.NewDownloadDropletCmd(&cmd.DownloadDropletCmdInitiliazer{})
	cmd.Start()
}
