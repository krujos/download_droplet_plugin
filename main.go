package main

import "github.com/krujos/download_droplet_plugin/command"

func main() {
	cmd := command.NewDownloadDropletCmd(&command.DownloadDropletCmdInitiliazer{})
	cmd.Start()
}
